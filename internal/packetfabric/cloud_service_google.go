package packetfabric

import "fmt"

const serviceGoogleMktConnReqURI = "/v2/services/third-party/hosted/google"
const serviceGoogleHostedConnURI = "/v2/services/cloud/hosted/google"
const serviceGoogleDedicatedConnURI = "/v2/services/cloud/dedicated/google"
const servicesGoogleCloudRouterConnURI = "/v2.1/services/cloud-routers/%s/connections/google"

// Struct representation: https://docs.packetfabric.com/api/v2/redoc/#operation/post_google_marketplace_cloud
type GoogleMktCloudConn struct {
	RoutingID                string `json:"routing_id,omitempty"`
	Market                   string `json:"market,omitempty"`
	Description              string `json:"description,omitempty"`
	GooglePairingKey         string `json:"google_pairing_key,omitempty"`
	GoogleVlanAttachmentName string `json:"google_vlan_attachment_name,omitempty"`
	AccountUUID              string `json:"account_uuid,omitempty"`
	Pop                      string `json:"pop,omitempty"`
	Speed                    string `json:"speed,omitempty"`
	ServiceUUID              string `json:"service_uuid,omitempty"`
}

type GoogleCloudRouterConn struct {
	AccountUUID              string `json:"account_uuid,omitempty"`
	MaybeNat                 bool   `json:"maybe_nat,omitempty"`
	GooglePairingKey         string `json:"google_pairing_key,omitempty"`
	GoogleVlanAttachmentName string `json:"google_vlan_attachment_name,omitempty"`
	Description              string `json:"description,omitempty"`
	Pop                      string `json:"pop,omitempty"`
	Speed                    string `json:"speed,omitempty"`
	PublishedQuoteLineUUID   string `json:"published_quote_line_uuid,omitempty"`
}

type GoogleMktCloudConnCreateResp struct {
	VcRequestUUID  string       `json:"vc_request_uuid,omitempty"`
	VcCircuitID    string       `json:"vc_circuit_id,omitempty"`
	FromCustomer   FromCustomer `json:"from_customer,omitempty"`
	ToCustomer     ToCustomer   `json:"to_customer,omitempty"`
	Status         string       `json:"status,omitempty"`
	RequestType    string       `json:"request_type,omitempty"`
	Text           string       `json:"text,omitempty"`
	Bandwidth      Bandwidth    `json:"bandwidth,omitempty"`
	RateLimitIn    int          `json:"rate_limit_in,omitempty"`
	RateLimitOut   int          `json:"rate_limit_out,omitempty"`
	ServiceName    string       `json:"service_name,omitempty"`
	AllowUntaggedZ bool         `json:"allow_untagged_z,omitempty"`
	TimeCreated    string       `json:"time_created,omitempty"`
	TimeUpdated    string       `json:"time_updated,omitempty"`
}

// Struct representation: https://docs.packetfabric.com/api/v2/redoc/#operation/google_hosted_connection_post
type GoogleReqHostedConn struct {
	AccountUUID              string `json:"account_uuid,omitempty"`
	GooglePairingKey         string `json:"google_pairing_key,omitempty"`
	GoogleVlanAttachmentName string `json:"google_vlan_attachment_name,omitempty"`
	Description              string `json:"description,omitempty"`
	Port                     string `json:"port,omitempty"`
	Vlan                     int    `json:"vlan,omitempty"`
	SrcSvlan                 int    `json:"src_svlan,omitempty"`
	Pop                      string `json:"pop,omitempty"`
	Speed                    string `json:"speed,omitempty"`
	PublishedQuoteLineUUID   string `json:"published_quote_line_uuid,omitempty"`
}

// Struct representation: https://docs.packetfabric.com/api/v2/redoc/#operation/google_dedicated_connection_post
type GoogleReqDedicatedConn struct {
	AccountUUID            string `json:"account_uuid,omitempty"`
	Description            string `json:"description,omitempty"`
	Zone                   string `json:"zone,omitempty"`
	Pop                    string `json:"pop,omitempty"`
	SubscriptionTerm       int    `json:"subscription_term,omitempty"`
	ServiceClass           string `json:"service_class,omitempty"`
	Autoneg                bool   `json:"autoneg,omitempty"`
	Speed                  string `json:"speed,omitempty"`
	Loa                    string `json:"loa,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
}

func (c *PFClient) CreateRequestHostedGoogleMktConn(googleConn GoogleMktCloudConn) (*GoogleMktCloudConnCreateResp, error) {
	expectedResp := &GoogleMktCloudConnCreateResp{}
	_, err := c.sendRequest(serviceGoogleMktConnReqURI, postMethod, googleConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) CreateGoogleCloudRouterConn(googleConn GoogleCloudRouterConn, cID string) (*CloudRouterConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(servicesGoogleCloudRouterConnURI, cID)
	expectedResp := &CloudRouterConnectionReadResponse{}
	_, err := c.sendRequest(formatedURI, postMethod, googleConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) CreateRequestHostedGoogleConn(googleConn GoogleReqHostedConn) (*CloudServiceConnCreateResp, error) {
	expectedResp := &CloudServiceConnCreateResp{}
	_, err := c.sendRequest(serviceGoogleHostedConnURI, postMethod, googleConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) CreateRequestDedicatedGoogleConn(googleConn GoogleReqDedicatedConn) (*CloudServiceConnCreateResp, error) {
	expectedResp := &CloudServiceConnCreateResp{}
	_, err := c.sendRequest(serviceGoogleDedicatedConnURI, postMethod, googleConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
