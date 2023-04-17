package packetfabric

import (
	"errors"
	"fmt"
)

const FlexBandwidthURI = "/v2/flex-bandwidth"

// This struct represents a Flex Bandwidth
// https://docs.packetfabric.com/api/v2/swagger/#/Flex%20Bandwidth/create_flex_bandwidth
type FlexBandwidth struct {
	Description      string `json:"description"`
	AccountUUID      string `json:"account_uuid"`
	SubscriptionTerm int    `json:"subscription_term"`
	Capacity         string `json:"capacity"`
	PONumber         string `json:"po_number,omitempty"`
}

// This struct represents a Flex Bandwidth create response
// https://docs.packetfabric.com/api/v2/swagger/#/Flex%20Bandwidth/create_flex_bandwidth
type FlexBandwidthResponse struct {
	FlexBandwidthID       string `json:"flex_bandwidth_id"`
	AccountUUID           string `json:"account_uuid"`
	Description           string `json:"description"`
	SubscriptionTerm      int    `json:"subscription_term"`
	CapacityMbps          int    `json:"capacity_mbps"`
	UsedCapacityMbps      int    `json:"used_capacity_mbps"`
	AvailableCapacityMbps int    `json:"available_capacity_mbps"`
	PONumber              string `json:"po_number,omitempty"`
	TimeCreated           string `json:"time_created"`
	TimeUpdated           string `json:"time_updated"`
}

// This struct represents a Flex Bandwidth delete response
// https://docs.packetfabric.com/api/v2/swagger/#/Flex%20Bandwidth/delete_flex_bandwidth
type FlexBandwidthDelResp struct {
	Message string `json:"message"`
}

// This function represents the Action to create a new Flex Bandwidth
// https://docs.packetfabric.com/api/v2/swagger/#/Flex%20Bandwidth/create_flex_bandwidth
func (c *PFClient) CreateFlexBandwidth(flexBand FlexBandwidth) (*FlexBandwidthResponse, error) {
	resp := &FlexBandwidthResponse{}
	_, err := c.sendRequest(FlexBandwidthURI, postMethod, flexBand, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing Flex Bandwidth by ID
// https://docs.packetfabric.com/api/v2/swagger/#/Flex%20Bandwidth/get_flex_bandwidth_by_id
func (c *PFClient) ReadFlexBandwidth(flexID string) (*FlexBandwidthResponse, error) {
	formatedURI := fmt.Sprintf("%s/%s", FlexBandwidthURI, flexID)
	resp := &FlexBandwidthResponse{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Delete an existing Flex Bandwidth
// https://docs.packetfabric.com/api/v2/swagger/#/Flex%20Bandwidth/delete_flex_bandwidth
func (c *PFClient) DeleteFlexBandwidth(flexID string) (*FlexBandwidthDelResp, error) {
	if flexID == "" {
		return nil, errors.New(errorMsg)
	}
	formatedURI := fmt.Sprintf("%s/%s", FlexBandwidthURI, flexID)
	expectedResp := &FlexBandwidthDelResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}
