package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOutboundCrossConnects() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOutboundCrossConnectsRead,
		Schema: map[string]*schema.Schema{
			PfOutboundCrossConnects: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: PfOutboundCrossConnectsDescription2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPort:                     schemaStringComputed(PfPortDescription2),
						PfSite:                     schemaStringComputed(PfSiteDescription5),
						PfDocumentUuid:             schemaStringComputed(PfDocumentUuidDescription),
						PfOutboundCrossConnectId:   schemaStringComputed(PfOutboundCrossConnectIdDescription),
						PfObccStatus:               schemaStringComputed(PfObccStatusDescription),
						PfDescription:              schemaStringComputed(PfOutboundCrossConnectsDescription),
						PfUserDescription:          schemaStringComputed(PfUserDescriptionDescription),
						PfDestinationName:          schemaStringComputed(PfDestinationNameDescription),
						PfDestinationCircuitId:     schemaStringComputed(PfDestinationCircuitIdDescription),
						PfPanel:                    schemaStringComputed(PfPanelDescription),
						PfModule:                   schemaStringComputed(PfModuleDescription),
						PfPosition:                 schemaStringComputed(PfPositionDescription),
						PfDataCenterCrossConnectId: schemaStringComputed(PfDataCenterCrossConnectIdDescription),
						PfCircuitId:                schemaStringComputed(PfCircuitIdDescription4),
						PfProgress:                 schemaIntComputed(PfProgressDescription),
						PfDeleted:                  schemaBoolComputed(PfDeletedDescription4),
						PfZLocCfa:                  schemaStringComputed(PfZLocCfaDescription),
						PfTimeCreated:              schemaStringComputed(PfTimeCreatedDescription7),
						PfTimeUpdated:              schemaStringComputed(PfTimeUpdatedDescription7),
					},
				},
			},
		},
	}
}

func dataSourceOutboundCrossConnectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	crossConns, err := c.ListOutboundCrossConnects()
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set(PfOutboundCrossConnects, flattenOutboundCrossConnects(crossConns)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenOutboundCrossConnects(crossConns *[]packetfabric.OutboundCrossConnectResp) []interface{} {
	fields := stringsToMap(PfPort, PfSite, PfDocumentUuid, PfOutboundCrossConnectId, PfObccStatus, PfDescription, PfUserDescription, PfDestinationName, PfDestinationCircuitId, PfPanel, PfModule, PfPosition, PfDataCenterCrossConnectId, PfProgress, PfDeleted, PfZLocCfa, PfCircuitId, PfTimeCreated, PfTimeUpdated)
	if crossConns != nil {
		flattens := make([]interface{}, len(*crossConns))
		for i, crossConn := range *crossConns {
			flattens[i] = structToMap(&crossConn, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}
