package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIpamAsns() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIpamAsnsRead,
		Schema: map[string]*schema.Schema{
			"ipam_asn": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asn_byte_type": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"asn": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIpamAsnsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ipamAsns, err := c.ReadIpamAsns()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ipam_asns", flattenIpamAsns(&ipamAsns))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenIpamAsns(ipamAsns *[]packetfabric.IpamAsn) []interface{} {
	if ipamAsns != nil {
		flattens := make([]interface{}, len(*ipamAsns))
		for i, ipamAsn := range *ipamAsns {
			flatten := make(map[string]interface{})
			flatten["asn_byte_type"] = ipamAsn.AsnByteType
			flatten["asn"] = ipamAsn.Asn
			flatten["time_created"] = ipamAsn.TimeCreated
			flatten["time_updated"] = ipamAsn.TimeUpdated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
