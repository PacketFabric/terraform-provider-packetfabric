package provider

import (
	"context"
	"fmt"
	//"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudRouterQuickConnect() *schema.Resource {
	return &schema.Resource{
		Timeouts:      schema10MinuteTimeouts(),
		CreateContext: resourceCloudRouterQuickConnectCreate,
		ReadContext:   resourceCloudRouterQuickConnectRead,
		UpdateContext: resourceCloudRouterQuickConnectUpdate,
		DeleteContext: resourceCloudRouterQuickConnectDelete,
		Schema: map[string]*schema.Schema{
			PfId:                  schemaStringComputed(PfImportCircuitIdDescription),
			PfRouteSetCircuitId:   schemaStringComputed(PfRouteSetCircuitIdDescription),
			PfCrCircuitId:         schemaStringRequiredNewNotEmpty(PfCrCircuitIdDescription),
			PfConnectionCircuitId: schemaStringRequiredNewNotEmpty(PfConnectionCircuitIdDescription),
			PfServiceUuid:         schemaStringRequiredNewValidate(PfServiceUuidDescription3, validation.IsUUID),
			PfImportFilters: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: PfImportFiltersDescription2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPrefix:          schemaStringOptionalComputedNotEmpty(PfPrefixDescription5),
						PfMatchType:       schemaStringOptionalComputedValidate(PfMatchTypeDescription5, validateMatchType()),
						PfLocalPreference: schemaIntOptionalDefault(PfLocalPreferenceDescription6, 0),
					},
				},
			},
			PfReturnFilters: {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPrefix:    schemaStringRequiredNotEmpty(PfPrefixDescription6),
						PfMatchType: schemaStringOptionalValidateDefault(PfMatchTypeDescription, validateMatchType(), PfMatchTypeDefault),
						PfAsPrepend: schemaIntOptionalValidateDefault(PfAsPrependDescription6, validateAsPrepend(), 0),
						PfMed:       schemaIntOptionalDefault(PfMedDescription8, 0),
					},
				},
			},
			PfIsDefunct:        schemaBoolComputed(PfIsDefunctDescription),
			PfState:            schemaStringComputed(PfStateDescriptionD),
			PfConnectionSpeed:  schemaStringComputed(PfConnectionSpeedDescription),
			PfSubscriptionTerm: schemaIntOptionalValidateDefault(PfSubscriptionTermDescription2, validateSubscriptionTerm(), 1),
		},
		Importer: &schema.ResourceImporter{
			StateContext: QuickConnectImportStatePassthroughContext,
		},
	}
}

func resourceCloudRouterQuickConnectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	crCID := d.Get(PfCrCircuitId).(string)
	connCID := d.Get(PfConnectionCircuitId).(string)
	var diags diag.Diagnostics
	var importFilters []packetfabric.QuickConnectImportFilters
	// if import_filters not set by the user, use the marketplace service route set
	if _, ok := d.GetOk(PfImportFilters); !ok {
		if serviceUUID, ok := d.GetOk(PfServiceUuid); ok {
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
	_ = setResourceDataKeys(d, resp, PfRouteSetCircuitId, PfIsDefunct, PfState, PfConnectionSpeed, PfSubscriptionTerm)
	if importFilters != nil {
		importFilters := flattenImportFiltersConfiguration(importFilters)
		_ = d.Set(PfImportFilters, importFilters)
	}

	d.SetId(resp.ImportCircuitID)
	return diags
}

func resourceCloudRouterQuickConnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)

	var crCIDStr, connCIDStr string
	if crCID, ok := d.GetOk(PfCrCircuitId); ok {
		crCIDStr = crCID.(string)
	}
	if connCID, ok := d.GetOk(PfConnectionCircuitId); ok {
		connCIDStr = connCID.(string)
	}
	resp, err := c.GetCloudRouterQuickConnect(crCIDStr, connCIDStr, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = setResourceDataKeys(d, resp, PfRouteSetCircuitId, PfIsDefunct, PfState, PfConnectionSpeed, PfServiceUuid, PfSubscriptionTerm)

		returnFilters := flattenReturnFiltersConfiguration(resp.ReturnFilters)
		if err := d.Set(PfReturnFilters, returnFilters); err != nil {
			return diag.Errorf("error setting 'return_filters': %s", err)
		}
		if resp.ImportFilters != nil {
			importFilters := flattenImportFiltersConfiguration(resp.ImportFilters)
			_ = d.Set(PfImportFilters, importFilters)
		}
	}
	return diag.Diagnostics{}
}

func resourceCloudRouterQuickConnectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	crCID := d.Get(PfCrCircuitId).(string)
	connCID := d.Get(PfConnectionCircuitId).(string)

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
				Summary:  MessageIgnoredImportFilter,
				Detail:   fmt.Sprintf("The import filter with prefix %s does not exist in the provider's configuration and will be ignored.", localFilter.Prefix),
			})
		}
	}

	// If return filters are updated by the user
	if d.HasChange(PfReturnFilters) {
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
	if crCID, ok := d.GetOk(PfCrCircuitId); ok {
		crCIDStr = crCID.(string)
	}
	if connCID, ok := d.GetOk(PfConnectionCircuitId); ok {
		connCIDStr = connCID.(string)
	}
	warningMsg, err := c.DeleteCloudRouterQuickConnect(crCIDStr, connCIDStr, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if warningMsg != PfEmptyString {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  MessageCRQCDelete,
			Detail:   warningMsg,
		})
	}
	return diags
}

func extractCloudRouterQuickConnect(d *schema.ResourceData, importFilters []packetfabric.QuickConnectImportFilters) packetfabric.CloudRouterQuickConnect {
	quickConnect := packetfabric.CloudRouterQuickConnect{}
	if serviceUUID, ok := d.GetOk(PfServiceUuid); ok {
		quickConnect.ServiceUUID = serviceUUID.(string)
	}
	if subscriptionTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
		quickConnect.SubscriptionTerm = subscriptionTerm.(int)
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
	if subscriptionTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
		quickConnect.SubscriptionTerm = subscriptionTerm.(int)
	}
	return quickConnect
}

func extractImportFilters(d *schema.ResourceData) []packetfabric.QuickConnectImportFilters {
	if importFilters, ok := d.GetOk(PfImportFilters); ok {
		extractedFilters := make([]packetfabric.QuickConnectImportFilters, 0)
		for _, filter := range importFilters.(*schema.Set).List() {
			extractedFilters = append(extractedFilters, packetfabric.QuickConnectImportFilters{
				Prefix:          filter.(map[string]interface{})[PfPrefix].(string),
				MatchType:       filter.(map[string]interface{})[PfMatchType].(string),
				LocalPreference: filter.(map[string]interface{})[PfLocalPreference].(int),
			})
		}
		return extractedFilters
	}
	return make([]packetfabric.QuickConnectImportFilters, 0)
}

func extractReturnFilters(d *schema.ResourceData) []packetfabric.QuickConnectReturnFilters {
	if returnFilters, ok := d.GetOk(PfReturnFilters); ok {
		extractedFilters := make([]packetfabric.QuickConnectReturnFilters, 0)
		for _, filter := range returnFilters.(*schema.Set).List() {
			extractedFilters = append(extractedFilters, packetfabric.QuickConnectReturnFilters{
				Prefix:    filter.(map[string]interface{})[PfPrefix].(string),
				MatchType: filter.(map[string]interface{})[PfMatchType].(string),
				AsPrepend: filter.(map[string]interface{})[PfAsPrepend].(int),
				Med:       filter.(map[string]interface{})[PfMed].(int),
			})
		}
		return extractedFilters
	}
	return make([]packetfabric.QuickConnectReturnFilters, 0)
}

func flattenReturnFiltersConfiguration(prefixes []packetfabric.QuickConnectReturnFilters) []map[string]interface{} {
	fields := stringsToMap(PfPrefix, PfMatchType, PfAsPrepend, PfMed)
	result := make([]map[string]interface{}, len(prefixes))
	for i, prefix := range prefixes {
		result[i] = structToMap(&prefix, fields)
	}
	return result
}

func flattenImportFiltersConfiguration(prefixes []packetfabric.QuickConnectImportFilters) []map[string]interface{} {
	fields := stringsToMap(PfPrefix, PfMatchType, PfLocalPreference)
	result := make([]map[string]interface{}, len(prefixes))
	for i, prefix := range prefixes {
		result[i] = structToMap(&prefix, fields)
	}
	return result
}

func mapToQuickConnectImportFilters(data []map[string]interface{}) []packetfabric.QuickConnectImportFilters {
	result := make([]packetfabric.QuickConnectImportFilters, len(data))
	for i, item := range data {
		result[i] = packetfabric.QuickConnectImportFilters{
			Prefix:          item[PfPrefix].(string),
			MatchType:       item[PfMatchType].(string),
			LocalPreference: item[PfLocalPreference].(int),
		}
	}
	return result
}
