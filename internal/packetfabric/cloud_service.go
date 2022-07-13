package packetfabric

import "fmt"

const backDeleteURI = "/v2/services/%s"

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
