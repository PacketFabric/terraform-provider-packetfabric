package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAwsHostedMktConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		CreateContext: resourceAwsHostedMktConnCreate,
		ReadContext:   resourceAwsHostedMktConnRead,
		UpdateContext: resourceAwsHostedMktConnUpdate,
		DeleteContext: resourceAwsHostedMktConnDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			"aws_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The AWS account ID to connect with. Must be 12 characters long.",
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
				Description:  "The description of the AWS Marketplace Cloud connection.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired location for the new AWS Hosted Connection.",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port to connect to AWS.",
			},
			"vlan": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid VLAN range is from 4-4094, inclusive.",
			},
			"src_svlan": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Valid S-VLAN range is from 4-4094, inclusive.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The desired zone of the new connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The desired speed of the new connection.\n\t\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsHostedMktConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	awsHosted := extractAwsHostedMktConn(d)
	resp, err := c.CreateAwsHostedConn(awsHosted)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.CustomerUUID)
	return diags
}

func resourceAwsHostedMktConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceAwsHostedMktConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceAwsHostedMktConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceProvisionDelete(ctx, d, m)
}

func extractAwsHostedMktConn(d *schema.ResourceData) packetfabric.HostedAwsConnection {
	hostedMkt := packetfabric.HostedAwsConnection{}
	if awsAccountID, ok := d.GetOk("aws_account_id"); ok {
		hostedMkt.AwsAccountID = awsAccountID.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedMkt.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		hostedMkt.Description = description.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		hostedMkt.Pop = pop.(string)
	}
	if port, ok := d.GetOk("port"); ok {
		hostedMkt.Port = port.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		hostedMkt.Vlan = vlan.(int)
	}
	if srcSvlan, ok := d.GetOk("src_svlan"); ok {
		hostedMkt.SrcSvlan = srcSvlan.(int)
	}
	if zone, ok := d.GetOk("zone"); ok {
		hostedMkt.Zone = zone.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedMkt.Speed = speed.(string)
	}
	return hostedMkt
}
