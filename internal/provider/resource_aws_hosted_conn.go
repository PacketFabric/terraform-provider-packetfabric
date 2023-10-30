package provider

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAwsRequestHostConn() *schema.Resource {
	return &schema.Resource{
		Timeouts:      schemaTimeouts(60, 10, 60, 60),
		CreateContext: resourceAwsReqHostConnCreate,
		UpdateContext: resourceAwsReqHostConnUpdate,
		ReadContext:   resourceAwsReqHostConnRead,
		DeleteContext: resourceAwsServicesDelete,
		Schema: map[string]*schema.Schema{
			PfId:                        schemaStringComputedPlain(),
			PfAwsAccountId:              schemaAwsAccountId(PfAwsAccountIdDescription3),
			PfAccountUuid:               schemaAccountUuid(PfAccountUuidDescription2),
			PfDescription:               schemaStringRequiredNotEmpty(PfConnectionDescription),
			PfPop:                       schemaStringRequiredNewNotEmpty(PfPopDescription7),
			PfPort:                      schemaStringRequiredNewNotEmpty(PfPortDescription3),
			PfVlan:                      schemaIntRequiredNewValidate(PfVlanDescription, validateVlan()),
			PfSrcSvlan:                  schemaIntOptionalNewValidate(PfSrcSvlanDescription, validateVlan()),
			PfZone:                      schemaStringRequiredNewNotEmpty(PfZoneDescription),
			PfSpeed:                     schemaStringRequiredValidate(PfSpeedDescriptionK, validateSpeed()),
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
						PfMtu:             schemaIntOptionalValidateDefault(PfMtuDescription2, validateMtu(), 1500),
						PfAwsVifType:      schemaStringRequiredValidate(PfAwsVifTypeDescription, validateVifType()),
						PfBgpSettings: {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfCustomerAsn:        schemaIntRequired(PfCustomerAsnDescription),
									PfL3Address:          schemaStringOptionalNotEmpty(PfL3AddressDescription2),
									PfRemoteAddress:      schemaStringOptionalNotEmpty(PfRemoteAddressDescription2),
									PfAddressFamily:      schemaStringOptionalValidateDefault(PfAddressFamilyDescription2, validateIpVersion(), PfAddressFamilyDefault),
									PfMd5:                schemaStringOptionalNotEmpty(PfMd5Description2),
									PfAdvertisedPrefixes: schemaStringListOptional(PfAdvertisedPrefixesDescription),
								},
							},
						},
						PfAwsGateways: {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    2,
							Description: PfAwsGatewaysDescription,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfType:            schemaStringRequiredValidate(PfTypeDescription6, validateVifType()),
									PfName:            schemaStringOptionalNotEmpty(PfNameDescription4),
									PfId:              schemaStringOptionalNotEmpty(PfIdDescription2),
									PfAsn:             schemaIntOptional(PfAsnDescription),
									PfVpcId:           schemaStringOptionalNotEmpty(PfVpcIdDescription),
									PfSubnetIds:       schemaStringListOptional(PfSubnetIdsDescription),
									PfAllowedPrefixes: schemaStringListOptional(PfAllowedPrefixesDescription),
								},
							},
						},
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
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsReqHostConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	reqConn := extractAwsReqConn(d)
	expectedResp, err := c.CreateAwsHostedConn(reqConn)
	if err != nil {
		return diag.FromErr(err)
	}
	// Cloud Everywhere: if cloud_circuit_id is null display error
	if expectedResp.CloudCircuitID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  MessageMissingHostedLocation,
			Detail:   MessageMissingHostedLocationDetail,
		})
		return diags
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			dedicatedConns, err := c.GetCurrentCustomersHosted()
			if dedicatedConns != nil && err == nil && len(dedicatedConns) > 0 {
				for _, conn := range dedicatedConns {
					if expectedResp.UUID == conn.UUID && conn.State == PfActive {
						expectedResp.CloudCircuitID = conn.CloudCircuitID
						ticker.Stop()
						createOk <- true
					}
				}
			}
		}
	}()
	<-createOk
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(expectedResp.CloudCircuitID)

	if _, ok := d.GetOk(PfCloudSettings); !ok {
		time.Sleep(90 * time.Second) // wait for the connection to show on AWS
		resp, err := c.GetCloudConnInfo(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		if resp.CloudProviderConnectionID == PfEmptyString || resp.Settings.VlanIDPf == 0 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  MessageIncompleteCR,
				Detail:   MessageIncompleteCRDetail,
			})
			return diags
		} else {
			_ = d.Set(PfCloudProviderConnectionId, resp.CloudProviderConnectionID)
			_ = d.Set(PfVlanIdPf, resp.Settings.VlanIDPf)
		}
	}

	if labels, ok := d.GetOk(PfLabels); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}
	return diags
}

func resourceAwsReqHostConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = setResourceDataKeys(d, resp, PfAccountUuid, PfDescription, PfSpeed, PfPoNumber)
		_ = d.Set(PfVlan, resp.Settings.VlanIDCust)
		_ = d.Set(PfPop, resp.CloudProvider.Pop)
		_ = d.Set(PfAwsAccountId, resp.Settings.AwsAccountID)

		if _, ok := d.GetOk(PfCloudSettings); ok {
			cloudSettings := mapStruct(&resp.CloudSettings, PfCredentialsUuid, PfAwsRegion, PfMtu, PfAwsVifType)
			cloudSettings[PfBgpSettings] = mapStruct(&resp.CloudSettings.BgpSettings, PfCustomerAsn, PfAddressFamily)

			agFields := stringsToMap(PfType, PfName, PfId, PfAsn, PfVpcId, PfSubnetIds, PfAllowedPrefixes)
			awsGateways := make([]map[string]interface{}, len(resp.CloudSettings.AwsGateways))
			for i, gateway := range resp.CloudSettings.AwsGateways {
				awsGateways[i] = structToMap(&gateway, agFields)
			}
			cloudSettings[PfAwsGateways] = awsGateways
			_ = d.Set(PfCloudSettings, cloudSettings)
		} else {
			if resp.CloudProviderConnectionID == "" || resp.Settings.VlanIDPf == 0 {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  MessageIncompleteCR,
					Detail:   MessageIncompleteCRDetail,
				})
				return diags
			} else {
				_ = d.Set(PfCloudProviderConnectionId, resp.CloudProviderConnectionID)
				_ = d.Set(PfVlanIdPf, resp.Settings.VlanIDPf)
			}
		}
	}
	resp2, err2 := c.GetBackboneByVcCID(d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	if resp2 != nil {
		_ = d.Set(PfPort, resp2.Interfaces[0].PortCircuitID) // Port A
		if _, ok := d.GetOk(PfSrcSvlan); ok {
			if resp2.Interfaces[0].Svlan != 0 {
				_ = d.Set(PfSrcSvlan, resp2.Interfaces[0].Svlan) // Port A if ENNI
			}
		}
		_ = d.Set(PfZone, resp2.Interfaces[1].Zone) // Port Z
	}

	if _, ok := d.GetOk(PfLabels); ok {
		labels, err3 := getLabels(c, d.Id())
		if err3 != nil {
			return diag.FromErr(err3)
		}
		_ = d.Set(PfLabels, labels)
	}

	etl, err4 := c.GetEarlyTerminationLiability(d.Id())
	if err4 != nil {
		return diag.FromErr(err4)
	}
	if etl > 0 {
		_ = d.Set(PfEtl, etl)
	}

	return diags
}

func resourceAwsReqHostConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesHostedUpdate(ctx, d, m)
}

func extractAwsReqConn(d *schema.ResourceData) packetfabric.HostedAwsConnection {
	hostedAwsConn := packetfabric.HostedAwsConnection{}
	if awsAccountID, ok := d.GetOk(PfAwsAccountId); ok {
		hostedAwsConn.AwsAccountID = awsAccountID.(string)
	}
	if accountUUID, ok := d.GetOk(PfAccountUuid); ok {
		hostedAwsConn.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk(PfDescription); ok {
		hostedAwsConn.Description = description.(string)
	}
	if pop, ok := d.GetOk(PfPop); ok {
		hostedAwsConn.Pop = pop.(string)
	}
	if port, ok := d.GetOk(PfPort); ok {
		hostedAwsConn.Port = port.(string)
	}
	if vlan, ok := d.GetOk(PfVlan); ok {
		hostedAwsConn.Vlan = vlan.(int)
	}
	if srcSvlan, ok := d.GetOk(PfSrcSvlan); ok {
		hostedAwsConn.SrcSvlan = srcSvlan.(int)
	}
	if zone, ok := d.GetOk(PfZone); ok {
		hostedAwsConn.Zone = zone.(string)
	}
	if speed, ok := d.GetOk(PfSpeed); ok {
		hostedAwsConn.Speed = speed.(string)
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		hostedAwsConn.PONumber = poNumber.(string)
	}
	if cloudSettingsList, ok := d.GetOk(PfCloudSettings); ok {
		cs := cloudSettingsList.([]interface{})[0].(map[string]interface{})
		hostedAwsConn.CloudSettings = &packetfabric.CloudSettings{}
		hostedAwsConn.CloudSettings.CredentialsUUID = cs[PfCredentialsUuid].(string)
		if awsRegion, ok := cs[PfAwsRegion]; ok {
			hostedAwsConn.CloudSettings.AwsRegion = awsRegion.(string)
		}
		if mtu, ok := cs[PfMtu]; ok {
			hostedAwsConn.CloudSettings.Mtu = mtu.(int)
		}
		hostedAwsConn.CloudSettings.AwsVifType = cs[PfAwsVifType].(string)
		if bgpSettings, ok := cs[PfBgpSettings]; ok {
			bgpSettingsMap := bgpSettings.([]interface{})[0].(map[string]interface{})
			hostedAwsConn.CloudSettings.BgpSettings = &packetfabric.BgpSettings{}
			hostedAwsConn.CloudSettings.BgpSettings.CustomerAsn = bgpSettingsMap[PfCustomerAsn].(int)
			hostedAwsConn.CloudSettings.BgpSettings.AddressFamily = bgpSettingsMap[PfAddressFamily].(string)
		}
		if awsGateways, ok := cs[PfAwsGateways]; ok {
			hostedAwsConn.CloudSettings.AwsGateways = extractAwsGateways(awsGateways.([]interface{}))
		}
	}
	return hostedAwsConn
}

func extractAwsGateways(gateways []interface{}) []packetfabric.AwsGateway {
	var awsGateways []packetfabric.AwsGateway
	for _, gw := range gateways {
		gateway := gw.(map[string]interface{})

		subnetIDsInterface, subnetIDsExist := gateway[PfSubnetIds].([]interface{})
		var subnetIDs []string
		if subnetIDsExist {
			subnetIDs = make([]string, len(subnetIDsInterface))
			for i, elem := range subnetIDsInterface {
				subnetIDs[i] = elem.(string)
			}
		}

		allowedPrefixesInterface, allowedPrefixesExist := gateway[PfAllowedPrefixes].([]interface{})
		var allowedPrefixes []string
		if allowedPrefixesExist {
			allowedPrefixes = make([]string, len(allowedPrefixesInterface))
			for i, elem := range allowedPrefixesInterface {
				allowedPrefixes[i] = elem.(string)
			}
		}

		awsGateway := packetfabric.AwsGateway{}

		if t, ok := gateway[PfType].(string); ok {
			awsGateway.Type = t
		}
		if name, ok := gateway[PfName].(string); ok {
			awsGateway.Name = name
		}
		if id, ok := gateway[PfId].(string); ok {
			awsGateway.ID = id
		}
		if asn, ok := gateway[PfAsn].(int); ok {
			awsGateway.Asn = asn
		}
		if vpcID, ok := gateway[PfVpcId].(string); ok {
			awsGateway.VpcID = vpcID
		}

		awsGateway.SubnetIDs = subnetIDs
		awsGateway.AllowedPrefixes = allowedPrefixes

		awsGateways = append(awsGateways, awsGateway)
	}
	return awsGateways
}
