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
		DeleteContext: resourceGoogleHostedMktConnDelete,
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
			"google_pairing_key": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Google pairing key to use for this connection. This is provided when you create the VLAN attachment from the Google Cloud console.",
			},
			"google_vlan_attachment_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The name you used for your VLAN attachment in Google.",
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
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the hosted connection should be provisioned (the cloud on-ramp).",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\", \"100Mbps\", \"200Mbps\", \"300Mbps\", \"400Mbps\", \"500Mbps\", \"1Gbps\", \"2Gbps\", \"5Gbps\", \"10Gbps\"]",
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

func resourceGoogleHostedMktConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		Summary:  "Google Hosted marketplace delete result",
		Detail:   msg,
	})
	d.SetId("")
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
