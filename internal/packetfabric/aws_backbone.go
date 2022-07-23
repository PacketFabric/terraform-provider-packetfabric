package packetfabric

import "fmt"

const awsBackStatusURI = "/v2.1/services/%s/status"

func (c *PFClient) GetAwsBackboneState(vcCircuitID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(awsBackStatusURI, vcCircuitID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
