package packetfabric

import "fmt"

const azureBackboneURI = "/v2/services/backbone"
const azureHostedMktReqURI = "/v2/services/third-party/hosted/azure"
const azureMktProvisionReqURI = "/v2/services/request/%s/provision/hosted"
const azureExpressRouteURI = "/v2/services/cloud/hosted/azure"
const azureExpressRouteDedicatedURI = "/v2/services/cloud/dedicated/azure"

// Struct representation: https://docs.packetfabric.com/api/v2/redoc/#operation/post_service_backbone
type AzureBackbone struct {
	Description  string                    `json:"description,omitempty"`
	Bandwidth    AzureBackboneBandwidth    `json:"bandwidth,omitempty"`
	Interfaces   []AzureBackboneInterfaces `json:"interfaces,omitempty"`
	RateLimitIn  int                       `json:"rate_limit_in,omitempty"`
	RateLimitOut int                       `json:"rate_limit_out,omitempty"`
	Epl          bool                      `json:"epl,omitempty"`
}
type AzureBackboneBandwidth struct {
	AccountUUID      string `json:"account_uuid,omitempty"`
	SubscriptionTerm int    `json:"subscription_term,omitempty"`
	LonghaulType     string `json:"longhaul_type,omitempty"`
	Speed            string `json:"speed,omitempty"`
}
type AzureBackboneInterfaces struct {
	PortCircuitID string `json:"port_circuit_id,omitempty"`
	Vlan          int    `json:"vlan,omitempty"`
	Untagged      bool   `json:"untagged,omitempty"`
}

type AzureBackboneCreateResp struct {
	VcCircuitID  string                      `json:"vc_circuit_id,omitempty"`
	CustomerUUID string                      `json:"customer_uuid,omitempty"`
	State        string                      `json:"state,omitempty"`
	ServiceType  string                      `json:"service_type,omitempty"`
	ServiceClass string                      `json:"service_class,omitempty"`
	Mode         string                      `json:"mode,omitempty"`
	Connected    bool                        `json:"connected,omitempty"`
	Bandwidth    AzureBackboneBandwidth      `json:"bandwidth,omitempty"`
	Description  string                      `json:"description,omitempty"`
	RateLimitIn  int                         `json:"rate_limit_in,omitempty"`
	RateLimitOut int                         `json:"rate_limit_out,omitempty"`
	TimeCreated  string                      `json:"time_created,omitempty"`
	TimeUpdated  string                      `json:"time_updated,omitempty"`
	Interfaces   []AzureInterfacesCreateResp `json:"interfaces,omitempty"`
}

type AzureInterfacesCreateResp struct {
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

// Struct representation: https://docs.packetfabric.com/api/v2/redoc/#operation/post_aws_marketplace_cloud
type AzureHostedMktReq struct {
	RoutingID       string `json:"routing_id,omitempty"`
	Market          string `json:"market,omitempty"`
	Description     string `json:"description,omitempty"`
	AzureServiceKey string `json:"azure_service_key,omitempty"`
	AccountUUID     string `json:"account_uuid,omitempty"`
	Zone            string `json:"zone,omitempty"`
	Speed           string `json:"speed,omitempty"`
	ServiceUUID     string `json:"service_uuid,omitempty"`
}

type AzureHostedMktReqResp struct {
	VcRequestUUID  string                  `json:"vc_request_uuid,omitempty"`
	VcCircuitID    string                  `json:"vc_circuit_id,omitempty"`
	FromCustomer   AzureFromCustomer       `json:"from_customer,omitempty"`
	ToCustomer     AzureToCustomer         `json:"to_customer,omitempty"`
	Status         string                  `json:"status,omitempty"`
	RequestType    string                  `json:"request_type,omitempty"`
	Text           string                  `json:"text,omitempty"`
	Bandwidth      AzureHostedMktBandwidth `json:"bandwidth,omitempty"`
	RateLimitIn    int                     `json:"rate_limit_in,omitempty"`
	RateLimitOut   int                     `json:"rate_limit_out,omitempty"`
	ServiceName    string                  `json:"service_name,omitempty"`
	AllowUntaggedZ bool                    `json:"allow_untagged_z,omitempty"`
	TimeCreated    string                  `json:"time_created,omitempty"`
	TimeUpdated    string                  `json:"time_updated,omitempty"`
}
type AzureFromCustomer struct {
	CustomerUUID     string `json:"customer_uuid,omitempty"`
	Name             string `json:"name,omitempty"`
	ContactFirstName string `json:"contact_first_name,omitempty"`
	ContactLastName  string `json:"contact_last_name,omitempty"`
	ContactEmail     string `json:"contact_email,omitempty"`
	ContactPhone     string `json:"contact_phone,omitempty"`
}
type AzureToCustomer struct {
	CustomerUUID string `json:"customer_uuid,omitempty"`
	Name         string `json:"name,omitempty"`
}
type AzureHostedMktBandwidth struct {
	AccountUUID      string `json:"account_uuid,omitempty"`
	SubscriptionTerm int    `json:"subscription_term,omitempty"`
	LonghaulType     string `json:"longhaul_type,omitempty"`
	Speed            string `json:"speed,omitempty"`
}

// Struct representation: https://docs.packetfabric.com/api/v2/redoc/#operation/provision_marketplace_cloud
type AzureProvisionMktReq struct {
	Provider    string                  `json:"provider,omitempty"`
	Interface   AzureProvisionInterface `json:"interface,omitempty"`
	Description string                  `json:"description,omitempty"`
}
type AzureProvisionInterface struct {
	PortCircuitID string `json:"port_circuit_id,omitempty"`
	Vlan          int    `json:"vlan,omitempty"`
}

type AzureProvisionMktReqResp struct {
	VcCircuitID  string                      `json:"vc_circuit_id,omitempty"`
	CustomerUUID string                      `json:"customer_uuid,omitempty"`
	State        string                      `json:"state,omitempty"`
	ServiceType  string                      `json:"service_type,omitempty"`
	ServiceClass string                      `json:"service_class,omitempty"`
	Mode         string                      `json:"mode,omitempty"`
	Connected    bool                        `json:"connected,omitempty"`
	Description  string                      `json:"description,omitempty"`
	RateLimitIn  int                         `json:"rate_limit_in,omitempty"`
	RateLimitOut int                         `json:"rate_limit_out,omitempty"`
	TimeCreated  string                      `json:"time_created,omitempty"`
	TimeUpdated  string                      `json:"time_updated,omitempty"`
	Interfaces   []AzureInterfacesCreateResp `json:"interfaces,omitempty"`
}

// Struct representation: https://docs.packetfabric.com/api/v2/redoc/#operation/azure_hosted_connection_post
type AzureExpressRoute struct {
	AzureServiceKey        string `json:"azure_service_key,omitempty"`
	AccountUUID            string `json:"account_uuid,omitempty"`
	Description            string `json:"description,omitempty"`
	Port                   string `json:"port,omitempty"`
	VlanPrivate            int    `json:"vlan_private,omitempty"`
	VlanMicrosoft          int    `json:"vlan_microsoft,omitempty"`
	SrcSvlan               int    `json:"src_svlan,omitempty"`
	Speed                  string `json:"speed,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
}

type AzureExpressRouteCreateResp struct {
	UUID            string          `json:"uuid,omitempty"`
	CustomerUUID    string          `json:"customer_uuid,omitempty"`
	UserUUID        string          `json:"user_uuid,omitempty"`
	ServiceProvider string          `json:"service_provider,omitempty"`
	PortType        string          `json:"port_type,omitempty"`
	Speed           string          `json:"speed,omitempty"`
	ServiceClass    string          `json:"service_class,omitempty"`
	Description     string          `json:"description,omitempty"`
	State           string          `json:"state,omitempty"`
	CloudCircuitID  interface{}     `json:"cloud_circuit_id,omitempty"`
	Billing         AzureBilling    `json:"billing,omitempty"`
	Components      AzureComponents `json:"components,omitempty"`
	Settings        AzureSettings   `json:"settings,omitempty"`
}
type AzureBilling struct {
	AccountUUID string `json:"account_uuid,omitempty"`
}
type AzureComponents struct {
	IfdPortCircuitIDCust int `json:"ifd_port_circuit_id_cust,omitempty"`
	IfdPortCircuitIDPf   int `json:"ifd_port_circuit_id_pf,omitempty"`
}
type AzureSettings struct {
	VlanIDPrivate   int    `json:"vlan_id_private,omitempty"`
	VlanIDMicrosoft int    `json:"vlan_id_microsoft,omitempty"`
	AzureServiceKey string `json:"azure_service_key,omitempty"`
}

// Struct representation: https://docs.packetfabric.com/api/v2/redoc/#operation/azure_dedicated_connection_post
type AzureExpressRouteDedicated struct {
	AccountUUID            string `json:"account_uuid,omitempty"`
	Description            string `json:"description,omitempty"`
	Zone                   string `json:"zone,omitempty"`
	Pop                    string `json:"pop,omitempty"`
	SubscriptionTerm       int    `json:"subscription_term,omitempty"`
	ServiceClass           string `json:"service_class,omitempty"`
	Speed                  string `json:"speed,omitempty"`
	Loa                    string `json:"loa,omitempty"`
	Encapsulation          string `json:"encapsulation,omitempty"`
	PortCategory           string `json:"port_category,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
}

type AzureExpressRouteDedicatedCreateResp struct {
	CustomerUUID      string               `json:"customer_uuid,omitempty"`
	UserUUID          string               `json:"user_uuid,omitempty"`
	ServiceProvider   string               `json:"service_provider,omitempty"`
	PortType          string               `json:"port_type,omitempty"`
	ServiceClass      string               `json:"service_class,omitempty"`
	Description       string               `json:"description,omitempty"`
	State             string               `json:"state,omitempty"`
	Speed             string               `json:"speed,omitempty"`
	CloudCircuitID    string               `json:"cloud_circuit_id,omitempty"`
	TimeCreated       string               `json:"time_created,omitempty"`
	TimeUpdated       string               `json:"time_updated,omitempty"`
	AzurePortCategory string               `json:"azure_port_category,omitempty"`
	Settings          AzureExpressSettings `json:"settings,omitempty"`
}
type AzureExpressSettings struct {
	ZoneDest      string `json:"zone_dest,omitempty"`
	Encapsulation string `json:"encapsulation,omitempty"`
}

func (c *PFClient) CreateAzureBackbone(azureBackbone AwsBackbone) (*AwsBackboneResp, error) {
	azureBackboneResp := &AwsBackboneResp{}
	_, err := c.sendRequest(azureBackboneURI, postMethod, azureBackbone, azureBackboneResp)
	if err != nil {
		return nil, err
	}
	return azureBackboneResp, nil
}

func (c *PFClient) CreateAzureHostedMktRequest(azureMktReq AzureHostedMktReq) (*AzureHostedMktReqResp, error) {
	azureMktReqResp := &AzureHostedMktReqResp{}
	_, err := c.sendRequest(azureHostedMktReqURI, postMethod, azureMktReq, azureMktReqResp)
	if err != nil {
		return nil, err
	}
	return azureMktReqResp, nil
}

func (c *PFClient) CreateMktProvisionReq(azureMktProvision ServiceAwsMktConn, vcRequestUUID string) (*MktConnProvisionResp, error) {
	azureMktProvisionResp := &MktConnProvisionResp{}
	azureMktProvision.Provider = "azure"
	formatedURI := fmt.Sprintf(azureMktProvisionReqURI, vcRequestUUID)
	_, err := c.sendRequest(formatedURI, postMethod, azureMktProvision, azureMktProvisionResp)
	if err != nil {
		return nil, err
	}
	return azureMktProvisionResp, nil
}

func (c *PFClient) CreateAzureExpressRoute(azureExpressRoute AzureExpressRoute) (*AzureExpressRouteCreateResp, error) {
	expressRouteResp := &AzureExpressRouteCreateResp{}
	formatedURI := fmt.Sprintf(azureExpressRouteURI)
	_, err := c.sendRequest(formatedURI, postMethod, azureExpressRoute, expressRouteResp)
	if err != nil {
		return nil, err
	}
	return expressRouteResp, nil
}

func (c *PFClient) CreateAzureExpressRouteDedicated(azureExpressDedicated AzureExpressRouteDedicated) (*AzureExpressRouteDedicatedCreateResp, error) {
	expressRouteResp := &AzureExpressRouteDedicatedCreateResp{}
	_, err := c.sendRequest(azureExpressRouteDedicatedURI, postMethod, azureExpressDedicated, expressRouteResp)
	if err != nil {
		return nil, err
	}
	return expressRouteResp, nil
}
