package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGoogleCloudRouterConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGoogleCloudRouterConnCreate,
		ReadContext:   resourceGoogleCloudRouterConnRead,
		UpdateContext: resourceGoogleCloudRouterConnUpdate,
		DeleteContext: resourceGoogleCloudRouterConnDelete,
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
			"circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
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
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set this to true if you intend to use NAT on this connection. ",
			},
			"maybe_dnat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set this to true if you intend to use DNAT on this connection. ",
			},
			"google_pairing_key": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Google pairing key to use for this connection. This is generated when you create your Google Cloud VLAN attachment.",
			},
			"google_vlan_attachment_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Google VLAN attachment name.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The POP in which you want to provision the connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The desired speed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"labels": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Label value linked to an object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceGoogleCloudRouterConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	googleRoute := extractGoogleRouteConn(d)
	if cid, ok := d.GetOk("circuit_id"); ok {
		resp, err := c.CreateGoogleCloudRouterConn(googleRoute, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		createOkCh := make(chan bool)
		defer close(createOkCh)
		fn := func() (*packetfabric.ServiceState, error) {
			return c.GetCloudConnectionStatus(cid.(string), resp.CloudCircuitID)
		}
		go c.CheckServiceStatus(createOkCh, fn)
		if !<-createOkCh {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.CloudCircuitID)

			if labels, ok := d.GetOk("labels"); ok {
				diagnostics, created := createLabels(c, d.Id(), labels)
				if !created {
					return diagnostics
				}
			}
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Circuit ID not present",
			Detail:   "Please provide a valid Circuit ID.",
		})
	}
	return diags
}

func resourceGoogleCloudRouterConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		resp, err := c.ReadCloudRouterConnection(cid.(string), cloudConnCID.(string))
		if err != nil {
			diags = diag.FromErr(err)
			return diags
		}

		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("circuit_id", resp.CloudRouterCircuitID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("pop", resp.Pop)
		_ = d.Set("google_pairing_key", resp.CloudSettings.GooglePairingKey)
		_ = d.Set("google_vlan_attachment_name", resp.CloudSettings.GoogleVlanAttachmentName)
		_ = d.Set("po_number", resp.PONumber)

		if resp.CloudSettings.PublicIP != "" {
			_ = d.Set("is_public", true)
		} else {
			_ = d.Set("is_public", false)
		}
		// unsetFields: published_quote_line_uuid
	}

	labels, err2 := getLabels(c, d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	_ = d.Set("labels", labels)
	return diags
}

func resourceGoogleCloudRouterConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceGoogleCloudRouterConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractGoogleRouteConn(d *schema.ResourceData) packetfabric.GoogleCloudRouterConn {
	googleRoute := packetfabric.GoogleCloudRouterConn{}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		googleRoute.AccountUUID = accountUUID.(string)
	}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		googleRoute.MaybeNat = maybeNat.(bool)
	}
	if maybeDNat, ok := d.GetOk("maybe_dnat"); ok {
		googleRoute.MaybeDNat = maybeDNat.(bool)
	}
	if pairingKey, ok := d.GetOk("google_pairing_key"); ok {
		googleRoute.GooglePairingKey = pairingKey.(string)
	}
	if vlanAttName, ok := d.GetOk("google_vlan_attachment_name"); ok {
		googleRoute.GoogleVlanAttachmentName = vlanAttName.(string)
	}
	if desc, ok := d.GetOk("description"); ok {
		googleRoute.Description = desc.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		googleRoute.Pop = pop.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		googleRoute.Speed = speed.(string)
	}
	if publishedQuoteLine, ok := d.GetOk("published_quote_line_uuid"); ok {
		googleRoute.PublishedQuoteLineUUID = publishedQuoteLine.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		googleRoute.PONumber = poNumber.(string)
	}
	return googleRoute
}
