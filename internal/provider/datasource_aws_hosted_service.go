package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


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
