package packetfabric

import (
	"encoding/json"
	"fmt"
	"testing"
)

const _routingID = "PF-1RI-OQ85"
const _serviceUUID = "7138cc00-c611-4dec-a05e-5c4b1cae13c0"
const _vcRequestUUID = "8f630c69-5fb4-4159-a3cc-a3c182f88ab3"
const _vcCircuitID = "PF-DC-SMF-PDX-12345"
const _customerUUID = "e7eefd45-cb13-4c62-b229-e5bbc1362123"
const _userUUID = "7c4d2d7d-8620-4fb3-967a-4a621082cf1f"
const _portCircuitID = "PF-AP-LAX1-1002"
const _awsServiceConnDesc = "AWS Service Connection"

func Test_CreateAwsHostedMkt(t *testing.T) {
	var payload ServiceAws
	var expectedResp AwsHostedMktResp
	_ = json.Unmarshal(_buildFakeCreateAwsHostedMkt(), &payload)
	_ = json.Unmarshal(_buildFakeCreateAwsHostedMktResp(), &expectedResp)
	cTest.runFakeHttpServer(_callCreateAwsHostedMkt, payload, expectedResp, _buildFakeCreateAwsHostedMktResp(), "test-create-aws-hosted-mkt", t)
}

func Test_CreateAwsProvisionReq(t *testing.T) {
	var payload ServiceAwsMktConn
	var expectedResp MktConnProvisionResp
	_ = json.Unmarshal(_buildFakeCreateAwsProvisionReq(), &payload)
	_ = json.Unmarshal(_buildFakeCreateAwsProvisionReqResp(), &expectedResp)
	cTest.runFakeHttpServer(_callCreateAwsProvisionReq, payload, expectedResp, _buildFakeCreateAwsProvisionReqResp(), "test-create-aws-provision-req", t)
}

func Test_CreateAwsHostedConn(t *testing.T) {
	var payload HostedAwsConnection
	var expectedResp HostedConnectionResp
	_ = json.Unmarshal(_buildFakeCreateAwsHostedConn(), &payload)
	_ = json.Unmarshal(_buildFakeCreateAwsHostedConnResp(), &expectedResp)
	cTest.runFakeHttpServer(_callCreateAwsHostedConn, payload, expectedResp, _buildFakeCreateAwsHostedConnResp(), "test-create-aws-hosted-conn", t)
}

func Test_CreateDedicadedAWSConn(t *testing.T) {
	var payload DedicatedAwsConn
	var expectedResp AwsDedicatedConnCreateResp
	_ = json.Unmarshal(_buildFakeCreateDedicadedAWSConn(), &payload)
	_ = json.Unmarshal(_buildFakeCreateDedicadedAWSConnResp(), &expectedResp)
	cTest.runFakeHttpServer(_callCreateDedicadedAWSConn, payload, expectedResp, _buildFakeCreateDedicadedAWSConnResp(), "test-create-dedicated-aws-conn", t)
}

func Test_UpdateAwsServiceConn(t *testing.T) {
	var expectedResp HostedConnectionResp
	_ = json.Unmarshal(_buildFakeUpdateAwsServiceConnResp(), &expectedResp)
	cTest.runFakeHttpServer(_callUpdateAwsServiceConn, _awsServiceConnDesc, expectedResp, _buildFakeUpdateAwsServiceConnResp(), "test-update-aws-services-conn", t)
}

func Test_GetCloudConnInfo(t *testing.T) {
	var expectedResp AwsCloudConnInfo
	_ = json.Unmarshal(_buildFakeGetCloudConnInfoResp(), &expectedResp)
	cTest.runFakeHttpServer(_callGetCloudConnInfo, nil, expectedResp, _buildFakeGetCloudConnInfoResp(), "test-get-cloud-conn-info", t)
}

func Test_GetCurrentCustomersHosted(t *testing.T) {
	var expectedResp []CloudConnCurrentCustomers
	_ = json.Unmarshal(_buildFakeGetCurrentCustomersHostedResp(), &expectedResp)
	cTest.runFakeHttpServer(_callGetCurrentCustomersHosted, nil, expectedResp, _buildFakeGetCurrentCustomersHostedResp(), "test-get-current-customers-hosted", t)
}

func Test_callGetHostedCloudConnRequestsSent(t *testing.T) {
	var expectedResp []AwsHostedMktResp
	_ = json.Unmarshal(_buildFakeGetHostedCloudConnRequestsSent(), &expectedResp)
	cTest.runFakeHttpServer(_callGetHostedCloudConnRequestsSent, nil, expectedResp, _buildFakeGetHostedCloudConnRequestsSent(), "test-get-hosted-cloud-conn-requests", t)
}

func Test_GetCurrentCustomersDedicated(t *testing.T) {
	var expectedResp []DedicatedConnResp
	_ = json.Unmarshal(_buildFakeGetCurrentCustomersDedicatedResp(), &expectedResp)
	cTest.runFakeHttpServer(_callGetCurrentCustomersDedicated, nil, expectedResp, _buildFakeGetCurrentCustomersDedicatedResp(), "test-get-current-customers-dedicated", t)
}

func Test_GetCloudServiceStatus(t *testing.T) {
	var expectedResp ServiceState
	_ = json.Unmarshal(_buildMockGetCloudServiceStatusResp(), &expectedResp)
	cTest.runFakeHttpServer(_callGetCloudServiceStatus, nil, expectedResp, _buildMockGetCloudServiceStatusResp(), "test-get-cloud-service-status", t)
}

func _callCreateAwsHostedMkt(payload interface{}) (interface{}, error) {
	return cTest.CreateAwsHostedMkt(payload.(ServiceAws))
}

func _callCreateAwsProvisionReq(payload interface{}) (interface{}, error) {
	return cTest.CreateAwsProvisionReq(payload.(ServiceAwsMktConn), _vcRequestUUID)
}

func _callCreateAwsHostedConn(payload interface{}) (interface{}, error) {
	return cTest.CreateAwsHostedConn(payload.(HostedAwsConnection))
}

func _callCreateDedicadedAWSConn(payload interface{}) (interface{}, error) {
	return cTest.CreateDedicadedAWSConn(payload.(DedicatedAwsConn))
}

func _callUpdateAwsServiceConn(payload interface{}) (interface{}, error) {
	return cTest.UpdateServiceConn(_awsServiceConnDesc, _cloudCircuitID)
}

func _callGetCloudConnInfo(payload interface{}) (interface{}, error) {
	return cTest.GetCloudConnInfo(_circuitIdMock)
}

func _callGetCurrentCustomersHosted(payload interface{}) (interface{}, error) {
	return cTest.GetCurrentCustomersHosted()
}

func _callGetHostedCloudConnRequestsSent(payload interface{}) (interface{}, error) {
	return cTest.GetHostedCloudConnRequestsSent()
}

func _callGetCurrentCustomersDedicated(payload interface{}) (interface{}, error) {
	return cTest.GetCurrentCustomersDedicated()
}

func _callGetCloudServiceStatus(payload interface{}) (interface{}, error) {
	return cTest.GetCloudServiceStatus(_cloudCircuitID)
}

func _callDeleteCloudService(payload interface{}) (interface{}, error) {
	return nil, cTest.DeleteCloudService(payload.(string))
}

// PAYLOADS

func _buildFakeCreateAwsHostedMkt() []byte {
	return []byte(fmt.Sprintf(`{
		"routing_id": "%s",
		"market": "ATL",
		"description": "My AWS Marketplace Cloud connection",
		"aws_account_id": "02345678910",
		"account_uuid": "%s",
		"pop": "DAL1",
		"zone": "A",
		"speed": "100Mbps",
		"service_uuid": "%s"
	  }`, _routingID, _serviceUUID, _accountUUID))
}

func _buildFakeCreateAwsProvisionReq() []byte {
	return []byte(fmt.Sprintf(`{
		"provider": "aws",
		"interface": {
		  "port_circuit_id": "%s",
		  "vlan": 6
		},
		"description": "500Mbps connection to Vandelay Industries in SAC1"
	  }`, _portCircuitID))
}

func _buildFakeCreateAwsHostedConn() []byte {
	return []byte(fmt.Sprintf(`{
		"aws_account_id": "%s",
		"account_uuid": "%s",
		"description": "AWS Hosted connection for Foo Corp",
		"pop": "DAL1",
		"port": "%s",
		"vlan": 6,
		"src_svlan": 100,
		"zone": "A",
		"speed": "100Mbps"
	  }`, _awsAccountID, _accountUUID, _portCircuitID))
}

func _buildFakeCreateDedicadedAWSConn() []byte {
	return []byte(fmt.Sprintf(`{
		"aws_region": "us-west-1",
		"account_uuid": "%s",
		"description": "%s",
		"zone": "A",
		"pop": "DAL1",
		"subscription_term": 1,
		"service_class": "longhaul",
		"autoneg": false,
		"speed": "1Gbps",
		"should_create_lag": true,
		"loa": "SSBhbSBhIFBERg=="
	  }`, _awsAccountUUID, _awsServiceConnDesc))
}

func _buildFakeGetHostedCloudConnRequestsSent() []byte {
	return []byte(fmt.Sprintf(`[
		{
		  "vc_request_uuid": "%s",
		  "vc_circuit_id": "%s",
		  "from_customer": {
			"customer_uuid": "%s",
			"name": "Vandelay Industries",
			"contact_first_name": "James",
			"contact_last_name": "Bond",
			"contact_email": "user@example.com",
			"contact_phone": "111-111-1111"
		  },
		  "to_customer": {
			"customer_uuid": "%s",
			"name": "Vandelay Industries"
		  },
		  "status": "pending",
		  "request_type": "marketplace",
		  "text": "Vandelay Industries would like to connect with you in Los Angeles",
		  "bandwidth": {
			"account_uuid": "%s",
			"subscription_term": 1,
			"longhaul_type": "dedicated",
			"speed": "50Mbps"
		  },
		  "rate_limit_in": 1000,
		  "rate_limit_out": 1000,
		  "service_name": "Example Service",
		  "allow_untagged_z": true,
		  "time_created": "2016-01-11T08:30:00+00:00",
		  "time_updated": "2016-01-11T08:30:00+00:00"
		}
	  ]`, _vcRequestUUID, _vcCircuitID, _customerUUID, _customerUUID, _accountUUID))
}

func _buildFakeUpdateAwsServiceConn() []byte {
	return []byte(fmt.Sprintf(`{
		"description": "%s"
	  }`, _awsServiceConnDesc))
}

// RESPONSES

func _buildFakeCreateAwsHostedMktResp() []byte {
	return []byte(fmt.Sprintf(`{
		"vc_request_uuid": "a375494c-5a61-47ce-b727-879a5407eac2",
		"vc_circuit_id": "PF-DC-SMF-PDX-12345",
		"from_customer": {
		  "customer_uuid": "f11ca343-ac2f-4b92-a66b-d4a58d827654",
		  "name": "Vandelay Industries",
		  "contact_first_name": "James",
		  "contact_last_name": "Bond",
		  "contact_email": "user@example.com",
		  "contact_phone": "111-111-1111"
		},
		"to_customer": {
		  "customer_uuid": "f11ca343-ac2f-4b92-a66b-d4a58d827654",
		  "name": "Vandelay Industries"
		},
		"status": "pending",
		"request_type": "marketplace",
		"text": "Vandelay Industries would like to connect with you in Los Angeles",
		"bandwidth": {
		  "account_uuid": "%s",
		  "subscription_term": 1,
		  "longhaul_type": "dedicated",
		  "speed": "50Mbps"
		},
		"rate_limit_in": 1000,
		"rate_limit_out": 1000,
		"service_name": "Example Service",
		"allow_untagged_z": true,
		"time_created": "2016-01-11T08:30:00+00:00",
		"time_updated": "2016-01-11T08:30:00+00:00"
	  }`, _accountUUID))
}

func _buildFakeCreateAwsProvisionReqResp() []byte {
	return []byte(fmt.Sprintf(`{
		"vc_circuit_id": "PF-BC-DA1-DA1-1234567",
		"customer_uuid": "%s",
		"state": "Pending",
		"service_type": "backbone",
		"service_class": "metro",
		"mode": "epl",
		"connected": true,
		"description": "DA1 to DA1 (name)",
		"rate_limit_in": 1000,
		"rate_limit_out": 1000,
		"time_created": "2019-08-24T14:15:22Z",
		"time_updated": "2019-08-24T14:15:22Z",
		"interfaces": [
		  {
			"port_circuit_id": "PF-AP-LAX1-1234",
			"pop": "LAS1",
			"site": "SW-LAS1",
			"site_name": "Switch Las Vegas 7",
			"speed": "1Gbps",
			"media": "LX",
			"zone": "A",
			"description": "User provided description",
			"vlan": 6,
			"untagged": false,
			"provisioning_status": "provisioning",
			"admin_status": "string",
			"operational_status": "string",
			"customer_uuid": "%s",
			"customer_name": "string",
			"region": "US",
			"is_cloud": false,
			"is_ptp": false,
			"time_created": "2020-09-10T14:11:50.075143Z",
			"time_updated": "2020-09-10T14:11:50.075143Z"
		  }
		]
	  }`, _customerUUID, _customerUUID))
}

func _buildFakeCreateAwsHostedConnResp() []byte {
	return []byte(fmt.Sprintf(`{
		"customer_uuid": "%s",
		"user_uuid": "%s",
		"service_provider": "aws",
		"port_type": "hosted",
		"service_class": "longhaul",
		"description": "Hosted-connection-Foo_Corp",
		"state": "requested",
		"speed": "100Mbps"
	  }`, _customerUUID, _userUUID))
}

func _buildFakeCreateDedicadedAWSConnResp() []byte {
	return []byte(fmt.Sprintf(`{
		"customer_uuid": "%s",
		"user_uuid": "%s",
		"service_provider": "aws",
		"port_type": "hosted",
		"service_class": "longhaul",
		"description": "%s.",
		"state": "requested",
		"speed": "1Gbps",
		"cloud_circuit_id": "%s",
		"time_created": "2019-08-24T14:15:22Z",
		"time_updated": "2019-08-24T14:15:22Z"
	  }`, _customerUUID, _userUUID, _cloudConnDesc, _cloudCircuitID))
}

func _buildFakeUpdateAwsServiceConnResp() []byte {
	return []byte(fmt.Sprintf(`{
		"customer_uuid": "%s",
		"user_uuid": "%s",
		"service_provider": "aws",
		"port_type": "hosted",
		"service_class": "longhaul",
		"description": "%s",
		"state": "requested",
		"speed": "100Mbps"
	  }`, _customerUUID, _userUUID, _awsServiceConnDesc))
}

func _buildFakeGetCloudConnInfoResp() []byte {
	return []byte(fmt.Sprintf(`{
		"cloud_circuit_id": "%s",
		"customer_uuid": "%s",
		"user_uuid": "%s",
		"state": "active",
		"service_provider": "aws",
		"service_class": "longhaul",
		"port_type": "hosted",
		"speed": "1Gbps",
		"description": "%s",
		"cloud_provider": {
		  "pop": "LAX1",
		  "region": "us-west-1"
		},
		"time_created": "2019-08-24T14:15:22Z",
		"time_updated": "2019-08-24T14:15:22Z",
		"pop": "LAS1",
		"site": "Switch Las Vegas 7"
	  }`, _cloudCircuitID, _customerUUID, _userUUID, _awsServiceConnDesc))
}

func _buildFakeGetCurrentCustomersHostedResp() []byte {
	return []byte(fmt.Sprintf(`[
		{
		  "is_cloud_router_connection": true,
		  "cloud_circuit_id": "%s",
		  "customer_uuid": "%s",
		  "user_uuid": "%s",
		  "state": "active",
		  "service_provider": "aws",
		  "service_class": "longhaul",
		  "port_type": "hosted",
		  "speed": "1Gbps",
		  "description": "%s",
		  "cloud_provider": {
			"pop": "LAX1",
			"region": "us-west-1"
		  },
		  "time_created": "2019-08-24T14:15:22Z",
		  "time_updated": "2019-08-24T14:15:22Z",
		  "interfaces": [
			{
			  "port_circuit_id": "%s",
			  "pop": "LAS1",
			  "site": "SW-LAS1",
			  "site_name": "Switch Las Vegas 7",
			  "speed": "1Gbps",
			  "media": "LX",
			  "zone": "A",
			  "description": "User provided description",
			  "vlan": 6,
			  "untagged": false,
			  "provisioning_status": "provisioning",
			  "admin_status": "string",
			  "operational_status": "string",
			  "customer_uuid": "%s",
			  "customer_name": "string",
			  "region": "US",
			  "is_cloud": false,
			  "is_ptp": false,
			  "time_created": "2020-09-10T14:11:50.075143Z",
			  "time_updated": "2020-09-10T14:11:50.075143Z"
			}
		  ]
		}
	  ]`, _cloudCircuitID, _customerUUID, _userUUID, _awsServiceConnDesc, _portCircuitID, _customerUUID))
}

func _buildFakeGetCurrentCustomersDedicatedResp() []byte {
	return []byte(fmt.Sprintf(`[
		{
		  "cloud_circuit_id": "%s",
		  "customer_uuid": "%s",
		  "user_uuid": "%s",
		  "state": "active",
		  "service_provider": "aws",
		  "service_class": "longhaul",
		  "port_type": "hosted",
		  "speed": "1Gbps",
		  "description": "%s",
		  "cloud_provider": {
			"pop": "LAX1",
			"region": "us-west-1"
		  },
		  "time_created": "2019-08-24T14:15:22Z",
		  "time_updated": "2019-08-24T14:15:22Z",
		  "pop": "LAS1",
		  "site": "Switch Las Vegas 7"
		}
	  ]`, _cloudCircuitID, _customerUUID, _userUUID, _awsServiceConnDesc))
}

func _buildMockGetCloudServiceStatusResp() []byte {
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
