package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOutboundCrossConnect() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOutboundCrossConnectRead,
		Schema: map[string]*schema.Schema{
			"outbound_cross_connects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Outbound Cross Connects.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Port.",
						},
						"site": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Site.",
						},
						"document_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Document UUID.",
						},
						"outbound_cross_connect_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Outbound Cross Connect ID.",
						},
						"obcc_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect OBCC Status.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Description.",
						},
						"user_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect User description.",
						},
						"destination_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Destination name.",
						},
						"destination_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Destination CID.",
						},
						"panel": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Panel.",
						},
						"module": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Module.",
						},
						"position": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Position.",
						},
						"data_center_cross_connect_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Data Center Cross Connect ID.",
						},
						"progress": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Outbound Cross Connect Progress.",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The Outbound Cross Connect delete state.",
						},
						"z_loc_cfa": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Panel/module/position.",
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Time created.",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Outbound Cross Connect Time updated.",
						},
					},
				},
			},
		},
	}
}

func dataSourceOutboundCrossConnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	crossConns, err := c.GetOutboundCrossConnects()
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("outbound_cross_connects", flattenOutboundCrossConnects(crossConns)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenOutboundCrossConnects(crossConns *[]packetfabric.OutboundCrossConnectResp) []interface{} {
	if crossConns != nil {
		flattens := make([]interface{}, len(*crossConns), len(*crossConns))
		for i, crossConn := range *crossConns {
			flatten := make(map[string]interface{})
			flatten["port"] = crossConn.Port
			flatten["site"] = crossConn.Site
			flatten["document_uuid"] = crossConn.DocumentUUID
			flatten["outbound_cross_connect_id"] = crossConn.OutboundCrossConnectID
			flatten["obcc_status"] = crossConn.ObccStatus
			flatten["description"] = crossConn.Description
			flatten["user_description"] = crossConn.UserDescription
			flatten["destination_name"] = crossConn.DestinationName
			flatten["destination_circuit_id"] = crossConn.DestinationCircuitID
			flatten["panel"] = crossConn.Panel
			flatten["module"] = crossConn.Module
			flatten["position"] = crossConn.Position
			flatten["data_center_cross_connect_id"] = crossConn.DataCenterCrossConnectID
			flatten["progress"] = crossConn.Progress
			flatten["deleted"] = crossConn.Deleted
			flatten["z_loc_cfa"] = crossConn.ZLocCfa
			flatten["time_created"] = crossConn.TimeCreated
			flatten["time_updated"] = crossConn.TimeUpdated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
