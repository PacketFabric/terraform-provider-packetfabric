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
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		CreateContext: resourceCreateAwsHostedMkt,
		ReadContext:   resourceThirdPartyVirtualCircuitRead,
		DeleteContext: resourceDeleteAwsHostedMkt,
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
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},
			"aws_account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_AWS_ACCOUNT_ID", nil),
				ValidateFunc: validation.StringIsNotEmpty,
				Description: "The AWS account ID to connect with. Must be 12 characters long. " +
					"Can also be set with the PF_AWS_ACCOUNT_ID environment variable.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the hosted connection should be provisioned (the cloud on-ramp).",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired zone of the new connection",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\", \"100Mbps\", \"200Mbps\", \"300Mbps\", \"400Mbps\", \"500Mbps\", \"1Gbps\", \"2Gbps\", \"5Gbps\", \"10Gbps\"]",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
	if resp != nil {
		d.SetId(resp.VcRequestUUID)
	}
	return diags
}

func resourceDeleteAwsHostedMkt(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	msg, err := c.DeleteHostedMktConnection(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	diags = make(diag.Diagnostics, 0)
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "AWS Hosted marketplace delete result",
		Detail:   msg,
	})
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
