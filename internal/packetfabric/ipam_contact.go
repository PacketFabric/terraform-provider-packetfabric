package packetfabric

import (
	"errors"
	"fmt"
)

const IpamContactURI = "/ipam/contact"

// This struct represents a IPAM contact
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/contact_post
type IpamContact struct {
	UUID        string `json:"uuid,omitempty"`
	ContactName string `json:"contact_name"`
	OrgName     string `json:"org_name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	ArinOrgId   string `json:"arin_org_id,omitempty"`
	ApnicOrgId  string `json:"apnic_org_id,omitempty"`
	RipeOrgId   string `json:"ripe_org_id,omitempty"`
}

// This struct represents a IPAM contact delete response
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/contact_delete
type IpamContactDeleteResponse struct {
	UUID    string `json:"uuid"`
	Message string `json:"message"`
}

// This function represents the Action to create a new ipam contact
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/contact_post
func (c *PFClient) CreateIpamContact(ipamContact IpamContact) (*IpamContact, error) {
	resp := &IpamContact{}
	_, err := c.sendRequest(IpamContactURI, postMethod, ipamContact, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing IPAM contact by ID
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/contact_get_by_login
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
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/contact
func (c *PFClient) ReadIpamContacts() ([]IpamContact, error) {
	resp := make([]IpamContact, 0)
	_, err := c.sendRequest(IpamContactURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to update an existing IPAM Contact
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/contact_patch
func (c *PFClient) UpdateIpamContact(ipamContact IpamContact) (*IpamContact, error) {
	formatedURI := fmt.Sprintf("%s/%s", IpamContactURI, ipamContact.UUID)
	resp := &IpamContact{}
	_, err := c.sendRequest(formatedURI, patchMethod, &ipamContact, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Delete an existing IPAM contact
// https://docs.packetfabric.com/api/v2/swagger/#/ipam/contact_delete_by_login
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
