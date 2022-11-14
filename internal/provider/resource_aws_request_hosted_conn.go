package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The AWS account ID to connect with. Must be 12 characters long.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID for the billing account that should be billed.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the hosted connection should be provisioned (the cloud on-ramp).",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the PacketFabric port you want to connect to AWS. This starts with \"PF-AP-\".",
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
				Description:  "The speed of the new connection.\n\n\tAvailable: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
	d.SetId(resp.CloudCircuitID)
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
	hostedAwsConn := packetfabric.HostedAwsConnection{}
	if awsAccountID, ok := d.GetOk("aws_account_id"); ok {
		hostedAwsConn.AwsAccountID = awsAccountID.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedAwsConn.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		hostedAwsConn.Description = description.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		hostedAwsConn.Pop = pop.(string)
	}
	if port, ok := d.GetOk("port"); ok {
		hostedAwsConn.Port = port.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		hostedAwsConn.Vlan = vlan.(int)
	}
	if srcSvlan, ok := d.GetOk("src_svlan"); ok {
		hostedAwsConn.SrcSvlan = srcSvlan.(int)
	}
	if zone, ok := d.GetOk("zone"); ok {
		hostedAwsConn.Zone = zone.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedAwsConn.Speed = speed.(string)
	}
	return hostedAwsConn
}
