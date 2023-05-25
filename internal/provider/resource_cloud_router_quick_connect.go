package provider

import (
	"context"
	"fmt"
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
				Type:        schema.TypeString,
				Description: "The Circuit ID of this Cloud Router Import.",
				Computed:    true,
			},
			"route_set_circuit_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Circuit ID of the Route Set selected for this Cloud Router Import.",
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
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "This is set by the service provider.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The route prefix that you will be importing from the Quick Connect. This is set by the service provider.",
						},
						"match_type": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger"}, true),
							Description:  "The match type for the imported prefix. This is set by the service provider.\n\n\tEnum: `\"exact\"` `\"orlonger\"` ",
						},
						"local_preference": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The local preference to apply to the prefix.\n\n\tAvailable range is 1 through 4294967295. ",
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
							Optional:     true,
							Default:      "exact",
							ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger"}, true),
							Description:  "The match type of this prefix.\n\n\tEnum: `\"exact\"` `\"orlonger\"` ",
						},
						"as_prepend": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntBetween(1, 5),
							Description:  "The AS prepend to apply to the exported/returned prefix.\n\n\tAvailable range is 1 through 5. ",
						},
						"med": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The MED to apply to the exported/returned prefix.\n\n\tAvailable range is 1 through 4294967295. ",
						},
					},
				},
			},
			"is_defunct": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the Quick Connect is defunct. This typically happens when the provider removes the service.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Shows the state of this import.\n\n\tEnum: `\"pending\"` `\"rejected\"` `\"provisioning\"`  `\"active\"`  `\"deleting\"`  `\"inactive\"`",
			},
			"connection_speed": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The speed of the target cloud router connection.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: QuickConnectImportStatePassthroughContext,
		},
	}
}

func resourceCloudRouterQuickConnectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	crCID := d.Get("cr_circuit_id").(string)
	connCID := d.Get("connection_circuit_id").(string)
	var diags diag.Diagnostics
	var importFilters []packetfabric.QuickConnectImportFilters
	// if import_filters not set by the user, use the marketplace service route set
	if _, ok := d.GetOk("import_filters"); !ok {
		if serviceUUID, ok := d.GetOk("service_uuid"); ok {
			resp2, err2 := c.GetMarketPlaceServiceRouteSet(serviceUUID.(string))
			if err2 != nil {
				return diag.FromErr(err2)
			}
			importFilters = mapToQuickConnectImportFilters(flattenImportFiltersConfiguration(resp2.Prefixes))
		}
	}
	quickConnect := extractCloudRouterQuickConnect(d, importFilters)
	resp, err := c.CreateCloudRouterQuickConnect(crCID, connCID, quickConnect)
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}
	_ = d.Set("route_set_circuit_id", resp.RouteSetCircuitID)
	_ = d.Set("is_defunct", resp.IsDefunct)
	_ = d.Set("state", resp.State)
	_ = d.Set("connection_speed", resp.ConnectionSpeed)
	if importFilters != nil {
		importFilters := flattenImportFiltersConfiguration(importFilters)
		_ = d.Set("import_filters", importFilters)
	}

	d.SetId(resp.ImportCircuitID)
	return diags
}

func resourceCloudRouterQuickConnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)

	var crCIDStr, connCIDStr string
	if crCID, ok := d.GetOk("cr_circuit_id"); ok {
		crCIDStr = crCID.(string)
	}
	if connCID, ok := d.GetOk("connection_circuit_id"); ok {
		connCIDStr = connCID.(string)
	}
	resp, err := c.GetCloudRouterQuickConnect(crCIDStr, connCIDStr, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("route_set_circuit_id", resp.RouteSetCircuitID)
		_ = d.Set("is_defunct", resp.IsDefunct)
		_ = d.Set("state", resp.State)
		_ = d.Set("connection_speed", resp.ConnectionSpeed)
		_ = d.Set("service_uuid", resp.ServiceUUID)

		returnFilters := flattenReturnFiltersConfiguration(resp.ReturnFilters)
		if err := d.Set("return_filters", returnFilters); err != nil {
			return diag.Errorf("error setting 'return_filters': %s", err)
		}
		if resp.ImportFilters != nil {
			importFilters := flattenImportFiltersConfiguration(resp.ImportFilters)
			_ = d.Set("import_filters", importFilters)
		}
	}
	return diag.Diagnostics{}
}

func resourceCloudRouterQuickConnectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	crCID := d.Get("cr_circuit_id").(string)
	connCID := d.Get("connection_circuit_id").(string)

	// Get the current cloud router quick connect configuration from the provider
	currentQuickConn, err := c.GetCloudRouterQuickConnect(crCID, connCID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Extract the local configuration
	localQuickConn := extractCloudRouterQuickConnectUpdate(d)
	// Copy the local configuration for further modifications
	updateQuickConn := localQuickConn

	// Make a map from currentImportFilters for easy lookup
	currentFiltersMap := make(map[string]packetfabric.QuickConnectImportFilters)
	for _, filter := range currentQuickConn.ImportFilters {
		currentFiltersMap[filter.Prefix] = filter
	}

	// Iterate over localQuickConn.ImportFilters, and update localPreference for existing ones
	for i, localFilter := range updateQuickConn.ImportFilters {
		_, exists := currentFiltersMap[localFilter.Prefix]

		if exists {
			// If the filter exists, update its localPreference
			updateQuickConn.ImportFilters[i].LocalPreference = localFilter.LocalPreference
		} else {
			// If the filter doesn't exist, show a warning
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Ignored Import Filter",
				Detail:   fmt.Sprintf("The import filter with prefix %s does not exist in the provider's configuration and will be ignored.", localFilter.Prefix),
			})
		}
	}

	// If return filters are updated by the user
	if d.HasChange("return_filters") {
		// The return filters are changed, set the new return filters from user
		updateQuickConn.ReturnFilters = localQuickConn.ReturnFilters
	}

	// Update the cloud router quick connect with the new settings
	if err := c.UpdateCloudRouterQuickConnect(crCID, connCID, d.Id(), updateQuickConn); err != nil {
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
	warningMsg, err := c.DeleteCloudRouterQuickConnect(crCIDStr, connCIDStr, d.Id())
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

func extractCloudRouterQuickConnect(d *schema.ResourceData, importFilters []packetfabric.QuickConnectImportFilters) packetfabric.CloudRouterQuickConnect {
	quickConnect := packetfabric.CloudRouterQuickConnect{}
	if serviceUUID, ok := d.GetOk("service_uuid"); ok {
		quickConnect.ServiceUUID = serviceUUID.(string)
	}
	if len(importFilters) > 0 {
		quickConnect.ImportFilters = importFilters
	} else {
		quickConnect.ImportFilters = extractImportFilters(d)
	}
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
				Prefix:          filter.(map[string]interface{})["prefix"].(string),
				MatchType:       filter.(map[string]interface{})["match_type"].(string),
				LocalPreference: filter.(map[string]interface{})["local_preference"].(int),
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
				AsPrepend: filter.(map[string]interface{})["as_prepend"].(int),
				Med:       filter.(map[string]interface{})["med"].(int),
			})
		}
		return extractedFilters
	}
	return make([]packetfabric.QuickConnectReturnFilters, 0)
}

func flattenReturnFiltersConfiguration(prefixes []packetfabric.QuickConnectReturnFilters) []map[string]interface{} {
	result := make([]map[string]interface{}, len(prefixes))
	for i, prefix := range prefixes {
		data := make(map[string]interface{})
		data["prefix"] = prefix.Prefix
		data["match_type"] = prefix.MatchType
		data["as_prepend"] = prefix.AsPrepend
		data["med"] = prefix.Med
		result[i] = data
	}
	return result
}

func flattenImportFiltersConfiguration(prefixes []packetfabric.QuickConnectImportFilters) []map[string]interface{} {
	result := make([]map[string]interface{}, len(prefixes))
	for i, prefix := range prefixes {
		data := map[string]interface{}{
			"prefix":           prefix.Prefix,
			"match_type":       prefix.MatchType,
			"local_preference": prefix.LocalPreference,
		}
		result[i] = data
	}
	return result
}

func mapToQuickConnectImportFilters(data []map[string]interface{}) []packetfabric.QuickConnectImportFilters {
	result := make([]packetfabric.QuickConnectImportFilters, len(data))
	for i, item := range data {
		result[i] = packetfabric.QuickConnectImportFilters{
			Prefix:          item["prefix"].(string),
			MatchType:       item["match_type"].(string),
			LocalPreference: item["local_preference"].(int),
		}
	}
	return result
}
