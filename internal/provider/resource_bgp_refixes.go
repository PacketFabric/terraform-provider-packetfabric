package provider

import (
	"context"
	"fmt"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBgpPrefixes() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBgpPrefixesCreate,
		ReadContext:   resourceBgpPrefixesRead,
		UpdateContext: resourceBgpPrefixesUpdate,
		DeleteContext: resourceBgpPrefixesDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Unique ID of Prefixes after create.",
			},
			"bgp_settings_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "UUID of BGP settings to fetch prefixes for.",
			},
			"prefixes": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The list of prefixes to be created.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp_prefix_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UUID of the instance.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The actual IP Prefix of this instance.",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Whether this prefix is in or out.",
						},
						"order": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The Order of this Prefix agains the others.",
						},
					},
				},
			},
		},
	}
}

func resourceBgpPrefixesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var bgpSessionUUID string
	var bgpPrefixes []packetfabric.BgpPrefix
	if sessionUUID, ok := d.GetOk("bgp_settings_uuid"); !ok {
		return diag.Errorf("please provide a valid BGP Session UUID")
	} else {
		bgpSessionUUID = sessionUUID.(string)
	}
	if prefixes := extractBgpSessionPrefixes(d); len(prefixes) <= 0 {
		return diag.Errorf("please provide a valid list of prefixes")
	} else {
		bgpPrefixes = prefixes
	}
	var diags diag.Diagnostics
	resp, err := c.CreateBgpSessionPrefixes(bgpPrefixes, bgpSessionUUID)
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func resourceBgpPrefixesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var bgpSettingsUUID string
	if settingsUUID, ok := d.GetOk("bgp_settings_uuid"); !ok {
		return diag.Errorf("please provide a valid BGP Settings UUID")
	} else {
		bgpSettingsUUID = settingsUUID.(string)
	}
	var diags diag.Diagnostics
	if prefixes, err := c.ReadBgpSessionPrefixes(bgpSettingsUUID); len(prefixes) <= 0 && err != nil {
		return diag.FromErr(err)
	} else {
		_ = d.Set("prefixes", _flattenPrefixes(prefixes))
	}
	return diags
}

func resourceBgpPrefixesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Prefixes update not available at this moment.",
		Detail:   "Prefixes can only be updated through through `cloud_router_bgp_session` resource.",
	})
	return diags
}

func resourceBgpPrefixesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	diags := make(diag.Diagnostics, 0)
	var bgpSettingsUUID string
	if settingsUUID, ok := d.GetOk("bgp_settings_uuid"); !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "There's no `bgp_settings_uuid`",
			Detail:   "There's no `bgp_settings_uuid` created yet, please make sure the terraform plan was applyed.",
		})
		return diags
	} else {
		bgpSettingsUUID = settingsUUID.(string)
	}
	currentPrefixes, err := c.ReadBgpSessionPrefixes(bgpSettingsUUID)
	if err != nil || len(currentPrefixes) <= 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "There are not BGP Prefixes",
			Detail:   fmt.Sprintf("There are no BGP prefixes created for %s. Please make sure the terraform plan was applyed.", bgpSettingsUUID),
		})
		return diags
	}
	prefixUUIDs := make([]string, 0)
	for _, prefix := range currentPrefixes {
		prefixUUIDs = append(prefixUUIDs, prefix.BgpPrefixUUID)
	}
	if _, err := c.DeleteBgpPrefixes(prefixUUIDs, bgpSettingsUUID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Prefixes cannot be deleted",
			Detail:   err.Error(),
		})
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Prefixes deleted with success.",
		})
	}
	return diags
}

func extractBgpSessionPrefixes(d *schema.ResourceData) []packetfabric.BgpPrefix {
	if prefixes, ok := d.GetOk("prefixes"); ok {
		sessionPrefixes := make([]packetfabric.BgpPrefix, 0)
		for _, pref := range prefixes.(*schema.Set).List() {
			sessionPrefixes = append(sessionPrefixes, packetfabric.BgpPrefix{
				Prefix: pref.(map[string]interface{})["prefix"].(string),
				Type:   pref.(map[string]interface{})["type"].(string),
				Order:  pref.(map[string]interface{})["order"].(int),
			})
		}
		return sessionPrefixes
	}
	return make([]packetfabric.BgpPrefix, 0)
}

func _flattenPrefixes(prefixes []packetfabric.BgpSessionResponse) []interface{} {
	flattens := make([]interface{}, len(prefixes), len(prefixes))
	for i, prefix := range prefixes {
		flatten := make(map[string]interface{})
		flatten["bgp_prefix_uuid"] = prefix.BgpPrefixUUID
		flatten["prefix"] = prefix.Prefix
		flatten["type"] = prefix.Type
		flatten["order"] = prefix.Order
		flattens[i] = prefix
	}
	return flattens
}
