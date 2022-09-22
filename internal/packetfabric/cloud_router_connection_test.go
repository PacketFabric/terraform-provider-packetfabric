package packetfabric

import (
	"encoding/json"
	"fmt"
	"testing"
)

const _circuitIdMock = " PF-L3-CUST-2"
const _cloudCircuitID = "PF-AP-LAX1-1002"
const _cloudConnCid = "PF-CC-LAX-NYC-0192345"
const _cloudConnDesc = "New Super Cool Cloud Router Connection"
const _cloudConnUpdateDesc = "Updated Cloud Router Connection"
const _cloudConnUUID = "095be615-a8ad-4c33-8e9c-c7612fbf6c9f"
const _cloudConnCustomerUUID = "e7eefd45-cb13-4c62-b229-e5bbc1362123"
const _cloudConnUserUUID = "7c4d2d7d-8620-4fb3-967a-4a621082cf1f"
const _cloudConnBillingUUID = "a2115890-ed02-4795-a6dd-c485bec3529c"
const _awsAccountID = "723804547887"
const _accountUUID = "847548f7-9cde-4fe5-8751-32ff19825b7e"

var _clConnectionCreateResp = AwsConnectionCreateResponse{
	PublicIP:        "",
	UUID:            _cloudConnUUID,
	CustomerUUID:    _cloudConnCustomerUUID,
	UserUUID:        _cloudConnUserUUID,
	ServiceProvider: "aws",
	PortType:        "hosted",
	Settings: AwsSettings{
		AwsRegion:       "",
		AwsHostedType:   "",
		AwsConnectionID: "",
		AwsAccountID:    "",
	},
	CloudCircuitID: _cloudCircuitID,
	AccountUUID:    _cloudConnBillingUUID,
	ServiceClass:   "metro",
	Description:    _cloudConnDesc,
	State:          "Requested",
	Billing: AwsBilling{
		AccountUUID:      "",
		SubscriptionTerm: 0,
	},
	Speed: "1Gbps",
	Components: AwsComponents{
		IfdPortCircuitIDCust: "",
		IfdPortCircuitIDPf:   "",
	},
}

var _clConnUpdateExpectedResp = make([]CloudRouterConnectionReadResponse, 0)

func _buildConnUpdateExpectedResp() {
	_clConnUpdateExpectedResp = append(_clConnUpdateExpectedResp, CloudRouterConnectionReadResponse{
		PortType:                  "hosted",
		ConnectionType:            "cloud_hosted",
		PortCircuitID:             "PF-AE-1234",
		PendingDelete:             true,
		Deleted:                   true,
		Speed:                     "1Gbps",
		State:                     "Requested",
		CloudCircuitID:            "PF-AP-LAX1-1002",
		AccountUUID:               _cloudConnBillingUUID,
		ServiceClass:              "metro",
		ServiceProvider:           "aws",
		ServiceType:               "cr_connection",
		Description:               _cloudConnUpdateDesc,
		UUID:                      _cloudConnUUID,
		CloudProviderConnectionID: "dxcon-fgadaaa1",
		UserUUID:                  _cloudConnUserUUID,
		CustomerUUID:              _cloudConnCustomerUUID,
		TimeCreated:               "2019-08-24T14:15:22Z",
		TimeUpdated:               "2019-08-24T14:15:22Z",
		Pop:                       "LAX1",
		Site:                      "us-west-1",
		BgpState:                  "string",
		CloudRouterCircuitID:      "PF-L3-CUST-2001",
		NatCapable:                true,
	})
}

type MessageResp struct {
	Message string `json:"message"`
}

type MockedReadParams struct {
	CircuitID   string
	CloudConnID string
}

func init() {
	_buildConnUpdateExpectedResp()
}

func Test_CreateAwsConn(t *testing.T) {
	var payload AwsConnection
	var expectedResp AwsConnectionCreateResponse
	_ = json.Unmarshal(_buildMockCloudRouterConnectionCreate(), &payload)
	_ = json.Unmarshal(_buildMockCloudRouterCreateResp(), &expectedResp)
	cTest.runFakeHttpServer(_callCreateAwsConn, payload, expectedResp, _buildMockCloudRouterCreateResp(), "-test-create-aws-conn", t)
}

func Test_ReadCloudRouterConnection(t *testing.T) {
	readParamsPayload := MockedReadParams{
		CircuitID:   _circuitIdMock,
		CloudConnID: _cloudConnCid,
	}
	cTest.runFakeHttpServer(_callReadAwsConn, readParamsPayload, _clConnUpdateExpectedResp[0], buildMockCloudRouterReadResp(_cloudConnUpdateDesc), "aws-cloud-router-conn-read", t)
}

func Test_UpdateCloudRouterConnection(t *testing.T) {
	var expectedResp CloudRouterConnectionReadResponse
	payload := DescriptionUpdate{
		Description: _cloudConnUpdateDesc,
	}
	_ = json.Unmarshal(_buildMockCloudRouterConnResp(_cloudConnUpdateDesc), &expectedResp)
	cTest.runFakeHttpServer(_callUpdateAwsConn, payload, expectedResp, _buildMockCloudRouterConnResp(_cloudConnUpdateDesc), "aws-cloud-router-conn-update", t)
}

func Test_GetCloudConnectionStatus(t *testing.T) {
	var expectedResp ServiceState
	_ = json.Unmarshal(_buildMockCloudRouterConnStatus(), &expectedResp)
	cTest.runFakeHttpServer(_callGetClouConnectionStatus, nil, expectedResp, _buildMockCloudRouterConnStatus(), "aws-cloud-router-conn-get-status", t)
}

func Test_ListAwsRouterConnections(t *testing.T) {
	var expectedResp []CloudRouterConnectionReadResponse
	_ = json.Unmarshal(_buildMockCloudRouterConnResps(), &expectedResp)
	cTest.runFakeHttpServer(_callListAwsRouterConnections, nil, expectedResp, _buildMockCloudRouterConnResps(), "aws-cloud-router-conns", t)

}

func Test_DeleteAwsConnection(t *testing.T) {
	var expectedResp ConnectionDeleteResp
	_ = json.Unmarshal(_buildConnDeleteResp(), &expectedResp)
	cTest.runFakeHttpServer(_callDeleteAwsConn, nil, expectedResp, _buildConnDeleteResp(), "test-delete-aws-connection", t)
}

func _callCreateAwsConn(payload interface{}) (interface{}, error) {
	return cTest.CreateAwsConnection(payload.(AwsConnection), _circuitIdMock)
}

func _callReadAwsConn(payload interface{}) (interface{}, error) {
	return cTest.ReadAwsConnection(payload.(MockedReadParams).CircuitID, payload.(MockedReadParams).CloudConnID)
}

func _callGetClouConnectionStatus(payload interface{}) (interface{}, error) {
	return cTest.GetCloudConnectionStatus(_circuitIdMock, _cloudConnCid)
}

func _callListAwsRouterConnections(payload interface{}) (interface{}, error) {
	return cTest.ListAwsRouterConnections(_circuitIdMock)
}

func _callUpdateAwsConn(payload interface{}) (interface{}, error) {
	return cTest.UpdateAwsConnection(_circuitIdMock, _cloudConnCid, payload.(DescriptionUpdate))
}

func _callDeleteAwsConn(payload interface{}) (interface{}, error) {
	return cTest.DeleteAwsConnection(_circuitIdMock, _cloudConnCid)
}

func _buildMockCloudRouterConnectionCreate() []byte {
	return []byte(fmt.Sprintf(`{
		"aws_account_id": "%s",
		"account_uuid": "%s",
		"maybe_nat": false,
		"description": "New AWS Cloud Router Connection",
		"pop": "LAX1",
		"zone": "c",
		"is_public": true,
		"speed": "50Mbps"
	  }`, _awsAccountID, _accountUUID))
}

func _buildMockCloudRouterCreateResp() []byte {
	return []byte(fmt.Sprintf(`{
		"uuid": "3342b8c7-34c3-4fa7-b819-009be3f2dcf8",
		"customer_uuid": "%s",
		"user_uuid": "%s",
		"service_provider": "aws",
		"port_type": "hosted",
		"deleted": false,
		"cloud_circuit_id": "%s",
		"account_uuid": "%s",
		"service_class": "metro",
		"description": "New AWS Cloud Router Connection",
		"state": "requested",
		"settings": {
			"vlan_id_pf": 104,
			"vlan_id_cust": null,
			"aws_region": "us-west-1",
			"aws_hosted_type": "hosted-connection",
			"aws_connection_id": "dxlag-fgl9ffux",
			"aws_account_id": "%s",
			"public_ip": "185.161.1.152/31",
			"nat_public_ip": null
		},
		"billing": {
			"account_uuid": "%s",
			"subscription_term": 1,
			"speed": "50Mbps"
		},
		"components": {
			"vc_id": 38243,
			"ifd_port_circuit_id_pf": "PF-AE-LAX1-37056"
		},
		"is_cloud_router_connection": true,
		"speed": "50Mbps"
	}`, _customerUUID, _userUUID, _cloudCircuitID, _accountUUID, _awsAccountID, _accountUUID))
}

func _buildMockCloudRouterConnResp(description string) []byte {
	return []byte(fmt.Sprintf(`{
		"port_type": "hosted",
		"connection_type": "cloud_hosted",
		"port_circuit_id": "PF-AE-1234",
		"pending_delete": true,
		"deleted": true,
		"speed": "1Gbps",
		"state": "Requested",
		"cloud_circuit_id": "PF-AP-LAX1-1002",
		"account_uuid": "a2115890-ed02-4795-a6dd-c485bec3529c",
		"service_class": "metro",
		"service_provider": "aws",
		"service_type": "cr_connection",
		"description": "%s",
		"uuid": "%s",
		"cloud_provider_connection_id": "dxcon-fgadaaa1",
		"cloud_settings": {
		  "aws_region": "",
		  "aws_hosted_type": "",
		  "aws_connection_id": "",
		  "aws_account_id": ""
		},
		"user_uuid": "7c4d2d7d-8620-4fb3-967a-4a621082cf1f",
		"customer_uuid": "e7eefd45-cb13-4c62-b229-e5bbc1362123",
		"time_created": "2019-08-24T14:15:22Z",
		"time_updated": "2019-08-24T14:15:22Z",
		"cloud_provider": {
		  "pop": "",
		  "region": ""
		},
		"pop": "LAX1",
		"site": "us-west-1",
		"bgp_state": "string",
		"cloud_router_circuit_id": "PF-L3-CUST-2001",
		"nat_capable": true
	  }`, description, _cloudConnUUID))
}

func _buildMockCloudRouterConnUpdateResp(description string) []byte {
	return []byte(fmt.Sprintf(`[
		{
		  "port_type": "hosted",
		  "connection_type": "cloud_hosted",
		  "port_circuit_id": "PF-AE-1234",
		  "pending_delete": true,
		  "deleted": true,
		  "speed": "1Gbps",
		  "state": "Requested",
		  "cloud_circuit_id": "PF-AP-LAX1-1002",
		  "account_uuid": "%s",
		  "service_class": "metro",
		  "service_provider": "aws",
		  "service_type": "cr_connection",
		  "description": "%s",
		  "uuid": "%s",
		  "cloud_provider_connection_id": "dxcon-fgadaaa1",
		  "cloud_settings": {},
		  "user_uuid": "%s",
		  "customer_uuid": "%s",
		  "time_created": "%s",
		  "time_updated": "%s",
		  "cloud_provider": {
			"pop": "LAX1",
			"site": "us-west-1"
		  },
		  "pop": "LAX1",
		  "site": "us-west-1",
		  "bgp_state": "string",
		  "cloud_router_circuit_id": "PF-L3-CUST-2001",
		  "nat_capable": true
		}
	  ]`, _awsAccountUUID, description, _cloudConnUUID, _cloudConnUserUUID, _cloudConnCustomerUUID, _createdTime, _updatedTime))
}

func buildMockCloudRouterReadResp(description string) []byte {
	return []byte(`{
		"port_type": "hosted",
		"connection_type": "cloud_hosted",
		"port_circuit_id": "PF-AE-1234",
		"pending_delete": true,
		"deleted": true,
		"speed": "1Gbps",
		"state": "Requested",
		"cloud_circuit_id": "PF-AP-LAX1-1002",
		"account_uuid": "a2115890-ed02-4795-a6dd-c485bec3529c",
		"service_class": "metro",
		"service_provider": "aws",
		"service_type": "cr_connection",
		"description": "Updated Cloud Router Connection",
		"uuid": "095be615-a8ad-4c33-8e9c-c7612fbf6c9f",
		"cloud_provider_connection_id": "dxcon-fgadaaa1",
		"cloud_settings": {
		  "aws_region": "",
		  "aws_hosted_type": "",
		  "aws_connection_id": "",
		  "aws_account_id": ""
		},
		"user_uuid": "7c4d2d7d-8620-4fb3-967a-4a621082cf1f",
		"customer_uuid": "e7eefd45-cb13-4c62-b229-e5bbc1362123",
		"time_created": "2019-08-24T14:15:22Z",
		"time_updated": "2019-08-24T14:15:22Z",
		"cloud_provider": {
		  "pop": "",
		  "region": ""
		},
		"pop": "LAX1",
		"site": "us-west-1",
		"bgp_state": "string",
		"cloud_router_circuit_id": "PF-L3-CUST-2001",
		"nat_capable": true
	  }`)
}

func _buildMockCloudRouterConnStatus() []byte {
	return []byte(fmt.Sprintf(`{
		"circuit_id": "%s",
		"status": {
		  "object": {
			"state": "string",
			"deleted": true
		  },
		  "current": {
			"state": "ACTIVATING",
			"description": "Still activating"
		  },
		  "last_workflow": {
			"name": "string",
			"root": "f8b25186-636c-4f70-846e-7b7e4e31ba72",
			"current": "b6f6f63e-cd8b-43b2-a95a-2087cb7f8e11",
			"state": "COMPLETED",
			"current_name": "COMPLETED",
			"prev_state": "BILLING_ADD:BILLING_ADD_WORKFLOW",
			"failures": [
			  "Error message"
			],
			"is_final": true,
			"progress": {
			  "position": 7,
			  "steps": 7
			},
			"states": [
			  {
				"state": "Some State",
				"description": "State Desc"
			  }
			]
		  }
		}
	  }`, _circuitIdMock))
}

func _buildMockCloudRouterConnResps() []byte {
	return []byte(`[{
		"port_type": "hosted",
		"connection_type": "cloud_hosted",
		"port_circuit_id": "PF-AE-1234",
		"pending_delete": true,
		"deleted": true,
		"speed": "1Gbps",
		"state": "Requested",
		"cloud_circuit_id": "PF-AP-LAX1-1002",
		"account_uuid": "a2115890-ed02-4795-a6dd-c485bec3529c",
		"service_class": "metro",
		"service_provider": "aws",
		"service_type": "cr_connection",
		"description": "Updated Cloud Router Connection",
		"uuid": "095be615-a8ad-4c33-8e9c-c7612fbf6c9f",
		"cloud_provider_connection_id": "dxcon-fgadaaa1",
		"cloud_settings": {
		  "aws_region": "",
		  "aws_hosted_type": "",
		  "aws_connection_id": "",
		  "aws_account_id": ""
		},
		"user_uuid": "7c4d2d7d-8620-4fb3-967a-4a621082cf1f",
		"customer_uuid": "e7eefd45-cb13-4c62-b229-e5bbc1362123",
		"time_created": "2019-08-24T14:15:22Z",
		"time_updated": "2019-08-24T14:15:22Z",
		"cloud_provider": {
		  "pop": "",
		  "region": ""
		},
		"pop": "LAX1",
		"site": "us-west-1",
		"bgp_state": "string",
		"cloud_router_circuit_id": "PF-L3-CUST-2001",
		"nat_capable": true
	  }]`)
}
func _buildConnDeleteResp() []byte {
	return []byte(`{
		"message": "Cloud router connection deleted."
	  }`)
}
