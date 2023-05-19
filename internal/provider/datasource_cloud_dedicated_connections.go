package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceDedicatedCloudConnections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDedicatedConRead,
		Schema: map[string]*schema.Schema{
			"dedicated_connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique PF circuit ID for this connection",
						},
						"customer_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UUID for the customer this connection belongs to",
						},
						"user_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UUID for the user this connection belongs to",
						},
						"service_provider": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service provider of the connection\n\t\tEnum: [\"aws\" \"azure\" \"packet\" \"google\" \"ibm\" \"salesforce\" \"webex\"]",
						},
						"port_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The port type for the given port\n\t\tEnum: [\"hosted\" \"dedicated\"]",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is connection deleted?",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time connection was last updated",
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time of connection creation",
						},
						"cloud_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique PF circuit ID for this connection",
						},
						"account_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PacketFabric account UUID. The contact that will be billed.",
						},
						"cloud_provider": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pop": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The desired location for the new AWS Hosted Connection.\n\t\tExample: DAL1",
									},
									"site": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Site name",
									},
								},
							},
						},
						"pop": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The desired location for the new AWS Hosted Connection.\n\t\tExample: DAL1",
						},
						"site": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Site name",
						},
						"service_class": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service class for the given port, either long haul or metro.\n\t\tEnum: [\"longhaul\",\"metro\"]",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of this connection.\n\t\tExample: AWS Hosted connection for Foo Corp",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the connection\n\t\tEnum: [\"active\" \"deleting\" \"inactive\" \"pending\" \"requested\"]",
						},
						"settings_aws_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region short name",
						},
						"settings": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aws_region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region that the new connection will connect to.\n\t\tExample: us-west-1",
									},
									"zone_dest": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The desired AWS Availability zone of the new connection.\n\t\tExample: \"A\"",
									},
									"autoneg": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the port auto-negotiates or not, this is currently only possible with 1Gbps ports and the request will fail if specified with 10Gbps.",
									},
								},
							},
						},
						"is_cloud_router_connection": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not this is a Cloud Router hosted connection.",
						},
						"speed": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The desired speed of the new connection.\n\t\tEnum: []\"1gps\", \"10gbps\"]",
						},
					},
				},
			},
		},
	}
}

// https://docs.packetfabric.com/api/v2/redoc/#operation/get_connections_dedicated_list
func dataSourceDedicatedConRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	sessions, err := c.GetCurrentCustomersDedicated()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("dedicated_connections", flattenDedicatedConns(&sessions))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenDedicatedConns(conns *[]packetfabric.DedicatedConnResp) []interface{} {
	if conns != nil {
		flattens := make([]interface{}, len(*conns), len(*conns))
		for i, conn := range *conns {
			flatten := make(map[string]interface{})
			flatten["uuid"] = conn.UUID
			flatten["port_type"] = conn.PortType
			flatten["deleted"] = conn.Deleted
			flatten["speed"] = conn.Speed
			flatten["state"] = conn.State
			flatten["cloud_circuit_id"] = conn.CloudCircuitID
			flatten["account_uuid"] = conn.AccountUUID
			flatten["cloud_provider"] = flattenCloudServiceProvider(&conn.CloudProvider)
			flatten["service_class"] = conn.ServiceClass
			flatten["service_provider"] = conn.ServiceProvider
			flatten["description"] = conn.Description
			flatten["user_uuid"] = conn.UserUUID
			flatten["customer_uuid"] = conn.CustomerUUID
			flatten["time_created"] = conn.TimeCreated
			flatten["time_updated"] = conn.TimeUpdated
			flatten["settings"] = flattenCloudServiceSettings(&conn.Settings)
			flatten["is_cloud_router_connection"] = conn.IsCloudRouterConnection
			flatten["pop"] = conn.Pop
			flatten["site"] = conn.Site
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenCloudServiceProvider(provider *packetfabric.CloudServiceProvider) []interface{} {
	flattens := make([]interface{}, 0)
	if provider != nil {
		flatten := make(map[string]interface{})
		flatten["pop"] = provider.Pop
		flatten["site"] = provider.Site
		flattens = append(flattens, flatten)
	}
	return flattens
}

func flattenCloudServiceSettings(settings *packetfabric.CloudServiceSettings) []interface{} {
	flattens := make([]interface{}, 0)
	if settings != nil {
		flatten := make(map[string]interface{})
		flatten["aws_region"] = settings.AwsRegion
		flatten["zone_dest"] = settings.ZoneDest
		flatten["autoneg"] = settings.Autoneg
		flattens = append(flattens, flatten)
	}
	return flattens
}
