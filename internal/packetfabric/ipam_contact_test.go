package packetfabric

import (
	"encoding/json"
	"testing"
)

func Test_CreateIpamContact(t *testing.T) {
	expectedPayload := IpamContact{}
	expectedResp := IpamContact{}
	jsonPayload := []byte(`{
		"uuid":         "88888f88-8f88-8f8f-888f-8f888f8f8f8f",
		"contact_name": "Jane Smith",
		"org_name":     "Terraform Test Organization Name",
		"address":      "1234 Peachtree St, Atlanta, GA",
		"phone":        "123-456-7890",
		"email":        "jane.smith@test.com",
		"arin_org_id":  "Optional ARIN Organization ID",
		"apnic_org_id": "Optional APNIC Organization ID",
		"ripe_org_id":  "Optional RIPE Organization ID"
	}`)
	jsonResponse := jsonPayload
	if err := json.Unmarshal(jsonPayload, &expectedPayload); err != nil {
		t.Fatalf("Failed to unmarshal payload: %s", err)
	}
	if err := json.Unmarshal(jsonResponse, &expectedResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err)
	}
	cTest.runFakeHttpServer(_callCreateIpamContact, expectedPayload, expectedResp, jsonResponse, "IpamContact-create", t)
}

func _callCreateIpamContact(payload interface{}) (interface{}, error) {
	return cTest.CreateIpamContact(payload.(IpamContact))
}
