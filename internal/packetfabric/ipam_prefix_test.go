package packetfabric

import (
	"encoding/json"
	"testing"
)

func Test_CreateIpamPrefix(t *testing.T) {
	expectedPayload := IpamPrefix{}
	expectedResp := IpamPrefix{}
	jsonPayload := []byte(`{
		"prefix_uuid":        "11111a11-1a11-1a1a-111a-1a111a1a1a1a",
		"prefix":             "1.2.3.0/33",
		"state":              "complete",
		"length":             33,
		"version":            4,
		"bgp_region":         "Antarctica",
		"admin_contact_uuid": "11111a11-1a11-1a1a-111a-1a111a1a1a1a",
		"tech_contact_uuid":  "22222b22-2b22-2b2b-222b-2b222b2b2b2b",
		"ipj_details": {
			"currently_used_prefixes": [
				{
					"prefix":        "128.192.1.0/24",
					"ips_in_use":    33,
					"description":   "Optional description",
					"isp_name":      "Optional ISP name",
					"will_renumber": true
				},
				{
					"prefix":        "127.0.127.0/24",
					"ips_in_use":    88
				}
			],
			"planned_prefixes": [
				{
					"prefix":      "8.8.8.0/24",
					"description": "Another optional description",
					"location":    "Optional Location",
					"usage_30d":   2,
					"usage_3m":    0,
					"usage_6m":    2,
					"usage_1y":    3
				},
				{
					"prefix":    "4.4.4.0/24",
					"usage_30d": 2,
					"usage_3m":  0,
					"usage_6m":  2,
					"usage_1y":  3
				}
			]
		}
	}`)
	jsonResponse := jsonPayload
	if err := json.Unmarshal(jsonPayload, &expectedPayload); err != nil {
		t.Fatalf("Failed to unmarshal payload: %s", err)
	}
	if err := json.Unmarshal(jsonResponse, &expectedResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err)
	}
	cTest.runFakeHttpServer(_callCreateIpamPrefix, expectedPayload, expectedResp, jsonResponse, "IpamPrefix-create", t)
}

func _callCreateIpamPrefix(payload interface{}) (interface{}, error) {
	return cTest.CreateIpamPrefix(payload.(IpamPrefix))
}
