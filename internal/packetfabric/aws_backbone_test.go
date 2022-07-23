package packetfabric

import (
	"encoding/json"
	"fmt"
	"testing"
)

const _awsBackboneDesc = "Packet Fabric AWS Backbone"
const _awsAccountUUID = "847548f7-9cde-4fe5-8751-32ff19825b7e"
const _awsPortCircuitIDOne = "PF-AP-SAC1-1000"
const _awsPortCircuitIDTwo = "PF-AP-LAS1-2000"

func Test_CreateBackbone(t *testing.T) {
	expectedPayload := Backbone{}
	expectedResp := BackboneResp{}
	_ = json.Unmarshal(_buildCreateBackbonePayload(_awsBackboneDesc), &expectedPayload)
	_ = json.Unmarshal(_buildCreateBackBoneResp(_awsBackboneDesc, _awsAccountID), &expectedResp)
	cTest.runFakeHttpServer(_callCreateBackbone, expectedPayload, expectedResp, _buildCreateBackBoneResp(_awsBackboneDesc, _awsAccountID), "aws-backbone-create", t)
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
	  }`, description, _awsAccountUUID, _awsPortCircuitIDOne, _awsPortCircuitIDTwo))
}

func _buildCreateBackBoneResp(description, awsAccountID string) []byte {
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
	  }`, description, _awsAccountUUID, _awsPortCircuitIDOne, _awsPortCircuitIDTwo))
}
