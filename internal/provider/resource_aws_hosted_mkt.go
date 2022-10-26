package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAwsHostedMkt() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreateAwsHostedMkt,
		UpdateContext: resourceUpdateAwsHostedMkt,
		ReadContext:   resourceReadAwsHostedMkt,
		DeleteContext: resourceDeleteAwsHostedMkt,
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
				Description:  "The routing ID of the marketplace provider that will be receiving this request.\n\n\tExample: TR-1RI-OQ85",
			},
			"market": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The market code (e.g. \"ATL\" or \"DAL\") in which you would like the marketplace provider to provision their side of the connection.\n\n\tIf the marketplace provider has services published in the marketplace, you can use the PacketFabric portal to see which POPs they are in. Simply remove the number from the POP to get the market code (e.g. if they offer services in \"DAL5\", enter \"DAL\" for the market).",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The UUID for the billing account that should be billed. This is your billing account, not the marketplace provider's.",
			},
			"aws_account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The AWS account ID to connect with. Must be 12 characters long.",
			},
			"pop": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the hosted connection should be provisioned (the cloud on-ramp).",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired zone of the new connection",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\", \"100Mbps\", \"200Mbps\", \"300Mbps\", \"400Mbps\", \"500Mbps\", \"1Gbps\", \"2Gbps\", \"5Gbps\", \"10Gbps\"]",
			},
		},
	}
}

func resourceCreateAwsHostedMkt(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	hostedMkt := extractHostedMkt(d)
	resp, err := c.CreateAwsHostedMkt(hostedMkt)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.VcRequestUUID)
	return diags
}

func resourceReadAwsHostedMkt(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceUpdateAwsHostedMkt(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceDeleteAwsHostedMkt(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	err := c.DeleteHostedMktConnection(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func extractHostedMkt(d *schema.ResourceData) packetfabric.ServiceAws {
	hostedMkt := packetfabric.ServiceAws{}
	if routingID, ok := d.GetOk("routing_id"); ok {
		hostedMkt.RoutingID = routingID.(string)
	}
	if market, ok := d.GetOk("market"); ok {
		hostedMkt.Market = market.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedMkt.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		hostedMkt.Description = description.(string)
	}
	if awsAccountID, ok := d.GetOk("aws_account_id"); ok {
		hostedMkt.AwsAccountID = awsAccountID.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		hostedMkt.Pop = pop.(string)
	}
	if zone, ok := d.GetOk("zone"); ok {
		hostedMkt.Zone = zone.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedMkt.Speed = speed.(string)
	}
	return hostedMkt
}
