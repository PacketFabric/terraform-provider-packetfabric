package provider

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePointToPoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePointToPointCreate,
		ReadContext:   resourcePointToPointRead,
		UpdateContext: resourcePointToPointUpdate,
		DeleteContext: resourcePointToPointDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ptp_uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The point-to-point connection UUID.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"1Gbps", "10Gbps", "40Gbps", "100Gbps"}, true),
				Description:  "The capacity for this connection.\n\n\tEnum: [\"1Gbps\" \"10Gbps\" \"40Gbps\" \"100Gbps\"]",
			},
			"media": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(pointToPointMediaOptions(), true),
				Description:  "Optic media type.\n\n\tEnum: [\"LX\" \"EX\" \"ZX\" \"LR\" \"ER\" \"ER DWDM\" \"ZR\" \"ZR DWDM\" \"LR4\" \"ER4\" \"CWDM4\" \"LR4\" \"ER4 Lite\"]",
			},
			"endpoints": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pop": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Point of presence in which the port should be located.",
						},
						"zone": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Availability zone of the port.",
						},
						"customer_site_code": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Unique site code of the customer's equipment.",
						},
						"autoneg": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Only applicable to 1Gbps ports. Controls whether auto negotiation is on (true) or off (false). The request will fail if specified with ports greater than 1Gbps.",
						},
						"loa": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringIsBase64,
							Description:  "A base64 encoded string of a PDF of a LOA.",
						},
						"port_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The circuit ID for the port. This starts with \"PF-AP-\"",
						},
					},
				},
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},
			"subscription_term": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
				Description:  "Duration of the subscription in months\n\n\tEnum [\"1\" \"12\" \"24\" \"36\"]",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Label value linked to an object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"etl": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Early Termination Liability (ETL) fees apply when terminating a service before its term ends. ETL is prorated to the remaining contract days.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePointToPointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ptpService := extractPtpService(d)
	resp, err := c.CreatePointToPointService(ptpService)
	if err != nil {
		return diag.FromErr(err)
	}

	host := os.Getenv("PF_HOST")
	testingInLab := strings.Contains(host, "api.dev")

	if testingInLab {
		time.Sleep(time.Duration(30) * time.Second)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "In the dev environment, need to add a delay before checking the status.",
		})
	}

	if err2 := checkPtpStatus(c, resp.PtpCircuitID); err2 != nil {
		return diag.FromErr(err2)
	}
	if resp != nil {
		_ = d.Set("ptp_uuid", resp.PtpUUID)
		d.SetId(resp.PtpCircuitID)

		if labels, ok := d.GetOk("labels"); ok {
			diagnostics, created := createLabels(c, d.Id(), labels)
			if !created {
				return diagnostics
			}
		}
	}
	return diags
}

func resourcePointToPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.ReadPointToPoint(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("account_uuid", resp.Billing.AccountUUID)
		_ = d.Set("ptp_uuid", resp.PtpUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("media", resp.Media)
		_ = d.Set("subscription_term", resp.Billing.SubscriptionTerm)
		if _, ok := d.GetOk("po_number"); ok {
			_ = d.Set("po_number", resp.PONumber)
		}

		if len(resp.Interfaces) == 2 {
			interface1 := make(map[string]interface{})
			interface1["pop"] = resp.Interfaces[0].Pop
			interface1["zone"] = resp.Interfaces[0].Zone
			interface1["customer_site_code"] = resp.Interfaces[0].CustomerSiteCode
			interface1["port_circuit_id"] = resp.Interfaces[0].PortCircuitID

			interface2 := make(map[string]interface{})
			interface2["pop"] = resp.Interfaces[1].Pop
			interface2["zone"] = resp.Interfaces[1].Zone
			interface2["customer_site_code"] = resp.Interfaces[1].CustomerSiteCode
			interface2["port_circuit_id"] = resp.Interfaces[1].PortCircuitID

			endpoints := []interface{}{interface1, interface2}
			_ = d.Set("endpoints", endpoints)
		}
	}
	// unsetFields: loa, autoneg, published_quote_line_uuid
	if _, ok := d.GetOk("labels"); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set("labels", labels)
	}

	etl, err3 := c.GetEarlyTerminationLiability(d.Id())
	if err3 != nil {
		return diag.FromErr(err3)
	}
	if etl > 0 {
		_ = d.Set("etl", etl)
	}
	return diags
}

func resourcePointToPointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	if d.HasChanges([]string{"po_number", "description"}...) {
		updatePointToPointData := packetfabric.UpdatePointToPointData{}
		desc, ok := d.GetOk("description")
		if !ok {
			return diag.FromErr(errors.New("please enter a description"))
		}
		updatePointToPointData.Description = desc.(string)

		poNumber, ok := d.GetOk("po_number")
		if !ok {
			return diag.FromErr(errors.New("please enter a purchase order number"))
		}
		updatePointToPointData.PONumber = poNumber.(string)
		ptpUuid := d.Get("ptp_uuid").(string) // must use the UUID to update the PTP
		if _, err := c.UpdatePointToPoint(ptpUuid, updatePointToPointData); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("subscription_term") {
		if subTerm, ok := d.GetOk("subscription_term"); ok {
			billing := packetfabric.BillingUpgrade{
				SubscriptionTerm: subTerm.(int),
			}
			cID := d.Get("ptp_circuit_id").(string)
			if _, err := c.ModifyBilling(cID, billing); err != nil {
				return diag.FromErr(err)
			}
			_ = d.Set("subscription_term", subTerm.(int))
		} else {
			return diag.Errorf("please provide a subscription term")
		}
	}

	if d.HasChange("labels") {
		labels := d.Get("labels")
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}
	return diags
}

func resourcePointToPointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	host := os.Getenv("PF_HOST")
	testingInLab := strings.Contains(host, "api.dev")

	if testingInLab {
		endpoints := d.Get("endpoints").([]interface{})
		for _, v := range endpoints {
			endpoint := v.(map[string]interface{})
			portCircuitID := endpoint["port_circuit_id"].(string)
			if toggleErr := _togglePortStatus(c, false, portCircuitID); toggleErr != nil {
				return diag.FromErr(toggleErr)
			}
		}
		time.Sleep(time.Duration(180) * time.Second)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "In the dev environment, ports are disabled prior to deletion.",
		})
	}
	etlDiags, err := addETLWarning(c, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	diags = append(diags, etlDiags...)
	ptpUuid := d.Get("ptp_uuid").(string) // must use the UUID to delete the PTP
	if err := c.DeletePointToPointService(ptpUuid); err != nil {
		return diag.FromErr(err)
	} else {
		if err := checkPtpStatus(c, d.Id()); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return diags
}

func extractPtpService(d *schema.ResourceData) packetfabric.PointToPoint {
	ptpService := packetfabric.PointToPoint{}
	if description, ok := d.GetOk("description"); ok {
		ptpService.Description = description.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		ptpService.Speed = speed.(string)
	}
	if media, ok := d.GetOk("media"); ok {
		ptpService.Media = media.(string)
	}
	if endpoints, ok := d.GetOk("endpoints"); ok {
		edps := make([]packetfabric.Endpoints, 0)
		for _, endpoint := range endpoints.([]interface{}) {
			edps = append(edps, packetfabric.Endpoints{
				Pop:              endpoint.(map[string]interface{})["pop"].(string),
				Zone:             endpoint.(map[string]interface{})["zone"].(string),
				CustomerSiteCode: endpoint.(map[string]interface{})["customer_site_code"].(string),
				Autoneg:          endpoint.(map[string]interface{})["autoneg"].(bool),
				Loa:              endpoint.(map[string]interface{})["loa"].(string),
				PortCircuitID:    endpoint.(map[string]interface{})["port_circuit_id"].(string),
			})
		}
		ptpService.Endpoints = edps
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		ptpService.AccountUUID = accountUUID.(string)
	}
	if subTerm, ok := d.GetOk("subscription_term"); ok {
		ptpService.SubscriptionTerm = subTerm.(int)
	}
	if quote, ok := d.GetOk("published_quote_line_uuid"); ok {
		ptpService.PublishedQuoteLineUUID = quote.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		ptpService.PONumber = poNumber.(string)
	}
	return ptpService
}

func pointToPointMediaOptions() []string {
	return []string{"LX", "EX", "ZX", "LR",
		"ER", "ER DWDM", "ZR",
		"ZE DWDM", "LR4", "ER4",
		"CWDM4", "LR4", "ER4 Lite"}
}

func checkPtpStatus(c *packetfabric.PFClient, cid string) error {
	statusOk := make(chan bool)
	defer close(statusOk)

	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetPointToPointStatus(cid)
	}
	go c.CheckServiceStatus(statusOk, fn)
	time.Sleep(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	if !<-statusOk {
		return fmt.Errorf("failed to retrieve the status for %s", cid)
	}
	return nil
}
