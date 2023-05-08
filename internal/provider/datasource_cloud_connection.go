package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionRead,
		Schema: map[string]*schema.Schema{
			"circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The circuit ID of the connection associated with the BGP session. This starts with \"PF-L3-CON-\".",
			},
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
							Type:     schema.TypeString,
							Computed: true,
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
	}
}

func dataSourceCloudConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid Circuit ID")
	}
	connCID, ok := d.GetOk("connection_id")
	if !ok {
		return diag.Errorf("please provide a valid Cloud Router Connection ID")
	}
	awsConn, err := c.ReadCloudRouterConnection(cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	// Flatten the connection information into a map.
	connInfoMap := flattenCloudConnection(awsConn)

	// Update the resource data with the flattened map.
	for k, v := range connInfoMap {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(connCID.(string) + "-data")
	return diags
}

func flattenCloudConnection(conn *packetfabric.CloudRouterConnectionReadResponse) map[string]interface{} {
	connInfoMap := map[string]interface{}{
		"uuid":                         conn.UUID,
		"port_type":                    conn.PortType,
		"connection_type":              conn.ConnectionType,
		"port_circuit_id":              conn.PortCircuitID,
		"pending_delete":               conn.PendingDelete,
		"deleted":                      conn.Deleted,
		"speed":                        conn.Speed,
		"state":                        conn.State,
		"cloud_circuit_id":             conn.CloudCircuitID,
		"account_uuid":                 conn.AccountUUID,
		"service_class":                conn.ServiceClass,
		"service_provider":             conn.ServiceProvider,
		"service_type":                 conn.ServiceType,
		"description":                  conn.Description,
		"cloud_provider_connection_id": conn.CloudProviderConnectionID,
		"user_uuid":                    conn.UserUUID,
		"customer_uuid":                conn.CustomerUUID,
		"time_created":                 conn.TimeCreated,
		"time_updated":                 conn.TimeUpdated,
		"pop":                          conn.Pop,
		"site":                         conn.Site,
		"cloud_router_name":            conn.CloudRouterName,
		"cloud_router_asn":             conn.CloudRouterASN,
		"cloud_router_circuit_id":      conn.CloudRouterCircuitID,
		"nat_capable":                  conn.NatCapable,
		"dnat_capable":                 conn.DNatCapable,
		"zone":                         conn.Zone,
		"vlan":                         conn.Vlan,
	}
	connInfoMap["cloud_settings"] = flattenCloudSettings(&conn.CloudSettings)
	connInfoMap["cloud_provider"] = flattenCloudProvider(&conn.CloudProvider)
	connInfoMap["bgp_state_list"] = flattenBgpStateList(&conn.BgpStateList)

	return connInfoMap
}
