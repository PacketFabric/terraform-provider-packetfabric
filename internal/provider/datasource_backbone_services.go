package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceBackboneServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: backboneServiceRead,
		Schema: map[string]*schema.Schema{
			"backbone_services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: map[string]*schema.Schema{
					"vc_circuit_id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The VC Circuit ID.",
					},
					"customer_uuid": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The Customer UUID.",
					},
					"state": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The service state.",
					},
					"service_type": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The service type.",
					},
					"service_class": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The service class.",
					},
					"mode": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The service mode.",
					},
					"connected": {
						Type:        schema.TypeBool,
						Computed:    true,
						Description: "Current connection status.",
					},
					"bandwidth": {
						Type:        schema.TypeSet,
						Optional:    true,
						Description: "Backbone service bandwidth",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"account_uuid": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The UUID of the PacketFabric contact that will be billed.\n\t\tExample: a2115890-ed02-4795-a6dd-c485bec3529c",
								},
								"longhaul_type": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Dedicated (no limits or additional charges), usage-based (per transfered GB) pricing model or hourly billing\n\t\tEnum: [\"dedicated\" \"usage\" \"hourly\"]",
								},
								"subscription_term": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Subscription term in months. Not applicable for hourly billing.\n\t\tEnum: [\"1\" \"12\" \"24\" \"36\"]",
								},
								"speed": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The desired speed of the new connection.\n\t\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
								},
							},
						},
					},
					"description": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The service description.",
					},
					"rate_limit_in": {
						Type:        schema.TypeInt,
						Computed:    true,
						Description: "The rate limit in.",
					},
					"rate_limit_out": {
						Type:        schema.TypeInt,
						Computed:    true,
						Description: "The rate limit out.",
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
					"interfaces": {
						Type:     schema.TypeList,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"port_circuit_id": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The port circuit ID.",
								},
								"pop": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface POP.",
								},
								"site": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface site.",
								},
								"site_name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface site name.",
								},
								"customer_site_code": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface customer site code.",
								},
								"customer_site_name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The customer site name.",
								},
								"speed": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface speed.",
								},
								"media": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The media size.",
								},
								"zone": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface zone.",
								},
								"description": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface description.",
								},
								"vlan": {
									Type:        schema.TypeInt,
									Computed:    true,
									Description: "The interface vlan.",
								},
								"untagged": {
									Type:        schema.TypeBool,
									Computed:    true,
									Description: "The interface untagged state.",
								},
								"provisioning_status": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Interface provisioning status.",
								},
								"admin_status": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface admin status.",
								},
								"operational_status": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface operational status.",
								},
								"customer_uuid": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface customer UUID.",
								},
								"customer_name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface customer name.",
								},
								"region": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface region.",
								},
								"is_cloud": {
									Type:        schema.TypeBool,
									Computed:    true,
									Description: "Interface cloud state.",
								},
								"is_ptp": {
									Type:        schema.TypeBool,
									Computed:    true,
									Description: "Interface PTP state.",
								},
								"time_created": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface creation time.",
								},
								"time_updated": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The interface updated time.",
								},
							},
						},
					},
				},
			},
		},
	}
}

func backboneServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCID, ok := d.GetOk("vc_circuit_id"); ok {
		service, err := c.GetBackboneByVcCID(vcCID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		if setErr := d.Set("backbone_services", flattenBackboneService(service)); setErr != nil {
			return diag.FromErr(setErr)
		}
	}
	return diags
}

func flattenBackboneService(service *packetfabric.BackboneResp) []interface{} {
	flattens := make([]interface{}, 0)
	if service != nil {
		flatten := make(map[string]interface{})
		flatten["vc_circuit_id"] = service.VcCircuitID
		flatten["customer_uuid"] = service.CustomerUUID
		flatten["state"] = service.State
		flatten["service_type"] = service.ServiceType
		flatten["service_class"] = service.ServiceClass
		flatten["mode"] = service.Mode
		flatten["connected"] = service.Connected
		flatten["bandwidth"] = flattenBandwidth(&service.Bandwidth)
		flatten["description"] = service.Description
		flatten["rate_limit_in"] = service.RateLimitIn
		flatten["rate_limit_out"] = service.RateLimitOut
		flatten["time_created"] = service.TimeCreated
		flatten["time_updated"] = service.TimeUpdated
		flatten["interfaces"] = flattenBackBoneInterfaces(&service.Interfaces)
	}
	return flattens
}

func flattenBackBoneInterfaces(interfs *[]packetfabric.BackboneInterfResp) []interface{} {
	flattens := make([]interface{}, 0)
	if interfs != nil {
		for i, interf := range *interfs {
			flatten := make(map[string]interface{})
			flatten["port_circuit_id"] = interf.PortCircuitID
			flatten["pop"] = interf.Pop
			flatten["site"] = interf.Site
			flatten["site_name"] = interf.SiteName
			flatten["customer_site_code"] = interf.CustomerSiteCode
			flatten["customer_site_name"] = interf.CustomerSiteName
			flatten["speed"] = interf.Speed
			flatten["media"] = interf.Media
			flatten["zone"] = interf.Zone
			flatten["description"] = interf.Description
			flatten["vlan"] = interf.Vlan
			flatten["untagged"] = interf.Untagged
			flatten["provisioning_status"] = interf.ProvisioningStatus
			flatten["admin_status"] = interf.AdminStatus
			flatten["operational_status"] = interf.OperationalStatus
			flatten["customer_uuid"] = interf.CustomerUUID
			flatten["customer_name"] = interf.CustomerName
			flatten["region"] = interf.Region
			flatten["is_cloud"] = interf.IsCloud
			flatten["is_ptp"] = interf.IsPtp
			flatten["time_created"] = interf.TimeCreated
			flatten["time_updated"] = interf.TimeUpdated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
