package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIBMCloudRouteConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCloudRouteConnCreate,
		ReadContext:   resourceIBMCloudRouteConnRead,
		UpdateContext: resourceIBMCloudRouteConnUpdate,
		DeleteContext: resourceIBMCloudRouteConnDelete,
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
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set this to true if you intend to use NAT on this connection. Defaults: false",
			},
			"maybe_dnat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set this to true if you intend to use DNAT on this connection. Defaults: false",
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
			"ibm_account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_IBM_ACCOUNT_ID", nil),
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description: "Your IBM account ID. " +
					"Can also be set with the PF_IBM_ACCOUNT_ID environment variable.",
			},
			"ibm_bgp_asn": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Enter an ASN to use with your BGP session. This should be the same ASN you used for your Cloud Router.",
			},
			"ibm_bgp_cer_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The IP address in CIDR format for the PacketFabric-side router in the BGP session. If you do not specify an address, IBM will assign one on your behalf.",
			},
			"ibm_bgp_ibm_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The IP address in CIDR format for the IBM-side router in the BGP session. If you do not specify an address, IBM will assign one on your behalf. See the documentation for information on which IP ranges are allowed.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of this connection. This will appear as the connection name from the IBM side.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which you want to provision the connection (the on-ramp).",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired availability zone of the connection.",
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
			"gateway_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IBM Gateway ID.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceIBMCloudRouteConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		ibmRouter := extractIBMCloudRouterConn(d)
		resp, err := c.CreateIBMCloudRouteConn(ibmRouter, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		// Skip status check as the status will show active when the connection request has been accepted
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

func resourceIBMCloudRouteConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		_ = d.Set("ibm_account_id", resp.CloudSettings.AccountID)
		_ = d.Set("ibm_bgp_asn", resp.CloudSettings.BgpAsn)
		if _, ok := d.GetOk("ibm_bgp_cer_cidr"); ok {
			_ = d.Set("ibm_bgp_cer_cidr", resp.CloudSettings.BgpCerCidr)
		}
		if _, ok := d.GetOk("ibm_bgp_ibm_cidr"); ok {
			_ = d.Set("ibm_bgp_ibm_cidr", resp.CloudSettings.BgpIbmCidr)
		}
		_ = d.Set("description", resp.Description)
		_ = d.Set("pop", resp.CloudProvider.Pop)
		_ = d.Set("speed", resp.Speed)
		if _, ok := d.GetOk("zone"); ok {
			_ = d.Set("zone", resp.Zone)
		}
		if _, ok := d.GetOk("po_number"); ok {
			_ = d.Set("po_number", resp.PONumber)
		}
		if resp.CloudSettings.GatewayID == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Incomplete Cloud Information",
				Detail:   "The gateway_id is currently unavailable.",
			})
			return diags
		} else {
			_ = d.Set("gateway_id", resp.CloudSettings.GatewayID)
		}
		// unsetFields: published_quote_line_uuid
	}

	if _, ok := d.GetOk("labels"); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set("labels", labels)
	}
	return diags
}

func resourceIBMCloudRouteConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceIBMCloudRouteConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractIBMCloudRouterConn(d *schema.ResourceData) packetfabric.IBMCloudRouterConn {
	ibmRouter := packetfabric.IBMCloudRouterConn{}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		ibmRouter.MaybeNat = maybeNat.(bool)
	}
	if maybeDNat, ok := d.GetOk("maybe_dnat"); ok {
		ibmRouter.MaybeDNat = maybeDNat.(bool)
	}
	if ibmAccountID, ok := d.GetOk("ibm_account_id"); ok {
		ibmRouter.IbmAccountID = ibmAccountID.(string)
	}
	if ibmBgpAsn, ok := d.GetOk("ibm_bgp_asn"); ok {
		ibmRouter.IbmBgpAsn = ibmBgpAsn.(int)
	}
	if ibmBgpCerCidr, ok := d.GetOk("ibm_bgp_cer_cidr"); ok {
		ibmRouter.IbmBgpCerCidr = ibmBgpCerCidr.(string)
	}
	if ibmBgpIbmCidr, ok := d.GetOk("ibm_bgp_ibm_cidr"); ok {
		ibmRouter.IbmBgpIbmCidr = ibmBgpIbmCidr.(string)
	}
	if desc, ok := d.GetOk("description"); ok {
		ibmRouter.Description = desc.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		ibmRouter.AccountUUID = accountUUID.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		ibmRouter.Pop = pop.(string)
	}
	if zone, ok := d.GetOk("zone"); ok {
		ibmRouter.Zone = zone.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		ibmRouter.Speed = speed.(string)
	}
	if publishedQuote, ok := d.GetOk("published_quote_line_uuid"); ok {
		ibmRouter.PublishedQuoteLineUUID = publishedQuote.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		ibmRouter.PONumber = poNumber.(string)
	}
	return ibmRouter
}
