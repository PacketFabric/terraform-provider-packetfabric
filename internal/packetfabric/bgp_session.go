package packetfabric

import (
	"fmt"
	"net/http"
	"time"
)

const bgpSessionURI = "/v2/bgp-settings/%s/prefixes"
const bgpSessionPrefixesURI = "/v2/bgp-settings/%s/prefixes"
const bgpSessionCloudRouterURI = "/v2/services/cloud-routers/%s/connections/%s/bgp"
const bgpSessionSettingsURI = "/v2/bgp-settings"
const bgpSessionSettingsByUUIDURI = "/v2/services/cloud-routers/%s/connections/%s/bgp/%s"

// This struct represents a Bgp Session for an existing Cloud Router connection
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_bgp_create
type BgpSession struct {
	Md5             string      `json:"md5"`
	L3Address       string      `json:"l3_address,omitempty"`
	PrimarySubnet   string      `json:"primary_subnet,omitempty"`
	SecondarySubnet string      `json:"secondary_subnet,omitempty"`
	AddressFamily   string      `json:"address_family"`
	RemoteAddress   string      `json:"remote_address"`
	RemoteAsn       int         `json:"remote_asn"`
	MultihopTTL     int         `json:"multihop_ttl"`
	LocalPreference int         `json:"local_preference,omitempty"`
	Med             int         `json:"med,omitempty"`
	Community       int         `json:"community,omitempty"`
	AsPrepend       int         `json:"as_prepend,omitempty"`
	Orlonger        bool        `json:"orlonger"`
	BfdInterval     int         `json:"bfd_interval,omitempty"`
	BfdMultiplier   int         `json:"bfd_multiplier,omitempty"`
	Disabled        bool        `json:"disabled,omitempty"`
	Prefixes        []BgpPrefix `json:"prefixes,omitempty"`
}

type BgpSessionUpdate struct {
	AddressFamily   string               `json:"address_family"`
	BgpSettingsUUID string               `json:"bgp_settings_uuid"`
	Disabled        bool                 `json:"disabled"`
	MultihopTTL     int                  `json:"multihop_ttl"`
	Orlonger        bool                 `json:"orlonger"`
	RemoteAddress   string               `json:"remote_address"`
	RemoteAsn       int                  `json:"remote_asn"`
	L3Address       string               `json:"l3_address"`
	PrimarySubnet   string               `json:"primary_subnet"`
	SecondarySubnet string               `json:"secondary_subnet"`
	Prefixes        []BgpSessionResponse `json:"prefixes"`
}

type BgpNat struct {
	PreNatSources []interface{} `json:"pre_nat_sources,omitempty"`
	PoolPrefixes  []interface{} `json:"pool_prefixes,omitempty"`
}

type BgpPrefix struct {
	Prefix          string `json:"prefix,omitempty"`
	MatchType       string `json:"match_type,omitempty"`
	AsPrepend       int    `json:"as_prepend,omitempty"`
	Med             int    `json:"med,omitempty"`
	LocalPreference int    `json:"local_preference,omitempty"`
	Type            string `json:"type,omitempty"`
	Order           int    `json:"order,omitempty"`
}

type BgpSessionCreateResp struct {
	BgpSettingsUUID string `json:"bgp_settings_uuid"`
	AddressFamily   string `json:"address_family"`
	RemoteAddress   string `json:"remote_address"`
	RemoteAsn       int    `json:"remote_asn"`
	MultihopTTL     int    `json:"multihop_ttl"`
	LocalPreference int    `json:"local_preference"`
	Community       string `json:"community"`
	AsPrepend       int    `json:"as_prepend"`
	Med             int    `json:"med"`
	Md5             string `json:"md5"`
	Orlonger        bool   `json:"orlonger"`
	BfdInterval     int    `json:"bfd_interval"`
	BfdMultiplier   int    `json:"bfd_multiplier"`
	Disabled        bool   `json:"disabled"`
	Nat             struct {
		PreNatSources []string `json:"pre_nat_sources"`
		PoolPrefixes  []string `json:"pool_prefixes"`
	} `json:"nat"`
	BgpState    string `json:"bgp_state"`
	TimeCreated string `json:"time_created"`
	TimeUpdated string `json:"time_updated"`
}

type BgpSessionBySettingsUUID struct {
	BgpSettingsUUID string      `json:"bgp_settings_uuid"`
	AddressFamily   string      `json:"address_family"`
	RemoteAddress   string      `json:"remote_address"`
	RemoteAsn       int         `json:"remote_asn"`
	MultihopTTL     int         `json:"multihop_ttl"`
	LocalPreference int         `json:"local_preference"`
	Md5             string      `json:"md5"`
	Med             int         `json:"med"`
	L3Address       string      `json:"l3_address,omitempty"`
	PrimarySubnet   string      `json:"primary_subnet,omitempty"`
	SecondarySubnet string      `json:"secondary_subnet,omitempty"`
	Community       interface{} `json:"community"`
	AsPrepend       int         `json:"as_prepend"`
	Orlonger        bool        `json:"orlonger"`
	BfdInterval     int         `json:"bfd_interval"`
	BfdMultiplier   int         `json:"bfd_multiplier"`
	Disabled        bool        `json:"disabled"`
	BgpState        string      `json:"bgp_state"`
	Subnet          interface{} `json:"subnet"`
	PublicIP        string      `json:"public_ip"`
}

// This struct represents a Bgp Session create response
// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_create
type BgpSessionResponse struct {
	BgpPrefixUUID string `json:"bgp_prefix_uuid"`
	Prefix        string `json:"prefix"`
	Type          string `json:"type"`
	Order         int    `json:"order"`
}

type BgpSessionAssociatedResp struct {
	BgpSettingsUUID string `json:"bgp_settings_uuid"`
	AddressFamily   string `json:"address_family"`
	RemoteAddress   string `json:"remote_address"`
	RemoteAsn       int    `json:"remote_asn"`
	MultihopTTL     int    `json:"multihop_ttl"`
	LocalPreference int    `json:"local_preference"`
	Community       string `json:"community"`
	AsPrepend       int    `json:"as_prepend"`
	Med             int    `json:"med"`
	Orlonger        bool   `json:"orlonger"`
	BfdInterval     int    `json:"bfd_interval"`
	BfdMultiplier   int    `json:"bfd_multiplier"`
	Disabled        bool   `json:"disabled"`
	TimeCreated     string `json:"time_created"`
	TimeUpdated     string `json:"time_updated"`
}

type BgpDeleteMessage struct {
	Message string `json:"message"`
}

// This function represents the Action to Create a Bgp Session using an existing Bgp Settigs UUID
// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_create
func (c *PFClient) CreateBgpSession(bgpSession BgpSession, cID, connID string) (*BgpSessionCreateResp, error) {
	formatedURI := fmt.Sprintf(bgpSessionCloudRouterURI, cID, connID)
	expectedResp := &BgpSessionCreateResp{}
	_, err := c.sendRequest(formatedURI, postMethod, bgpSession, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) CreateBgpSessionPrefixes(prefixes []BgpPrefix, bgpSessionUUID string) ([]BgpSessionResponse, error) {
	formatedURI := fmt.Sprintf(bgpSessionPrefixesURI, bgpSessionUUID)
	expectedResp := make([]BgpSessionResponse, 0)
	_, err := c.sendRequest(formatedURI, postMethod, prefixes, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) ReadBgpSessionPrefixes(bgpSettingsUUID string) ([]BgpSessionResponse, error) {
	formatedURI := fmt.Sprintf(bgpSessionPrefixesURI, bgpSettingsUUID)
	expectedResp := make([]BgpSessionResponse, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// This function represents the Action to Retrieve a list of Bgp Sessions by Bgp Settings UUID
// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_list
func (c *PFClient) ReadBgpSession(bgpSetUUID string) ([]BgpSessionResponse, error) {
	formatedURI := fmt.Sprintf(bgpSessionURI, bgpSetUUID)
	expectedResp := make([]BgpSessionResponse, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)

	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetBgpSessionBy(cID, cloudConnID, bgpSettingsUUID string) (*BgpSessionBySettingsUUID, error) {
	formatedURI := fmt.Sprintf(bgpSessionSettingsByUUIDURI, cID, cloudConnID, bgpSettingsUUID)
	expectedResp := &BgpSessionBySettingsUUID{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}

// This function represents the Action to Update a given Cloud Router BGP session
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_bgp_update
func (c *PFClient) UpdateBgpSession(bgpSession BgpSession, cID, connCID string) (*http.Response, *BgpSessionCreateResp, error) {
	formatedURI := fmt.Sprintf(bgpSessionCloudRouterURI, cID, connCID)
	expectedResp := &BgpSessionCreateResp{}
	resp, err := c.sendRequest(formatedURI, putMethod, bgpSession, expectedResp)
	if err != nil {
		return nil, nil, err
	}
	return resp.(*http.Response), expectedResp, err
}

func (c *PFClient) DeleteBgpPrefixes(prefixesUUID []string, bgpSettingsUUID string) ([]BgpSessionResponse, error) {
	formatedURI := fmt.Sprintf(bgpSessionPrefixesURI, bgpSettingsUUID)
	expectedResp := make([]BgpSessionResponse, 0)
	_, err := c.sendRequest(formatedURI, deleteMethod, prefixesUUID, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// This function represents the Action to Delete the list of existing Bgp Sessions by Bgp Settings UUID
// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_delete
func (c *PFClient) DisableBgpSession(bgpSession *BgpSessionUpdate, cID, cloudConnCID string) error {
	formatedURI := fmt.Sprintf(bgpSessionCloudRouterURI, cID, cloudConnCID)
	_, err := c.sendRequest(formatedURI, putMethod, bgpSession, nil)
	// Adding sleep time to avoid concurrent overlay.
	time.Sleep(10 * time.Second)
	if err != nil {
		return err
	}
	return nil
}

// This function represents the Action to Delete a single BGP Session by a Circuit ID,
// Cloud Connection Circuit ID and BGP Settings UUID
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_bgp_delete_by_uuid
func (c *PFClient) DeleteBgpSession(cID, cloudConnCID, bgpSettingsUUID string) (*BgpDeleteMessage, error) {
	formatedURI := fmt.Sprintf(bgpSessionSettingsByUUIDURI, cID, cloudConnCID, bgpSettingsUUID)
	expectedResp := &BgpDeleteMessage{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	// Adding sleep time to avoid concurrent overlay.
	time.Sleep(10 * time.Second)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// This function represents the Action to Return a list of Bgp settings instances associated with the current Account.
// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_session_settings_list
func (c *PFClient) ListBgpSessions() ([]BgpSessionAssociatedResp, error) {
	expectedResp := make([]BgpSessionAssociatedResp, 0)
	_, err := c.sendRequest(bgpSessionSettingsURI, getMethod, nil, &expectedResp)

	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (current *BgpSessionBySettingsUUID) BuildNewBgpSessionInstance() *BgpSessionUpdate {
	return &BgpSessionUpdate{
		AddressFamily: current.AddressFamily,
		Disabled:      current.Disabled,
		MultihopTTL:   current.MultihopTTL,
		Orlonger:      current.Orlonger,
		RemoteAddress: current.RemoteAddress,
		RemoteAsn:     current.RemoteAsn,
	}
}
