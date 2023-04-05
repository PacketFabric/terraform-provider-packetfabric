package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudServicesConnInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudServicesRead,
		Schema: map[string]*schema.Schema{
			"cloud_circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique PF circuit ID for this connection\n\t\tExample: PF-AP-LAX1-1002",
			},
			"customer_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The UUID for the customer this connection belongs to.",
			},
			"user_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The UUID for the user this connection belongs to.",
			},
			"state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The state of the connection.\n\t\tEnum: [ \"active\", \"deleting\", \"inactive\", \"pending\", \"requested\" ]",
			},
			"service_provider": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The service provider of the connection\n\t\tEnum: [ \"aws\", \"azure\", \"packet\", \"google\", \"ibm\", \"salesforce\", \"webex\" ]",
			},
			"service_class": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The service class for the given port, either long haul or metro.\n\t\tEnum: [ \"longhaul\", \"metro\" ]",
			},
			"port_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port type for the given port.\n\t\tEnum: [ \"hosted\", \"dedicated\" ]",
			},
			"speed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The desired speed of the connection.\n\t\tEnum: [ \"50Mbps\", \"100Mbps\", \"200Mbps\", \"300Mbps\", \"400Mbps\", \"500Mbps\", \"1Gbps\", \"2Gbps\", \"5Gbps\", \"10Gbps\" ]",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this connection.\n\t\tExample: AWS connection for Foo Corp.",
			},
			"cloud_provider_pop": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Point of Presence for the cloud provider location.\n\t\tExample: DAL1",
			},
			"cloud_provider_region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Region short name.\n\t\tExample: us-west-1",
			},
			"settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_id_pf": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vlan_id_cust": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"svlan_id_cust": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"aws_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_hosted_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_connection_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"google_vlan_attachment_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"google_pairing_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan_id_private": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vlan_id_microsoft": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"azure_service_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"azure_service_tag": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"azure_connection_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"oracle_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vc_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_cross_connect_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_compartment_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gateway_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bgp_asn": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bgp_cer_cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bgp_ibm_cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"cloud_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_id_pf": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vlan_id_cust": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"svlan_id_cust": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"aws_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_hosted_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_connection_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"credentials_uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"aws_dx_location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_dx_bandwidth": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_dx_jumbo_frame_capable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"aws_dx_aws_device": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_dx_aws_logical_device_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_dx_has_logical_redundancy": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"aws_dx_mac_sec_capable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"aws_dx_encryption_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_vif_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_vif_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_vif_bgp_peer_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_vif_direct_connect_gw_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"google_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"google_project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"google_vlan_attachment_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"google_edge_availability_domain": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"google_dataplane_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"google_interface_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"google_pairing_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"google_cloud_router_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloud_state": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aws_dx_connection_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"aws_dx_port_encryption_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"aws_vif_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"google_interconnect_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"google_interconnect_admin_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"bgp_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"bgp_settings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertised_prefixes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"customer_asn": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"remote_asn": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"l3_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remote_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"address_family": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"md5": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"customer_router_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remote_router_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"google_advertise_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"google_keepalive_interval": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"google_advertised_ip_ranges": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"time_created": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Date and time of connection creation",
			},
			"time_updated": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Date and time connection was last updated",
			},
			"pop": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Point of Presence for the connection.\n\t\tExample: LAS1",
			},
			"site": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Site name\n\t\tExample: SwitchNAP Las Vegas 7",
			},
			"is_awaiting_onramp": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not this connection is waiting for an onramp to be available before provisioning.",
			},
		},
	}
}

// https://docs.packetfabric.com/api/v2/redoc/#operation/get_connection
func dataSourceCloudServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	circuitID, ok := d.GetOk("cloud_circuit_id")
	if !ok {
		return diag.Errorf("cloud_circuit_id is required")
	}

	connInfo, err := c.GetCloudConnInfo(circuitID.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	// Flatten the connection information into a map.
	connInfoMap := flattenCloudConnInfo(connInfo)

	// Update the resource data with the flattened map.
	for k, v := range connInfoMap {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(connInfo.CloudCircuitID)

	return diags
}

func flattenCloudConnInfo(connInfo *packetfabric.CloudConnInfo) map[string]interface{} {
	connInfoMap := map[string]interface{}{
		"cloud_circuit_id":      connInfo.CloudCircuitID,
		"customer_uuid":         connInfo.CustomerUUID,
		"user_uuid":             connInfo.UserUUID,
		"state":                 connInfo.State,
		"service_provider":      connInfo.ServiceProvider,
		"service_class":         connInfo.ServiceClass,
		"port_type":             connInfo.PortType,
		"speed":                 connInfo.Speed,
		"description":           connInfo.Description,
		"cloud_provider_pop":    connInfo.CloudProvider.Pop,
		"cloud_provider_region": connInfo.CloudProvider.Region,
		"time_created":          connInfo.TimeCreated,
		"time_updated":          connInfo.TimeUpdated,
		"pop":                   connInfo.Pop,
		"site":                  connInfo.Site,
		"is_awaiting_onramp":    connInfo.IsAwaitingOnramp,
	}

	settingsList := make([]interface{}, 0)
	if connInfo.Settings != nil {
		settingsList = append(settingsList, flattenCloudConnInfoSettings(connInfo.Settings))
	}
	connInfoMap["settings"] = settingsList

	cloudSettingsList := make([]interface{}, 0)
	if connInfo.CloudSettings != nil {
		cloudSettingsList = append(cloudSettingsList, flattenCloudConnInfoCloudSettings(connInfo.CloudSettings))
	}
	connInfoMap["cloud_settings"] = cloudSettingsList

	return connInfoMap
}

func flattenCloudConnInfoSettings(settings *packetfabric.Settings) map[string]interface{} {
	return map[string]interface{}{
		"vlan_id_pf":                  settings.VlanIDPf,
		"vlan_id_cust":                settings.VlanIDCust,
		"svlan_id_cust":               settings.SvlanIDCust,
		"aws_region":                  settings.AwsRegion,
		"aws_hosted_type":             settings.AwsHostedType,
		"aws_account_id":              settings.AwsAccountID,
		"aws_connection_id":           settings.AwsConnectionID,
		"google_vlan_attachment_name": settings.GoogleVlanAttachmentName,
		"google_pairing_key":          settings.GooglePairingKey,
		"vlan_id_private":             settings.VlanIDPrivate,
		"vlan_id_microsoft":           settings.VlanIDMicrosoft,
		"azure_service_key":           settings.AzureServiceKey,
		"azure_service_tag":           settings.AzureServiceTag,
		"oracle_region":               settings.OracleRegion,
		"vc_ocid":                     settings.VcOcid,
		"port_cross_connect_ocid":     settings.PortCrossConnectOcid,
		"port_compartment_ocid":       settings.PortCompartmentOcid,
		"account_id":                  settings.AccountID,
		"gateway_id":                  settings.GatewayID,
		"port_id":                     settings.PortID,
		"name":                        settings.Name,
		"bgp_asn":                     settings.BgpAsn,
		"bgp_cer_cidr":                settings.BgpCerCidr,
		"bgp_ibm_cidr":                settings.BgpIbmCidr,
	}
}

func flattenCloudConnInfoCloudSettings(cloudSettings *packetfabric.CloudSettingsHosted) map[string]interface{} {
	cloudSettingsMap := map[string]interface{}{
		"vlan_id_pf":                      cloudSettings.VlanIDPf,
		"vlan_id_cust":                    cloudSettings.VlanIDCust,
		"svlan_id_cust":                   cloudSettings.SvlanIDCust,
		"aws_region":                      cloudSettings.AwsRegion,
		"aws_hosted_type":                 cloudSettings.AwsHostedType,
		"aws_account_id":                  cloudSettings.AwsAccountID,
		"aws_connection_id":               cloudSettings.AwsConnectionID,
		"credentials_uuid":                cloudSettings.CredentialsUUID,
		"mtu":                             cloudSettings.Mtu,
		"aws_dx_location":                 cloudSettings.AwsDxLocation,
		"aws_dx_bandwidth":                cloudSettings.AwsDxBandwidth,
		"aws_dx_jumbo_frame_capable":      cloudSettings.AwsDxJumboFrameCapable,
		"aws_dx_aws_device":               cloudSettings.AwsDxAWSDevice,
		"aws_dx_aws_logical_device_id":    cloudSettings.AwsDxAWSLogicalDeviceID,
		"aws_dx_has_logical_redundancy":   cloudSettings.AwsDxHasLogicalRedundancy,
		"aws_dx_mac_sec_capable":          cloudSettings.AwsDxMacSecCapable,
		"aws_dx_encryption_mode":          cloudSettings.AwsDxEncryptionMode,
		"aws_vif_type":                    cloudSettings.AwsVifType,
		"aws_vif_id":                      cloudSettings.AwsVifID,
		"aws_vif_bgp_peer_id":             cloudSettings.AwsVifBGPPeerID,
		"aws_vif_direct_connect_gw_id":    cloudSettings.AwsVifDirectConnectGwID,
		"google_region":                   cloudSettings.GoogleRegion,
		"google_project_id":               cloudSettings.GoogleProjectID,
		"google_vlan_attachment_name":     cloudSettings.GoogleVlanAttachmentName,
		"google_edge_availability_domain": cloudSettings.GoogleEdgeAvailabilityDomain,
		"google_dataplane_version":        cloudSettings.GoogleDataplaneVersion,
		"google_interface_name":           cloudSettings.GoogleInterfaceName,
		"google_pairing_key":              cloudSettings.GooglePairingKey,
		"google_cloud_router_name":        cloudSettings.GoogleCloudRouterName,
	}

	cloudStateList := make([]interface{}, 0)
	if cloudSettings.CloudState != nil {
		cloudStateList = append(cloudStateList, flattenCloudConnInfoCloudState(cloudSettings.CloudState))
	}
	cloudSettingsMap["cloud_state"] = cloudStateList

	bgpSettingsList := make([]interface{}, 0)
	if cloudSettings.BgpSettings != nil {
		bgpSettingsList = append(bgpSettingsList, flattenCloudConnInfoBGPSettings(cloudSettings.BgpSettings))
	}
	cloudSettingsMap["bgp_settings"] = bgpSettingsList

	return cloudSettingsMap
}

func flattenCloudConnInfoBGPSettings(bgpSettings *packetfabric.BgpSettings) map[string]interface{} {
	return map[string]interface{}{
		"customer_asn":                bgpSettings.CustomerAsn,
		"l3_address":                  bgpSettings.L3Address,
		"remote_address":              bgpSettings.RemoteAddress,
		"address_family":              bgpSettings.AddressFamily,
		"md5":                         bgpSettings.Md5,
		"advertised_prefixes":         bgpSettings.AdvertisedPrefixes,
		"customer_router_ip":          bgpSettings.CustomerRouterIp,
		"remote_router_ip":            bgpSettings.RemoteRouterIp,
		"google_advertise_mode":       bgpSettings.GoogleAdvertiseMode,
		"google_keepalive_interval":   bgpSettings.GoogleKeepaliveInterval,
		"google_advertised_ip_ranges": bgpSettings.GoogleAdvertisedIPRanges,
	}
}

func flattenCloudConnInfoCloudState(cloudState *packetfabric.CloudStateHosted) map[string]interface{} {
	return map[string]interface{}{
		"aws_dx_connection_state":           cloudState.AwsDxConnectionState,
		"aws_dx_port_encryption_status":     cloudState.AwsDxPortEncryptionStatus,
		"aws_vif_state":                     cloudState.AwsVifState,
		"bgp_state":                         cloudState.BgpState,
		"google_interconnect_state":         cloudState.GoogleInterconnectState,
		"google_interconnect_admin_enabled": cloudState.GoogleInterconnectAdminEnabled,
	}
}
