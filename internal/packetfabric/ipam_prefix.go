package packetfabric

import (
	"errors"
	"fmt"
)

const IpamPrefixURI = "/ipam/prefix"
const IpamPrefixConfirmationURI = "/ipam/prefix/%s/confirm" // <prefix_uuid>

type IpamPrefix struct {
	UUID             string      `json:"uuid,omitempty"`   // set by the client, not user or api
	Prefix           string      `json:"prefix,omitempty"` // set by the client, not user or api
	Length           int         `json:"length"`
	Version          int         `json:"version" validate:"oneof=4 6" default:"4"`
	BgpRegion        string      `json:"bgp_region,omitempty"`
	AdminContactUuid string      `json:"admin_contact_uuid,omitempty"`
	TechContactUuid  string      `json:"tech_contact_uuid,omitempty"`
	IpjDetails       *IpjDetails `json:"ipj_details,omitempty"`
}

type IpjDetails struct {
	CurrentlyUsedPrefixes []IpamCurrentlyUsedPrefixes `json:"currently_used_prefixes"`
	PlannedPrefixes       []IpamPlannedPrefixes       `json:"planned_prefixes"`
}

type IpamCurrentlyUsedPrefixes struct {
	Prefix       string `json:"prefix"`
	IpsInUse     int    `json:"ips_in_use"`
	Description  string `json:"description,omitempty"`
	IspName      string `json:"isp_name,omitempty"`
	WillRenumber bool   `json:"will_renumber,omitempty"`
}

type IpamPlannedPrefixes struct {
	Prefix      string `json:"prefix"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Usage30d    int    `json:"usage_30d"`
	Usage3m     int    `json:"usage_3m"`
	Usage6m     int    `json:"usage_6m"`
	Usage1y     int    `json:"usage_1y"`
}

type IpamPrefixCreateResponse struct {
	PrefixUuid string `json:"prefix_uuid"`
	Prefix     string `json:"prefix"`
	BgpRegion  string `json:"bgp_region,omitempty"`
}
type IpamPrefixDeleteResponse struct {
	Message string `json:"message"`
}

type IpamPrefixConfirmation struct {
	PrefixUuid       string      `json:"prefix_uuid,omitempty"` // set by user, used by client, ignored by api
	AdminContactUuid string      `json:"admin_contact_uuid"`
	TechContactUuid  string      `json:"tech_contact_uuid"`
	IpjDetails       *IpjDetails `json:"ipj_details,omitempty"`
}

type IpamPrefixConfirmationCreationResponse struct {
	Message string `json:"message"`
}

type IpamPrefixConfirmationDeleteResponse struct {
	Message string `json:"message"`
}

/////////////////////////////////////////////////////////////////////////////

// This function represents the Action to create a new ipam prefix
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix
func (c *PFClient) CreateIpamPrefix(ipamPrefix IpamPrefix) (*IpamPrefix, error) {
	resp := &IpamPrefixCreateResponse{}
	_, err := c.sendRequest(IpamPrefixURI, postMethod, ipamPrefix, &resp)
	if err != nil {
		return nil, err
	}
	return c.ReadIpamPrefix(resp.PrefixUuid)
}

// This function represents the Action to Retrieve an existing IPAM prefix by ID
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix
func (c *PFClient) ReadIpamPrefix(ipamPrefixID string) (*IpamPrefix, error) {
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixURI, ipamPrefixID)
	resp := &IpamPrefix{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to update an existing IPAM prefix
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix
func (c *PFClient) UpdateIpamPrefix(ipamPrefix IpamPrefix) (*IpamPrefix, error) {
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixURI, ipamPrefix.UUID)
	resp := &IpamPrefix{}
	_, err := c.sendRequest(formatedURI, patchMethod, &ipamPrefix, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Delete an existing IPAM prefix
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix
func (c *PFClient) DeleteIpamPrefix(ipamPrefixID string) (*IpamPrefixDeleteResponse, error) {
	if ipamPrefixID == "" {
		return nil, errors.New(errorMsg)
	}
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixURI, ipamPrefixID)
	expectedResp := &IpamPrefixDeleteResponse{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}

/////////////////////////////////////////////////////////////////////////////

// This function represents the Action to create a new ipam prefix confirmation
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix_confirmation
func (c *PFClient) CreateIpamPrefixConfirmation(ipamPrefixConfirmation IpamPrefixConfirmation) (*IpamPrefixConfirmationCreationResponse, error) {
	resp := &IpamPrefixConfirmationCreationResponse{}
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixConfirmationURI, ipamPrefixConfirmation.PrefixUuid)
	_, err := c.sendRequest(formatedURI, postMethod, ipamPrefixConfirmation, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing IPAM prefix confirmation by ID
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix_confirmation
func (c *PFClient) ReadIpamPrefixConfirmation(ipamPrefixConfirmationID string) (*IpamPrefixConfirmation, error) {
	resp := &IpamPrefixConfirmation{}
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixConfirmationURI, ipamPrefixConfirmationID)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to update an existing IPAM prefix confirmation
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix_confirmation
func (c *PFClient) UpdateIpamPrefixConfirmation(ipamPrefixConfirmation IpamPrefixConfirmation) (*IpamPrefixConfirmation, error) {
	resp := &IpamPrefixConfirmation{}
	formatedURI := fmt.Sprintf(IpamPrefixConfirmationURI, ipamPrefixConfirmation.PrefixUuid)
	_, err := c.sendRequest(formatedURI, patchMethod, &ipamPrefixConfirmation, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Delete an existing IPAM prefix confirmation
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix_confirmation
func (c *PFClient) DeleteIpamPrefixConfirmation(ipamPrefixConfirmationID string) (*IpamPrefixConfirmationDeleteResponse, error) {
	if ipamPrefixConfirmationID == "" {
		return nil, errors.New(errorMsg)
	}
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixConfirmationURI, ipamPrefixConfirmationID)
	expectedResp := &IpamPrefixConfirmationDeleteResponse{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}
