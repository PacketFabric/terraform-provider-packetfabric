package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourcePointToPoints() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourcePointToPointsRead,
		Schema: map[string]*schema.Schema{
			"point_to_points": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ptp_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PTP UUID",
						},
						"ptp_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The PTP Circuit ID.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The PTP description",
						},
						"speed": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The PTP speed.",
						},
						"media": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The PTP media type.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The PTP state.",
						},
						"billing": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_uuid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The billing account UUID.",
									},
									"subscription_term": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The subscription term.",
									},
									"contracted_speed": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The contracted speed.",
									},
								},
							},
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The PTP time created.",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The PTP time updated.",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is PTP deleted.",
						},
						"service_class": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service class for the associated VC of this PTP.",
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
		},
	}
}

func datasourcePointToPointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ptps, err := c.ListPointToPoints()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("point_to_points", flattenPointToPoints(ptps)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenPointToPoints(ptps *[]packetfabric.PointToPointResp) []interface{} {
	if ptps != nil {
		flattens := make([]interface{}, len(*ptps))
		for i, ptp := range *ptps {
			flatten := make(map[string]interface{})
			flatten["ptp_uuid"] = ptp.PtpUUID
			flatten["ptp_circuit_id"] = ptp.PtpCircuitID
			flatten["description"] = ptp.Description
			flatten["speed"] = ptp.Speed
			flatten["media"] = ptp.Media
			flatten["state"] = ptp.State
			flatten["billing"] = flattenBilling(&ptp.Billing)
			flatten["time_created"] = ptp.TimeCreated
			flatten["time_updated"] = ptp.TimeUpdated
			flatten["deleted"] = ptp.Deleted
			flatten["service_class"] = ptp.ServiceClass
			flatten["interfaces"] = flattenPtpInterf(&ptp.Interfaces)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenBilling(billing *packetfabric.Billing) []interface{} {
	flattens := make([]interface{}, 0)
	if billing != nil {
		flatten := make(map[string]interface{})
		flatten["account_uuid"] = billing.AccountUUID
		flatten["subscription_term"] = billing.SubscriptionTerm
		flatten["contracted_speed"] = billing.ContractedSpeed
		flattens = append(flattens, flatten)
	}
	return flattens
}

func flattenPtpInterf(interfs *[]packetfabric.Interfaces) []interface{} {
	if interfs != nil {
		flattens := make([]interface{}, len(*interfs), len(*interfs))
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
