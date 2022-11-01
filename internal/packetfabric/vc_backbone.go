package packetfabric

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

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

func (c *PFClient) IsBackboneComplete(vcCircuitID string) (result bool) {
	status, err := c.GetBackboneState(vcCircuitID)
	debugLog := make(map[string]interface{})
	debugLog["status"] = status
	debugLog["error"] = err
	tflog.Debug(c.Ctx, fmt.Sprintf("\n### BACKLOG STATUS: VCCID [%s] ###", vcCircuitID), debugLog)
	if err == nil && status.Status.LastWorkflow.CurrentState != "COMPLETE" {
		result = true
		return
	}
	if err != nil {
		result = false
	}
	return
}
