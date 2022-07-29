package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceLinkAggregationGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceLinkAggregationGroupsRead,
		Schema:      interfacesSchema(),
	}
}

func datasourceLinkAggregationGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	interfs, err := c.GetLAGInterfaces(d.Id())
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
