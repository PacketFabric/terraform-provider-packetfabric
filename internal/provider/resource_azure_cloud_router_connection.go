package provider

import (
	"context"
	"errors"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAzureExpressRouteConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAzureExpressRouteConnCreate,
		ReadContext:   resourceAzureExpressRouteConnRead,
		UpdateContext: resourceAzureExpressRouteConnUpdate,
		DeleteContext: resourceAzureExpressRouteConnDelete,
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
			"azure_service_key": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The Service Key provided by Microsoft Azure when you created your ExpressRoute circuit.",
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
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
			"is_public": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether PacketFabric should allocate a public IP address for this connection. Set this to true if you intend to set up peering with Microsoft public services (such as Microsoft 365). ",
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
			"azure_connection_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Azure connection type.\n\t\tExample: primary or seconday",
			},
			"vlan_id_private": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The private peering vlan.",
			},
			"vlan_id_microsoft": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The microsoft peering vlan.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceAzureExpressRouteConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	expressRoute := extractAzureExpressRouteConn(d)
	if cid, ok := d.GetOk("circuit_id"); ok {
		resp, err := c.CreateAzureExpressRouteConn(expressRoute, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		if err := checkCloudRouterConnectionStatus(c, cid.(string), resp.CloudCircuitID); err != nil {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.CloudCircuitID)

			time.Sleep(90 * time.Second) // wait for the connection to show on AWS
			resp, err := c.ReadCloudRouterConnection(cid.(string), resp.CloudCircuitID)
			if err != nil {
				diags = diag.FromErr(err)
				return diags
			}

			if resp.CloudSettings.AzureConnectionType == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Incomplete Cloud Information",
					Detail:   "The azure_connection_type is currently unavailable.",
				})
				return diags
			} else {
				_ = d.Set("azure_connection_type", resp.CloudSettings.AzureConnectionType)
			}
			if resp.CloudSettings.VlanPrivate == 0 {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Incomplete Cloud Information",
					Detail:   "The vlan_id_private is currently unavailable.",
				})
				return diags
			} else {
				_ = d.Set("vlan_id_private", resp.CloudSettings.VlanIDPf)
			}
			if resp.CloudSettings.VlanMicrosoft == 0 {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Incomplete Cloud Information",
					Detail:   "The vlan_id_microsoft are currently unavailable.",
				})
				return diags
			} else {
				_ = d.Set("vlan_id_microsoft", resp.CloudSettings.VlanIDPf)
			}

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

func resourceAzureExpressRouteConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	circuitID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}
	cloudConnCID := d.Get("id")
	resp, err := c.ReadCloudRouterConnection(circuitID.(string), cloudConnCID.(string))
	if err != nil {
		diags = diag.FromErr(err)
		return diags
	}

	_ = d.Set("account_uuid", resp.AccountUUID)
	_ = d.Set("circuit_id", resp.CloudRouterCircuitID)
	_ = d.Set("description", resp.Description)
	_ = d.Set("speed", resp.Speed)
	_ = d.Set("azure_service_key", resp.CloudSettings.AzureServiceKey)
	if _, ok := d.GetOk("po_number"); ok {
		_ = d.Set("po_number", resp.PONumber)
	}

	if resp.CloudSettings.PublicIP != "" {
		_ = d.Set("is_public", true)
	} else {
		_ = d.Set("is_public", false)
	}

	if resp.CloudSettings.AzureConnectionType == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Incomplete Cloud Information",
			Detail:   "The azure_connection_type is currently unavailable.",
		})
		return diags
	} else {
		_ = d.Set("azure_connection_type", resp.CloudSettings.AzureConnectionType)
	}
	if resp.CloudSettings.VlanPrivate == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Incomplete Cloud Information",
			Detail:   "The vlan_id_private is currently unavailable.",
		})
		return diags
	} else {
		_ = d.Set("vlan_id_private", resp.CloudSettings.VlanIDPf)
	}
	if resp.CloudSettings.VlanMicrosoft == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Incomplete Cloud Information",
			Detail:   "The vlan_id_microsoft are currently unavailable.",
		})
		return diags
	} else {
		_ = d.Set("vlan_id_microsoft", resp.CloudSettings.VlanIDPf)
	}
	// unsetFields: published_quote_line_uuid

	if _, ok := d.GetOk("labels"); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set("labels", labels)
	}
	return diags
}

func resourceAzureExpressRouteConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceAzureExpressRouteConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractAzureExpressRouteConn(d *schema.ResourceData) packetfabric.AzureExpressRouteConn {
	expressRoute := packetfabric.AzureExpressRouteConn{}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		expressRoute.MaybeNat = maybeNat.(bool)
	}
	if azureServiceKey, ok := d.GetOk("azure_service_key"); ok {
		expressRoute.AzureServiceKey = azureServiceKey.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		expressRoute.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		expressRoute.Description = description.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		expressRoute.Speed = speed.(string)
	}
	if isPublic, ok := d.GetOk("is_public"); ok {
		expressRoute.IsPublic = isPublic.(bool)
	}
	if publishedQuoteLine, ok := d.GetOk("published_quote_line_uuid"); ok {
		expressRoute.PublishedQuoteLineUUID = publishedQuoteLine.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		expressRoute.PONumber = poNumber.(string)
	}
	return expressRoute
}
