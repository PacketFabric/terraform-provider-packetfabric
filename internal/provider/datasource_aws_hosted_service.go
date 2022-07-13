package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceProvisionRequested() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHostedServiceRead,
		Schema: map[string]*schema.Schema{
			"hosted_service_requests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vc_request_uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UUID of the service request",
						},
						"from_customer": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"customer_uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The UUID for the customer associated with this Virtual Circuit",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customer Name",
									},
									"market": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The market that the VC will be requested in.\n\t\tExample: ATL",
									},
									"market_description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the AWS Marketplace Cloud connection.\n\t\tExample: My AWS Marketplace Cloud connection",
									},
									"contact_first_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customer contact first name",
									},
									"contact_last_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customer contact last name",
									},
									"contact_email": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customer contact email",
									},
									"contact_phone": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customer contact phone",
									},
								},
							},
						},
						"to_customer": {
							Type: schema.TypeSet,

							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"customer_uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The UUID for the customer this connection belongs to",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customer Name",
									},
									"market": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The market that the VC will be requested in",
									},
									"market_description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the AWS Marketplace Cloud connection.\n\t\tExample: My AWS Marketplace Cloud connection",
									},
								},
							},
						},
						"text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "General purpose text field",
						},
						"status": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Provisioning status of the port.\n\t\tEnum: [\"Requested\" \"Active\"]",
						},
						"vc_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Mode of operation of the point to point virtual circuit/\n\t\tEnum: [\"epl\" \"evpl\" \"evpl-untagged\" \"mixed\"]",
						},
						"request_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of service being requested.\n\t\tEnum: [\"internet_exchange\" \"legacy_azure\" \"marketplace\" \"marketplace_cloud_aws\" \"marketplace_cloud_azure\" \"marketplace_cloud_google\"]",
						},
						"bandwidth": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The UUID of the PacketFabric contact that will be billed.\n\t\tExample: a2115890-ed02-4795-a6dd-c485bec3529c",
									},
									"longhaul_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Dedicated (no limits or additional charges), usage-based (per transfered GB) pricing model or hourly billing\n\t\tEnum: [\"dedicated\" \"usage\" \"hourly\"]",
									},
									"subscription_term": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Subscription term in months. Not applicable for hourly billing.\n\t\tEnum: [\"1\" \"12\" \"24\" \"36\"]",
									},
									"speed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The desired speed of the new connection.\n\t\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
									},
								},
							},
						},
						"time_created": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date and time of connection creation",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date and time connection was last updated",
						},
						"allow_untagged_z": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, the accepting customer can choose to make this VC untagged. This will only be False if there is only one logical interface on the requesting customer's port and that single logical interface is untagged.",
						},
					},
				},
			},
		},
	}
}

func dataSourceHostedServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	services, err := c.GetHostedCloudConnRequestsSent()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("hosted_service_requests", flattenHostedServiceRequests(&services)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenHostedServiceRequests(services *[]packetfabric.AwsHostedMktResp) []interface{} {
	if services != nil {
		flattens := make([]interface{}, len(*services), len(*services))
		for i, service := range *services {
			flatten := make(map[string]interface{})
			flatten["vc_request_uuid"] = service.VcRequestUUID
			flatten["from_customer"] = flattenFromCustomer(&service.FromCustomer)
			flatten["to_customer"] = flattenToCustomer(&service.ToCustomer)
			flatten["text"] = service.Text
			flatten["status"] = service.Status
			flatten["vc_mode"] = service.VcMode
			flatten["request_type"] = service.RequestType
			flatten["bandwidth"] = flattenBandwidth(&service.Bandwidth)
			flatten["time_created"] = service.TimeCreated
			flatten["time_updated"] = service.TimeUpdated
			flatten["allow_untagged_z"] = service.AllowUntaggedZ
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
