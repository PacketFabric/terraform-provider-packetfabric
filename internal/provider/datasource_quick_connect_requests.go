package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceQuickConnectRequests() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceQuickConnectRequestsRead,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"sent", "received"}, true),
				Description:  "The VC request type. (sent/received)",
			},
			"quick_connect_requests": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of quick connect requests.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"import_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Circuit ID of this Cloud Router Import.",
						},
						"cloud_router_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Circuit ID of the source Cloud Router",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The customer that initiated this Cloud Router Import Request.",
						},
						"service_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service UUID of the third-party service associated with the Cloud Router.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shows the state of this import.",
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time created.",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time of last update.",
						},
						"request_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the Cloud Router Import Request.",
						},
						"import_filters": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The import filters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The prefix of the Import Filter.",
									},
									"match_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The match type of the Import Filter.",
									},
									"local_preference": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The localpref of the Import Filter.",
									},
								},
							},
						},
						"return_filters": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The return filters",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The prefix of the Return Filter.",
									},
									"match_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Quick Connect prefix match type.",
									},
									"as_prepend": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The Quick Connect prefix as prepend.",
									},
									"med": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The Quick Connect prefix med.",
									},
								},
							},
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func datasourceQuickConnectRequestsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	reqType := d.Get("type").(string)
	reqs, err := c.GetCloudRouterRequests(reqType)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("quick_connect_requests", flattenQuickConnectRequests(&reqs)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags

}

func flattenQuickConnectRequests(requests *[]packetfabric.CloudRouterRequest) []interface{} {
	if requests != nil {
		flattens := make([]interface{}, len(*requests))
		for i, request := range *requests {
			flatten := make(map[string]interface{})
			flatten["import_circuit_id"] = request.ImportCircuitID
			flatten["cloud_router_circuit_id"] = request.CloudRouterCircuitID
			flatten["customer_name"] = request.CustomerName
			flatten["service_uuid"] = request.ServiceUUID
			flatten["state"] = request.State
			flatten["time_created"] = request.TimeCreated
			flatten["time_updated"] = request.TimeUpdated
			flatten["request_type"] = request.RequestType
			flatten["import_filters"] = flattenImportFilters(request.ImportFilters)
			flatten["return_filters"] = flattenReturnFilters(request.ReturnFilters)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenImportFilters(filters []packetfabric.ImportFilters) []interface{} {
	flattens := make([]interface{}, len(filters))
	for i, filter := range filters {
		flatten := make(map[string]interface{})
		flatten["prefix"] = filter.Prefix
		flatten["match_type"] = filter.MatchType
		flatten["local_preference"] = filter.Localpref
		flattens[i] = flatten
	}
	return flattens
}

func flattenReturnFilters(filters []packetfabric.ReturnFilters) []interface{} {
	flattens := make([]interface{}, len(filters))
	for i, filter := range filters {
		flatten := make(map[string]interface{})
		flatten["prefix"] = filter.Prefix
		flatten["match_type"] = filter.MatchType
		flatten["as_prepend"] = filter.Asprepend
		flatten["med"] = filter.Med
		flattens[i] = flatten
	}
	return flattens
}
