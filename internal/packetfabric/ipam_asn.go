package packetfabric

import (
	"errors"
	"fmt"
)

const IpamAsnURI = "/v2/services/ipam/asns"

// This struct represents a ipam asns
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_asn_post
type IpamAsn struct {
	AsnByteType int    `json:"asn_byte_type,omitempty"` // write
	Asn         int    `json:"asn,omitempty"`           // read
	TimeCreated string `json:"time_created,omitempty"`  // read
	TimeUpdated string `json:"time_updated,omitempty"`  // read
}

// This struct represents a ipam asns delete response
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_asn_delete
type IpamAsnDeleteResponse struct {
	Message string `json:"message"`
}

// This function represents the Action to create a new ipam asns
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_asn_post
func (c *PFClient) CreateIpamAsn(ipamAsn IpamAsn) (*IpamAsn, error) {
	resp := &IpamAsn{}
	_, err := c.sendRequest(IpamAsnURI, postMethod, ipamAsn, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing ipam asns by ID
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_asn_get
func (c *PFClient) ReadIpamAsn(asn int) (*IpamAsn, error) {
	formatedURI := fmt.Sprintf("%s/%d", IpamAsnURI, asn)
	resp := &IpamAsn{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing ipam asns by ID
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_asn_get_list
func (c *PFClient) ReadIpamAsns() ([]IpamAsn, error) {
	resp := make([]IpamAsn, 0)
	_, err := c.sendRequest(IpamAsnURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to update an existing IPAM Contact
// This operation is not a supported use case
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_asn_patch
func (c *PFClient) UpdateIpamAsn(ipamAsn IpamAsn) (*IpamAsn, error) {
	return nil, fmt.Errorf("IPAM ASNs Update Unsupported operation")
}

// This function represents the Action to Delete an existing ipam asns
// https://docs.packetfabric.net/openapi/index.html#/IPAM/ipam_asn_delete
func (c *PFClient) DeleteIpamAsn(asn int) (*IpamAsnDeleteResponse, error) {
	if asn == 0 {
		return nil, errors.New("IPAM ASN must be non-zero")
	}
	formatedURI := fmt.Sprintf("%s/%d", IpamAsnURI, asn)
	expectedResp := &IpamAsnDeleteResponse{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}
