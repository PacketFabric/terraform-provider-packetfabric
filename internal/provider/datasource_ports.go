package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceInterfaces() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceInterfacesRead,
		Schema: map[string]*schema.Schema{
			"interfaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"autoneg": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "TRUE when Interface Autoneg is enabled.",
						},
						"port_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Circuit ID.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port State.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Status.",
						},
						"speed": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Speed.",
						},
						"media": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Media type.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Zone.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Region.",
						},
						"market": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Market.",
						},
						"market_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Market description.",
						},
						"pop": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port POP.",
						},
						"site": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Site.",
						},
						"site_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Site code.",
						},
						"operational_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Operational status.",
						},
						"admin_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Admin status.",
						},
						"mtu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Interface Port MTU.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Description.",
						},
						"vc_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port VC Mode.",
						},
						"is_lag": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "TRUE when Interface Port is LAG.",
						},
						"is_lag_member": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "TRUE when Interface Port is LAG member.",
						},
						"is_cloud": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "TRUE when Interface Port is Cloud.",
						},
						"is_ptp": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "TRUE when Interface Port is Point to Point.",
						},
						"is_nni": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "TRUE when Interface Port is NNI.",
						},
						"lag_interval": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port LAG interval.",
						},
						"member_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Interface Port Member Count.",
						},
						"parent_lag_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Parent LAG Circuit ID.",
						},
						"account_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Account UUID.",
						},
						"subscription_term": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Interface Port Subscription term.",
						},
						"disabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "TRUE when Interface Port is diabled.",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Customer name.",
						},
						"customer_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Customer UUID.",
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port time created.",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port time updated.",
						},
					},
				},
			},
		},
	}
}

func datasourceInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	interfs, err := c.ListPorts()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("interfaces", flattenInterfaces(interfs))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenInterfaces(interfs *[]packetfabric.InterfaceReadResp) []interface{} {
	if interfs != nil {
		flattens := make([]interface{}, len(*interfs), len(*interfs))
		for i, interf := range *interfs {
			flatten := make(map[string]interface{})
			flatten["autoneg"] = interf.Autoneg
			flatten["port_circuit_id"] = interf.PortCircuitID
			flatten["state"] = interf.State
			flatten["status"] = interf.Status
			flatten["speed"] = interf.Speed
			flatten["media"] = interf.Media
			flatten["zone"] = interf.Zone
			flatten["region"] = interf.Region
			flatten["market"] = interf.Market
			flatten["market_description"] = interf.MarketDescription
			flatten["pop"] = interf.Pop
			flatten["site"] = interf.Site
			flatten["site_code"] = interf.SiteCode
			flatten["operational_status"] = interf.OperationalStatus
			flatten["admin_status"] = interf.AdminStatus
			flatten["mtu"] = interf.Mtu
			flatten["description"] = interf.Description
			flatten["vc_mode"] = interf.VcMode
			flatten["is_lag"] = interf.IsLag
			flatten["is_lag_member"] = interf.IsLagMember
			flatten["is_cloud"] = interf.IsCloud
			flatten["is_ptp"] = interf.IsPtp
			flatten["is_nni"] = interf.IsNni
			flatten["lag_interval"] = interf.LagInterval
			flatten["member_count"] = interf.MemberCount
			flatten["parent_lag_circuit_id"] = interf.ParentLagCircuitID
			flatten["account_uuid"] = interf.AccountUUID
			flatten["subscription_term"] = interf.SubscriptionTerm
			flatten["disabled"] = interf.Disabled
			flatten["customer_name"] = interf.CustomerName
			flatten["customer_uuid"] = interf.CustomerUUID
			flatten["time_created"] = interf.TimeCreated
			flatten["time_updated"] = interf.TimeUpdated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
