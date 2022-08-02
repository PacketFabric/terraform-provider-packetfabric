package packetfabric

import "fmt"

const backboneURI = "/v2/services/backbone"
const backDeleteURI = "/v2/services/%s"

type Backbone struct {
	Description  string              `json:"description"`
	Bandwidth    BackboneBandwidth   `json:"bandwidth"`
	Interfaces   []BackBoneInterface `json:"interfaces"`
	RateLimitIn  int                 `json:"rate_limit_in"`
	RateLimitOut int                 `json:"rate_limit_out"`
	Epl          bool                `json:"epl"`
}

type BackboneBandwidth struct {
	AccountUUID      string `json:"account_uuid"`
	SubscriptionTerm int    `json:"subscription_term"`
	Speed            string `json:"speed"`
}

type BackBoneInterface struct {
	PortCircuitID string `json:"port_circuit_id"`
	Vlan          int    `json:"vlan"`
	Untagged      bool   `json:"untagged"`
}

type BackboneResp struct {
	VcCircuitID  string            `json:"vc_circuit_id"`
	CustomerUUID string            `json:"customer_uuid"`
	State        string            `json:"state"`
	ServiceType  string            `json:"service_type"`
	ServiceClass string            `json:"service_class"`
	Mode         string            `json:"mode"`
	Connected    bool              `json:"connected"`
	Bandwidth    BackboneBandwidth `json:"bandwidth"`
	Description  string            `json:"description"`
	RateLimitIn  int               `json:"rate_limit_in"`
	RateLimitOut int               `json:"rate_limit_out"`
	TimeCreated  string            `json:"time_created"`
	TimeUpdated  string            `json:"time_updated"`
	Interfaces   []struct {
		PortCircuitID      string `json:"port_circuit_id"`
		Pop                string `json:"pop"`
		Site               string `json:"site"`
		SiteName           string `json:"site_name"`
		Speed              string `json:"speed"`
		Media              string `json:"media"`
		Zone               string `json:"zone"`
		Description        string `json:"description"`
		Vlan               int    `json:"vlan"`
		Untagged           bool   `json:"untagged"`
		ProvisioningStatus string `json:"provisioning_status"`
		AdminStatus        string `json:"admin_status"`
		OperationalStatus  string `json:"operational_status"`
		CustomerUUID       string `json:"customer_uuid"`
		CustomerName       string `json:"customer_name"`
		Region             string `json:"region"`
		IsCloud            bool   `json:"is_cloud"`
		IsPtp              bool   `json:"is_ptp"`
		TimeCreated        string `json:"time_created"`
		TimeUpdated        string `json:"time_updated"`
	} `json:"interfaces"`
}

type FromCustomer struct {
	CustomerUUID      string `json:"customer_uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Market            string `json:"market,omitempty"`
	MarketDescription string `json:"market_description,omitempty"`
	ContactFirstName  string `json:"contact_first_name,omitempty"`
	ContactLastName   string `json:"contact_last_name,omitempty"`
	ContactEmail      string `json:"contact_email,omitempty"`
	ContactPhone      string `json:"contact_phone,omitempty"`
}
type ToCustomer struct {
	CustomerUUID      string `json:"customer_uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Market            string `json:"market,omitempty"`
	MarketDescription string `json:"market_description,omitempty"`
}
type Bandwidth struct {
	AccountUUID      string `json:"account_uuid,omitempty"`
	LonghaulType     string `json:"longhaul_type,omitempty"`
	SubscriptionTerm int    `json:"subscription_term,omitempty"`
	Speed            string `json:"speed,omitempty"`
}

type HostedConnectionResp struct {
	CustomerUUID    string `json:"customer_uuid"`
	UserUUID        string `json:"user_uuid"`
	ServiceProvider string `json:"service_provider"`
	PortType        string `json:"port_type"`
	ServiceClass    string `json:"service_class"`
	Description     string `json:"description"`
	State           string `json:"state"`
	Speed           string `json:"speed"`
}

type BackboneDeleteResp struct {
	Message string `json:"message"`
}

func (c *PFClient) CreateBackbone(backbone Backbone) (*BackboneResp, error) {
	backboneResp := &BackboneResp{}
	_, err := c.sendRequest(backboneURI, postMethod, backbone, backboneResp)
	if err != nil {
		return nil, err
	}
	return backboneResp, nil
}

func (c *PFClient) UpdateServiceConn(description, cloudCID string) (*HostedConnectionResp, error) {
	formatedURI := fmt.Sprintf(updateCloudConnURI, cloudCID)
	type UpdateServiceConn struct {
		Description string `json:"description"`
	}
	expectedResp := &HostedConnectionResp{}
	_, err := c.sendRequest(formatedURI, patchMethod, UpdateServiceConn{description}, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) DeleteBackbone(vcCircuitID string) (*BackboneDeleteResp, error) {
	formatedURI := fmt.Sprintf(backDeleteURI, vcCircuitID)
	expectedResp := &BackboneDeleteResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	deleteOk := make(chan bool)
	defer close(deleteOk)
	fn := func() (*ServiceState, error) {
		return c.GetAwsBackboneState(vcCircuitID)
	}
	go c.CheckServiceStatus(deleteOk, err, fn)
	if !<-deleteOk {
		return nil, err
	}
	return expectedResp, nil
}
