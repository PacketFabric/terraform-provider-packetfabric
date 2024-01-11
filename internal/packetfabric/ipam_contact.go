package packetfabric

import (
	"errors"
	"fmt"
)

const IpamContactURI = "/v2/services/ipam/contacts"

// This struct represents a IPAM contact
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_contact_post
type IpamContact struct {
	UUID        string `json:"uuid,omitempty"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	CountryCode string `json:"country_code"`
	ApnicOrgId  string `json:"apnic_org_id,omitempty"`
	ApnicRef    string `json:"apnic_ref,omitempty"`
	RipeOrgId   string `json:"ripe_org_id,omitempty"`
	RipeRef     string `json:"ripe_ref,omitempty"`
	TimeCreated string `json:"time_created,omitempty"`
	TimeUpdated string `json:"time_updated,omitempty"`
}

// This struct represents a IPAM contact delete response
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_contact_delete
type IpamContactDeleteResponse struct {
	Message string `json:"message"`
}

// This function represents the Action to create a new ipam contact
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_contact_post
func (c *PFClient) CreateIpamContact(ipamContact IpamContact) (*IpamContact, error) {
	resp := &IpamContact{}
	_, err := c.sendRequest(IpamContactURI, postMethod, ipamContact, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing IPAM contact by ID
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_contact_get
func (c *PFClient) ReadIpamContact(ipamContactID string) (*IpamContact, error) {
	formatedURI := fmt.Sprintf("%s/%s", IpamContactURI, ipamContactID)
	resp := &IpamContact{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing IPAM contact by ID
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_contact_get_list
func (c *PFClient) ReadIpamContacts() ([]IpamContact, error) {
	resp := make([]IpamContact, 0)
	_, err := c.sendRequest(IpamContactURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to update an existing IPAM Contact
// This operation is not a supported use case
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_contact_patch
func (c *PFClient) UpdateIpamContact(ipamContact IpamContact) (*IpamContact, error) {
	return nil, fmt.Errorf("Unsupported operation")
}

// This function represents the Action to Delete an existing IPAM contact
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_contact_delete
func (c *PFClient) DeleteIpamContact(ipamContactID string) (*IpamContactDeleteResponse, error) {
	if ipamContactID == "" {
		return nil, errors.New(errorMsg)
	}
	formatedURI := fmt.Sprintf("%s/%s", IpamContactURI, ipamContactID)
	expectedResp := &IpamContactDeleteResponse{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}
