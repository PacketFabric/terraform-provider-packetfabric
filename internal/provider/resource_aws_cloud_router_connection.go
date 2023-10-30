package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterConnectionAws() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterConnectionAwsCreate,
		ReadContext:   resourceRouterConnectionAwsRead,
		UpdateContext: resourceRouterConnectionAwsUpdate,
		DeleteContext: resourceRouterConnectionAwsDelete,
		Timeouts:      schemaTimeouts(60, 10, 60, 60),
		Schema: map[string]*schema.Schema{
			PfId:                        schemaStringComputedPlain(),
			PfCircuitId:                 schemaStringRequiredNew(PfCircuitIdDescription),
			PfBgpSettingsUuid:           schemaStringOptionalComputed(PfBgpSettingsUuidDescription3),
			PfAwsAccountId:              schemaAwsAccountId(PfAwsAccountIdDescription3),
			PfAccountUuid:               schemaAccountUuid(PfAccountUuidDescription2),
			PfMaybeNat:                  schemaBoolOptionalNew(PfMaybeNatDescription),
			PfMaybeDnat:                 schemaBoolOptionalNew(PfMaybeDnatDescription),
			PfDescription:               schemaStringRequired(PfConnectionDescription),
			PfPop:                       schemaStringRequiredNew(PfPopDescription5),
			PfZone:                      schemaStringRequiredNew(PfZoneDescription),
			PfIsPublic:                  schemaBoolOptionalNewDefault(PfIsPublicDescription, false),
			PfSpeed:                     schemaStringRequired(PfSpeedDescriptionH),
			PfSubscriptionTerm:          schemaIntOptionalValidateDefault(PfSubscriptionTermDescription2, validateSubscriptionTerm(), 1),
			PfPublishedQuoteLineUuid:    schemaStringOptionalNewValidate(PfPublishedQuoteLineUuidDescription, validation.IsUUID),
			PfPoNumber:                  schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfLabels:                    schemaStringSetOptional(PfLabelsDescription),
			PfCloudProviderConnectionId: schemaStringComputed(PfCloudProviderConnectionIdDescription),
			PfVlanIdPf:                  schemaIntComputed(PfVlanIdPfDescription),
			PfCloudSettings: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: PfCloudSettingsDescription,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfCredentialsUuid: schemaStringRequiredValidate(PfCredentialsUuidDescription, validation.IsUUID),
						PfAwsRegion:       schemaStringOptionalNotEmpty(PfAwsRegionDescription),
						PfMtu:             schemaIntOptionalValidateDefault(PfMtuDescription2, validateMtu(), PfMtuDefault),
						PfAwsVifType:      schemaStringRequiredValidate(PfAwsVifTypeDescription, validateVifType()),
						PfAwsGateways: {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    2,
							Description: PfAwsGatewaysDescription,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfType:      schemaStringRequiredValidate(PfTypeDescription6, validateVifType()),
									PfName:      schemaStringOptionalNotEmpty(PfNameDescription4),
									PfId:        schemaStringOptionalNotEmpty(PfIdDescription2),
									PfAsn:       schemaIntOptional(PfAsnDescription),
									PfVpcId:     schemaStringOptionalNotEmpty(PfVpcIdDescription),
									PfSubnetIds: schemaStringListOptional(PfSubnetIdsDescription),
								},
							},
						},
						PfBgpSettings: resourceRouterConnectionAwsBgpSettings(),
					},
				},
			},
			PfEtl: schemaFloatComputed(PfEtlDescription),
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, d *schema.ResourceDiff, m interface{}) error {
				if d.Id() == PfEmptyString {
					return nil
				}
				if _, ok := d.GetOk(PfCloudSettings); !ok {
					return nil
				}

				attributes := []string{
					"cloud_settings.0.aws_region",
					"cloud_settings.0.aws_vif_type",
					"cloud_settings.0.aws_gateways",
				}

				for _, attribute := range attributes {
					oldRaw, newRaw := d.GetChange(attribute)
					if oldRaw != nil && !reflect.DeepEqual(oldRaw, newRaw) {
						return fmt.Errorf("updating %s in-place is not supported, delete and recreate the resource with the updated values", attribute)
					}
				}
				return nil
			},
		),
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceRouterConnectionAwsBgpSettings() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				PfMd5:             schemaStringOptionalNotEmpty(PfMd5Description),
				PfL3Address:       schemaStringOptionalNotEmpty(PfL3AddressDescription2),
				PfRemoteAddress:   schemaStringOptionalNotEmpty(PfRemoteAddressDescription2),
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
										PfPrivatePrefix:     schemaStringRequiredValidate(PfPrivatePrefixDescription, validateIPAddressWithPrefix),
										PfPublicPrefix:      schemaStringRequiredValidate(PfPublicPrefixDescription, validateIPAddressWithPrefix),
										PfConditionalPrefix: schemaStringOptionalValidate(PfConditionalPrefixDescription, validateIPAddressWithPrefix),
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
							PfPrefix:          schemaStringRequiredValidate(PfPrefixDescription2, validateIPAddressWithPrefix),
							PfMatchType:       schemaStringOptionalValidate(PfMatchTypeDescription, validateMatchType()),
							PfAsPrepend:       schemaIntOptionalValidateDefault(PfAsPrependDescription2, validateAsPrepend(), 0),
							PfMed:             schemaIntOptionalDefault(PfMedDescription5, 0),
							PfLocalPreference: schemaIntOptionalDefault(PfLocalPreferenceDescription4, 0),
							PfType:            schemaStringRequiredValidate(PfTypeDescription8, validateIO()),
						},
					},
				},
			},
		},
	}
}

func resourceRouterConnectionAwsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	awsConn := extractAwsConnection(d)

	if cid, ok := d.GetOk(PfCircuitId); ok {
		cloudSettingsList := d.Get(PfCloudSettings).([]interface{})
		if len(cloudSettingsList) != 0 {
			cloudSettings := cloudSettingsList[0].(map[string]interface{})
			bgpSettingsList := cloudSettings[PfBgpSettings].([]interface{})
			if len(bgpSettingsList) != 0 {
				bgpSettings := bgpSettingsList[0].(map[string]interface{})
				prefixesValue := bgpSettings[PfPrefixes]
				prefixesSet := prefixesValue.(*schema.Set)
				prefixesList := prefixesSet.List()
				if err := validatePrefixes(prefixesList); err != nil {
					return diag.FromErr(err)
				}
			}
		}
		resp, err := c.CreateAwsConnection(awsConn, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		if err := checkCloudRouterConnectionStatus(c, cid.(string), resp.CloudCircuitID); err != nil {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.CloudCircuitID)
			_ = d.Set(PfSubscriptionTerm, resp.Billing.SubscriptionTerm)

			if _, ok := d.GetOk(PfCloudSettings); !ok {
				time.Sleep(90 * time.Second) // wait for the connection to show on AWS
				resp, err := c.ReadCloudRouterConnection(cid.(string), resp.CloudCircuitID)
				if err != nil {
					diags = diag.FromErr(err)
					return diags
				}

				if resp.CloudProviderConnectionID == PfEmptyString || resp.CloudSettings.VlanIDPf == 0 {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  MessageIncompleteCR,
						Detail:   MessageIncompleteCRDetail,
					})
					return diags
				} else {
					_ = d.Set(PfCloudProviderConnectionId, resp.CloudProviderConnectionID)
					_ = d.Set(PfVlanIdPf, resp.CloudSettings.VlanIDPf)
				}
			}

			if _, ok := d.GetOk(PfCloudSettings); ok {
				// Extract the BGP settings UUID
				resp, err := c.ReadCloudRouterConnection(cid.(string), resp.CloudCircuitID)
				if err != nil {
					diags = diag.FromErr(err)
					return diags
				}
				if len(resp.BgpStateList) > 0 {
					_ = d.Set(PfBgpSettingsUuid, resp.BgpStateList[0].BgpSettingsUUID)
				}
			}
			if labels, ok := d.GetOk(PfLabels); ok {
				diagnostics, created := createLabels(c, d.Id(), labels)
				if !created {
					return diagnostics
				}
			}
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  MessageMissingCircuitId,
			Detail:   MessageMissingCircuitIdDetail,
		})
	}
	return diags
}

func resourceRouterConnectionAwsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	circuitID, ok := d.GetOk(PfCircuitId)
	if !ok {
		return diag.FromErr(errors.New(MessageMissingCircuitIdDetail))
	}

	cloudConnCID := d.Get(PfId)
	resp, err := c.ReadCloudRouterConnection(circuitID.(string), cloudConnCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	_ = setResourceDataKeys(d, resp, PfAccountUuid, PfSubscriptionTerm, PfDescription, PfSpeed, PfZone, PfPoNumber)
	_ = d.Set(PfCircuitId, resp.CloudRouterCircuitID)
	_ = d.Set(PfPop, resp.CloudProvider.Pop)
	_ = d.Set(PfAwsAccountId, resp.CloudSettings.AwsAccountID)

	if resp.CloudSettings.PublicIP != PfEmptyString {
		_ = d.Set(PfIsPublic, true)
	} else {
		_ = d.Set(PfIsPublic, false)
	}

	if _, ok := d.GetOk(PfCloudSettings); ok {
		// Extract the BGP settings UUID
		var bgpSettingsUUID string
		if len(resp.BgpStateList) > 0 {
			bgpSettingsUUID = resp.BgpStateList[0].BgpSettingsUUID
			_ = d.Set(PfBgpSettingsUuid, bgpSettingsUUID)
		}
		bgp, err := c.GetBgpSessionBy(circuitID.(string), cloudConnCID.(string), bgpSettingsUUID)
		if err != nil {
			return diag.FromErr(errors.New(MessageBgpSessionNotFound))
		}
		cloudSettings := make(map[string]interface{})
		cloudSettings[PfCredentialsUuid] = resp.CloudSettings.CredentialsUUID
		cloudSettings[PfAwsRegion] = resp.CloudSettings.AwsRegion
		if _, ok := d.GetOk("cloud_settings.0.mtu"); ok {
			cloudSettings[PfMtu] = resp.CloudSettings.Mtu
		}
		cloudSettings[PfAwsVifType] = resp.CloudSettings.AwsVifType
		bgpSettings := make(map[string]interface{})
		if bgp != nil {
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.remote_asn"); ok {
				bgpSettings[PfRemoteAsn] = bgp.RemoteAsn
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.l3_address"); ok {
				bgpSettings[PfL3Address] = bgp.L3Address
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.remote_address"); ok {
				bgpSettings[PfRemoteAddress] = bgp.RemoteAddress
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.disabled"); ok {
				bgpSettings[PfDisabled] = bgp.Disabled
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.orlonger"); ok {
				bgpSettings[PfOrlonger] = bgp.Orlonger
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.md5"); ok {
				bgpSettings[PfMd5] = bgp.Md5
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.med"); ok {
				bgpSettings[PfMed] = bgp.Med
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.as_prepend"); ok {
				bgpSettings[PfAsPrepend] = bgp.AsPrepend
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.local_preference"); ok {
				bgpSettings[PfLocalPreference] = bgp.LocalPreference
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.bfd_interval"); ok {
				bgpSettings[PfBfdInterval] = bgp.BfdInterval
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.bfd_multiplier"); ok {
				bgpSettings[PfBfdMultiplier] = bgp.BfdMultiplier
			}
			if bgp.Nat != nil {
				nat := flattenNatConfiguration(bgp.Nat)
				bgpSettings[PfNat] = nat
			}
			prefixes := flattenPrefixConfiguration(bgp.Prefixes)
			bgpSettings[PfPrefixes] = prefixes
		}
		cloudSettings[PfBgpSettings] = bgpSettings

		awsGateways := make([]map[string]interface{}, len(resp.CloudSettings.AwsGateways))
		awsGatewayFields := stringsToMap(PfType, PfName, PfId, PfAsn, PfVpcId, PfSubnetIds, PfAllowedPrefixes)
		for i, gateway := range resp.CloudSettings.AwsGateways {
			awsGateways[i] = structToMap(&gateway, awsGatewayFields)
		}
		cloudSettings[PfAwsGateways] = awsGateways
		_ = d.Set(PfCloudSettings, cloudSettings)
	} else {
		if resp.CloudProviderConnectionID == PfEmptyString || resp.CloudSettings.VlanIDPf == 0 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  MessageIncompleteCR,
				Detail:   MessageIncompleteCRDetail,
			})
			return diags
		} else {
			_ = d.Set(PfCloudProviderConnectionId, resp.CloudProviderConnectionID)
			_ = d.Set(PfVlanIdPf, resp.CloudSettings.VlanIDPf)
		}
	}
	// unsetFields: published_quote_line_uuid

	if _, ok := d.GetOk(PfLabels); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set(PfLabels, labels)
	}

	etl, err := c.GetEarlyTerminationLiability(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if etl > 0 {
		_ = d.Set(PfEtl, etl)
	}
	return diags
}

func resourceRouterConnectionAwsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceRouterConnectionAwsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractAwsConnection(d *schema.ResourceData) packetfabric.AwsConnection {
	awsConn := packetfabric.AwsConnection{}
	if awsAccountID, ok := d.GetOk(PfAwsAccountId); ok {
		awsConn.AwsAccountID = awsAccountID.(string)
	}
	if accountUUID, ok := d.GetOk(PfAccountUuid); ok {
		awsConn.AccountUUID = accountUUID.(string)
	}
	if maybeNat, ok := d.GetOk(PfMaybeNat); ok {
		awsConn.MaybeNat = maybeNat.(bool)
	}
	if maybeDNat, ok := d.GetOk(PfMaybeDnat); ok {
		awsConn.MaybeDNat = maybeDNat.(bool)
	}
	if description, ok := d.GetOk(PfDescription); ok {
		awsConn.Description = description.(string)
	}
	if pop, ok := d.GetOk(PfPop); ok {
		awsConn.Pop = pop.(string)
	}
	if zone, ok := d.GetOk(PfZone); ok {
		awsConn.Zone = zone.(string)
	}
	if isPublic, ok := d.GetOk(PfIsPublic); ok {
		awsConn.IsPublic = isPublic.(bool)
	}
	if subscriptionTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
		awsConn.SubscriptionTerm = subscriptionTerm.(int)
	}
	if speed, ok := d.GetOk(PfSpeed); ok {
		awsConn.Speed = speed.(string)
	}
	if publishedQuoteLineUUID, ok := d.GetOk(PfPublishedQuoteLineUuid); ok {
		awsConn.PublishedQuoteLineUUID = publishedQuoteLineUUID.(string)
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		awsConn.PONumber = poNumber.(string)
	}
	if cloudSettingsList, ok := d.GetOk(PfCloudSettings); ok {
		cs := cloudSettingsList.([]interface{})[0].(map[string]interface{})
		awsConn.CloudSettings = extractAwsRouterCloudConnSettings(cs)
	}
	return awsConn
}

func extractAwsRouterCloudConnSettings(cs map[string]interface{}) *packetfabric.CloudSettings {
	cloudSettings := &packetfabric.CloudSettings{}
	cloudSettings.CredentialsUUID = cs[PfCredentialsUuid].(string)

	if awsRegion, ok := cs[PfAwsRegion]; ok {
		cloudSettings.AwsRegion = awsRegion.(string)
	}
	cloudSettings.AwsVifType = cs[PfAwsVifType].(string)
	if awsGateways, ok := cs[PfAwsGateways]; ok {
		cloudSettings.AwsGateways = extractAwsGateways(awsGateways.([]interface{}))
	}
	if mtu, ok := cs[PfMtu]; ok {
		cloudSettings.Mtu = mtu.(int)
	}
	if bgpSettings, ok := cs[PfBgpSettings]; ok {
		bgpSettingsMap := bgpSettings.([]interface{})[0].(map[string]interface{})
		cloudSettings.BgpSettings = extractRouterConnBgpSettings(bgpSettingsMap)
	}
	return cloudSettings
}
