package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		ReadContext:   resourceAwsServicesRead,
		UpdateContext: resourceAwsServicesUpdate,
		DeleteContext: resourceAwsServicesDelete,
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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The UUID of the contact that will be billed.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of the AWS Marketplace Cloud connection.",
			},
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The desired location for the new AWS Hosted Connection.",
			},
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The port to connect to AWS.",
			},
			"vlan": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Valid VLAN range is from 4-4094, inclusive.",
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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The desired speed of the new connection.\n\t\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
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

func extractAwsHostedMktConn(d *schema.ResourceData) packetfabric.HostedAwsConnection {
	hostedMkt := packetfabric.HostedAwsConnection{
		AwsAccountID: d.Get("aws_account_id").(string),
		AccountUUID:  d.Get("account_uuid").(string),
		Description:  d.Get("description").(string),
		Pop:          d.Get("pop").(string),
		Port:         d.Get("port").(string),
		Vlan:         d.Get("vlan").(int64),
		SrcSvlan:     d.Get("src_vlan").(int64),
		Zone:         d.Get("zone").(string),
		Speed:        d.Get("speed").(string),
	}
	return hostedMkt
}
