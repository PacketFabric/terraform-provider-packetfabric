package packetfabric

import (
	"errors"
	"fmt"
)

const cloudRouterURI = "/v2/services/cloud-routers"

const errorMsg = "Please provide a valid Account UUID."

// This struct represents a Cloud Router
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_create
type CloudRouter struct {
	Asn         int      `json:"asn,omitempty"`
	Name        string   `json:"name"`
	AccountUUID string   `json:"account_uuid"`
	Regions     []string `json:"regions,omitempty"`
	Capacity    string   `json:"capacity"`
}

// This struct represents a Cloud Router create response
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_create
type CloudRouterResponse struct {
	CircuitID   string   `json:"circuit_id"`
	AccountUUID string   `json:"account_uuid"`
	Asn         int      `json:"asn"`
	Name        string   `json:"name"`
	Capacity    string   `json:"capacity"`
	Regions     []Region `json:"regions"`
	TimeCreated string   `json:"time_created"`
	TimeUpdated string   `json:"time_updated"`
	PONumber    string   `json:"po_number"`
}

// This struct represents a Cloud Router Region
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_create
type Region struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CloudRouterUpdate struct {
	Name     string   `json:"name,omitempty"`
	Regions  []string `json:"regions,omitempty"`
	Capacity string   `json:"capacity,omitempty"`
}

// This struct represents a Cloud Router delete response
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_router_delete
type CloudRouterDelResp struct {
	Message string `json:"message"`
}

// This function represents the Action to create a new Cloud Router
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_create
func (c *PFClient) CreateCloudRouter(router CloudRouter) (*CloudRouterResponse, error) {
	resp := &CloudRouterResponse{}
	_, err := c.sendRequest(cloudRouterURI, postMethod, router, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing Cloud Router by Circut ID
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_router_get
func (c *PFClient) ReadCloudRouter(cID string) (*CloudRouterResponse, error) {
	formatedURI := fmt.Sprintf("%s/%s", cloudRouterURI, cID)
	resp := &CloudRouterResponse{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action tp update an existing Cloud Router
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_patch
func (c *PFClient) UpdateCloudRouter(router CloudRouterUpdate, cID string) (*CloudRouterResponse, error) {
	formatedURI := fmt.Sprintf("%s/%s", cloudRouterURI, cID)
	resp := &CloudRouterResponse{}
	_, err := c.sendRequest(formatedURI, patchMethod, router, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Delete an existing Cloud Router
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_router_delete
func (c *PFClient) DeleteCloudRouter(cID string) (*CloudRouterDelResp, error) {
	if cID == "" {
		return nil, errors.New(errorMsg)
	}
	formatedURI := fmt.Sprintf("%s/%s", cloudRouterURI, cID)
	expectedResp := &CloudRouterDelResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}

// This function represents the Action to retrieve the list of existing Cloud Routers
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_list
func (c *PFClient) ListCloudRouters() ([]CloudRouterResponse, error) {
	expectedResp := make([]CloudRouterResponse, 0)
	_, err := c.sendRequest(cloudRouterURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
