package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBgpSession() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBgpSessionRead,
		Schema: map[string]*schema.Schema{
			PfCircuitId:    schemaStringRequired(PfCircuitIdDescription),
			PfConnectionId: schemaStringRequired(PfConnectionIdDescription),
			PfBgpSessions: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfBgpSettingsUuid: schemaStringComputed(PfBgpSettingsUuidDescription2),
						PfAddressFamily:   schemaStringComputed(PfAddressFamilyDescription),
						PfRemoteAddress:   schemaStringComputed(PfRemoteAddressDescription),
						PfRemoteAsn:       schemaIntComputed(PfRemoteAsnDescription2),
						PfMultihopTtl:     schemaIntComputed(PfMultihopTtlDescription),
						PfLocalPreference: schemaIntComputed(PfLocalPreferenceDescription2),
						PfAsPrepend:       schemaIntComputed(PfAsPrependDescription3),
						PfMed:             schemaIntComputed(PfMedDescription3),
						PfL3Address:       schemaStringComputed(PfL3AddressDescription),
						PfOrlonger:        schemaBoolComputed(PfOrlongerDescription2),
						PfBfdInterval:     schemaIntComputed(PfBfdIntervalDescription2),
						PfBfdMultiplier:   schemaIntComputed(PfBfdMultiplierDescription2),
						PfDisabled:        schemaBoolComputed(PfDisabledDescription2),
						PfNat: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfPreNatSources: schemaStringListComputedNotEmpty(PfPreNatSourcesDescription2),
									PfPoolPrefixes:  schemaStringListComputedNotEmpty(PfPoolPrefixesDescription2),
									PfDirection:     schemaStringComputed(PfDirectionDescription2),
									PfNatType:       schemaStringComputed(PfNatTypeDescription2),
									PfDnatMappings: {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												PfPrivatePrefix:     schemaStringComputed(PfPrivatePrefixDescription2),
												PfPublicPrefix:      schemaStringComputed(PfPublicPrefixDescription2),
												PfConditionalPrefix: schemaStringComputed(PfConditionalPrefixDescription2),
											},
										},
									},
								},
							},
						},
						PfBgpState:    schemaStringComputed(PfBgpStateDescription),
						PfTimeCreated: schemaStringComputed(PfTimeCreatedDescription4),
						PfTimeUpdated: schemaStringComputed(PfTimeUpdatedDescription2),
						PfPrefixes: {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: PfPrefixesDescription,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfBgpPrefixUuid:   schemaStringComputed(PfBgpPrefixUuidDescription3),
									PfPrefix:          schemaStringComputed(PfPrefixDescription),
									PfMatchType:       schemaStringComputed(PfMatchTypeDescription2),
									PfAsPrepend:       schemaIntComputed(PfAsPrependDescription4),
									PfMed:             schemaIntComputed(PfMedDescription6),
									PfLocalPreference: schemaIntComputed(PfLocalPreferenceDescription3),
									PfType:            schemaStringComputed(PfTypeDescription),
									PfOrder: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: PfOrderDescription,
										Deprecated:  PfDeprecatedField,
									},
								},
							},
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
	cID, ok := d.GetOk(PfCircuitId)
	if !ok {
		return diag.Errorf(MessageMissingCircuitIdDetail)
	}
	connCID, ok := d.GetOk(PfConnectionId)
	if !ok {
		return diag.Errorf(MesssageCRCIdRequired)
	}
	sessions, err := c.ListBgpSessions(cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfBgpSessions, flattenBgpSessions(&sessions))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(connCID.(string) + "-data")
	return diags
}

func flattenBgpSessions(sessions *[]packetfabric.BgpSessionAssociatedResp) []interface{} {
	fields := stringsToMap(PfBgpSettingsUuid, PfAddressFamily, PfRemoteAddress, PfRemoteAsn, PfMultihopTtl, PfLocalPreference, PfAsPrepend, PfL3Address, PfMed, PfOrlonger, PfBfdInterval, PfBfdMultiplier, PfDisabled, PfBgpState, PfTimeCreated, PfTimeUpdated)

	if sessions != nil {
		flattens := make([]interface{}, len(*sessions), len(*sessions))

		for i, session := range *sessions {
			flatten := structToMap(session, fields)
			flatten[PfPrefixes] = flattenBgpSessionsPrefixes(&session.Prefixes)
			flatten[PfNat] = flattenBgpSessionsNat(session.Nat)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenBgpSessionsPrefixes(prefixes *[]packetfabric.BgpPrefix) []interface{} {
	fields := stringsToMap(PfBgpPrefixUuid, PfPrefix, PfMatchType, PfAsPrepend, PfMed, PfLocalPreference, PfType)
	if prefixes != nil {
		flattens := make([]interface{}, len(*prefixes), len(*prefixes))
		for i, prefix := range *prefixes {
			flattens[i] = structToMap(prefix, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenBgpSessionsNat(nat *packetfabric.BgpNat) []interface{} {
	fields := stringsToMap(PfPreNatSources, PfPoolPrefixes, PfDirection, PfNatType)
	flattens := make([]interface{}, 0)
	if nat != nil {
		flatten := structToMap(nat, fields)
		flatten[PfDnatMappings] = flattenBgpSessionsDnat(&nat.DnatMappings)
		flattens = append(flattens, flatten)
	}
	return flattens
}

func flattenBgpSessionsDnat(dnats *[]packetfabric.BgpDnatMapping) []interface{} {
	fields := stringsToMap(PfPrivatePrefix, PfPublicPrefix, PfConditionalPrefix)
	if dnats != nil {
		flattens := make([]interface{}, len(*dnats), len(*dnats))
		for i, dnat := range *dnats {
			flattens[i] = structToMap(dnat, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}
