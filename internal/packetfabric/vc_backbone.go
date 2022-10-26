package packetfabric

import "fmt"

const vcBackStatusURI = "/v2.1/services/%s/status"

func (c *PFClient) GetBackboneState(vcCircuitID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(vcBackStatusURI, vcCircuitID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
