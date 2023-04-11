package packetfabric

import (
	"fmt"
	"time"
)

const serviceAwsURI = "/v2/services/third-party/hosted/aws"
const provisionMktConnURI = "/v2/services/requests/%s/provision/hosted"
const hostedConnURI = "/v2/services/cloud/hosted/aws"
const hostedMktService = "/v2/services/cloud/%s"
const hostedMktServiceRequestsURI = "/v2/services/requests/%s"
const dedicatedConnURI = "/v2/services/cloud/dedicated/aws"
const updateCloudConnHostedURI = "/v2/services/cloud/hosted/%s"
const updateCloudConnDedicatedURI = "/v2/services/cloud/dedicated/%s"
const cloudServiceStatusURI = "/v2.1/services/cloud/connections/%s/status"
const cloudServicesURI = "/v2/services/cloud/%s"
const cloudConnectionInfoURI = "/v2/services/cloud/connections/%s"
const cloudConnectionCurrentCustomersURI = "/v2/services/cloud/connections/hosted"
const cloudConnectionCurrentCustmersDedicatedURI = "/v2/services/cloud/connections/dedicated"
const cloudConnectionHostedRequestsSentURI = "/v2/services/requests?type=%s"
const routerConfigURI = "/v2/services/cloud/connections/%s/router-config?router_type=%s"

type ServiceAws struct {
	RoutingID    string `json:"routing_id,omitempty"`
	Market       string `json:"market,omitempty"`
	Description  string `json:"description,omitempty"`
	AwsAccountID string `json:"aws_account_id,omitempty"`
	AccountUUID  string `json:"account_uuid,omitempty"`
	Pop          string `json:"pop,omitempty"`
	Zone         string `json:"zone,omitempty"`
	Speed        string `json:"speed,omitempty"`
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

type ServiceAwsMktConn struct {
	Provider    string           `json:"provider"`
	Interface   ServiceAwsInterf `json:"interface"`
	Description string           `json:"description"`
}

type ServiceAwsInterf struct {
	PortCircuitID string `json:"port_circuit_id,omitempty"`
	Vlan          int    `json:"vlan,omitempty"`
	VlanPrivate   int    `json:"vlan_private,omitempty"`
	VlanMicrosoft int    `json:"vlan_microsoft,omitempty"`
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
	PortCircuitID      string `json:"port_circuit_id,omitempty"`
	Pop                string `json:"pop,omitempty"`
	Site               string `json:"site,omitempty"`
	SiteName           string `json:"site_name,omitempty"`
	Speed              string `json:"speed,omitempty"`
	Media              string `json:"media,omitempty"`
	Zone               string `json:"zone,omitempty"`
	Description        string `json:"description,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Svlan              int    `json:"svlan,omitempty"`
	Untagged           bool   `json:"untagged"`
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
	CustomerSiteCode   string `json:"customer_site_code,omitempty"`
	CustomerSiteName   string `json:"customer_site_name,omitempty"`
}

type HostedAwsConnection struct {
	AwsAccountID  string         `json:"aws_account_id,omitempty"`
	AccountUUID   string         `json:"account_uuid,omitempty"`
	Description   string         `json:"description,omitempty"`
	Pop           string         `json:"pop,omitempty"`
	Port          string         `json:"port,omitempty"`
	Vlan          int            `json:"vlan,omitempty"`
	SrcSvlan      int            `json:"src_svlan,omitempty"`
	Zone          string         `json:"zone,omitempty"`
	Speed         string         `json:"speed,omitempty"`
	PONumber      string         `json:"po_number,omitempty"`
	CloudSettings *CloudSettings `json:"cloud_settings,omitempty"`
}

// used for both Hosted Cloud and Cloud Router Connections
type CloudSettings struct {
	AccountID                    string       `json:"account_id,omitempty"`
	AwsAccountID                 string       `json:"aws_account_id,omitempty"`
	AwsConnectionID              string       `json:"aws_connection_id,omitempty"`
	AwsDxAWSDevice               string       `json:"aws_dx_aws_device,omitempty"`
	AwsDxAWSLogicalDeviceID      string       `json:"aws_dx_aws_logical_device_id,omitempty"`
	AwsDxBandwidth               string       `json:"aws_dx_bandwidth,omitempty"`
	AwsDxEncryptionMode          string       `json:"aws_dx_encryption_mode,omitempty"`
	AwsDxHasLogicalRedundancy    bool         `json:"aws_dx_has_logical_redundancy,omitempty"`
	AwsDxJumboFrameCapable       bool         `json:"aws_dx_jumbo_frame_capable,omitempty"`
	AwsDxLocation                string       `json:"aws_dx_location,omitempty"`
	AwsDxMacSecCapable           bool         `json:"aws_dx_mac_sec_capable,omitempty"`
	AwsGateways                  []AwsGateway `json:"aws_gateways,omitempty"`
	AwsHostedType                string       `json:"aws_hosted_type,omitempty"`
	AwsRegion                    string       `json:"aws_region,omitempty"`
	AwsVifBGPPeerID              string       `json:"aws_vif_bgp_peer_id,omitempty"`
	AwsVifDirectConnectGwID      string       `json:"aws_vif_direct_connect_gw_id,omitempty"`
	AwsVifID                     string       `json:"aws_vif_id,omitempty"`
	AwsVifType                   string       `json:"aws_vif_type,omitempty"`
	AzureConnectionType          string       `json:"azure_connection_type,omitempty"`
	AzureServiceKey              string       `json:"azure_service_key,omitempty"`
	AzureServiceTag              int          `json:"azure_service_tag,omitempty"`
	BgpAsn                       int          `json:"bgp_asn,omitempty"`
	BgpCerCidr                   string       `json:"bgp_cer_cidr,omitempty"`
	BgpIbmCidr                   string       `json:"bgp_ibm_cidr,omitempty"`
	BgpSettings                  *BgpSettings `json:"bgp_settings,omitempty"`
	CloudState                   *CloudState  `json:"cloud_state,omitempty"`
	CredentialsUUID              string       `json:"credentials_uuid,omitempty"`
	GatewayID                    string       `json:"gateway_id,omitempty"`
	GoogleCloudRouterName        string       `json:"google_cloud_router_name,omitempty"`
	GoogleDataplaneVersion       int          `json:"google_dataplane_version,omitempty"`
	GoogleEdgeAvailabilityDomain int          `json:"google_edge_availability_domain,omitempty"`
	GoogleInterfaceName          string       `json:"google_interface_name,omitempty"`
	GooglePairingKey             string       `json:"google_pairing_key,omitempty"`
	GoogleProjectID              string       `json:"google_project_id,omitempty"`
	GoogleRegion                 string       `json:"google_region,omitempty"`
	GoogleVPCName                string       `json:"google_vpc_name,omitempty"`
	GoogleVlanAttachmentName     string       `json:"google_vlan_attachment_name,omitempty"`
	Mtu                          int          `json:"mtu,omitempty"`
	Name                         string       `json:"name,omitempty"`
	NatPublicIP                  string       `json:"nat_public_ip,omitempty"`
	OracleRegion                 string       `json:"oracle_region,omitempty"`
	PortCompartmentOcid          string       `json:"port_compartment_ocid,omitempty"`
	PortCrossConnectOcid         string       `json:"port_cross_connect_ocid,omitempty"`
	PortID                       string       `json:"port_id,omitempty"`
	PublicIP                     string       `json:"public_ip,omitempty"`
	SvlanIDCust                  int          `json:"svlan_id_cust,omitempty"`
	VcOcid                       string       `json:"vc_ocid,omitempty"`
	VlanIDCust                   int          `json:"vlan_id_cust,omitempty"`
	VlanIDMicrosoft              int          `json:"vlan_id_microsoft,omitempty"`
	VlanIDPf                     int          `json:"vlan_id_pf,omitempty"`
	VlanPrivate                  int          `json:"vlan_id_private,omitempty"`
}

type CloudState struct {
	AwsDxConnectionState           string `json:"aws_dx_connection_state,omitempty"`
	AwsDxPortEncryptionStatus      string `json:"aws_dx_port_encryption_status,omitempty"`
	AwsVifState                    string `json:"aws_vif_state,omitempty"`
	BgpState                       string `json:"bgp_state,omitempty"`
	GoogleInterconnectState        string `json:"google_interconnect_state,omitempty"`
	GoogleInterconnectAdminEnabled bool   `json:"google_interconnect_admin_enabled,omitempty"`
}

type BgpSettings struct {
	AddressFamily            string      `json:"address_family,omitempty"`
	AdvertisedPrefixes       []string    `json:"advertised_prefixes,omitempty"`
	AsPrepend                int         `json:"as_prepend,omitempty"`
	BfdInterval              int         `json:"bfd_interval,omitempty"`
	BfdMultiplier            int         `json:"bfd_multiplier,omitempty"`
	Community                int         `json:"community,omitempty"`
	CustomerAsn              int         `json:"customer_asn,omitempty"`
	CustomerRouterIp         string      `json:"customer_router_ip,omitempty"`
	Disabled                 bool        `json:"disabled,omitempty"`
	GoogleAdvertiseMode      string      `json:"google_advertise_mode,omitempty"`
	GoogleAdvertisedIPRanges []string    `json:"google_advertised_ip_ranges,omitempty"`
	GoogleKeepaliveInterval  int         `json:"google_keepalive_interval,omitempty"`
	L3Address                string      `json:"l3_address,omitempty"`
	LocalPreference          int         `json:"local_preference,omitempty"`
	Md5                      string      `json:"md5,omitempty"`
	Med                      int         `json:"med,omitempty"`
	MultihopTTL              int         `json:"multihop_ttl,omitempty"`
	Nat                      *BgpNat     `json:"nat,omitempty"`
	Orlonger                 bool        `json:"orlonger,omitempty"`
	Prefixes                 []BgpPrefix `json:"prefixes,omitempty"`
	PrimarySubnet            string      `json:"primary_subnet,omitempty"`
	RemoteAddress            string      `json:"remote_address,omitempty"`
	RemoteAsn                int         `json:"remote_asn,omitempty"`
	RemoteRouterIp           string      `json:"remote_router_ip,omitempty"`
	SecondarySubnet          string      `json:"secondary_subnet,omitempty"`
}

type AwsGateway struct {
	Type            string   `json:"type,omitempty"`
	Name            string   `json:"name,omitempty"`
	ID              string   `json:"id,omitempty"`
	Asn             int      `json:"asn,omitempty"`
	VpcID           string   `json:"vpc_id,omitempty"`
	SubnetIDs       []string `json:"subnet_ids,omitempty"`
	AllowedPrefixes []string `json:"allowed_prefixes,omitempty"`
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
	PONumber         string      `json:"po_number,omitempty"`
}

type CloudServiceConnCreateResp struct {
	UUID                    string      `json:"uuid,omitempty"`
	CustomerUUID            string      `json:"customer_uuid,omitempty"`
	UserUUID                string      `json:"user_uuid,omitempty"`
	ServiceProvider         string      `json:"service_provider,omitempty"`
	PortType                string      `json:"port_type,omitempty"`
	Deleted                 bool        `json:"deleted,omitempty"`
	CloudCircuitID          string      `json:"cloud_circuit_id,omitempty"`
	AccountUUID             string      `json:"account_uuid,omitempty"`
	CustomerSiteCode        interface{} `json:"customer_site_code,omitempty"`
	CustomerSiteName        interface{} `json:"customer_site_name,omitempty"`
	ServiceClass            string      `json:"service_class,omitempty"`
	Description             string      `json:"description,omitempty"`
	State                   string      `json:"state,omitempty"`
	Settings                Settings    `json:"settings,omitempty"`
	Billing                 Billing     `json:"billing,omitempty"`
	Components              Components  `json:"components,omitempty"`
	IsCloudRouterConnection bool        `json:"is_cloud_router_connection,omitempty"`
	IsAwaitingOnramp        bool        `json:"is_awaiting_onramp,omitempty"`
	AzurePortCategory       string      `json:"azure_port_category,omitempty"`
	Speed                   string      `json:"speed,omitempty"`
}
type Settings struct {
	VlanIDPf                 int         `json:"vlan_id_pf,omitempty"`
	VlanIDCust               int         `json:"vlan_id_cust,omitempty"`
	SvlanIDCust              interface{} `json:"svlan_id_cust,omitempty"`
	VlanIDPrivate            int         `json:"vlan_id_private,omitempty"`
	VlanIDMicrosoft          int         `json:"vlan_id_microsoft,omitempty"`
	VcIDPrivate              int         `json:"vc_id_private,omitempty"`
	SvlanIDCustomer          interface{} `json:"svlan_id_customer,omitempty"`
	AzureServiceKey          string      `json:"azure_service_key,omitempty"`
	AzureServiceTag          int         `json:"azure_service_tag,omitempty"`
	AzureEncapsulation       string      `json:"encapsulation,omitempty"`
	GooglePairingKey         string      `json:"google_pairing_key,omitempty"`
	GoogleVlanAttachmentName string      `json:"google_vlan_attachment_name,omitempty"`
	AwsRegion                string      `json:"aws_region,omitempty"`
	AwsHostedType            string      `json:"aws_hosted_type,omitempty"`
	AwsConnectionID          string      `json:"aws_connection_id,omitempty"`
	AwsAccountID             string      `json:"aws_account_id,omitempty"`
	ZoneDest                 string      `json:"zone_dest,omitempty"`
	Autoneg                  bool        `json:"autoneg,omitempty"`
	OracleRegion             string      `json:"oracle_region,omitempty"`
	VcOcid                   string      `json:"vc_ocid,omitempty"`
	PortCrossConnectOcid     string      `json:"port_cross_connect_ocid,omitempty"`
	PortCompartmentOcid      string      `json:"port_compartment_ocid,omitempty"`
	AccountID                string      `json:"account_id,omitempty"`
	GatewayID                string      `json:"gateway_id,omitempty"`
	PortID                   string      `json:"port_id,omitempty"`
	Name                     string      `json:"name,omitempty"`
	BgpAsn                   int         `json:"bgp_asn,omitempty"`
	BgpCerCidr               string      `json:"bgp_cer_cidr,omitempty"`
	BgpIbmCidr               string      `json:"bgp_ibm_cidr,omitempty"`
}
type Billing struct {
	AccountUUID      string `json:"account_uuid,omitempty"`
	SubscriptionTerm int    `json:"subscription_term,omitempty"`
	Speed            string `json:"speed,omitempty"`
	ContractedSpeed  string `json:"contracted_speed,omitempty"`
}
type Components struct {
	IfdPortCircuitIDCust string `json:"ifd_port_circuit_id_cust,omitempty"`
	VcIDMicrosoft        int    `json:"vc_id_microsoft,omitempty"`
	VcIDPrivate          int    `json:"vc_id_private,omitempty"`
}

type DedicatedConnResp struct {
	UUID                    string               `json:"uuid"`
	CustomerUUID            string               `json:"customer_uuid"`
	UserUUID                string               `json:"user_uuid"`
	ServiceProvider         string               `json:"service_provider"`
	PortType                string               `json:"port_type"`
	Deleted                 bool                 `json:"deleted"`
	TimeUpdated             string               `json:"time_updated"`
	TimeCreated             string               `json:"time_created"`
	CloudCircuitID          string               `json:"cloud_circuit_id"`
	AccountUUID             string               `json:"account_uuid"`
	CloudProvider           CloudServiceProvider `json:"cloud_provider"`
	Pop                     string               `json:"pop"`
	Site                    string               `json:"site"`
	ServiceClass            string               `json:"service_class"`
	Description             string               `json:"description"`
	State                   string               `json:"state"`
	Settings                CloudServiceSettings `json:"settings"`
	SubscriptionTerm        int                  `json:"subscription_term"`
	IsCloudRouterConnection bool                 `json:"is_cloud_router_connection"`
	Speed                   string               `json:"speed"`
}

type CloudServiceProvider struct {
	Pop  string `json:"pop"`
	Site string `json:"site"`
}

type CloudServiceSettings struct {
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

type HostedConnResp struct {
	UUID                    string             `json:"uuid,omitempty"`
	CustomerUUID            string             `json:"customer_uuid,omitempty"`
	UserUUID                string             `json:"user_uuid,omitempty"`
	ServiceProvider         string             `json:"service_provider,omitempty"`
	PortType                string             `json:"port_type,omitempty"`
	Deleted                 bool               `json:"deleted,omitempty"`
	TimeUpdated             string             `json:"time_updated,omitempty"`
	TimeCreated             string             `json:"time_created,omitempty"`
	CloudCircuitID          string             `json:"cloud_circuit_id,omitempty"`
	AccountUUID             string             `json:"account_uuid,omitempty"`
	CloudProvider           CloudProvider      `json:"cloud_provider,omitempty"`
	ServiceClass            string             `json:"service_class,omitempty"`
	Description             string             `json:"description,omitempty"`
	State                   string             `json:"state,omitempty"`
	IsCloudRouterConnection bool               `json:"is_cloud_router_connection,omitempty"`
	IsAwaitingOnramp        bool               `json:"is_awaiting_onramp,omitempty"`
	Speed                   string             `json:"speed,omitempty"`
	Interfaces              []HostedInterfaces `json:"interfaces,omitempty"`
}

type HostedInterfaces struct {
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

type RouterConfig struct {
	CloudCircuitID string `json:"cloud_circuit_id"`
	RouterType     string `json:"router_type"`
	RouterConfig   string `json:"router_config"`
}

type CloudConnInfo struct {
	UUID                    string         `json:"uuid,omitempty"`
	CloudCircuitID          string         `json:"cloud_circuit_id,omitempty"`
	CustomerUUID            string         `json:"customer_uuid,omitempty"`
	AccountUUID             string         `json:"account_uuid,omitempty"`
	UserUUID                string         `json:"user_uuid,omitempty"`
	State                   string         `json:"state,omitempty"`
	ServiceProvider         string         `json:"service_provider,omitempty"`
	ServiceClass            string         `json:"service_class,omitempty"`
	PortType                string         `json:"port_type,omitempty"`
	Speed                   string         `json:"speed,omitempty"`
	Deleted                 bool           `json:"deleted,omitempty"`
	Description             string         `json:"description,omitempty"`
	CloudProvider           CloudProvider  `json:"cloud_provider,omitempty"`
	Settings                *Settings      `json:"settings,omitempty"`
	CloudSettings           *CloudSettings `json:"cloud_settings,omitempty"`
	SubscriptionTerm        int            `json:"subscription_term,omitempty"`
	TimeCreated             string         `json:"time_created,omitempty"`
	TimeUpdated             string         `json:"time_updated,omitempty"`
	Pop                     string         `json:"pop,omitempty"`
	Site                    string         `json:"site,omitempty"`
	CustomerSiteName        string         `json:"customer_site_name,omitempty"`
	CustomerSiteCode        string         `json:"customer_site_code,omitempty"`
	IsAwaitingOnramp        bool           `json:"is_awaiting_onramp,omitempty"`
	IsCloudRouterConnection bool           `json:"is_cloud_router_connection,omitempty"`
	AzurePortCategory       string         `json:"azure_port_category,omitempty"`
	PONumber                string         `json:"po_number,omitempty"`
}

type UpdateServiceConn struct {
	Description   string         `json:"description,omitempty"`
	PONumber      string         `json:"po_number,omitempty"`
	CloudSettings *CloudSettings `json:"cloud_settings,omitempty"`
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

func (c *PFClient) CreateAwsHostedConn(hostedConn HostedAwsConnection) (*CloudServiceConnCreateResp, error) {
	expectedResp := &CloudServiceConnCreateResp{}
	_, err := c.sendRequest(hostedConnURI, postMethod, hostedConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) CreateDedicadedAWSConn(dedicatedConn DedicatedAwsConn) (*CloudServiceConnCreateResp, error) {
	expectedResp := &CloudServiceConnCreateResp{}
	_, err := c.sendRequest(dedicatedConnURI, postMethod, dedicatedConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) GetCloudConnInfo(cID string) (*CloudConnInfo, error) {
	formatedURI := fmt.Sprintf(cloudConnectionInfoURI, cID)
	resp := &CloudConnInfo{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *PFClient) GetCurrentCustomersHosted() ([]HostedConnResp, error) {
	expectedResp := make([]HostedConnResp, 0)
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

func (c *PFClient) GetCurrentCustomersDedicated() ([]DedicatedConnResp, error) {
	expectedResp := make([]DedicatedConnResp, 0)
	_, err := c.sendRequest(cloudConnectionCurrentCustmersDedicatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) DeleteCloudService(cloudCID string) error {
	formatedURI := fmt.Sprintf(cloudServicesURI, cloudCID)
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, nil)
	// Upon requested on issue #157
	time.Sleep(30 * time.Second)
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
	return c._deleteMktService(vcRequestUUID, hostedMktService)
}

func (c *PFClient) UpdateServiceHostedConn(cloudCID string, updateServiceConnData UpdateServiceConn) (*CloudServiceConnCreateResp, error) {
	formatedURI := fmt.Sprintf(updateCloudConnHostedURI, cloudCID)
	expectedResp := &CloudServiceConnCreateResp{}
	_, err := c.sendRequest(formatedURI, patchMethod, updateServiceConnData, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) UpdateServiceDedicatedConn(cloudCID string, updateServiceConnData UpdateServiceConn) (*CloudServiceConnCreateResp, error) {
	formatedURI := fmt.Sprintf(updateCloudConnDedicatedURI, cloudCID)
	expectedResp := &CloudServiceConnCreateResp{}
	_, err := c.sendRequest(formatedURI, patchMethod, updateServiceConnData, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

// Status can be [ pending, provisioned, rejected ]
// if rejected or provisioned, we skip the delete, if pending, we do the delete as you already implemented it.
func (c *PFClient) DeleteHostedMktConnection(vcRequestUUID string) (message string, err error) {
	vcReq, e := c.GetVCRequest(vcRequestUUID)
	if e != nil {
		err = e
		return
	}
	if vcReq == nil {
		message = "The Marketplace connection request has been either accepted or rejected."
		return message, err
	}
	if vcReq.Status == "provisioned" {
		message = `The Z side has approved the request and provisioned the connection.
		Please import the new resource created to manage it with 
		Terraform and update your Terraform configuration.`
	}
	if vcReq.Status == "rejected" {
		message = "the Z side has rejected the request. Remove the resource from Terraform state and resubmit your request as needed"
	}
	if vcReq.Status == "pending" {
		err = c._deleteMktService(vcRequestUUID, hostedMktServiceRequestsURI)
	}
	return message, err
}

func (c *PFClient) _deleteMktService(vcRequestUUID, uri string) error {
	type DeleteReason struct {
		DeleteReason string `json:"delete_reason"`
	}
	formatedURI := fmt.Sprintf(uri, vcRequestUUID)
	reason := DeleteReason{
		DeleteReason: "Delete requested by PacketFabric terraform plugin.",
	}
	_, err := c.sendRequest(formatedURI, deleteMethod, &reason, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *PFClient) GetRouterConfiguration(cloudCircuitID, routerType string) (*RouterConfig, error) {
	formattedURI := fmt.Sprintf(routerConfigURI, cloudCircuitID, routerType)
	expectedResp := &RouterConfig{}

	_, err := c.sendRequest(formattedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}

	return expectedResp, nil
}
