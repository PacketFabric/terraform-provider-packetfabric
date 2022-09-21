package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBgpSession() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBgpSessionRead,
		Schema: map[string]*schema.Schema{
			"bgp_sessions": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp_settings_uuid": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The UUID of the instance.\n\t\tExample: 3d78949f-1396-4163-b0ca-3eba3592abcd",
						},
						"address_family": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Whether this instance is IPv4 or IPv6.\n\t\tEnum: \"v4\" \"v6\"",
						},
						"remote_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The cloud-side address of the instance.",
						},
						"remote_asn": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The cloud-side ASN of the instance.",
						},
						"multihop_ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The TTL of this session.\n\t\tDefaults to 1.",
						},
						"local_preference": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The preference for this instance. Deprecated.",
						},
						"community": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The BGP community for this instance. Deprecated.",
						},
						"as_prepend": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The BGP prepend value for this instance. Deprecated.",
						},
						"med": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The Multi-Exit Discriminator of this instance. Deprecated.",
						},
						"orlonger": {
							Type:        schema.TypeBool,
							Computed:    true,
							Optional:    true,
							Description: "Whether to use exact match or longer for all prefixes.",
						},
						"bfd_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "Minimum interval, in microseconds, for transmitting BFD Control packets.\n\t\tAvailable range is 3 through 30000.",
						},
						"bfd_multiplier": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The number of BFD Control packets not received by a neighbor that causes the session to be declared down.\n\t\tAvailable range is 2 through 16.",
						},
						"disabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Optional:    true,
							Description: "Whether this BGP session is disabled.\n\t\tDefault \"false\"",
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Time the instance was created.",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Time the instance was last updated.",
						},
					},
				},
			},
		},
	}
}

func dataSourceBgpSessionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	sessions, err := c.ListBgpSessions()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("bgp_sessions", flattenBgpSessions(&sessions))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenBgpSessions(sessions *[]packetfabric.BgpSessionAssociatedResp) []interface{} {
	if sessions != nil {
		flattens := make([]interface{}, len(*sessions), len(*sessions))

		for i, session := range *sessions {
			flatten := make(map[string]interface{})
			flatten["bgp_settings_uuid"] = session.BgpSettingsUUID
			flatten["address_family"] = session.AddressFamily
			flatten["remote_address"] = session.RemoteAddress
			flatten["remote_asn"] = session.RemoteAsn
			flatten["multihop_ttl"] = session.MultihopTTL
			flatten["local_preference"] = session.LocalPreference
			flatten["community"] = session.Community
			flatten["as_prepend"] = session.AsPrepend
			flatten["med"] = session.Med
			flatten["orlonger"] = session.Orlonger
			flatten["bfd_interval"] = session.BfdInterval
			flatten["bfd_multiplier"] = session.BfdMultiplier
			flatten["disabled"] = session.Disabled
			flatten["time_created"] = session.TimeCreated
			flatten["time_updated"] = session.TimeUpdated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
