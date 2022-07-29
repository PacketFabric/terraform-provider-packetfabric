package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGoogleHostedMktConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGoogleHostedMktConnCreate,
		ReadContext:   resourceGoogleHostedMktConnRead,
		UpdateContext: resourceGoogleHostedMktConnUpdate,
		DeleteContext: resourceGoogleHostedMktConnDelete,
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
				Description:  "The routing ID of the customer to whom this VC will be connected.",
			},
			"market": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The market that the VC will be requested in.",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of the Google Marketplace Cloud connection.",
			},
			"google_pairing_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Google pairing key to use for this connection.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID of the contact that will be billed.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired location for the new Google Hosted Connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired speed of the new connection.\n\t\tEnum: [ \"50Mbps\", \"100Mbps\", \"200Mbps\", \"300Mbps\", \"400Mbps\", \"500Mbps\", \"1Gbps\", \"2Gbps\", \"5Gbps\", \"10Gbps\" ]",
			},
		},
	}
}

func resourceGoogleHostedMktConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	hostedGoogle := extractGoogleHostedMkt(d)
	resp, err := c.CreateRequestHostedGoogleMktConn(hostedGoogle)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.VcRequestUUID)
	return diags
}

func resourceGoogleHostedMktConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceGoogleHostedMktConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceGoogleHostedMktConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func extractGoogleHostedMkt(d *schema.ResourceData) packetfabric.GoogleMktCloudConn {
	hostedGoogle := packetfabric.GoogleMktCloudConn{}
	if routingID, ok := d.GetOk("routing_id"); ok {
		hostedGoogle.RoutingID = routingID.(string)
	}
	if market, ok := d.GetOk("market"); ok {
		hostedGoogle.Market = market.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		hostedGoogle.Description = description.(string)
	}
	if pairingKey, ok := d.GetOk("google_pairing_key"); ok {
		hostedGoogle.GooglePairingKey = pairingKey.(string)
	}
	if vlanAttachment, ok := d.GetOk("google_vlan_attachment_name"); ok {
		hostedGoogle.GoogleVlanAttachmentName = vlanAttachment.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedGoogle.AccountUUID = accountUUID.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		hostedGoogle.Pop = pop.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedGoogle.Speed = speed.(string)
	}
	return hostedGoogle
}
