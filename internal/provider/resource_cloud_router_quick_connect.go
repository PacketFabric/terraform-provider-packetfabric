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
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Cloud Couter Circuit ID.",
			},
			"connection_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The cloud router connection circuit ID.",
			},
			"import_circuit_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The cloud router quick connection import CID.",
			},
			"service_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The service UUID associated with the cloud router quick connect.",
			},
			"import_filters": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The import filters prefix.",
						},
						"match_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger", "longer"}, true),
							Description:  "The match type of this prefix.",
						},
						"local_preference": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The import filters local preference.",
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
							Description:  "The return filters prefix.",
						},
						"match_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger", "longer"}, true),
							Description:  "The match type of this prefix.",
						},
						"as_prepend": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The return filters as prepend.",
						},
						"med": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The return filters med.",
						},
					},
				},
			},
			"circuit_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cloud router quick connect CID.",
			},
			"route_set_circuit_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cloud router quick connect router set CID.",
			},
			"is_defunct": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whetheaver cloud router quick connect is defunct.",
			},
			"time_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cloud router quick connect time created.",
			},
			"time_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cloud router quick connect time updated.",
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
	_ = d.Set("circuit_id", resp.CircuitID)
	_ = d.Set("route_set_circuit_id", resp.RouteSetCircuitID)
	_ = d.Set("is_defunct", resp.IsDefunct)
	_ = d.Set("time_created", resp.TimeCreated)
	_ = d.Set("time_updated", resp.TimeUpdated)
	d.SetId(resp.CircuitID)
	return diags
}

func resourceCloudRouterQuickConnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceCloudRouterQuickConnectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if importCID, ok := d.GetOk("import_circuit_id"); ok {
		c := m.(*packetfabric.PFClient)
		c.Ctx = ctx
		crCID := d.Get("cr_circuit_id").(string)
		connCID := d.Get("connection_circuit_id").(string)
		quickConn := extractCloudRouterQuickConnectUpdate(d)
		if err := c.UpdateCloudRouterQuickConnect(crCID, connCID, importCID.(string), quickConn); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudRouterQuickConnectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	var crCIDStr, connCIDStr, importCIDStr string
	if crCID, ok := d.GetOk("cr_circuit_id"); ok {
		crCIDStr = crCID.(string)
	}
	if connCID, ok := d.GetOk("connection_circuit_id"); ok {
		connCIDStr = connCID.(string)
	}
	if importCID, ok := d.GetOk("import_circuit_id"); ok {
		importCIDStr = importCID.(string)
	}
	warningMsg, err := c.DeleteCloudRouterQuickConnect(d.Id(), crCIDStr, connCIDStr, importCIDStr)
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
