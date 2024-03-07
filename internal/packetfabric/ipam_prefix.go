package packetfabric

import (
	"errors"
	"fmt"
)

const IpamPrefixURI = "/v2/services/ipam/prefixes"

type IpamPrefix struct {
	Length               int         `json:"length"`                             // write
	Market               string      `json:"market,omitempty"`                   // write
	Family               string      `json:"address_family,omitempty"`           // write
	IpAddress            string      `json:"ip_address,omitempty"`               // read
	CircuitId            string      `json:"prefix_circuit_id,omitempty"`        // read
	LinkedCircuitId      string      `json:"linked_object_circuit_id,omitempty"` // read
	Type                 string      `json:"type,omitempty"`                     // read
	State                string      `json:"state,omitempty"`                    // read
	OrgId                string      `json:"org_id,omitempty"`                   // optional
	Address              string      `json:"address,omitempty"`                  // "
	City                 string      `json:"city,omitempty"`                     // "
	PostalCode           string      `json:"postal_code,omitempty"`              // "
	AdminIpamContactUuid string      `json:"admin_ipam_contact_uuid,omitempty"`
	TechIpamContactUuid  string      `json:"tech_ipam_contact_uuid,omitempty"`
	TimeCreated          string      `json:"time_created,omitempty"`
	TimeUpdated          string      `json:"time_updated,omitempty"`
	IpjDetails           *IpjDetails `json:"ipj_details,omitempty"`
}

type IpjDetails struct {
	CurrentPrefixes []IpamCurrentPrefixes `json:"current_prefixes"`
	PlannedPrefix   *IpamPlannedPrefix    `json:"planned_prefix"`
	RejectionReason string                `json:"rejection_reason,omitempty"` // read
}

type IpamCurrentPrefixes struct {
	Prefix       string `json:"prefix"`
	IpsInUse     int    `json:"ips_in_use"`
	Description  string `json:"description,omitempty"`
	IspName      string `json:"isp_name,omitempty"`
	WillRenumber bool   `json:"will_renumber,omitempty"`
}

type IpamPlannedPrefix struct {
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Usage30d    int    `json:"usage_30d"`
	Usage3m     int    `json:"usage_3m"`
	Usage6m     int    `json:"usage_6m"`
	Usage1y     int    `json:"usage_1y"`
}

type IpamPrefixDeleteResponse struct {
	Message string `json:"message"`
}

/////////////////////////////////////////////////////////////////////////////

// This function represents the Action to create a new ipam prefix
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix
func (c *PFClient) CreateIpamPrefix(ipamPrefix IpamPrefix) (*IpamPrefix, error) {
	resp := &IpamPrefix{}
	_, err := c.sendRequest(IpamPrefixURI, postMethod, ipamPrefix, &resp)
	if err != nil {
		return nil, err
	}
	return c.ReadIpamPrefix(resp.CircuitId)
}

// This function represents the Action to Retrieve an existing IPAM prefix by ID
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix
func (c *PFClient) ReadIpamPrefix(circuitId string) (*IpamPrefix, error) {
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixURI, circuitId)
	resp := &IpamPrefix{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
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
	return nil, fmt.Errorf("Ipam Prefix Update is an Unsupported operation")
}

// This function represents the Action to Delete an existing IPAM prefix
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/prefix
func (c *PFClient) DeleteIpamPrefix(circuitId string) (*IpamPrefixDeleteResponse, error) {
	if circuitId == "" {
		return nil, errors.New("Circuit ID required for delete operation")
	}
	formatedURI := fmt.Sprintf("%s/%s", IpamPrefixURI, circuitId)
	expectedResp := &IpamPrefixDeleteResponse{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}

/////////////////////////////////////////////////////////////////////////////
