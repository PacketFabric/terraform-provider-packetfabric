package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIpamPrefixes() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIpamPrefixesRead,
		Schema: map[string]*schema.Schema{
			"ipam_prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"circuit_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"org_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"city": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"postal_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"admin_ipam_contact_uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tech_ipam_contact_uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipj_details": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"current_prefixes": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"prefix": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"ips_in_use": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"description": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"isp_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"will_renumber": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
									"planned_prefix": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"description": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"usage_30d": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"usage_3m": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"usage_6m": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"usage_1y": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
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

func datasourceIpamPrefixesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ipamPrefixes, err := c.ReadIpamPrefixes()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ipam_prefixes", flattenIpamPrefixes(&ipamPrefixes))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenIpamPrefixes(ipamPrefixes *[]packetfabric.IpamPrefix) []interface{} {
	if ipamPrefixes != nil {
		flattens := make([]interface{}, len(*ipamPrefixes))
		for i, ipamPrefix := range *ipamPrefixes {
			flatten := make(map[string]interface{})
			flatten["ip_address"] = ipamPrefix.IpAddress
			flatten["circuit_id"] = ipamPrefix.CircuitId
			flatten["type"] = ipamPrefix.Type
			flatten["org_id"] = ipamPrefix.OrgId
			flatten["address"] = ipamPrefix.Address
			flatten["city"] = ipamPrefix.City
			flatten["postal_code"] = ipamPrefix.PostalCode
			flatten["admin_ipam_contact_uuid"] = ipamPrefix.AdminIpamContactUuid
			flatten["tech_ipam_contact_uuid"] = ipamPrefix.TechIpamContactUuid
			flatten["state"] = ipamPrefix.State
			flatten["ipj_details"] = flattenIpjDetails(ipamPrefix.IpjDetails)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
