package packetfabric

import "fmt"

const awsBackbonURI = "/v2/services/backbone"
const awsBackStatusURI = "/v2.1/services/%s/status"
const awsBackDeleteURI = "/v2/services/%s"

type AwsBackbone struct {
	Description  string                 `json:"description"`
	Bandwidth    AwsBackboneBandwidth   `json:"bandwidth"`
	Interfaces   []AwsBackBoneInterface `json:"interfaces"`
	RateLimitIn  int                    `json:"rate_limit_in"`
	RateLimitOut int                    `json:"rate_limit_out"`
	Epl          bool                   `json:"epl"`
}

type AwsBackboneBandwidth struct {
	AccountUUID      string `json:"account_uuid"`
	SubscriptionTerm int    `json:"subscription_term"`
	Speed            string `json:"speed"`
}

type AwsBackBoneInterface struct {
	PortCircuitID string `json:"port_circuit_id"`
	Vlan          int    `json:"vlan"`
	Untagged      bool   `json:"untagged"`
}

type AwsBackboneResp struct {
	VcCircuitID  string               `json:"vc_circuit_id"`
	CustomerUUID string               `json:"customer_uuid"`
	State        string               `json:"state"`
	ServiceType  string               `json:"service_type"`
	ServiceClass string               `json:"service_class"`
	Mode         string               `json:"mode"`
	Connected    bool                 `json:"connected"`
	Bandwidth    AwsBackboneBandwidth `json:"bandwidth"`
	Description  string               `json:"description"`
	RateLimitIn  int                  `json:"rate_limit_in"`
	RateLimitOut int                  `json:"rate_limit_out"`
	TimeCreated  string               `json:"time_created"`
	TimeUpdated  string               `json:"time_updated"`
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

type AwsBackboneDeleteResp struct {
	Message string `json:"message"`
}

func (c *PFClient) CreateAwsBackbone(awsBackbone AwsBackbone) (*AwsBackboneResp, error) {
	awsBackboneResp := &AwsBackboneResp{}
	_, err := c.sendRequest(awsBackbonURI, postMethod, awsBackbone, awsBackboneResp)
	if err != nil {
		return nil, err
	}
	createOk := make(chan bool)
	defer close(createOk)
	fn := func() (*ServiceState, error) {
		return c.GetAwsBackboneState(awsBackboneResp.VcCircuitID)
	}
	go c.CheckServiceStatus(createOk, err, fn)
	if !<-createOk {
		return nil, err
	}
	return awsBackboneResp, err
}

func (c *PFClient) GetAwsBackboneState(vcCircuitID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(awsBackStatusURI, vcCircuitID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) DeleteAwsBackbone(vcCircuitID string) (*AwsBackboneDeleteResp, error) {
	formatedURI := fmt.Sprintf(awsBackDeleteURI, vcCircuitID)
	expectedResp := &AwsBackboneDeleteResp{}
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
