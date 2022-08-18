package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGoogleRequestHostConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGoogleReqHostConnCreate,
		UpdateContext: resourceGoogeReqHostConnUpdate,
		ReadContext:   resourceGoogleReqHostConnRead,
		DeleteContext: resourceGoogleReqHostConnDelete,
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
			"google_pairing_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Google pairing key to use for this connection.",
			},
			"google_vlan_attachment_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Google Vlan attachment name.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of this connection.",
			},
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The port to connect to Google.",
			},
			"vlan": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Valid VLAN range is from 4-4094, inclusive.",
			},
			"src_svlan": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid S-VLAN range is from 4-4094, inclusive.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired location for the new Connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The desired speed of the new connection.\n\t\t Available: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
			},
		},
	}
}

func resourceGoogleReqHostConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	reqConn := extractGoogleReqConn(d)
	_, err := c.CreateRequestHostedGoogleConn(reqConn)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func resourceGoogleReqHostConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceGoogeReqHostConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceGoogleReqHostConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceBackboneDelete(ctx, d, m, c.DeleteBackbone)
}

func extractGoogleReqConn(d *schema.ResourceData) packetfabric.GoogleReqHostedConn {
	return packetfabric.GoogleReqHostedConn{
		AccountUUID:              d.Get("account_uuid").(string),
		GooglePairingKey:         d.Get("google_pairing_key").(string),
		GoogleVlanAttachmentName: d.Get("google_vlan_attachment_name").(string),
		Description:              d.Get("description").(string),
		Pop:                      d.Get("pop").(string),
		Port:                     d.Get("port").(string),
		Vlan:                     d.Get("vlan").(int),
		SrcSvlan:                 d.Get("src_svlan").(int),
		Speed:                    d.Get("speed").(string),
	}
}
