package packetfabric

import (
	"encoding/json"
	"fmt"
	"testing"
)

const _backboneDesc = "Packet Fabric AWS Backbone"
const _portCircuitIDOne = "PF-AP-SAC1-1000"
const _portCircuitIDTwo = "PF-AP-LAS1-2000"

func Test_CreateBackbone(t *testing.T) {
	expectedPayload := Backbone{}
	expectedResp := BackboneResp{}
	_ = json.Unmarshal(_buildCreateBackbonePayload(_backboneDesc), &expectedPayload)
	_ = json.Unmarshal(_buildCreateBackBoneResp(_backboneDesc, _awsAccountID), &expectedResp)
	cTest.runFakeHttpServer(_callCreateBackbone, expectedPayload, expectedResp, _buildCreateBackBoneResp(_backboneDesc, _awsAccountID), "backbone-create", t)
}

func _callCreateBackbone(payload interface{}) (interface{}, error) {
	return cTest.CreateBackbone(payload.(Backbone))
}

func _buildCreateBackbonePayload(description string) []byte {
	return []byte(fmt.Sprintf(`{
		"description": "%s",
		"bandwidth": {
		  "account_uuid": "%s",
		  "subscription_term": 1,
		  "longhaul_type": "dedicated",
		  "speed": "50Mbps"
		},
		"interfaces": [
		  {
			"port_circuit_id": "%s",
			"vlan": 6,
			"untagged": false
		  },
		  {
			"port_circuit_id": "%s",
			"vlan": 7,
			"untagged": false
		  }
		],
		"rate_limit_in": 1000,
		"rate_limit_out": 1000,
		"epl": false
	  }`, description, _accountUUID, _portCircuitIDOne, _portCircuitIDTwo))
}

func _buildCreateBackBoneResp(description, accountID string) []byte {
	return []byte(fmt.Sprintf(`{
		"description": "%s",
		"bandwidth": {
		  "account_uuid": "%s",
		  "subscription_term": 1,
		  "longhaul_type": "dedicated",
		  "speed": "50Mbps"
		},
		"interfaces": [
		  {
			"port_circuit_id": "%s",
			"vlan": 6,
			"untagged": false
		  },
		  {
			"port_circuit_id": "%s",
			"vlan": 7,
			"untagged": false
		  }
		],
		"rate_limit_in": 1000,
		"rate_limit_out": 1000,
		"epl": false
	  }`, description, _accountUUID, _portCircuitIDOne, _portCircuitIDTwo))
}
