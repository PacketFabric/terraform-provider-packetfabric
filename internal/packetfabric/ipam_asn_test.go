package packetfabric

import (
	"encoding/json"
	"testing"
)

func Test_CreateIpamAsn(t *testing.T) {
	expectedPayload := IpamAsn{}
	expectedResp := IpamAsn{}
	jsonPayload := []byte(`{
		"AsnByteType": 2
	}`)
	jsonResponse := jsonPayload
	if err := json.Unmarshal(jsonPayload, &expectedPayload); err != nil {
		t.Fatalf("Failed to unmarshal payload: %s", err)
	}
	if err := json.Unmarshal(jsonResponse, &expectedResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err)
	}
	cTest.runFakeHttpServer(_callCreateIpamAsn, expectedPayload, expectedResp, jsonResponse, "IpamAsn-create", t)
}

func _callCreateIpamAsn(payload interface{}) (interface{}, error) {
	return cTest.CreateIpamAsn(payload.(IpamAsn))
}
