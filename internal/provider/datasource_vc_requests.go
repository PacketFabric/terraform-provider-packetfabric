package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceVcRequests() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceVCRequests,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"sent", "received"}, true),
				Description:  "The VC request type. (sent/received)",
			},
			"vc_requests": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of VC Requests.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vc_request_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VC Request UUID.",
						},
						"vc_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VC Circuit ID.",
						},
						"from_customer": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"customer_uuid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The UUID for the customer associated with this Virtual Circuit",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Customer Name",
									},
									"market": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The market that the VC will be requested in.\n\t\tExample: ATL",
									},
									"market_description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the AWS Marketplace Cloud connection.\n\t\tExample: My AWS Marketplace Cloud connection",
									},
									"contact_first_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Customer contact first name",
									},
									"contact_last_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Customer contact last name",
									},
									"contact_email": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Customer contact email",
									},
									"contact_phone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Customer contact phone",
									},
								},
							},
						},
						"to_customer": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"customer_uuid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The UUID for the customer this connection belongs to",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Customer Name",
									},
									"market": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The market that the VC will be requested in",
									},
									"market_description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the AWS Marketplace Cloud connection.\n\t\tExample: My AWS Marketplace Cloud connection",
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The request status.",
						},
						"request_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The request type.",
						},
						"text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vc request text.",
						},
						"bandwidth": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_uuid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The UUID of the PacketFabric contact that will be billed.\n\t\tExample: a2115890-ed02-4795-a6dd-c485bec3529c",
									},
									"longhaul_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Dedicated (no limits or additional charges), usage-based (per transfered GB) pricing model or hourly billing\n\t\tEnum: [\"dedicated\" \"usage\" \"hourly\"]",
									},
									"subscription_term": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Subscription term in months. Not applicable for hourly billing.\n\t\tEnum: [\"1\" \"12\" \"24\" \"36\"]",
									},
									"speed": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The desired speed of the new connection.\n\t\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
									},
								},
							},
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
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VC Request service name.",
						},
						"allow_untagged_z": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If true, the accepting customer can choose to make this VC untagged. This will only be False if there is only one logical interface on the requesting customer's port and that single logical interface is untagged.",
						},
						"flex_bandwidth_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The flex bandwidth ID.",
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time of connection creation",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time connection was last updated",
						},
					},
				},
			},
		},
	}
}

func datasourceVCRequests(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	reqType := d.Get("type").(string)
	requests, err := c.GetVcRequestsByType(reqType)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("vc_requests", flattenVCRequests(&requests)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenVCRequests(requests *[]packetfabric.VcRequest) []interface{} {
	if requests != nil {
		flattens := make([]interface{}, len(*requests), len(*requests))
		for i, request := range *requests {
			flatten := make(map[string]interface{})
			flatten["vc_request_uuid"] = request.VcRequestUUID
			flatten["vc_circuit_id"] = request.VcCircuitID
			flatten["from_customer"] = flattenFromCustomer(&request.FromCustomer)
			flatten["to_customer"] = flattenToCustomer(&request.ToCustomer)
			flatten["status"] = request.Status
			flatten["request_type"] = request.RequestType
			flatten["text"] = request.Text
			flatten["bandwidth"] = flattenBandwidth(&request.Bandwidth)
			flatten["rate_limit_in"] = request.RateLimitIn
			flatten["rate_limit_out"] = request.RateLimitOut
			flatten["service_name"] = request.ServiceName
			flatten["allow_untagged_z"] = request.AllowUntaggedZ
			flatten["flex_bandwidth_id"] = request.FlexBandwidthID
			flatten["time_created"] = request.TimeCreated
			flatten["time_updated"] = request.TimeUpdated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenFromCustomer(fromCust *packetfabric.FromCustomer) []interface{} {
	flattens := make([]interface{}, 0)
	if fromCust != nil {
		flatten := make(map[string]interface{})
		flatten["customer_uuid"] = fromCust.CustomerUUID
		flatten["name"] = fromCust.Name
		flatten["market"] = fromCust.Market
		flatten["market_description"] = fromCust.MarketDescription
		flatten["contact_first_name"] = fromCust.ContactFirstName
		flatten["contact_last_name"] = fromCust.ContactLastName
		flatten["contact_email"] = fromCust.ContactEmail
		flatten["contact_phone"] = fromCust.ContactPhone
		flattens = append(flattens, flatten)
	}
	return flattens
}

func flattenToCustomer(toCust *packetfabric.ToCustomer) []interface{} {
	flattens := make([]interface{}, 0)
	if toCust != nil {
		flatten := make(map[string]interface{})
		flatten["customer_uuid"] = toCust.CustomerUUID
		flatten["name"] = toCust.Name
		flatten["market"] = toCust.Market
		flatten["market_description"] = toCust.MarketDescription
		flattens = append(flattens, flatten)
	}
	return flattens
}

func flattenBandwidth(bandw *packetfabric.Bandwidth) []interface{} {
	flattens := make([]interface{}, 0)
	if bandw != nil {
		flatten := make(map[string]interface{})
		flatten["account_uuid"] = bandw.AccountUUID
		flatten["longhaul_type"] = bandw.LonghaulType
		flatten["subscription_term"] = bandw.SubscriptionTerm
		flatten["speed"] = bandw.Speed
		flattens = append(flattens, flatten)
	}
	return flattens
}
