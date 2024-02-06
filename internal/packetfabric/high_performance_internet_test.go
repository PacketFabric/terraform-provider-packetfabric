package packetfabric

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_CreateHighPerformanceInternet(t *testing.T) {
    expectedPayload := HighPerformanceInternet{}
    expectedResp := HighPerformanceInternet{}
    jsonPayload := []byte(`{
		"account_uuid": "12345678-1234-1234-1234-123456789012",
		"circuit_id": "PF-L3-TRAN-12345",
		"description": "HPI for customer A",
		"market": "NYC",
		"port_circuit_id": "PF-AP-12345",
		"routing_type": "bgp",
		"speed": "1Gbps",
		"state": "active",
		"vlan": 4,
		"routing_configuration": {
			"static_routing_v4": {
				"l3_address": "string",
				"remote_address": "string",
				"static_routes": [
					{
						"prefix": "string"
					}
				]
			},
			"static_routing_v6": {
				"l3_address": "string",
				"remote_address": "string",
				"static_routes": [
					{
						"prefix": "string"
					}
				]
			},
			"bgp_v4": {
				"asn": 65000,
				"l3_address": "string",
				"remote_address": "string",
				"md5": "string",
				"prefixes": [
					{
						"prefix": "string",
						"local_preference": 100
					}
				]
			},
			"bgp_v6": {
				"asn": 65000,
				"l3_address": "string",
				"remote_address": "string",
				"md5": "bgp_v6_md5",
				"prefixes": [
					{
						"prefix": "string",
						"local_preference": 100
					},
					{
						"prefix": "string",
						"local_preference": 200
					}
				]
			}
		}
    }`)
    jsonResponse := jsonPayload
    if err := json.Unmarshal(jsonPayload, &expectedPayload); err != nil {
        t.Fatalf("Failed to unmarshal payload: %s", err)
    }
    if err := json.Unmarshal(jsonResponse, &expectedResp); err != nil {
        t.Fatalf("Failed to unmarshal response: %s", err)
    }

	jsonData, err := json.Marshal(expectedResp)
    if err == nil {
		s := string(jsonData)
		fmt.Printf("testing:%s\n", s)
    }
}
