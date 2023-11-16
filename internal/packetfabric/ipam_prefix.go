package packetfabric

import (
	"errors"
	"fmt"
)

const IpamPrefixURI = "/ipam/prefix"

type IpamPrefix struct {
	PrefixUuid       string      `json:"prefix_uuid,omitempty"` // set by the client, not user or api
	Prefix           string      `json:"prefix,omitempty"`      // set by the client, not user or api
	Length           int         `json:"length"`
	Version          int         `json:"version" validate:"oneof=4 6" default:"4"`
	BgpRegion        string      `json:"bgp_region,omitempty"`
	AdminContactUuid string      `json:"admin_contact_uuid,omitempty"`
	TechContactUuid  string      `json:"tech_contact_uuid,omitempty"`
	State            string      `json:"state,omitempty"`
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
	resp.PrefixUuid = ipamPrefixID
	return resp, nil
}

func (c *PFClient) ReadIpamPrefixes() ([]IpamPrefix, error) {
	resp := make([]IpamPrefix, 0)
	_, err := c.sendRequest(IpamPrefixURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to update an existing IPAM prefix
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix
func (c *PFClient) UpdateIpamPrefix(ipamPrefix IpamPrefix) (*IpamPrefix, error) {
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixURI, ipamPrefix.PrefixUuid)
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
		return nil, errors.New("IPAM Prefix UUID required for delete operation")
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
