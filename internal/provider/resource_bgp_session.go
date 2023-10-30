package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBgpSession() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBgpSessionCreate,
		ReadContext:   resourceBgpSessionRead,
		UpdateContext: resourceBgpSessionUpdate,
		DeleteContext: resourceBgpSessionDelete,
		Timeouts:      schema10MinuteTimeouts(),
		Schema: map[string]*schema.Schema{
			PfId:              schemaStringComputedPlain(),
			PfCircuitId:       schemaStringRequiredNew(PfCircuitIdDescription),
			PfConnectionId:    schemaStringRequiredNew(PfConnectionIdDescription),
			PfMd5:             schemaStringOptionalNotEmpty(PfMd5Description),
			PfL3Address:       schemaPrefixOptional(PfL3AddressDescription3),
			PfPrimarySubnet:   schemaPrefixOptional(PfPrimarySubnetDescription),
			PfSecondarySubnet: schemaPrefixOptional(PfSecondarySubnetDescription),
			PfAddressFamily:   schemaStringOptionalValidateDefault(PfAddressFamilyDescription3, validateIpVersion(), PfAddressFamily4v),
			PfRemoteAddress:   schemaPrefixOptional(PfRemoteAddressDescription3),
			PfRemoteAsn:       schemaIntRequired(PfRemoteAsnDescription3),
			PfMultihopTtl:     schemaIntOptionalValidateDefault(PfMultihopTtlDescription2, validateTTl(), 1),
			PfLocalPreference: schemaIntOptionalDefault(PfLocalPreferenceDescription, 0),
			PfMed:             schemaIntOptionalDefault(PfMedDescription4, 0),
			PfAsPrepend:       schemaIntOptionalValidateDefault(PfAsPrependDescription, validateAsPrepend(), 0),
			PfOrlonger:        schemaBoolOptionalDefault(PfOrlongerDescription, false),
			PfBfdInterval:     schemaIntOptionalValidateDefault(PfBfdIntervalDescription, validateBfdInterval(), 0),
			PfBfdMultiplier:   schemaIntOptionalValidateDefault(PfBfdMultiplierDescription, validateBfdMultiplier(), 0),
			PfDisabled:        schemaBoolOptionalDefault(PfDisabledDescription, false),
			PfNat: {
				Type:        schema.TypeSet,
				MaxItems:    1,
				Optional:    true,
				Description: PfNatDescription,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPreNatSources: schemaStringListOptionalDescribed(PfPreNatSourcesDescription3, PfPreNatSourcesDescription),
						PfPoolPrefixes:  schemaStringListOptionalDescribed(PfPoolPrefixesDescription, PfPreNatSourcesDescription),
						PfDirection:     schemaStringOptionalValidateDefault(PfDirectionDescription, validateDirection(), PfOutput),
						PfNatType:       schemaStringOptionalValidateDefault(PfNatTypeDescription, validateNatType(), PfNatTypeOverload),
						PfDnatMappings: {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: PfDnatMappingsDescription,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfPrivatePrefix:     schemaPrefixRequired(PfPrivatePrefixDescription),
									PfPublicPrefix:      schemaPrefixRequired(PfPublicPrefixDescription),
									PfConditionalPrefix: schemaPrefixOptional(PfConditionalPrefixDescription),
								},
							},
						},
					},
				},
			},
			PfPrefixes: {
				Type:        schema.TypeSet,
				Required:    true,
				Description: PfPrefixesDescription2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPrefix:          schemaPrefixRequired(PfPrefixDescription2),
						PfMatchType:       schemaStringOptionalValidate(PfMatchTypeDescription, validateMatchType()),
						PfAsPrepend:       schemaIntOptionalValidateDefault(PfAsPrependDescription2, validateAsPrepend(), 0),
						PfMed:             schemaIntOptionalDefault(PfMedDescription5, 0),
						PfLocalPreference: schemaIntOptionalDefault(PfLocalPreferenceDescription4, 0),
						PfType:            schemaStringRequiredValidate(PfTypeDescription8, validateIO()),
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: BgpImportStatePassthroughContext,
		},
	}
}

func resourceBgpSessionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk(PfCircuitId)
	if !ok {
		return diag.FromErr(errors.New(MessageMissingCircuitIdDetail))
	}
	connCID, ok := d.GetOk(PfConnectionId)
	if !ok {
		return diag.FromErr(errors.New(MesssageCRCIdRequired))
	}
	prefixesSet := d.Get(PfPrefixes).(*schema.Set)
	prefixesList := prefixesSet.List()
	if err := validatePrefixes(prefixesList); err != nil {
		return diag.FromErr(err)
	}
	session := extractBgpSessionCreate(d)
	resp, err := c.CreateBgpSession(session, cID.(string), connCID.(string))
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}
	if err := checkCloudRouterConnectionStatus(c, cID.(string), connCID.(string)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.BgpSettingsUUID)
	return diags
}

func resourceBgpSessionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	var cID, connCID, bgpSettingsUUID string

	if circuitID, ok := d.GetOk(PfCircuitId); !ok {
		return diag.FromErr(errors.New(MessageFailGetCRCId))
	} else {
		cID = circuitID.(string)
	}
	if connectionCID, ok := d.GetOk(PfConnectionId); !ok {
		return diag.FromErr(errors.New(MessageFailGetCRCCId))
	} else {
		connCID = connectionCID.(string)
	}
	if settingsUUID, ok := d.GetOk(PfId); !ok {
		return diag.FromErr(errors.New(MessageFailGetBgpSettings))
	} else {
		bgpSettingsUUID = settingsUUID.(string)
	}
	if diags != nil || len(diags) > 0 {
		return diags
	}
	bgp, err := c.GetBgpSessionBy(cID, connCID, bgpSettingsUUID)
	if err != nil {
		return diag.FromErr(errors.New(MessageBgpSessionNotFound))
	}

	_ = setResourceDataKeys(d, bgp, PfRemoteAsn, PfDisabled, PfOrlonger, PfAddressFamily, PfMultihopTtl)

	// If not Azure (Subnet empty)
	if bgp.Subnet == "" {
		_ = setResourceDataKeys(d, bgp, PfL3Address, PfRemoteAddress)
	} else {
		// If Azure will unset l3_address remote_address as those aren't in the BGP resource definition for Azure
		_ = d.Set(PfL3Address, nil)
		_ = d.Set(PfRemoteAddress, nil)
		// There is no way to know which Subnet is the primary or the secondary one
		// Display warning in case none of the is set in the state file
		if _, ok := d.GetOk(PfPrimarySubnet); !ok {
			if _, ok := d.GetOk(PfSecondarySubnet); !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  fmt.Sprintf("Manually set primary_subnet or secondary_subnet in Terraform state file using %s", bgp.Subnet),
				})
			}
		}
	}
	if bgp.Md5 != PfEmptyString {
		_ = d.Set(PfMd5, bgp.Md5)
	}
	_ = setResourceDataKeys(d, bgp, PfMed, PfAsPrepend, PfLocalPreference, PfBfdInterval, PfBfdMultiplier)

	if bgp.Nat != nil {
		nat := flattenNatConfiguration(bgp.Nat)
		if err := d.Set(PfNat, nat); err != nil {
			return diag.Errorf("error setting 'nat': %s", err)
		}
	}
	prefixes := flattenPrefixConfiguration(bgp.Prefixes)
	if err := d.Set(PfPrefixes, prefixes); err != nil {
		return diag.Errorf("error setting 'prefixes': %s", err)
	}

	return diags
}

func resourceBgpSessionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk(PfCircuitId)
	if !ok {
		return diag.FromErr(errors.New(MessageMissingCircuitIdDetail))
	}
	connCID, ok := d.GetOk(PfConnectionId)
	if !ok {
		return diag.FromErr(errors.New(MesssageCRCIdRequired))
	}
	if d.HasChange(PfPrimarySubnet) && d.HasChange(PfSecondarySubnet) {
		return diag.FromErr(errors.New(MessageCannotModifySubnets))
	}
	prefixesSet := d.Get(PfPrefixes).(*schema.Set)
	prefixesList := prefixesSet.List()
	if err := validatePrefixes(prefixesList); err != nil {
		return diag.FromErr(err)
	}
	session := extractBgpSessionUpdate(d)
	_, resp, err := c.UpdateBgpSession(session, cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	if err := checkCloudRouterConnectionStatus(c, cID.(string), connCID.(string)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.BgpSettingsUUID)
	return diags
}

func resourceBgpSessionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  MessageDeleteBgpSessionAndCRC,
	})
	d.SetId(PfEmptyString)
	return diags
}

func extractBgpSessionCreate(d *schema.ResourceData) packetfabric.BgpSession {
	bgpSession := packetfabric.BgpSession{}
	if l3Address, ok := d.GetOk(PfL3Address); ok {
		bgpSession.L3Address = l3Address.(string)
	}
	if primarySubnet, ok := d.GetOk(PfPrimarySubnet); ok {
		bgpSession.PrimarySubnet = primarySubnet.(string)
	}
	if secondarySubnet, ok := d.GetOk(PfSecondarySubnet); ok {
		bgpSession.SecondarySubnet = secondarySubnet.(string)
	}
	if addressFamily, ok := d.GetOk(PfAddressFamily); ok {
		bgpSession.AddressFamily = addressFamily.(string)
	}
	if remoteAddress, ok := d.GetOk(PfRemoteAddress); ok {
		bgpSession.RemoteAddress = remoteAddress.(string)
	}
	if remoteAsn, ok := d.GetOk(PfRemoteAsn); ok {
		bgpSession.RemoteAsn = remoteAsn.(int)
	}
	if multihopTTL, ok := d.GetOk(PfMultihopTtl); ok {
		bgpSession.MultihopTTL = multihopTTL.(int)
	}
	if localPreference, ok := d.GetOk(PfLocalPreference); ok {
		bgpSession.LocalPreference = localPreference.(int)
	}
	if med, ok := d.GetOk(PfMed); ok {
		bgpSession.Med = med.(int)
	}
	if asPrepend, ok := d.GetOk(PfAsPrepend); ok {
		bgpSession.AsPrepend = asPrepend.(int)
	}
	if orlonger, ok := d.GetOk(PfOrlonger); ok {
		bgpSession.Orlonger = orlonger.(bool)
	}
	if bfdInterval, ok := d.GetOk(PfBfdInterval); ok {
		bgpSession.BfdInterval = bfdInterval.(int)
	}
	if bfdMultiplier, ok := d.GetOk(PfBfdMultiplier); ok {
		bgpSession.BfdMultiplier = bfdMultiplier.(int)
	}
	if md5, ok := d.GetOk(PfMd5); ok {
		bgpSession.Md5 = md5.(string)
	}
	if nat, ok := d.GetOk(PfNat); ok {
		for _, nat := range nat.(*schema.Set).List() {
			bgpSession.Nat = extractConnBgpSessionNat(nat.(map[string]interface{}))
		}
	} else {
		bgpSession.Nat = nil
	}
	bgpSession.Prefixes = extractConnBgpSessionPrefixes(d)
	return bgpSession
}

func extractBgpSessionUpdate(d *schema.ResourceData) packetfabric.BgpSession {
	bgpSession := packetfabric.BgpSession{}
	if l3Address, ok := d.GetOk(PfL3Address); ok {
		bgpSession.L3Address = l3Address.(string)
	}
	// https://docs.packetfabric.com/api/v2/swagger/#/Cloud%20Router%20BGP%20Session%20Settings/cloud_routers_bgp_update
	// Azure BGP session Update: l3_address = Azure Subnet (primary or secondary)
	// set l3Address based on the values of primarySubnet and secondarySubnet when modified
	// This is a temporary solution until the BGP API is refactored.
	if d.HasChange(PfPrimarySubnet) {
		if primarySubnet, ok := d.GetOk(PfPrimarySubnet); ok {
			bgpSession.L3Address = primarySubnet.(string)
		}
	}
	if d.HasChange(PfSecondarySubnet) {
		if secondarySubnet, ok := d.GetOk(PfSecondarySubnet); ok {
			bgpSession.L3Address = secondarySubnet.(string)
		}
	}
	if d.HasChange(PfDisabled) {
		if disabled, ok := d.GetOk(PfDisabled); ok {
			bgpSession.Disabled = disabled.(bool)
		}
	}
	if addressFamily, ok := d.GetOk(PfAddressFamily); ok {
		bgpSession.AddressFamily = addressFamily.(string)
	}
	//remote_address not used for Azure
	if remoteAddress, ok := d.GetOk(PfRemoteAddress); ok {
		bgpSession.RemoteAddress = remoteAddress.(string)
	}
	if remoteAsn, ok := d.GetOk(PfRemoteAsn); ok {
		bgpSession.RemoteAsn = remoteAsn.(int)
	}
	if multihopTTL, ok := d.GetOk(PfMultihopTtl); ok {
		bgpSession.MultihopTTL = multihopTTL.(int)
	}
	if localPreference, ok := d.GetOk(PfLocalPreference); ok {
		bgpSession.LocalPreference = localPreference.(int)
	}
	if med, ok := d.GetOk(PfMed); ok {
		bgpSession.Med = med.(int)
	}
	if asPrepend, ok := d.GetOk(PfAsPrepend); ok {
		bgpSession.AsPrepend = asPrepend.(int)
	}
	if orlonger, ok := d.GetOk(PfOrlonger); ok {
		bgpSession.Orlonger = orlonger.(bool)
	}
	if bfdInterval, ok := d.GetOk(PfBfdInterval); ok {
		bgpSession.BfdInterval = bfdInterval.(int)
	}
	if bfdMultiplier, ok := d.GetOk(PfBfdMultiplier); ok {
		bgpSession.BfdMultiplier = bfdMultiplier.(int)
	}
	if md5, ok := d.GetOk(PfMd5); ok {
		bgpSession.Md5 = md5.(string)
	}
	if nat, ok := d.GetOk(PfNat); ok {
		for _, nat := range nat.(*schema.Set).List() {
			bgpSession.Nat = extractConnBgpSessionNat(nat.(map[string]interface{}))
		}
	} else {
		bgpSession.Nat = nil
	}
	bgpSession.Prefixes = extractConnBgpSessionPrefixes(d)
	return bgpSession
}

func extractConnBgpSessionPrefixes(d *schema.ResourceData) []packetfabric.BgpPrefix {
	if prefixes, ok := d.GetOk(PfPrefixes); ok {
		sessionPrefixes := make([]packetfabric.BgpPrefix, 0)
		for _, pref := range prefixes.(*schema.Set).List() {
			sessionPrefixes = append(sessionPrefixes, packetfabric.BgpPrefix{
				Prefix:          pref.(map[string]interface{})[PfPrefix].(string),
				MatchType:       pref.(map[string]interface{})[PfMatchType].(string),
				AsPrepend:       pref.(map[string]interface{})[PfAsPrepend].(int),
				Med:             pref.(map[string]interface{})[PfMed].(int),
				LocalPreference: pref.(map[string]interface{})[PfLocalPreference].(int),
				Type:            pref.(map[string]interface{})[PfType].(string),
			})
		}
		return sessionPrefixes
	}
	return make([]packetfabric.BgpPrefix, 0)
}

func extractConnBgpSessionNat(n map[string]interface{}) *packetfabric.BgpNat {
	nat := packetfabric.BgpNat{}
	if direction := n[PfDirection]; direction != nil {
		nat.Direction = direction.(string)
	}
	if natType := n[PfNatType]; natType != nil {
		nat.NatType = natType.(string)
	}
	nat.PreNatSources = extractPreNatSources(n[PfPreNatSources])
	nat.PoolPrefixes = extractPoolPrefixes(n[PfPoolPrefixes])
	nat.DnatMappings = extractConnBgpSessionDnat(n[PfDnatMappings].(*schema.Set))
	return &nat
}

func extractPreNatSources(d interface{}) []interface{} {
	if PreNatSources, ok := d.([]interface{}); ok {
		regs := make([]interface{}, 0)
		for _, reg := range PreNatSources {
			regs = append(regs, reg.(string))
		}
		return regs
	}
	return make([]interface{}, 0)
}

func extractPoolPrefixes(d interface{}) []interface{} {
	if PoolPrefixes, ok := d.([]interface{}); ok {
		regs := make([]interface{}, 0)
		for _, reg := range PoolPrefixes {
			regs = append(regs, reg.(string))
		}
		return regs
	}
	return make([]interface{}, 0)
}

func extractConnBgpSessionDnat(d *schema.Set) []packetfabric.BgpDnatMapping {
	sessionDnat := make([]packetfabric.BgpDnatMapping, 0)
	for _, dnat := range d.List() {
		sessionDnat = append(sessionDnat, packetfabric.BgpDnatMapping{
			PrivateIP:         dnat.(map[string]interface{})[PfPrivatePrefix].(string),
			PublicIP:          dnat.(map[string]interface{})[PfPublicPrefix].(string),
			ConditionalPrefix: dnat.(map[string]interface{})[PfConditionalPrefix].(string),
		})
	}
	return sessionDnat
}

func flattenNatConfiguration(nat *packetfabric.BgpNat) []interface{} {
	if nat == nil {
		return nil
	}

	data := mapStruct(&nat, PfPreNatSources, PfPoolPrefixes, PfDirection, PfNatType)

	if nat.DnatMappings != nil {
		data[PfDnatMappings] = flattenDnatMappings(nat.DnatMappings)
	}

	return []interface{}{data}
}

func flattenDnatMappings(dnatMappings []packetfabric.BgpDnatMapping) []interface{} {
	result := make([]interface{}, len(dnatMappings))
	for i, dnat := range dnatMappings {
		data := mapStruct(&dnat, PfConditionalPrefix)
		data[PfPrivatePrefix] = dnat.PrivateIP
		data[PfPublicPrefix] = dnat.PublicIP
		result[i] = data
	}
	return result
}

func flattenPrefixConfiguration(prefixes []packetfabric.BgpPrefix) []interface{} {
	fields := stringsToMap(PfPrefix, PfMatchType, PfAsPrepend, PfMed, PfLocalPreference, PfType)
	result := make([]interface{}, len(prefixes))
	for i, prefix := range prefixes {
		result[i] = structToMap(&prefix, fields)
	}
	return result
}
