package packetfabric

import (
	"errors"
	"fmt"
)

const cloudRouterQuickConnectURI = "/v2/services/cloud-routers/%s/connections/%s/third-party"
const cloudRouterQuickConnectByCIDURI = "/v2/services/cloud-routers/requests/%s"
const cloudRouterQuickConnectByImportCIDURI = "/v2/services/cloud-routers/%s/connections/%s/third-party/%s"

type CloudRouterQuickConnect struct {
	ServiceUUID   string                      `json:"service_uuid,omitempty"`
	ImportFilters []QuickConnectImportFilters `json:"import_filters,omitempty"`
	ReturnFilters []QuickConnectReturnFilters `json:"return_filters,omitempty"`
}

type CloudRouterQuickConnectUpdate struct {
	ImportFilters []QuickConnectImportFilters `json:"import_filters,omitempty"`
	ReturnFilters []QuickConnectReturnFilters `json:"return_filters,omitempty"`
}

type QuickConnectImportFilters struct {
	Prefix    string `json:"prefix,omitempty"`
	MatchType string `json:"match_type,omitempty"`
	Localpref int    `json:"localpref,omitempty"`
}
type QuickConnectReturnFilters struct {
	Prefix    string `json:"prefix,omitempty"`
	MatchType string `json:"match_type,omitempty"`
	Asprepend int    `json:"asprepend,omitempty"`
	Med       int    `json:"med,omitempty"`
	Localpref int    `json:"localpref,omitempty"`
}

type CloudRouterQuickConnectResp struct {
	CircuitID         string `json:"circuit_id,omitempty"`
	RouteSetCircuitID string `json:"route_set_circuit_id,omitempty"`
	ServiceUUID       string `json:"service_uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	IsDefunct         bool   `json:"is_defunct,omitempty"`
	State             string `json:"state,omitempty"`
	TimeCreated       string `json:"time_created,omitempty"`
	TimeUpdated       string `json:"time_updated,omitempty"`
}

func (c *PFClient) CreateCloudRouterQuickConnect(crCID, connCID string, quickConnect CloudRouterQuickConnect) (*CloudRouterQuickConnectResp, error) {
	formatedURI := fmt.Sprintf(cloudRouterQuickConnectURI, crCID, connCID)
	quickConnectResp := &CloudRouterQuickConnectResp{}
	if _, err := c.sendRequest(formatedURI, postMethod, quickConnect, quickConnectResp); err != nil {
		return nil, err
	}
	return quickConnectResp, nil
}

func (c *PFClient) GetCloudRouterQuickConnectState(circuitID string) (currentState string, err error) {
	type CloudRouterQuickConnectState struct {
		State string `json:"state"`
	}
	resp := &CloudRouterQuickConnectState{}
	if _, err = c.sendRequest(fmt.Sprintf(cloudRouterQuickConnectByCIDURI, circuitID), getMethod, nil, resp); err == nil {
		currentState = resp.State
	}
	return
}

func (c *PFClient) UpdateCloudRouterQuickConnect(crCID, connCID, importCID string, quickConnect CloudRouterQuickConnectUpdate) (err error) {
	formatedURI := fmt.Sprintf(cloudRouterQuickConnectByImportCIDURI, crCID, connCID, importCID)
	_, err = c.sendRequest(formatedURI, patchMethod, quickConnect, nil)
	return
}

func (c *PFClient) DeleteCloudRouterQuickConnect(cID, crCID, connCID, importCID string) (warningMessage string, err error) {
	state, err := c.GetCloudRouterQuickConnectState(cID)
	if err != nil {
		return "error", err
	}
	switch {
	case state == "pending":
		if cID == "" {
			err = errors.New("circuit ID cannot be empty")
			return
		}
		err = c._deletePendingCloudRouterQuickConnect(cID)
	case state == "active":
		if crCID == "" || connCID == "" || importCID == "" {
			err = errors.New("cloud router circuit ID, connection circuit ID or import circuit ID cannot be empty")
			return
		}
		err = c._deleteActiveCloudRouterQuickConnect(crCID, connCID, importCID)
	case state == "rejected":
		warningMessage = "the Z side has rejected the request. Remove the resource from Terraform state and resubmit your request as needed"
	case state == "innactive":
		warningMessage = "the cloud router quick connect is innactive and cannot be deleted"
	}
	return
}

func (c *PFClient) _deletePendingCloudRouterQuickConnect(cID string) (err error) {
	formatedURI := fmt.Sprintf(cloudRouterQuickConnectByCIDURI, cID)
	_, err = c.sendRequest(formatedURI, deleteMethod, nil, nil)
	return
}

func (c *PFClient) _deleteActiveCloudRouterQuickConnect(crID, connCID, importCID string) (err error) {
	formatedURI := fmt.Sprintf(cloudRouterQuickConnectByImportCIDURI, crID, connCID, importCID)
	_, err = c.sendRequest(formatedURI, deleteMethod, nil, nil)
	return
}
