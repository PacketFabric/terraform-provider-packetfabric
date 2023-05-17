package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAzureReqExpressDedicatedConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourceAzureReqExpressDedicatedConnCreate,
		ReadContext:   resourceAzureReqExpressDedicatedConnRead,
		UpdateContext: resourceAzureReqExpressDedicatedConnUpdate,
		DeleteContext: resourceAzureReqExpressDedicatedConnDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
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
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired zone of the new connection.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the dedicated port should be provisioned (the cloud on-ramp).",
			},
			"subscription_term": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
				Description:  "The billing term, in months, for this connection.\n\n\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
			},
			"service_class": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"longhaul", "metro"}, true),
				Description:  "The service class for the given port, either long haul or metro. Specify metro if the cloud on-ramp (the `pop`) is in the same market as the source ports (the ports to which you will be building out virtual circuits).\n\n\tEnum: [\"longhaul\" \"metro\"]",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"10Gbps", "100Gbps"}, true),
				Description:  "The capacity of the dedicated cloud port.\n\n\tEnum: [\"10Gbps\" \"100Gbps\"]",
			},
			"loa": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsBase64,
				Description:  "A base64 encoded string of a PDF for the LOA that you generated from the Azure portal",
			},
			"encapsulation": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"dot1q", "qinq"}, true),
				Description:  "Specify either QinQ and Dot1Q encapsulation.\n\n\tEnum: [\"dot1q\" \"qinq\"]",
			},
			"port_category": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"primary", "secondary"}, true),
				Description:  "Whether you intend to use this port for the primary or secondary connection in your ExpressRoute Direct circuit.\n\n\tEnum: [\"primary\" \"secondary\"]",
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

func resourceAzureReqExpressDedicatedConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	azureExpress := extractAzureExpressDedicatedConn(d)
	expectedResp, err := c.CreateAzureExpressRouteDedicated(azureExpress)
	if err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			dedicatedConns, err := c.GetCurrentCustomersDedicated()
			if dedicatedConns != nil && err == nil && len(dedicatedConns) > 0 {
				for _, conn := range dedicatedConns {
					if expectedResp.UUID == conn.UUID && conn.State == "active" {
						expectedResp.CloudCircuitID = conn.CloudCircuitID
						ticker.Stop()
						createOk <- true
					}
				}
			}
		}
	}()
	<-createOk
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(expectedResp.CloudCircuitID)

	if labels, ok := d.GetOk("labels"); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}
	return diags
}

func resourceAzureReqExpressDedicatedConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("pop", resp.Pop)
		_ = d.Set("subscription_term", resp.SubscriptionTerm)
		_ = d.Set("service_class", resp.ServiceClass)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("port_category", resp.AzurePortCategory)
		_ = d.Set("encapsulation", resp.Settings.AzureEncapsulation)
	}
	resp2, err2 := c.GetPortByCID(d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	if resp2 != nil {
		if _, ok := d.GetOk("zone"); ok {
			_ = d.Set("zone", resp2.Zone)
		}
		if _, ok := d.GetOk("po_number"); ok {
			_ = d.Set("po_number", resp2.PONumber)
		}
	}
	// unsetFields: loa

	if _, ok := d.GetOk("labels"); ok {
		labels, err3 := getLabels(c, d.Id())
		if err3 != nil {
			return diag.FromErr(err3)
		}
		_ = d.Set("labels", labels)
	}

	etl, err4 := c.GetEarlyTerminationLiability(d.Id())
	if err4 != nil {
		return diag.FromErr(err4)
	}
	if etl > 0 {
		_ = d.Set("etl", etl)
	}

	return diags
}

func resourceAzureReqExpressDedicatedConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesDedicatedUpdate(ctx, d, m)
}

func resourceAzureReqExpressDedicatedConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudSourceDelete(ctx, d, m, "Google Service Delete")
}

func extractAzureExpressDedicatedConn(d *schema.ResourceData) packetfabric.AzureExpressRouteDedicated {
	azureExpress := packetfabric.AzureExpressRouteDedicated{
		AccountUUID:            d.Get("account_uuid").(string),
		Description:            d.Get("description").(string),
		Zone:                   d.Get("zone").(string),
		Pop:                    d.Get("pop").(string),
		SubscriptionTerm:       d.Get("subscription_term").(int),
		ServiceClass:           d.Get("service_class").(string),
		Speed:                  d.Get("speed").(string),
		Loa:                    d.Get("loa").(string),
		Encapsulation:          d.Get("encapsulation").(string),
		PortCategory:           d.Get("port_category").(string),
		PublishedQuoteLineUUID: d.Get("published_quote_line_uuid").(string),
		PONumber:               d.Get("po_number").(string),
	}
	return azureExpress
}
