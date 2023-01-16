package provider

import (
	"context"

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
				Description: "The unique ID of prefixes after they are created.",
			},
			"bgp_settings_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "UUID of the BGP session for which you are fetching or setting prefixes.",
			},
			"prefixes": {
				Type:        schema.TypeSet,
				Required:    true,
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
							Description: "The actual IP prefix of this instance.",
						},
						"match_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The prefix match type",
						},
						"as_prepend": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The BGP prepend value of the bgp prefix",
						},
						"med": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The med of the bgp prefix",
						},
						"local_preference": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The local_preference of the bgp prefix",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Indicates whether the prefix is in or out",
						},
						"order": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The order of the bgp prefix against the others",
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
	sessionUUID, ok := d.GetOk("bgp_settings_uuid")
	if !ok {
		return diag.Errorf("please provide a valid BGP Session UUID")
	}
	bgpSessionUUID = sessionUUID.(string)
	prefixes := extractBgpSessionPrefixes(d)
	if len(prefixes) <= 0 {
		return diag.Errorf("please provide a valid list of prefixes")
	}
	bgpPrefixes = prefixes
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
	settingsUUID, ok := d.GetOk("bgp_settings_uuid")
	if !ok {
		return diag.Errorf("please provide a valid BGP Settings UUID")
	}
	bgpSettingsUUID = settingsUUID.(string)
	var diags diag.Diagnostics
	prefixes, err := c.ReadBgpSessionPrefixes(bgpSettingsUUID)
	if len(prefixes) <= 0 && err != nil {
		return diag.FromErr(err)
	}
	_ = d.Set("prefixes", _flattenPrefixes(prefixes))

	return diags
}

func resourceBgpPrefixesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "BGP Prefixes update are not supported via `packetfabric_cloud_router_bgp_prefixes` resource.",
		Detail:   "BGP Prefixes can only be updated through `packetfabric_cloud_router_bgp_session` resource.",
	})
	return diags
}

func resourceBgpPrefixesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	diags := make(diag.Diagnostics, 0)
	d.SetId("")
	// Skipping the BGP Prefix delete behavior due to the following issue: #140
	return diags
}

func extractBgpSessionPrefixes(d *schema.ResourceData) []packetfabric.BgpPrefix {
	if prefixes, ok := d.GetOk("prefixes"); ok {
		sessionPrefixes := make([]packetfabric.BgpPrefix, 0)
		for _, pref := range prefixes.(*schema.Set).List() {
			sessionPrefixes = append(sessionPrefixes, packetfabric.BgpPrefix{
				Prefix:          pref.(map[string]interface{})["prefix"].(string),
				MatchType:       pref.(map[string]interface{})["match_type"].(string),
				AsPrepend:       pref.(map[string]interface{})["as_prepend"].(int),
				Med:             pref.(map[string]interface{})["med"].(int),
				LocalPreference: pref.(map[string]interface{})["local_preference"].(int),
				Type:            pref.(map[string]interface{})["type"].(string),
				Order:           pref.(map[string]interface{})["order"].(int),
			})
		}
		return sessionPrefixes
	}
	return make([]packetfabric.BgpPrefix, 0)
}

func _flattenPrefixes(prefixes []packetfabric.BgpPrefix) []interface{} {
	flattens := make([]interface{}, len(prefixes), len(prefixes))
	for i, prefix := range prefixes {
		flatten := make(map[string]interface{})
		flatten["bgp_prefix_uuid"] = prefix.BgpPrefixUUID
		flatten["prefix"] = prefix.Prefix
		flatten["match_type"] = prefix.MatchType
		flatten["as_prepend"] = prefix.AsPrepend
		flatten["med"] = prefix.Med
		flatten["local_preference"] = prefix.LocalPreference
		flatten["type"] = prefix.Type
		flatten["order"] = prefix.Order
		flattens[i] = prefix
	}
	return flattens
}
