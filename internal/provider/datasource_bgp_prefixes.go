package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBgpPrefix() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBgpPrefixRead,
		Schema: map[string]*schema.Schema{
			"bgp_settings_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "UUID of BGP settings for prefixes.",
			},
			"bgp_prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp_prefix_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The UUID of the instance.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The actual IP Prefix of this instance.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Whether this prefix is in or out.",
						},
						"order": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The Order of this Prefix agains the others.",
						},
					},
				},
			},
		},
	}
}

func dataSourceBgpPrefixRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	var bgpSettingsUUID string
	settingsUUID, ok := d.GetOk("bgp_settings_uuid")
	if !ok {
		return diag.Errorf("please provide a valid BGP Settings UUID")
	}
	bgpSettingsUUID = settingsUUID.(string)
	prefixes, err := c.ReadBgpSessionPrefixes(bgpSettingsUUID)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("bgp_prefixes", flattenBgpPrefixes(&prefixes)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenBgpPrefixes(prefixes *[]packetfabric.BgpPrefix) []interface{} {
	if prefixes != nil {
		flattens := make([]interface{}, len(*prefixes), len(*prefixes))
		for i, prefix := range *prefixes {
			flatten := make(map[string]interface{})
			flatten["bgp_prefix_uuid"] = prefix.BgpPrefixUUID
			flatten["prefix"] = prefix.Prefix
			flatten["match_type"] = prefix.MatchType
			flatten["as_prepend"] = prefix.AsPrepend
			flatten["med"] = prefix.Med
			flatten["local_preference"] = prefix.LocalPreference
			flatten["type"] = prefix.Type
			flatten["order"] = prefix.Order
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
