package packetfabric

import (
	"fmt"
)

const serviceAwsURI = "/v2/services/third-party/hosted/aws"
const provisionMktConnURI = "/v2/services/requests/%s/provision/hosted"
const hostedConnURI = "/v2/services/cloud/hosted/aws"
const hostedMktService = "/v2/services/requests/%s"
const dedicatedConnURI = "/v2/services/cloud/dedicated/aws"
const updateCloudConnURI = "/v2/services/cloud/hosted/%s"
const cloudServiceStatusURI = "/v2.1/services/cloud/connections/%s/status"
const cloudServicesURI = "/v2/services/cloud/%s"
const cloudConnectionInfoURI = "/v2/services/cloud/connections/%s"
const cloudConnectionInstStatusURI = "/v2/services/cloud/connections/%s/status"
const cloudConnectionInstStatusOptsURI = "/v2/services/cloud/connections/%s/upgrade/options"
const cloudConnectionCurrentCustomersURI = "/v2/services/cloud/connections/hosted"
const cloudConnectionCurrentCustmersDedicatedURI = "/v2/services/cloud/connections/dedicated"
const cloudConnectionHostedRequestsSentURI = "/v2/services/requests?type=%s"

type ServiceAws struct {
	RoutingID    string `json:"routing_id"`
	Market       string `json:"market"`
	Description  string `json:"description"`
	AwsAccountID string `json:"aws_account_id"`
	AccountUUID  string `json:"account_uuid"`
	Pop          string `json:"pop"`
	Speed        string `json:"speed"`
}

// This struct represent the AWS Hosted Backbone Marketplace Cloud connection response
// https://docs.packetfabric.com/api/v2/redoc/#operation/post_aws_marketplace_cloud
type AwsHostedMktResp struct {
	VcRequestUUID  string       `json:"vc_request_uuid,omitempty"`
	FromCustomer   FromCustomer `json:"from_customer,omitempty"`
	ToCustomer     ToCustomer   `json:"to_customer,omitempty"`
	Text           string       `json:"text,omitempty"`
	Status         string       `json:"status,omitempty"`
	VcMode         string       `json:"vc_mode,omitempty"`
	RequestType    string       `json:"request_type,omitempty"`
	Bandwidth      Bandwidth    `json:"bandwidth,omitempty"`
	TimeCreated    string       `json:"time_created,omitempty"`
	TimeUpdated    string       `json:"time_updated,omitempty"`
	AllowUntaggedZ bool         `json:"allow_untagged_z,omitempty"`
}

type FromCustomer struct {
	CustomerUUID      string `json:"customer_uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Market            string `json:"market,omitempty"`
	MarketDescription string `json:"market_description,omitempty"`
	ContactFirstName  string `json:"contact_first_name,omitempty"`
	ContactLastName   string `json:"contact_last_name,omitempty"`
	ContactEmail      string `json:"contact_email,omitempty"`
	ContactPhone      string `json:"contact_phone,omitempty"`
}
type ToCustomer struct {
	CustomerUUID      string `json:"customer_uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Market            string `json:"market,omitempty"`
	MarketDescription string `json:"market_description,omitempty"`
}
type Bandwidth struct {
	AccountUUID      string `json:"account_uuid,omitempty"`
	LonghaulType     string `json:"longhaul_type,omitempty"`
	SubscriptionTerm int    `json:"subscription_term,omitempty"`
	Speed            string `json:"speed,omitempty"`
}

type ServiceAwsMktConn struct {
	Provider    string           `json:"provider"`
	Interface   ServiceAwsInterf `json:"interface"`
	Description string           `json:"description"`
}

type ServiceAwsInterf struct {
	PortCircuitID string `json:"port_circuit_id"`
	Vlan          int    `json:"vlan"`
}

type MktConnProvisionResp struct {
	VcCircuitID  string       `json:"vc_circuit_id"`
	CustomerUUID string       `json:"customer_uuid"`
	State        string       `json:"state"`
	ServiceType  string       `json:"service_type"`
	ServiceClass string       `json:"service_class"`
	Mode         string       `json:"mode"`
	Connected    bool         `json:"connected"`
	Description  string       `json:"description"`
	RateLimitIn  int          `json:"rate_limit_in"`
	RateLimitOut int          `json:"rate_limit_out"`
	TimeCreated  string       `json:"time_created"`
	TimeUpdated  string       `json:"time_updated"`
	Interfaces   []Interfaces `json:"interfaces"`
}

type Interfaces struct {
	PortCircuitID      string `json:"port_circuit_id"`
	Pop                string `json:"pop"`
	Site               string `json:"site"`
	SiteName           string `json:"site_name"`
	Speed              string `json:"speed"`
	Media              string `json:"media"`
	Zone               string `json:"zone"`
	Description        string `json:"description"`
	Vlan               int    `json:"vlan"`
	Untagged           bool   `json:"untagged"`
	ProvisioningStatus string `json:"provisioning_status"`
	AdminStatus        string `json:"admin_status"`
	OperationalStatus  string `json:"operational_status"`
	CustomerUUID       string `json:"customer_uuid"`
	CustomerName       string `json:"customer_name"`
	Region             string `json:"region"`
	IsCloud            bool   `json:"is_cloud"`
	IsPtp              bool   `json:"is_ptp"`
	TimeCreated        string `json:"time_created"`
	TimeUpdated        string `json:"time_updated"`
}

type HostedAwsConnection struct {
	AwsAccountID string `json:"aws_account_id"`
	AccountUUID  string `json:"account_uuid"`
	Description  string `json:"description"`
	Pop          string `json:"pop"`
	Port         string `json:"port"`
	Vlan         int    `json:"vlan"`
	SrcSvlan     int    `json:"src_svlan"`
	Zone         string `json:"zone"`
	Speed        string `json:"speed"`
}

type HostedAwsConnectionResp struct {
	CustomerUUID    string `json:"customer_uuid"`
	UserUUID        string `json:"user_uuid"`
	ServiceProvider string `json:"service_provider"`
	PortType        string `json:"port_type"`
	ServiceClass    string `json:"service_class"`
	Description     string `json:"description"`
	State           string `json:"state"`
	Speed           string `json:"speed"`
}

type DedicatedAwsConn struct {
	AwsRegion        string      `json:"aws_region"`
	AccountUUID      string      `json:"account_uuid"`
	Description      string      `json:"description"`
	Zone             string      `json:"zone"`
	Pop              string      `json:"pop"`
	SubscriptionTerm int         `json:"subscription_term"`
	ServiceClass     string      `json:"service_class"`
	AutoNeg          bool        `json:"autoneg"`
	Speed            string      `json:"speed"`
	ShouldCreateLag  bool        `json:"should_create_lag"`
	Loa              interface{} `json:"load"`
}

type AwsCloudConnInfo struct {
	CloudCircuitID  string `json:"cloud_circuit_id"`
	CustomerUUID    string `json:"customer_uuid"`
	UserUUID        string `json:"user_uuid"`
	State           string `json:"state"`
	ServiceProvider string `json:"service_provider"`
	ServiceClass    string `json:"service_class"`
	PortType        string `json:"port_type"`
	Speed           string `json:"speed"`
	Description     string `json:"description"`
	CloudProvider   struct {
		Pop    string `json:"pop"`
		Region string `json:"region"`
	} `json:"cloud_provider"`
	TimeCreated string `json:"time_created"`
	TimeUpdated string `json:"time_updated"`
	Pop         string `json:"pop"`
	Site        string `json:"site"`
}

type AwsDedicatedConnCreateResp struct {
	CustomerUUID    string `json:"customer_uuid"`
	UserUUID        string `json:"user_uuid"`
	ServiceProvider string `json:"service_provider"`
	PortType        string `json:"port_type"`
	ServiceClass    string `json:"service_class"`
	Description     string `json:"description"`
	State           string `json:"state"`
	Speed           string `json:"speed"`
	CloudCircuitID  string `json:"cloud_circuit_id"`
	TimeCreated     string `json:"time_created"`
	TimeUpdated     string `json:"time_updated"`
}

type AwsDedicatedConnResp struct {
	UUID                    string                  `json:"uuid"`
	CustomerUUID            string                  `json:"customer_uuid"`
	UserUUID                string                  `json:"user_uuid"`
	ServiceProvider         string                  `json:"service_provider"`
	PortType                string                  `json:"port_type"`
	Deleted                 bool                    `json:"deleted"`
	TimeUpdated             string                  `json:"time_updated"`
	TimeCreated             string                  `json:"time_created"`
	CloudCircuitID          string                  `json:"cloud_circuit_id"`
	AccountUUID             string                  `json:"account_uuid"`
	CloudProvider           AwsCloudServiceProvider `json:"cloud_provider"`
	Pop                     string                  `json:"pop"`
	Site                    string                  `json:"site"`
	ServiceClass            string                  `json:"service_class"`
	Description             string                  `json:"description"`
	State                   string                  `json:"state"`
	Settings                AwsCloudServiceSettings `json:"settings"`
	SubscriptionTerm        int                     `json:"subscription_term"`
	IsCloudRouterConnection bool                    `json:"is_cloud_router_connection"`
	Speed                   string                  `json:"speed"`
}

type AwsCloudServiceProvider struct {
	Pop  string `json:"pop"`
	Site string `json:"site"`
}

type CloudProvider struct {
	Pop    string `json:"pop"`
	Region string `json:"region"`
}

type AwsCloudServiceSettings struct {
	AwsRegion string `json:"aws_region"`
	ZoneDest  string `json:"zone_dest"`
	Autoneg   bool   `json:"autoneg"`
}

type CloudConnCurrentCustomers struct {
	IsCloudRouterConnection bool   `json:"is_cloud_router_connection"`
	CloudCircuitID          string `json:"cloud_circuit_id"`
	CustomerUUID            string `json:"customer_uuid"`
	UserUUID                string `json:"user_uuid"`
	State                   string `json:"state"`
	ServiceProvider         string `json:"service_provider"`
	ServiceClass            string `json:"service_class"`
	PortType                string `json:"port_type"`
	Speed                   string `json:"speed"`
	Description             string `json:"description"`
	CloudProvider           struct {
		Pop    string `json:"pop"`
		Region string `json:"region"`
	} `json:"cloud_provider"`
	TimeCreated string       `json:"time_created"`
	TimeUpdated string       `json:"time_updated"`
	Interfaces  []Interfaces `json:"interfaces"`
}

func (c *PFClient) CreateAwsHostedMkt(serviceAws ServiceAws) (*AwsHostedMktResp, error) {
	expectedResp := &AwsHostedMktResp{}
	_, err := c.sendRequest(serviceAwsURI, postMethod, serviceAws, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) CreateAwsProvisionReq(conn ServiceAwsMktConn, vcRequestUUID string) (*MktConnProvisionResp, error) {
	expectedResp := &MktConnProvisionResp{}
	formatedURI := fmt.Sprintf(provisionMktConnURI, vcRequestUUID)
	_, err := c.sendRequest(formatedURI, postMethod, conn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) CreateAwsHostedConn(hostedConn HostedAwsConnection) (*HostedAwsConnectionResp, error) {
	expectedResp := &HostedAwsConnectionResp{}
	_, err := c.sendRequest(hostedConnURI, postMethod, hostedConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) CreateDedicadedAWSConn(dedicatedConn DedicatedAwsConn) (*AwsDedicatedConnCreateResp, error) {
	expectedResp := &AwsDedicatedConnCreateResp{}
	_, err := c.sendRequest(dedicatedConnURI, postMethod, dedicatedConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) UpdateAwsServiceConn(description, cloudCID string) (*HostedAwsConnectionResp, error) {
	formatedURI := fmt.Sprintf(updateCloudConnURI, cloudCID)
	type UpdateServiceConn struct {
		Description string `json:"description"`
	}
	expectedResp := &HostedAwsConnectionResp{}
	_, err := c.sendRequest(formatedURI, patchMethod, UpdateServiceConn{description}, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) GetCloudConnInfo(cID string) (*AwsCloudConnInfo, error) {
	formatedURI := fmt.Sprintf(cloudConnectionInfoURI, cID)
	resp := &AwsCloudConnInfo{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *PFClient) GetCurrentCustomersHosted() ([]CloudConnCurrentCustomers, error) {
	expectedResp := make([]CloudConnCurrentCustomers, 0)
	_, err := c.sendRequest(cloudConnectionCurrentCustomersURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) GetHostedCloudConnRequestsSent() ([]AwsHostedMktResp, error) {
	formatedURI := fmt.Sprintf(cloudConnectionHostedRequestsSentURI, "sent")
	expectedResp := make([]AwsHostedMktResp, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetCurrentCustomersDedicated() ([]AwsDedicatedConnResp, error) {
	expectedResp := make([]AwsDedicatedConnResp, 0)
	_, err := c.sendRequest(cloudConnectionCurrentCustmersDedicatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) DeleteCloudService(cloudCID string) error {
	formatedURI := fmt.Sprintf(cloudServicesURI, cloudCID)
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *PFClient) GetCloudServiceStatus(cloudCID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(cloudServiceStatusURI, cloudCID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) DeleteRequestedHostedMktService(vcRequestUUID string) error {
	formatedURI := fmt.Sprintf(hostedMktService, vcRequestUUID)
	type DeleteReason struct {
		DeleteReason string `json:"delete_reason"`
	}
	reason := DeleteReason{
		DeleteReason: "Delete requested by PacketFabric terraform plugin.",
	}
	_, err := c.sendRequest(formatedURI, deleteMethod, &reason, nil)
	if err != nil {
		return err
	}
	return nil
}
