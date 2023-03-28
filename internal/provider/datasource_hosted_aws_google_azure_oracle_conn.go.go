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
						"aws_dx_aws_deviceV2": {
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

	// Retrieve the cloud circuit ID from the resource data.
	circuitID, ok := d.GetOk("cloud_circuit_id")
	if !ok {
		return diag.Errorf("cloud_circuit_id is required")
	}

	// Retrieve the connection information from the PacketFabric API.
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

	// Set the resource ID to the circuit ID.
	d.SetId(connInfo.CloudCircuitID)

	return diags
}

// flattenCloudConnInfo flattens a CloudConnectionInfo struct into a map.
func flattenCloudConnInfo(connInfo *packetfabric.CloudConnectionInfo) map[string]interface{} {
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

	// Flatten the settings list.
	settingsList := make([]interface{}, 0, len(connInfo.Settings))
	for _, s := range connInfo.Settings {
		settingsList = append(settingsList, flattenCloudConnInfoSettings(s))
	}
	connInfoMap["settings"] = settingsList

	// Flatten the cloud_settings list.
	cloudSettingsList := make([]interface{}, 0, len(connInfo.CloudSettings))
	for _, cs := range connInfo.CloudSettings {
		cloudSettingsList = append(cloudSettingsList, flattenCloudConnInfoCloudSettings(cs))
	}
	connInfoMap["cloud_settings"] = cloudSettingsList

	return connInfoMap
}

// flattenCloudConnInfoSettings flattens a CloudConnectionInfoSettings struct into a map.
func flattenCloudConnInfoSettings(settings *packetfabric.CloudConnectionInfoSettings) map[string]interface{} {
	return map[string]interface{}{
		"vlan_id_pf":       settings.VlanIDPF,
		"vlan_id_cust":     settings.VlanIDCust,
		"svlan_id_cust":    settings.SVlanIDCust,
		"aws_region":       settings.AWSRegion,
		"aws_hosted_type":  settings.AWSHostedType,
		"aws_account_id":   settings.AWSAccountID,
		"aws_connection_id": settings.AWSConnectionID,
	}
}

// flattenCloudConnInfoCloudSettings flattens a CloudConnectionInfoCloudSettings struct into a map.
func flattenCloudConnInfoCloudSettings(cloudSettings *packetfabric.CloudConnectionInfoCloudSettings) map[string]interface{} {
	cloudSettingsMap := map[string]interface{}{
		"vlan_id_pf":                  cloudSettings.VlanIDPf,
		"vlan_id_cust":                cloudSettings.VlanIDCust,
		"svlan_id_cust":               cloudSettings.SvlanIDCust,
		"aws_region":                  cloudSettings.AWSRegion,
		"aws_hosted_type":             cloudSettings.AWSHostedType,
		"aws_account_id":              cloudSettings.AWSAccountID,
		"aws_connection_id":           cloudSettings.AWSConnectionID,
		"credentials_uuid":            cloudSettings.CredentialsUUID,
		"mtu":                         cloudSettings.MTU,
		"aws_dx_location":             cloudSettings.AWSDXLocation,
		"aws_dx_bandwidth":            cloudSettings.AWSDXBandwidth,
		"aws_dx_jumbo_frame_capable":  cloudSettings.AWSDXJumboFrameCapable,
		"aws_dx_aws_device":           cloudSettings.AWSDXAWSDevice,
		"aws_dx_aws_deviceV2":         cloudSettings.AWSDXAWSDeviceV2,
		"aws_dx_aws_logical_device_id": cloudSettings.AWSDXAWSLogicalDeviceID,
		"aws_dx_has_logical_redundancy": cloudSettings.AWSDXHasLogicalRedundancy,
		"aws_dx_mac_sec_capable":      cloudSettings.AWSDXMACSecCapable,
		"aws_dx_encryption_mode":      cloudSettings.AWSDXEncryptionMode,
		"aws_vif_type":                cloudSettings.AWSVIFType,
		"aws_vif_id":                  cloudSettings.AWSVIFID,
		"aws_vif_bgp_peer_id":         cloudSettings.AWSVIFBGPPeerID,
		"aws_vif_direct_connect_gw_id": cloudSettings.AWSVIFDirectConnectGWID,
	}

	// Flatten the cloud_state list.
	cloudStateList := make([]interface{}, 0, len(cloudSettings.CloudState))
	for _, cs := range cloudSettings.CloudState {
		cloudStateList = append(cloudStateList, flattenCloudConnInfoCloudState(cs))
	}
	cloudSettingsMap["cloud_state"] = cloudStateList

	// Flatten the bgp_settings list.
	bgpSettingsList := make([]interface{}, 0, len(cloudSettings.BGPSettings))
	for _, bs := range cloudSettings.BGPSettings {
		bgpSettingsList = append(bgpSettingsList, flattenCloudConnInfoBGPSettings(bs))
	}
	cloudSettingsMap["bgp_settings"] = bgpSettingsList

	return cloudSettingsMap
}

// flattenCloudConnInfoCloudState flattens a CloudConnectionInfoCloudState struct into a map.
func flattenCloudConnInfoCloudState(cloudState *packetfabric.CloudConnectionInfoCloudState) map[string]interface{} {
	return map[string]interface{}{
		"aws_dx_connection_state":      cloudState.AWSDXConnectionState,
		"aws_dx_port_encryption_status": cloudState.AWSDXPortEncryptionStatus,
		"aws_vif_state":                cloudState.AWSVIFState,
		"bgp_state":                    cloudState.BGPState,
	}
}

// flattenCloudConnInfoBGPSettings flattens a CloudConnectionInfoBGPSettings struct into a map.
func flattenCloudConnInfoB

