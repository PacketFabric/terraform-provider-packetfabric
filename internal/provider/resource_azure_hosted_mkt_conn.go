package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAzureHostedMktConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreateAzureHostedMkt,
		ReadContext:   resourceAzureHostedMktRead,
		DeleteContext: resourceDeleteAzureHostedMkt,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"routing_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The routing ID of the marketplace provider that will be receiving this request.\n\n\tExample: TR-1RI-OQ85",
			},
			"market": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The market code (e.g. \"ATL\" or \"DAL\") in which you would like the marketplace provider to provision their side of the connection.\n\n\tIf the marketplace provider has services published in the marketplace, you can use the PacketFabric portal to see which POPs they are in. Simply remove the number from the POP to get the market code (e.g. if they offer services in \"DAL5\", enter \"DAL\" for the market).",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"azure_service_key": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Service Key provided by Microsoft Azure when you created your ExpressRoute circuit.",
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
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\", \"100Mbps\", \"200Mbps\", \"300Mbps\", \"400Mbps\", \"500Mbps\", \"1Gbps\", \"2Gbps\", \"5Gbps\", \"10Gbps\"]",
			},
			"service_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "UUID of the marketplace service being requested.",
			},
		},
	}
}

func resourceCreateAzureHostedMkt(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	hostedAzure := extractAzureHostedMkt(d)
	resp, err := c.CreateAzureHostedMktRequest(hostedAzure)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.VcRequestUUID)
	return diags
}

func resourceAzureHostedMktRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceDeleteAzureHostedMkt(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	vcRequestUUID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid VC Request UUID to delete")
	}
	msg, err := c.DeleteHostedMktConnection(vcRequestUUID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	diags = make(diag.Diagnostics, 0)
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Azure Hosted marketplace delete result",
		Detail:   msg,
	})
	d.SetId("")
	return diags
}

func extractAzureHostedMkt(d *schema.ResourceData) packetfabric.AzureHostedMktReq {
	hostedMkt := packetfabric.AzureHostedMktReq{}
	if routingID, ok := d.GetOk("routing_id"); ok {
		hostedMkt.RoutingID = routingID.(string)
	}
	if market, ok := d.GetOk("market"); ok {
		hostedMkt.Market = market.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		hostedMkt.Description = description.(string)
	}
	if serviceKey, ok := d.GetOk("azure_service_key"); ok {
		hostedMkt.AzureServiceKey = serviceKey.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedMkt.AccountUUID = accountUUID.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedMkt.Speed = speed.(string)
	}
	if serviceUUID, ok := d.GetOk("service_uuid"); ok {
		hostedMkt.ServiceUUID = serviceUUID.(string)
	}
	return hostedMkt
}
