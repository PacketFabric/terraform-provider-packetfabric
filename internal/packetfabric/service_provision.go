package packetfabric

import "fmt"

const requestServiceProvisionURI = "/v2/services/requests/%s/provision"
const requestHostedServiceProvisionURI = "/v2/services/requests/%s/provision/hosted"
const rejectServiceProvisionURI = "/v2/services/requests/%s/reject"

const (
	backboneService = "backbone"
	ixService       = "ix"
	cloudService    = "cloud"
)

type ServiceProvision struct {
	Provider    string    `json:"provider,omitempty"`
	Interface   Interface `json:"interface,omitempty"`
	Description string    `json:"description,omitempty"`
}

func (c *PFClient) RequestServiceProvision(vcRequestUUID, reqType string, provisionReq ServiceProvision) (*MktConnProvisionResp, error) {
	var formatedURI string
	switch reqType {
	case backboneService, ixService:
		formatedURI = fmt.Sprintf(requestServiceProvisionURI, vcRequestUUID)
	case cloudService:
		formatedURI = fmt.Sprintf(requestHostedServiceProvisionURI, vcRequestUUID)
	}
	expectedResp := &MktConnProvisionResp{}
	_, err := c.sendRequest(formatedURI, postMethod, provisionReq, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) RejectServiceRequest(vcRequestUUID string) (*VcRequest, error) {
	formatedURI := fmt.Sprintf(rejectServiceProvisionURI, vcRequestUUID)
	expectedResp := &VcRequest{}
	_, err := c.sendRequest(formatedURI, postMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
