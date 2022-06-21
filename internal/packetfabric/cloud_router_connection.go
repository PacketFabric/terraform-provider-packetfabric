package packetfabric

import (
	"errors"
	"fmt"
)

const awsConnectionURI = "/v2/services/cloud-routers/%s/connections/aws"
const awsConnectionListURI = "/v2/services/cloud-routers/%s/connections"
const awsConnectionByCidURI = "/v2/services/cloud-routers/%s/connections/%s"
const awsConnectionStatusURI = "/v2.1/services/cloud-routers/%s/connections/%s/status"

type AwsConnection struct {
	AwsAccountID string `json:"aws_account_id"`
	AccountUUID  string `json:"account_uuid"`
	MaybeNat     bool   `json:"maybe_nat"`
	Description  string `json:"description"`
	Pop          string `json:"pop"`
	Zone         string `json:"zone"`
	IsPublic     bool   `json:"is_public"`
	Speed        string `json:"speed"`
}

type AwsConnectionCreateResponse struct {
	PublicIP        string        `json:"public_ip"`
	UUID            string        `json:"uuid"`
	CustomerUUID    string        `json:"customer_uuid"`
	UserUUID        string        `json:"user_uuid"`
	ServiceProvider string        `json:"service_provider"`
	PortType        string        `json:"port_type"`
	Settings        AwsSettings   `json:"settings"`
	CloudCircuitID  string        `json:"cloud_circuit_id"`
	AccountUUID     string        `json:"account_uuid"`
	ServiceClass    string        `json:"service_class"`
	Description     string        `json:"description"`
	State           string        `json:"state"`
	Billing         AwsBilling    `json:"billing"`
	Speed           string        `json:"speed"`
	Components      AwsComponents `json:"components"`
}

type AwsConnectionReadResponse struct {
	PortType                  string           `json:"port_type"`
	PortCircuitID             string           `json:"port_circuit_id"`
	PendingDelete             bool             `json:"pending_delete"`
	State                     string           `json:"state"`
	CloudCircuitID            string           `json:"cloud_circuit_id"`
	Speed                     string           `json:"speed"`
	Deleted                   bool             `json:"deleted"`
	AccountUUID               string           `json:"account_uuid"`
	ServiceClass              string           `json:"service_class"`
	ServiceProvider           string           `json:"service_provider"`
	ServiceType               string           `json:"service_type"`
	Description               string           `json:"description"`
	UUID                      string           `json:"uuid"`
	CloudProviderConnectionID string           `json:"cloud_provider_connection_id"`
	CloudSettings             AwsCloudSettings `json:"cloud_settings"`
	NatCapable                bool             `json:"nat_capable"`
	BgpState                  interface{}      `json:"bgp_state"`
	CloudRouterCircuitID      string           `json:"cloud_router_circuit_id"`
	ConnectionType            string           `json:"connection_type"`
	UserUUID                  string           `json:"user_uuid"`
	CustomerUUID              string           `json:"customer_uuid"`
	TimeCreated               string           `json:"time_created"`
	TimeUpdated               string           `json:"time_updated"`
	CloudProvider             AwsCloudProvider `json:"cloud_provider"`
	Pop                       string           `json:"pop"`
	Site                      string           `json:"site"`
}
type AwsCloudSettings struct {
	VlanIDPf        int    `json:"vlan_id_pf"`
	VlanIDCust      int    `json:"vlan_id_cust"`
	AwsRegion       string `json:"aws_region"`
	AwsHostedType   string `json:"aws_hosted_type"`
	AwsConnectionID string `json:"aws_connection_id"`
	AwsAccountID    string `json:"aws_account_id"`
	PublicIP        string `json:"public_ip"`
	NatPublicIP     string `json:"nat_public_ip"`
}
type AwsCloudProvider struct {
	Pop  string `json:"pop"`
	Site string `json:"site"`
}

type AwsSettings struct {
	AwsRegion       string `json:"aws_region"`
	AwsHostedType   string `json:"aws_hosted_type"`
	AwsConnectionID string `json:"aws_connection_id"`
	AwsAccountID    string `json:"aws_account_id"`
}

type AwsBilling struct {
	AccountUUID      string `json:"account_uuid"`
	SubscriptionTerm int    `json:"subscription_term"`
}

type AwsComponents struct {
	IfdPortCircuitIDCust string `json:"ifd_port_circuit_id_cust"`
	IfdPortCircuitIDPf   string `json:"ifd_port_circuit_id_pf"`
}

type DescriptionUpdate struct {
	Description string `json:"description"`
}

type ConnectionDeleteResp struct {
	Message string `json:"message"`
}

func (c *PFClient) CreateAwsConnection(connection AwsConnection, circuitId string) (*AwsConnectionCreateResponse, error) {
	formatedURI := fmt.Sprintf(awsConnectionURI, circuitId)

	resp := &AwsConnectionCreateResponse{}
	_, err := c.sendRequest(formatedURI, postMethod, connection, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *PFClient) ReadAwsConnection(cID, connCid string) (*AwsConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(awsConnectionByCidURI, cID, connCid)

	resp := &AwsConnectionReadResponse{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) UpdateAwsConnection(cID, connCid string, description DescriptionUpdate) (*AwsConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(awsConnectionByCidURI, cID, connCid)

	resp := &AwsConnectionReadResponse{}
	_, err := c.sendRequest(formatedURI, patchMethod, description, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) DeleteAwsConnection(cID, connCid string) (*ConnectionDeleteResp, error) {
	formatedURI := fmt.Sprintf(awsConnectionByCidURI, cID, connCid)
	if cID == "" {
		return nil, errors.New(errorMsg)
	}

	routerConn, _ := c.ReadAwsConnection(cID, connCid)
	if routerConn == nil {
		return &ConnectionDeleteResp{Message: fmt.Sprintf("No cloud router connection to delete for %s", cID)}, nil
	}

	expectedResp := &ConnectionDeleteResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetCloudConnectionStatus(cID, connCID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(awsConnectionStatusURI, cID, connCID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil

}

func (c *PFClient) ListAwsRouterConnections(cID string) ([]AwsConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(awsConnectionListURI, cID)
	resp := make([]AwsConnectionReadResponse, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if len(resp) == 0 {
		return resp, nil
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
