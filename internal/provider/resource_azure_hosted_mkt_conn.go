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
		UpdateContext: resourceAzureHostedMktUpdate,
		DeleteContext: resourceDeleteAzureHostedMkt,
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
			"routing_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The routing ID of the customer to whom this VC will be connected.\n\t\tExample: PF-1RI-OQ85",
			},
			"market": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The market that the VC will be requested in.\n\t\tExample: ATL",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of the AWS Marketplace Cloud connection.\n\t\tExample: My AWS Marketplace Cloud connection",
			},
			"azure_service_key": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Service Key provided by Microsoft Azure.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The UUID of the contact that will be billed.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The desired zone of the new connection. Optional",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired speed of the new connection.\n\t\tEnum: [ \"50Mbps\", \"100Mbps\", \"200Mbps\", \"300Mbps\", \"400Mbps\", \"500Mbps\", \"1Gbps\", \"2Gbps\", \"5Gbps\", \"10Gbps\" ]",
			},
			"service_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
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

func resourceAzureHostedMktUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceDeleteAzureHostedMkt(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	vcRequestUUID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid VC Request UUID to delete")
	}
	err := c.DeleteRequestedHostedMktService(vcRequestUUID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
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
	if zone, ok := d.GetOk("pop"); ok {
		hostedMkt.Zone = zone.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedMkt.Speed = speed.(string)
	}
	if serviceUUID, ok := d.GetOk("service_uuid"); ok {
		hostedMkt.ServiceUUID = serviceUUID.(string)
	}
	return hostedMkt
}
