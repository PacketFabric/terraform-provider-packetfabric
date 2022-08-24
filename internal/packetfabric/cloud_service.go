package packetfabric

import (
	"fmt"

	"github.com/google/uuid"
)

const backboneURI = "/v2/services/backbone"
const backDeleteURI = "/v2/services/%s"
const cloudConnDeleteURI = "/v2/services/cloud/%s"
const mktProvisionReqURI = "/v2/services/requests/%s/provision/hosted"
const servicesURI = "/v2/services"

type Backbone struct {
	Description  string              `json:"description"`
	Bandwidth    Bandwidth           `json:"bandwidth"`
	Interfaces   []BackBoneInterface `json:"interfaces"`
	RateLimitIn  int                 `json:"rate_limit_in"`
	RateLimitOut int                 `json:"rate_limit_out"`
	Epl          bool                `json:"epl"`
}

type BackBoneInterface struct {
	PortCircuitID string `json:"port_circuit_id"`
	Vlan          int    `json:"vlan"`
	Untagged      bool   `json:"untagged"`
}

type BackboneResp struct {
	VcCircuitID  string    `json:"vc_circuit_id"`
	CustomerUUID string    `json:"customer_uuid"`
	State        string    `json:"state"`
	ServiceType  string    `json:"service_type"`
	ServiceClass string    `json:"service_class"`
	Mode         string    `json:"mode"`
	Connected    bool      `json:"connected"`
	Bandwidth    Bandwidth `json:"bandwidth"`
	Description  string    `json:"description"`
	RateLimitIn  int       `json:"rate_limit_in"`
	RateLimitOut int       `json:"rate_limit_out"`
	TimeCreated  string    `json:"time_created"`
	TimeUpdated  string    `json:"time_updated"`
	Interfaces   []struct {
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
	} `json:"interfaces"`
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

type HostedConnectionResp struct {
	CustomerUUID    string `json:"customer_uuid"`
	UserUUID        string `json:"user_uuid"`
	ServiceProvider string `json:"service_provider"`
	PortType        string `json:"port_type"`
	ServiceClass    string `json:"service_class"`
	Description     string `json:"description"`
	State           string `json:"state"`
	Speed           string `json:"speed"`
}

type BackboneDeleteResp struct {
	Message string `json:"message"`
}

type Services struct {
	VcCircuitID         string             `json:"vc_circuit_id,omitempty"`
	State               string             `json:"state,omitempty"`
	ServiceType         string             `json:"service_type,omitempty"`
	ServiceClass        string             `json:"service_class,omitempty"`
	Mode                string             `json:"mode,omitempty"`
	Connected           bool               `json:"connected,omitempty"`
	Bandwidth           ServiceBandwidth   `json:"bandwidth,omitempty"`
	Description         string             `json:"description,omitempty"`
	TimeCreated         string             `json:"time_created,omitempty"`
	TimeUpdated         string             `json:"time_updated,omitempty"`
	AggregateCapacityID interface{}        `json:"aggregate_capacity_id,omitempty"`
	AccountUUID         string             `json:"account_uuid,omitempty"`
	RateLimitIn         int                `json:"rate_limit_in,omitempty"`
	RateLimitOut        int                `json:"rate_limit_out,omitempty"`
	CustomerUUID        string             `json:"customer_uuid,omitempty"`
	Interfaces          []ServiceInterface `json:"interfaces,omitempty"`
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

func (c *PFClient) CreateBackbone(backbone Backbone) (*BackboneResp, error) {
	backboneResp := &BackboneResp{}
	_, err := c.sendRequest(backboneURI, postMethod, backbone, backboneResp)
	if err != nil {
		return nil, err
	}
	return backboneResp, nil
}

func (c *PFClient) UpdateServiceConn(description, cloudCID string) (*HostedConnectionResp, error) {
	formatedURI := fmt.Sprintf(updateCloudConnURI, cloudCID)
	type UpdateServiceConn struct {
		Description string `json:"description"`
	}
	expectedResp := &HostedConnectionResp{}
	_, err := c.sendRequest(formatedURI, patchMethod, UpdateServiceConn{description}, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
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

func (c *PFClient) DeleteBackbone(vcCircuitID string) (*BackboneDeleteResp, error) {
	return c._deleteService(vcCircuitID, backDeleteURI)
}

func (c *PFClient) DeleteCloudConn(vcCircuitID string) (*BackboneDeleteResp, error) {
	return c._deleteService(vcCircuitID, cloudConnDeleteURI)
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
		return c.GetAwsBackboneState(vcCircuitID)
	}
	go c.CheckServiceStatus(deleteOk, err, fn)
	if !<-deleteOk {
		return nil, err
	}
	return expectedResp, nil
}
