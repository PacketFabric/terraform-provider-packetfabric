package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudConnections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionsRead,
		Schema: map[string]*schema.Schema{
			"circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"cloud_connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The port type for the given port\n\t\t Enum: hosted, dedicated ",
						},
						"connection_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of the connection.\n\t\t Enum: cloud_hosted, cloud_dedicated, ipsec, packetfabric",
						},
						"port_circuit_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The circuit ID of the port to connect to the cloud router.\n\t\t Exampl \"PF-AE-1234\"",
						},
						"pending_delete": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not the connection is currently deleting.",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not the connection has been fully deleted.",
						},
						"speed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The speed of the connection.\n\t\tEnum: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps",
						},
						"state": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The state of the connection\n\t\tEnum: Requested, Active, Inactive, PendingDelete",
						},
						"cloud_circuit_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The unique PF circuit ID for this connection.\n\t\tExample: \"PF-AP-LAX1-1002\"",
						},
						"account_uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UUID of the PacketFabric contact that will be billed.\n\t\tExample: a2115890-ed02-4795-a6dd-c485bec3529c",
						},
						"service_class": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The service class of the connection.\n\t\tEnum: metro, longhaul",
						},
						"service_provider": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The service provider of the connection.\n\t\tEnum: aws, azure, packet, google, ibm, salesforce, webex",
						},
						"service_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of connection, this will currently always be cr_connection.\n\t\tEnum: cr_connection",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The description of this connection.",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UUID of the connection.",
						},
						"cloud_provider_connection_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud provider specific connection ID, eg. the Amazon connection ID of the cloud router connection.\n\t\tExample: dxcon-fgadaaa1",
						},
						"cloud_settings": {
							Type:     schema.TypeSet,
							Optional: true,
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
									"aws_connection_id": {
										Type:       schema.TypeString,
										Computed:   true,
										Deprecated: "This field is deprecated and will be removed in a future release.",
									},
									"aws_account_id": {
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
									"public_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"nat_public_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"user_uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UUID for the user this connection belongs to",
						},
						"customer_uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UUID for the customer this connection belongs to",
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
						"cloud_provider": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pop": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Point of Presence for the cloud provider location\n\t\tExample: LAX1",
									},
									"site": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Region short name\n\t\tExample: us-west-1",
									},
								},
							},
						},

						"pop": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Point of Presence for the cloud provider location\n\t\tExample: LAX1",
						},
						"site": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Region short name\n\t\tExample: us-west-1",
						},
						"bgp_state_list": {
							Type:        schema.TypeSet,
							Computed:    true,
							Optional:    true,
							Description: "A list of bgp sessions attached to the connection and their states.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bgp_settings_uuid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The UUID of the BGP Session",
									},
									"bgp_state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the BGP session\n\t\tEnum: established, configuring, fetching, etc.",
									},
								},
							},
						},
						"cloud_router_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the cloud router this connection is associated with.\n\t\tExample: Sample CR",
						},
						"cloud_router_asn": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The asn of the cloud router this connection is associated with.\n\t\tExample: 4556",
						},
						"cloud_router_circuit_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The circuit ID of the cloud router this connection is associated with.\n\t\tExample: PF-L3-CUST-2001",
						},
						"nat_capable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether this connection supports NAT",
						},
						"dnat_capable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether this connection supports DNAT",
						},
						"zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud router connection zone",
						},
						"vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The connection vlan for dedicated connections",
						},
						"desired_nat": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates the user's choice of NAT type",
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudConnectionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid Circuit ID")
	}
	awsConns, err := c.ListAwsRouterConnections(cID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("cloud_connections", flattenCloudConnnections(&awsConns))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(cID.(string) + "-data")
	return diags
}

func flattenCloudConnnections(conns *[]packetfabric.CloudRouterConnectionReadResponse) []interface{} {
	if conns != nil {
		flattens := make([]interface{}, len(*conns), len(*conns))
		for i, conn := range *conns {
			flatten := make(map[string]interface{})
			flatten["uuid"] = conn.UUID
			flatten["port_type"] = conn.PortType
			flatten["connection_type"] = conn.ConnectionType
			flatten["port_circuit_id"] = conn.PortCircuitID
			flatten["pending_delete"] = conn.PendingDelete
			flatten["deleted"] = conn.Deleted
			flatten["speed"] = conn.Speed
			flatten["state"] = conn.State
			flatten["cloud_circuit_id"] = conn.CloudCircuitID
			flatten["account_uuid"] = conn.AccountUUID
			flatten["service_class"] = conn.ServiceClass
			flatten["service_provider"] = conn.ServiceProvider
			flatten["service_type"] = conn.ServiceType
			flatten["description"] = conn.Description
			flatten["cloud_provider_connection_id"] = conn.CloudProviderConnectionID
			flatten["user_uuid"] = conn.UserUUID
			flatten["customer_uuid"] = conn.CustomerUUID
			flatten["time_created"] = conn.TimeCreated
			flatten["time_updated"] = conn.TimeUpdated
			flatten["cloud_settings"] = flattenCloudSettings(&conn.CloudSettings)
			flatten["cloud_provider"] = flattenCloudProvider(&conn.CloudProvider)
			flatten["pop"] = conn.Pop
			flatten["site"] = conn.Site
			flatten["bgp_state_list"] = flattenBgpStateList(&conn.BgpStateList)
			flatten["cloud_router_name"] = conn.CloudRouterName
			flatten["cloud_router_asn"] = conn.CloudRouterASN
			flatten["cloud_router_circuit_id"] = conn.CloudRouterCircuitID
			flatten["nat_capable"] = conn.NatCapable
			flatten["dnat_capable"] = conn.DNatCapable
			flatten["zone"] = conn.Zone
			flatten["vlan"] = conn.Vlan
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenCloudProvider(provider *packetfabric.CloudProvider) []interface{} {
	flattens := make([]interface{}, 0)
	if provider != nil {
		flatten := make(map[string]interface{})
		flatten["pop"] = provider.Pop
		flatten["site"] = provider.Site
		flattens = append(flattens, flatten)
	}
	return flattens
}

func flattenBgpStateList(BgpStateList *[]packetfabric.BgpStateObj) []interface{} {
	if BgpStateList != nil {
		flattens := make([]interface{}, len(*BgpStateList), len(*BgpStateList))

		for i, bgpStateObj := range *BgpStateList {
			flatten := make(map[string]interface{})
			flatten["bgp_settings_uuid"] = bgpStateObj.BgpSettingsUUID
			flatten["bgp_state"] = bgpStateObj.BgpState
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenCloudSettings(setts *packetfabric.CloudSettings) []interface{} {
	flattens := make([]interface{}, 0)
	if setts != nil {
		flatten := make(map[string]interface{})
		flatten["account_id"] = setts.AccountID
		flatten["aws_account_id"] = setts.AwsAccountID
		flatten["aws_connection_id"] = setts.AwsConnectionID
		flatten["aws_hosted_type"] = setts.AwsHostedType
		flatten["aws_region"] = setts.AwsRegion
		flatten["azure_connection_type"] = setts.AzureConnectionType
		flatten["azure_service_key"] = setts.AzureServiceKey
		flatten["azure_service_tag"] = setts.AzureServiceTag
		flatten["bgp_asn"] = setts.BgpAsn
		flatten["bgp_cer_cidr"] = setts.BgpCerCidr
		flatten["bgp_ibm_cidr"] = setts.BgpIbmCidr
		flatten["gateway_id"] = setts.GatewayID
		flatten["google_pairing_key"] = setts.GooglePairingKey
		flatten["google_vlan_attachment_name"] = setts.GoogleVlanAttachmentName
		flatten["name"] = setts.Name
		flatten["nat_public_ip"] = setts.NatPublicIP
		flatten["oracle_region"] = setts.OracleRegion
		flatten["port_compartment_ocid"] = setts.PortCompartmentOcid
		flatten["port_cross_connect_ocid"] = setts.PortCrossConnectOcid
		flatten["port_id"] = setts.PortID
		flatten["public_ip"] = setts.PublicIP
		flatten["svlan_id_cust"] = setts.SvlanIDCust
		flatten["vc_ocid"] = setts.VcOcid
		flatten["vlan_id_cust"] = setts.VlanIDCust
		flatten["vlan_id_microsoft"] = setts.VlanMicrosoft
		flatten["vlan_id_pf"] = setts.VlanIDPf
		flatten["vlan_id_private"] = setts.VlanPrivate
		flattens = append(flattens, flatten)
	}
	return flattens
}
