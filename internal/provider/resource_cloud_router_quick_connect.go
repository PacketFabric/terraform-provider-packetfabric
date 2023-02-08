package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudRouterQuickConnect() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		CreateContext: resourceCloudRouterQuickConnectCreate,
		ReadContext:   resourceCloudRouterQuickConnectRead,
		UpdateContext: resourceCloudRouterQuickConnectUpdate,
		DeleteContext: resourceCloudRouterQuickConnectDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cr_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of your Cloud Router.",
			},
			"connection_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the Cloud Router connection that will be importing the routes.",
			},
			"service_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The service UUID associated with the service provider's Quick Connect.",
			},
			"import_filters": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The route prefix that you will be importing from the Quick Connect. This is set by the service provider.",
						},
						"match_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger", "longer"}, true),
							Description:  "The match type for the imported prefix. This is set by the service provider.",
						},
						"local_preference": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The local preference to apply to the prefix.",
						},
					},
				},
			},
			"return_filters": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The prefix to export to the service provider that they will use for return traffic.",
						},
						"match_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger", "longer"}, true),
							Description:  "The match type of this prefix.\n\n\tEnum: `\"exact\"` `\"orlonger\"` `\"longer\"`",
						},
						"as_prepend": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The AS prepend to apply to the exported/returned prefix.",
						},
						"med": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The MED to apply to the exported/returned prefix.",
						},
					},
				},
			},
			"import_circuit_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Circuit ID of this Cloud Router Import.",
			},
			"route_set_circuit_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Circuit ID of the Route Set selected for this Cloud Router Import.",
			},
			"is_defunct": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the Quick Connect is defunct. This typically happens when the provider removes the service.",
			},
			"time_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Quick Connect time created.",
			},
			"time_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Quick Connect time updated.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudRouterQuickConnectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	crCID := d.Get("cr_circuit_id").(string)
	connCID := d.Get("connection_circuit_id").(string)
	var diags diag.Diagnostics
	quickConnect := extractCloudRouterQuickConnect(d)
	resp, err := c.CreateCloudRouterQuickConnect(crCID, connCID, quickConnect)
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}
	_ = d.Set("import_circuit_id", resp.ImportCircuitID)
	_ = d.Set("route_set_circuit_id", resp.RouteSetCircuitID)
	_ = d.Set("is_defunct", resp.IsDefunct)
	_ = d.Set("time_created", resp.TimeCreated)
	_ = d.Set("time_updated", resp.TimeUpdated)
	d.SetId(resp.ImportCircuitID)
	return diags
}

func resourceCloudRouterQuickConnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceCloudRouterQuickConnectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	crCID := d.Get("cr_circuit_id").(string)
	connCID := d.Get("connection_circuit_id").(string)
	quickConn := extractCloudRouterQuickConnectUpdate(d)
	if err := c.UpdateCloudRouterQuickConnect(crCID, connCID, d.Id(), quickConn); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceCloudRouterQuickConnectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	var crCIDStr, connCIDStr string
	if crCID, ok := d.GetOk("cr_circuit_id"); ok {
		crCIDStr = crCID.(string)
	}
	if connCID, ok := d.GetOk("connection_circuit_id"); ok {
		connCIDStr = connCID.(string)
	}
	warningMsg, err := c.DeleteCloudRouterQuickConnect(d.Id(), crCIDStr, connCIDStr, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if warningMsg != "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Cloud Router Quick Connect Delete",
			Detail:   warningMsg,
		})
	}
	return diags
}

func extractCloudRouterQuickConnect(d *schema.ResourceData) packetfabric.CloudRouterQuickConnect {
	quickConnect := packetfabric.CloudRouterQuickConnect{}
	if serviceUUID, ok := d.GetOk("service_uuid"); ok {
		quickConnect.ServiceUUID = serviceUUID.(string)
	}
	quickConnect.ImportFilters = extractImportFilters(d)
	quickConnect.ReturnFilters = extractReturnFilters(d)
	return quickConnect
}

func extractCloudRouterQuickConnectUpdate(d *schema.ResourceData) packetfabric.CloudRouterQuickConnectUpdate {
	quickConnect := packetfabric.CloudRouterQuickConnectUpdate{}
	quickConnect.ImportFilters = extractImportFilters(d)
	quickConnect.ReturnFilters = extractReturnFilters(d)
	return quickConnect
}

func extractImportFilters(d *schema.ResourceData) []packetfabric.QuickConnectImportFilters {
	if importFilters, ok := d.GetOk("import_filters"); ok {
		extractedFilters := make([]packetfabric.QuickConnectImportFilters, 0)
		for _, filter := range importFilters.(*schema.Set).List() {
			extractedFilters = append(extractedFilters, packetfabric.QuickConnectImportFilters{
				Prefix:    filter.(map[string]interface{})["prefix"].(string),
				MatchType: filter.(map[string]interface{})["match_type"].(string),
				Localpref: filter.(map[string]interface{})["local_preference"].(int),
			})
		}
		return extractedFilters
	}
	return make([]packetfabric.QuickConnectImportFilters, 0)
}

func extractReturnFilters(d *schema.ResourceData) []packetfabric.QuickConnectReturnFilters {
	if returnFilters, ok := d.GetOk("return_filters"); ok {
		extractedFilters := make([]packetfabric.QuickConnectReturnFilters, 0)
		for _, filter := range returnFilters.(*schema.Set).List() {
			extractedFilters = append(extractedFilters, packetfabric.QuickConnectReturnFilters{
				Prefix:    filter.(map[string]interface{})["prefix"].(string),
				MatchType: filter.(map[string]interface{})["match_type"].(string),
				Asprepend: filter.(map[string]interface{})["as_prepend"].(int),
				Med:       filter.(map[string]interface{})["med"].(int),
			})
		}
		return extractedFilters
	}
	return make([]packetfabric.QuickConnectReturnFilters, 0)
}
