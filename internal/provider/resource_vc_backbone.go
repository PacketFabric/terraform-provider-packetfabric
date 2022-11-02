package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcBackbone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVCBackboneCreate,
		UpdateContext: resourceVCBackboneUpdate,
		ReadContext:   resourceVCBackboneRead,
		DeleteContext: resourceVCBackboneDelete,
		Schema:        resourceBackbone(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceVCBackboneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceBackboneCreate(ctx, d, m, c.CreateBackbone)
}

func resourceVCBackboneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceVCBackboneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServiceSettingsUpdate(ctx, d, m)
}

func resourceVCBackboneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceBackboneDelete(ctx, d, m, c.DeleteBackbone)
}
