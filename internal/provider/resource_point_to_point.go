package provider

import (
	"context"
	"errors"
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
			"ptp_circuit_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The point-to-point connection ID.",
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
				Type:     schema.TypeSet,
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
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Label value linked to an object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
	time.Sleep(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	if err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			if c.IsPointToPointComplete(resp.PtpUUID) {
				ticker.Stop()
				createOk <- true
			}
		}
	}()
	<-createOk
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("ptp_circuit_id", resp.PtpCircuitID)
		d.SetId(resp.PtpUUID)

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
	ptpUUID := d.Get("id").(string)
	resp, err := c.ReadPointToPoint(ptpUUID)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("account_uuid", resp.Billing.AccountUUID)
		_ = d.Set("ptp_circuit_id", resp.PtpCircuitID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("media", resp.Media)
		_ = d.Set("subscription_term", resp.Billing.SubscriptionTerm)
		_ = d.Set("po_number", resp.PONumber)

		if len(resp.Interfaces) == 2 {
			interfaceA := make(map[string]interface{})
			interfaceA["pop"] = resp.Interfaces[0].Pop
			interfaceA["zone"] = resp.Interfaces[0].Zone
			interfaceA["customer_site_code"] = resp.Interfaces[0].CustomerSiteCode

			interfaceZ := make(map[string]interface{})
			interfaceZ["pop"] = resp.Interfaces[1].Pop
			interfaceZ["zone"] = resp.Interfaces[1].Zone
			interfaceZ["customer_site_code"] = resp.Interfaces[1].CustomerSiteCode

			endpoints := []interface{}{interfaceA, interfaceZ}
			_ = d.Set("endpoints", endpoints)
		}
	}
	// unsetFields: loa, autoneg, published_quote_line_uuid

	labels, err2 := getLabels(c, d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	_ = d.Set("labels", labels)
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

		if _, err := c.UpdatePointToPoint(d.Id(), updatePointToPointData); err != nil {
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
	if ptpUUID := d.Id(); ptpUUID != "" {
		if err := c.DeletePointToPointService(ptpUUID); err != nil {
			return diag.FromErr(err)
		} else {
			deleteOk := make(chan bool)
			defer close(deleteOk)
			ticker := time.NewTicker(30 * time.Second)
			go func() {
				for range ticker.C {
					if c.IsPointToPointDeleteComplete(ptpUUID) {
						ticker.Stop()
						deleteOk <- true
					}
				}
			}()
			<-deleteOk
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
		for _, endpoint := range endpoints.(*schema.Set).List() {
			edps = append(edps, packetfabric.Endpoints{
				Pop:              endpoint.(map[string]interface{})["pop"].(string),
				Zone:             endpoint.(map[string]interface{})["zone"].(string),
				CustomerSiteCode: endpoint.(map[string]interface{})["customer_site_code"].(string),
				Autoneg:          endpoint.(map[string]interface{})["autoneg"].(bool),
				Loa:              endpoint.(map[string]interface{})["loa"].(string),
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
