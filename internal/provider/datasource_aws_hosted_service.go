package provider

import (
	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
)

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
