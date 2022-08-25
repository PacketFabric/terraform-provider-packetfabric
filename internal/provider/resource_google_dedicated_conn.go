package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGoogleDedicatedConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourceGoogleDedicatedConnCreate,
		ReadContext:   resourceGoogleDedicatedConnRead,
		UpdateContext: resourceGoogleDedicatedConnUpdate,
		DeleteContext: resourceGoogleDedicatedConnDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID of the contact that will be billed.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of this connection.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The desired zone of the new connection.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired location for the new connection.",
			},
			"subscription_term": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
				Description:  "The billing term, in months, for this connection.",
			},
			"service_class": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"longhaul", "metro"}, false),
				Description:  "The service lass for the given port, either long haul or metro.",
			},
			"autoneg": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the port auto-negotiates or not, this is currently only possible with 1Gbps ports and the request will fail if specified with 10Gbps.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedGoogleDeicatedOptions(), false),
				Description:  "The desired speed of the new connection.",
			},
			"loa": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsBase64,
				Description:  "A base64 encoded string of a PDF of a LOA.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func speedGoogleDeicatedOptions() []string {
	return []string{"10Gbps", "100Gbps"}
}

func resourceGoogleDedicatedConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	dedicatedConn := extractGoogleDedicatedConn(d)
	expectedResp, err := c.CreateRequestDedicatedGoogleConn(dedicatedConn)
	if err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(10 * time.Second)
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
	return diags
}

func resourceGoogleDedicatedConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceGoogleDedicatedConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceGoogleDedicatedConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCID, ok := d.GetOk("id")
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Google Service Delete",
			Detail:   cloudCidNotFoundDetailsMsg,
		})
		return diags
	}
	err := c.DeleteCloudService(cloudCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	deleteOkCh := make(chan bool)
	defer close(deleteOkCh)
	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetCloudServiceStatus(cloudCID.(string))
	}
	go c.CheckServiceStatus(deleteOkCh, err, fn)
	if !<-deleteOkCh {
		return diag.FromErr(err)
	}
	return diags
}

func extractGoogleDedicatedConn(d *schema.ResourceData) packetfabric.GoogleReqDedicatedConn {
	dedicatedConn := packetfabric.GoogleReqDedicatedConn{}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		dedicatedConn.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		dedicatedConn.Description = description.(string)
	}
	if zone, ok := d.GetOk("zone"); ok {
		dedicatedConn.Zone = zone.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		dedicatedConn.Pop = pop.(string)
	}
	if subTerm, ok := d.GetOk("subscription_term"); ok {
		dedicatedConn.SubscriptionTerm = subTerm.(int)
	}
	if serviceClass, ok := d.GetOk("service_class"); ok {
		dedicatedConn.ServiceClass = serviceClass.(string)
	}
	if autoneg, ok := d.GetOk("autoneg"); ok {
		dedicatedConn.Autoneg = autoneg.(bool)
	}
	if speed, ok := d.GetOk("speed"); ok {
		dedicatedConn.Speed = speed.(string)
	}
	if loa, ok := d.GetOk("loa"); ok {
		dedicatedConn.Loa = loa.(string)
	}
	if publishedQuote, ok := d.GetOk("published_quote_line_uuid"); ok {
		dedicatedConn.PublishedQuoteLineUUID = publishedQuote.(string)
	}
	return dedicatedConn
}
