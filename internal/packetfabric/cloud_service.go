package packetfabric

import (
	"fmt"

	"github.com/google/uuid"
)

const backboneURI = "/v2/services/backbone"
const backboneByVCCIDURI = "/v2/services/%s"
const cloudConnDeleteURI = "/v2/services/cloud/%s"
const requestByVCUUIDURI = "/v2/services/requests/%s"
const mktProvisionReqURI = "/v2/services/requests/%s/provision/hosted"
const speedBurstURI = "/v2/services/%s/burst"
const vcRequestsURI = "/v2/services/requests"
const vcSentRequestsURI = "/v2/services/requests?type=%s"
const servicesURI = "/v2/services"
const serviceIxURI = "/v2/services/ix"
const thirdPartyVCURI = "/v2/services/third-party"
const acceptCloudRouterServiceURI = "/v2/services/cloud-routers/requests/%s/accept"
const rejectCloudRouterService = "/v2/services/cloud-routers/requests/%s/reject"

type Backbone struct {
	Description     string              `json:"description"`
	Bandwidth       Bandwidth           `json:"bandwidth"`
	Interfaces      []BackBoneInterface `json:"interfaces"`
	RateLimitIn     int                 `json:"rate_limit_in"`
	RateLimitOut    int                 `json:"rate_limit_out"`
	Epl             bool                `json:"epl"`
	FlexBandwidthID string              `json:"flex_bandwidth_id,omitempty"`
}

type BackBoneInterface struct {
	PortCircuitID string `json:"port_circuit_id"`
	Vlan          int    `json:"vlan,omitempty"`
	Untagged      bool   `json:"untagged,omitempty"`
}

type IxVirtualCircuit struct {
	RoutingID       string     `json:"routing_id,omitempty"`
	Market          string     `json:"market,omitempty"`
	Description     string     `json:"description,omitempty"`
	Asn             int        `json:"asn,omitempty"`
	RateLimitIn     int        `json:"rate_limit_in,omitempty"`
	RateLimitOut    int        `json:"rate_limit_out,omitempty"`
	Bandwidth       Bandwidth  `json:"bandwidth,omitempty"`
	Interface       Interfaces `json:"interface,omitempty"`
	FlexBandwidthID string     `json:"flex_bandwidth_id,omitempty"`
}

type ServiceSettingsUpdate struct {
	RateLimitIn  int          `json:"rate_limit_in,omitempty"`
	RateLimitOut int          `json:"rate_limit_out,omitempty"`
	Description  string       `json:"description,omitempty"`
	Interfaces   []Interfaces `json:"interfaces,omitempty"`
}

type BackboneResp struct {
	VcCircuitID  string               `json:"vc_circuit_id"`
	CustomerUUID string               `json:"customer_uuid"`
	State        string               `json:"state"`
	ServiceType  string               `json:"service_type"`
	ServiceClass string               `json:"service_class"`
	Mode         string               `json:"mode"`
	Connected    bool                 `json:"connected"`
	Bandwidth    Bandwidth            `json:"bandwidth"`
	Description  string               `json:"description"`
	RateLimitIn  int                  `json:"rate_limit_in"`
	RateLimitOut int                  `json:"rate_limit_out"`
	TimeCreated  string               `json:"time_created"`
	TimeUpdated  string               `json:"time_updated"`
	Interfaces   []BackboneInterfResp `json:"interfaces"`
}

type BackboneInterfResp struct {
	PortCircuitID      string `json:"port_circuit_id,omitempty"`
	Pop                string `json:"pop,omitempty"`
	Site               string `json:"site,omitempty"`
	SiteName           string `json:"site_name,omitempty"`
	CustomerSiteCode   string `json:"customer_site_code,omitempty"`
	CustomerSiteName   string `json:"customer_site_name,omitempty"`
	Speed              string `json:"speed,omitempty"`
	Media              string `json:"media,omitempty"`
	Zone               string `json:"zone,omitempty"`
	Description        string `json:"description,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Untagged           bool   `json:"untagged,omitempty"`
	ProvisioningStatus string `json:"provisioning_status,omitempty"`
	AdminStatus        string `json:"admin_status,omitempty"`
	OperationalStatus  string `json:"operational_status,omitempty"`
	CustomerUUID       string `json:"customer_uuid,omitempty"`
	CustomerName       string `json:"customer_name,omitempty"`
	Region             string `json:"region,omitempty"`
	IsCloud            bool   `json:"is_cloud,omitempty"`
	IsPtp              bool   `json:"is_ptp,omitempty"`
	TimeCreated        string `json:"time_created,omitempty"`
	TimeUpdated        string `json:"time_updated,omitempty"`
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

type BackboneDeleteResp struct {
	Message string `json:"message"`
}

type Services struct {
	VcCircuitID     string             `json:"vc_circuit_id,omitempty"`
	State           string             `json:"state,omitempty"`
	ServiceType     string             `json:"service_type,omitempty"`
	ServiceClass    string             `json:"service_class,omitempty"`
	Mode            string             `json:"mode,omitempty"`
	Connected       bool               `json:"connected,omitempty"`
	Bandwidth       ServiceBandwidth   `json:"bandwidth,omitempty"`
	Description     string             `json:"description,omitempty"`
	TimeCreated     string             `json:"time_created,omitempty"`
	TimeUpdated     string             `json:"time_updated,omitempty"`
	FlexBandwidthID string             `json:"flex_bandwidth_id,omitempty"`
	AccountUUID     string             `json:"account_uuid,omitempty"`
	RateLimitIn     int                `json:"rate_limit_in,omitempty"`
	RateLimitOut    int                `json:"rate_limit_out,omitempty"`
	CustomerUUID    string             `json:"customer_uuid,omitempty"`
	Interfaces      []ServiceInterface `json:"interfaces,omitempty"`
}
type ServiceBandwidth struct {
	AccountUUID      string `json:"account_uuid,omitempty"`
	SubscriptionTerm int    `json:"subscription_term,omitempty"`
	Speed            string `json:"speed,omitempty"`
}
type ServiceInterface struct {
	TimeCreated        string `json:"time_created,omitempty"`
	TimeUpdated        string `json:"time_updated,omitempty"`
	PortCircuitID      string `json:"port_circuit_id,omitempty"`
	Pop                string `json:"pop,omitempty"`
	Site               string `json:"site,omitempty"`
	SiteName           string `json:"site_name,omitempty"`
	Speed              string `json:"speed,omitempty"`
	Media              string `json:"media,omitempty"`
	Zone               string `json:"zone,omitempty"`
	Description        string `json:"description,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Untagged           bool   `json:"untagged,omitempty"`
	Svlan              int    `json:"svlan,omitempty"`
	ProvisioningStatus string `json:"provisioning_status,omitempty"`
	AdminStatus        string `json:"admin_status,omitempty"`
	OperationalStatus  string `json:"operational_status,omitempty"`
	CustomerName       string `json:"customer_name,omitempty"`
	CustomerUUID       string `json:"customer_uuid,omitempty"`
	Region             string `json:"region,omitempty"`
	IsCloud            bool   `json:"is_cloud,omitempty"`
	IsPtp              bool   `json:"is_ptp,omitempty"`
}

type VcRequest struct {
	VcRequestUUID   string       `json:"vc_request_uuid,omitempty"`
	VcCircuitID     string       `json:"vc_circuit_id,omitempty"`
	FromCustomer    FromCustomer `json:"from_customer,omitempty"`
	ToCustomer      ToCustomer   `json:"to_customer,omitempty"`
	Status          string       `json:"status,omitempty"`
	RequestType     string       `json:"request_type,omitempty"`
	Text            string       `json:"text,omitempty"`
	Bandwidth       Bandwidth    `json:"bandwidth,omitempty"`
	RateLimitIn     int          `json:"rate_limit_in,omitempty"`
	RateLimitOut    int          `json:"rate_limit_out,omitempty"`
	ServiceName     string       `json:"service_name,omitempty"`
	AllowUntaggedZ  bool         `json:"allow_untagged_z,omitempty"`
	FlexBandwidthID string       `json:"flex_bandwidth_id,omitempty"`
	TimeCreated     string       `json:"time_created,omitempty"`
	TimeUpdated     string       `json:"time_updated,omitempty"`
}

type ThirdPartyVC struct {
	RoutingID       string    `json:"routing_id,omitempty"`
	Market          string    `json:"market,omitempty"`
	Description     string    `json:"description,omitempty"`
	RateLimitIn     int       `json:"rate_limit_in,omitempty"`
	RateLimitOut    int       `json:"rate_limit_out,omitempty"`
	Bandwidth       Bandwidth `json:"bandwidth,omitempty"`
	Interface       Interface `json:"interface,omitempty"`
	ServiceUUID     string    `json:"service_uuid,omitempty"`
	FlexBandwidthID string    `json:"flex_bandwidth_id,omitempty"`
}

type CloudConnInfo struct {
	UUID                    string           `json:"uuid,omitempty"`
	CloudCircuitID          string           `json:"cloud_circuit_id,omitempty"`
	CustomerUUID            string           `json:"customer_uuid,omitempty"`
	AccountUUID             string           `json:"account_uuid,omitempty"`
	UserUUID                string           `json:"user_uuid,omitempty"`
	State                   string           `json:"state,omitempty"`
	ServiceProvider         string           `json:"service_provider,omitempty"`
	ServiceClass            string           `json:"service_class,omitempty"`
	PortType                string           `json:"port_type,omitempty"`
	Speed                   string           `json:"speed,omitempty"`
	Deleted                 bool             `json:"deleted,omitempty"`
	Description             string           `json:"description,omitempty"`
	CloudProvider           AwsCloudProvider `json:"cloud_provider,omitempty"`
	Settings                CloudSettings    `json:"settings,omitempty"`
	SubscriptionTerm        int              `json:"subscription_term,omitempty"`
	TimeCreated             string           `json:"time_created,omitempty"`
	TimeUpdated             string           `json:"time_updated,omitempty"`
	Pop                     string           `json:"pop,omitempty"`
	Site                    string           `json:"site,omitempty"`
	CustomerSiteName        string           `json:"customer_site_name,omitempty"`
	CustomerSiteCode        string           `json:"customer_site_code,omitempty"`
	IsAwaitingOnramp        bool             `json:"is_awaiting_onramp,omitempty"`
	IsCloudRouterConnection bool             `json:"is_cloud_router_connection,omitempty"`
}

// hosted and dedicated cloud
type UpdateServiceConn struct {
	Description string `json:"description"`
}

type ServiceMessage struct {
	Message string `json:"message"`
}

type ServiceRejectionReson struct {
	RejectionReason string `json:"rejection_reason"`
}

func (c *PFClient) CreateBackbone(backbone Backbone) (*BackboneResp, error) {
	backboneResp := &BackboneResp{}
	_, err := c.sendRequest(backboneURI, postMethod, backbone, backboneResp)
	if err != nil {
		return nil, err
	}
	return backboneResp, nil
}

func (c *PFClient) CreateIXVirtualCircuit(ixVc IxVirtualCircuit) (*VcRequest, error) {
	expectedResp := &VcRequest{}
	_, err := c.sendRequest(serviceIxURI, postMethod, ixVc, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// https://docs.packetfabric.com/api/v2/redoc/#tag/Services/operation/post_service_third_party
func (c *PFClient) CreateThirdPartyVC(thirdPartyVC ThirdPartyVC) (*VcRequest, error) {
	expectedResp := &VcRequest{}
	if _, err := c.sendRequest(thirdPartyVCURI, postMethod, thirdPartyVC, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// https://docs.packetfabric.com/api/v2/redoc/#tag/Services/operation/post_service_burst
func (c *PFClient) AddSpeedBurstToCircuit(vcCID, speed string) (*PortMessageResp, error) {
	expectedMsg := &PortMessageResp{}
	formatedURI := fmt.Sprintf(speedBurstURI, vcCID)
	type SpeedBurst struct {
		Speed string `json:"speed"`
	}
	if _, err := c.sendRequest(formatedURI, postMethod, SpeedBurst{Speed: speed}, expectedMsg); err != nil {
		return nil, err
	}
	return expectedMsg, nil
}

func (c *PFClient) GetBackboneByVcCID(vcCID string) (*BackboneResp, error) {
	formatedURI := fmt.Sprintf(backboneByVCCIDURI, vcCID)
	expectedResp := &BackboneResp{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetVCRequest(vcUUID string) (*VcRequest, error) {
	formatedURI := fmt.Sprintf(requestByVCUUIDURI, vcUUID)
	expectedResp := &VcRequest{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) UpdateServiceSettings(vcCID string, serviceSettings ServiceSettingsUpdate) (*BackboneResp, error) {
	formatedURI := fmt.Sprintf(backboneByVCCIDURI, vcCID)
	expectedResp := &BackboneResp{}
	_, err := c.sendRequest(formatedURI, patchMethod, serviceSettings, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) CreateMktProvisionReq(mktProvision ServiceAwsMktConn, vcRequestUUID, provider string) (*MktConnProvisionResp, error) {
	mktProvisionResp := &MktConnProvisionResp{}
	mktProvision.Provider = provider
	formatedURI := fmt.Sprintf(mktProvisionReqURI, vcRequestUUID)
	_, err := c.sendRequest(formatedURI, postMethod, mktProvision, mktProvisionResp)
	if err != nil {
		return nil, err
	}
	return mktProvisionResp, nil
}

func (c *PFClient) GetServices() ([]Services, error) {
	services := make([]Services, 0)
	_, err := c.sendRequest(servicesURI, getMethod, nil, &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (c *PFClient) GetVcRequests() ([]VcRequest, error) {
	return c._getVCRequests(vcRequestsURI)
}

func (c *PFClient) GetVcRequestsByType(reqType string) ([]VcRequest, error) {
	return c._getVCRequests(fmt.Sprintf(vcSentRequestsURI, reqType))
}

func (c *PFClient) _getVCRequests(uri string) ([]VcRequest, error) {
	requests := make([]VcRequest, 0)
	_, err := c.sendRequest(uri, getMethod, nil, &requests)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (c *PFClient) DeleteBackbone(vcCircuitID string) (*BackboneDeleteResp, error) {
	return c._deleteService(vcCircuitID, backboneByVCCIDURI)
}

func (c *PFClient) DeleteCloudConn(vcCircuitID string) (*BackboneDeleteResp, error) {
	return c._deleteService(vcCircuitID, cloudConnDeleteURI)
}

func (c *PFClient) DeleteService(vcCircuitID string) (*ServiceMessage, error) {
	formatedURI := fmt.Sprintf(backboneByVCCIDURI, vcCircuitID)
	expectedResp := &ServiceMessage{}
	if _, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// https://docs.packetfabric.com/api/v2/redoc/#tag/Services/operation/delete_service_burst
func (c *PFClient) DeleteSpeedBurst(vcCID string) (*PortMessageResp, error) {
	formatedURI := fmt.Sprintf(speedBurstURI, vcCID)
	expectedResp := &PortMessageResp{}
	if _, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// https://docs.packetfabric.com/api/v2/swagger/#/Services/delete_service_request
func (c *PFClient) DeleteVCRequest(vcUUID string) (*PortMessageResp, error) {
	formatedURI := fmt.Sprintf(requestByVCUUIDURI, vcUUID)
	type DeleteReason struct {
		DeleteReason string `json:"delete_reason"`
	}
	deleteReason := DeleteReason{
		DeleteReason: "Deleted from PF Terraform plugin.",
	}
	expectedResp := &PortMessageResp{}
	if _, err := c.sendRequest(formatedURI, deleteMethod, deleteReason, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) AcceptCloudRouterService(importCircuitID string) (*PortMessageResp, error) {
	formatedURI := fmt.Sprintf(acceptCloudRouterServiceURI, importCircuitID)
	response := &PortMessageResp{}
	if _, err := c.sendRequest(formatedURI, postMethod, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *PFClient) _deleteService(vcCircuitID, baseURI string) (*BackboneDeleteResp, error) {
	_, uuidParseErr := uuid.Parse(vcCircuitID)
	if uuidParseErr == nil {
		currentServices, servicesErr := c.GetServices()
		if servicesErr != nil {
			return nil, servicesErr
		}
		vcCircuitID = currentServices[0].VcCircuitID
	}
	expectedResp := &BackboneDeleteResp{}
	formatedURI := fmt.Sprintf(baseURI, vcCircuitID)
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	deleteOk := make(chan bool)
	defer close(deleteOk)
	fn := func() (*ServiceState, error) {
		return c.GetBackboneState(vcCircuitID)
	}
	go c.CheckServiceStatus(deleteOk, fn)
	if !<-deleteOk {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) RejectCloudRouterService(importCircuitID, rejectionReason string) (*ServiceRejectionReson, error) {
	formatedURI := fmt.Sprintf(rejectCloudRouterService, importCircuitID)
	rejectionResp := &ServiceRejectionReson{}
	type RejectionReasonMsg struct {
		RejectionReason string `json:"rejection_reason"`
	}
	reason := &RejectionReasonMsg{RejectionReason: rejectionReason}
	_, err := c.sendRequest(formatedURI, postMethod, reason, rejectionResp)
	if err != nil {
		return nil, err
	}
	return rejectionResp, nil
}
