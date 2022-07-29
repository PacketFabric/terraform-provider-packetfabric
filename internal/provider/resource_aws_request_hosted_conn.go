package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsRequestHostConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsReqHostConnCreate,
		UpdateContext: resourceAwsReqHostConnUpdate,
		ReadContext:   resourceAwsReqHostConnRead,
		DeleteContext: resourceAwsServicesDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
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
				Description: "The description of this connection.",
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
				Description: "The desired speed of the new connection.\n\t\t Available: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
			},
		},
	}
}

func resourceAwsReqHostConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	reqConn := extractReqConn(d)
	resp, err := c.CreateAwsHostedConn(reqConn)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Description)
	return diags
}

func resourceAwsReqHostConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceAwsReqHostConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func extractReqConn(d *schema.ResourceData) packetfabric.HostedAwsConnection {
	return packetfabric.HostedAwsConnection{
		AwsAccountID: d.Get("aws_account_id").(string),
		AccountUUID:  d.Get("account_uuid").(string),
		Description:  d.Get("description").(string),
		Pop:          d.Get("pop").(string),
		Port:         d.Get("port").(string),
		Vlan:         d.Get("vlan").(int),
		SrcSvlan:     d.Get("src_svlan").(int),
		Zone:         d.Get("zone").(string),
		Speed:        d.Get("speed").(string),
	}
}
