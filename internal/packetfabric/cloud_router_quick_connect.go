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
	Prefix          string `json:"prefix,omitempty"`
	MatchType       string `json:"match_type,omitempty"`
	LocalPreference int    `json:"local_preference,omitempty"`
}
type QuickConnectReturnFilters struct {
	Prefix          string `json:"prefix,omitempty"`
	MatchType       string `json:"match_type,omitempty"`
	AsPrepend       int    `json:"as_prepend,omitempty"`
	Med             int    `json:"med,omitempty"`
	LocalPreference int    `json:"local_preference,omitempty"`
	PendingApproval bool   `json:"pending_approval,omitempty"`
}

type CloudRouterQuickConnectResp struct {
	ImportCircuitID   string                      `json:"import_circuit_id,omitempty"`
	RouteSetCircuitID string                      `json:"route_set_circuit_id,omitempty"`
	ServiceUUID       string                      `json:"service_uuid,omitempty"`
	Name              string                      `json:"name,omitempty"`
	Description       string                      `json:"description,omitempty"`
	IsDefunct         bool                        `json:"is_defunct,omitempty"`
	State             string                      `json:"state,omitempty"`
	ConnectionSpeed   string                      `json:"connection_speed,omitempty"`
	ImportFilters     []QuickConnectImportFilters `json:"import_filters,omitempty"`
	ReturnFilters     []QuickConnectReturnFilters `json:"return_filters,omitempty"`
}

func (c *PFClient) CreateCloudRouterQuickConnect(crCID, connCID string, quickConnect CloudRouterQuickConnect) (*CloudRouterQuickConnectResp, error) {
	formatedURI := fmt.Sprintf(cloudRouterQuickConnectURI, crCID, connCID)
	quickConnectResp := &CloudRouterQuickConnectResp{}
	if _, err := c.sendRequest(formatedURI, postMethod, quickConnect, quickConnectResp); err != nil {
		return nil, err
	}
	return quickConnectResp, nil
}

func (c *PFClient) GetCloudRouterQuickConnect(crCID, connCID, importCID string) (*CloudRouterQuickConnectResp, error) {
	formatedURI := fmt.Sprintf(cloudRouterQuickConnectByImportCIDURI, crCID, connCID, importCID)
	quickConnectResp := &CloudRouterQuickConnectResp{}
	if _, err := c.sendRequest(formatedURI, getMethod, nil, &quickConnectResp); err != nil {
		return nil, err
	}
	return quickConnectResp, nil
}

func (c *PFClient) GetCloudRouterQuickConnectState(ImportCircuitID string) (currentState string, err error) {
	type CloudRouterQuickConnectState struct {
		State string `json:"state"`
	}
	resp := &CloudRouterQuickConnectState{}
	if _, err = c.sendRequest(fmt.Sprintf(cloudRouterQuickConnectByCIDURI, ImportCircuitID), getMethod, nil, resp); err == nil {
		currentState = resp.State
	}
	return
}

func (c *PFClient) UpdateCloudRouterQuickConnect(crCID, connCID, importCID string, quickConnect CloudRouterQuickConnectUpdate) (err error) {
	formatedURI := fmt.Sprintf(cloudRouterQuickConnectByImportCIDURI, crCID, connCID, importCID)
	_, err = c.sendRequest(formatedURI, patchMethod, quickConnect, nil)
	return
}

func (c *PFClient) DeleteCloudRouterQuickConnect(crCID, connCID, importCID string) (warningMessage string, err error) {
	state, err := c.GetCloudRouterQuickConnectState(importCID)
	if err != nil {
		return "error", err
	}
	switch {
	case state == "pending":
		if importCID == "" {
			err = errors.New("import circuit id cannot be empty")
			return
		}
		err = c._deletePendingCloudRouterQuickConnect(importCID)
	case state == "active":
		if crCID == "" || connCID == "" || importCID == "" {
			err = fmt.Errorf("import circuit id, cloud router circuit id, cloud router connection circuit id cannot be empty: id: %s; cr_circuit_id: %s; connection_circuit_id: %s", importCID, crCID, connCID)
			return
		}
		err = c._deleteActiveCloudRouterQuickConnect(crCID, connCID, importCID)
	case state == "rejected":
		warningMessage = "the Z side has rejected the request. Remove the resource from Terraform state and resubmit your request as needed"
	case state == "inactive":
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
