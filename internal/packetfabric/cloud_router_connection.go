package packetfabric

import (
	"errors"
	"fmt"
	"time"
)

const awsConnectionURI = "/v2/services/cloud-routers/%s/connections/aws"
const awsConnectionListURI = "/v2/services/cloud-routers/%s/connections"
const cloudRouterConnectionByCidURI = "/v2/services/cloud-routers/%s/connections/%s"
const awsConnectionStatusURI = "/v2.1/services/cloud-routers/%s/connections/%s/status"
const ibmCloudRouterConnectionByCidURI = "/v2.1/services/cloud-routers/%s/connections/ibm"
const ipsecCloudRouterConnectionByCidURI = "/v2/services/cloud-routers/%s/connections/ipsec"
const ipsecConnServiceByCidURI = "/v2/services/ipsec/%s"
const oracleCloudRouterConnectionByCidURI = "/v2/services/cloud-routers/%s/connections/oracle"

type AwsConnection struct {
	AwsAccountID           string `json:"aws_account_id,omitempty"`
	AccountUUID            string `json:"account_uuid,omitempty"`
	MaybeNat               bool   `json:"maybe_nat,omitempty"`
	Description            string `json:"description,omitempty"`
	Pop                    string `json:"pop,omitempty"`
	Zone                   string `json:"zone,omitempty"`
	IsPublic               bool   `json:"is_public,omitempty"`
	Speed                  string `json:"speed,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
}

type AwsConnectionCreateResponse struct {
	PublicIP        string              `json:"public_ip"`
	UUID            string              `json:"uuid"`
	CustomerUUID    string              `json:"customer_uuid"`
	UserUUID        string              `json:"user_uuid"`
	ServiceProvider string              `json:"service_provider"`
	PortType        string              `json:"port_type"`
	Settings        CloudRouterSettings `json:"settings"`
	CloudCircuitID  string              `json:"cloud_circuit_id"`
	AccountUUID     string              `json:"account_uuid"`
	ServiceClass    string              `json:"service_class"`
	Description     string              `json:"description"`
	State           string              `json:"state"`
	Billing         AwsBilling          `json:"billing"`
	Speed           string              `json:"speed"`
	Components      AwsComponents       `json:"components"`
}

type CloudRouterConnectionReadResponse struct {
	PortType                  string           `json:"port_type,omitempty"`
	PortCircuitID             string           `json:"port_circuit_id,omitempty"`
	PendingDelete             bool             `json:"pending_delete,omitempty"`
	State                     string           `json:"state,omitempty"`
	CloudCircuitID            string           `json:"cloud_circuit_id,omitempty"`
	Speed                     string           `json:"speed,omitempty"`
	Deleted                   bool             `json:"deleted,omitempty"`
	AccountUUID               string           `json:"account_uuid,omitempty"`
	ServiceClass              string           `json:"service_class,omitempty"`
	ServiceProvider           string           `json:"service_provider,omitempty"`
	ServiceType               string           `json:"service_type,omitempty"`
	Description               string           `json:"description,omitempty"`
	UUID                      string           `json:"uuid,omitempty"`
	CloudProviderConnectionID string           `json:"cloud_provider_connection_id,omitempty"`
	CloudSettings             CloudSettings    `json:"cloud_settings,omitempty"`
	NatCapable                bool             `json:"nat_capable,omitempty"`
	BgpState                  interface{}      `json:"bgp_state,omitempty"`
	BgpStateList              []BgpStateObj    `json:"bgp_state_list,omitempty"`
	CloudRouterName           string           `json:"cloud_router_name,omitempty"`
	CloudRouterASN            int              `json:"cloud_router_asn,omitempty"`
	CloudRouterCircuitID      string           `json:"cloud_router_circuit_id,omitempty"`
	ConnectionType            string           `json:"connection_type,omitempty"`
	UserUUID                  string           `json:"user_uuid,omitempty"`
	CustomerUUID              string           `json:"customer_uuid,omitempty"`
	TimeCreated               string           `json:"time_created,omitempty"`
	TimeUpdated               string           `json:"time_updated,omitempty"`
	CloudProvider             AwsCloudProvider `json:"cloud_provider,omitempty"`
	Pop                       string           `json:"pop,omitempty"`
	Site                      string           `json:"site,omitempty"`
}

type BgpStateObj struct {
	BgpSettingsUUID string `json:"bgp_settings_uuid,omitempty"`
	BgpState        string `json:"bgp_state,omitempty"`
}

type CloudSettings struct {
	VlanIDPf                 int    `json:"vlan_id_pf,omitempty"`
	VlanIDCust               int    `json:"vlan_id_cust,omitempty"`
	SvlanIDCust              int    `json:"svlan_id_cust,omitempty"`
	AwsRegion                string `json:"aws_region,omitempty"`
	AwsHostedType            string `json:"aws_hosted_type,omitempty"`
	AwsAccountID             string `json:"aws_account_id,omitempty"`
	AwsConnectionID          string `json:"aws_connection_id,omitempty"`
	GooglePairingKey         string `json:"google_pairing_key,omitempty"`
	GoogleVlanAttachmentName string `json:"google_vlan_attachment_name,omitempty"`
	VlanPrivate              int    `json:"vlan_id_private,omitempty"`
	VlanMicrosoft            int    `json:"vlan_id_microsoft,omitempty"`
	AzureServiceKey          string `json:"azure_service_key,omitempty"`
	AzureServiceTag          int    `json:"azure_service_tag,omitempty"`
	AzureConnectionType      string `json:"azure_connection_type,omitempty"`
	OracleRegion             string `json:"oracle_region,omitempty"`
	VcOcid                   string `json:"vc_ocid,omitempty"`
	PortCrossConnectOcid     string `json:"port_cross_connect_ocid,omitempty"`
	PortCompartmentOcid      string `json:"port_compartment_ocid,omitempty"`
	AccountID                string `json:"account_id,omitempty"`
	GatewayID                string `json:"gateway_id,omitempty"`
	PortID                   string `json:"port_id,omitempty"`
	Name                     string `json:"name,omitempty"`
	BgpAsn                   int    `json:"bgp_asn,omitempty"`
	BgpCerCidr               string `json:"bgp_cer_cidr,omitempty"`
	BgpIbmCidr               string `json:"bgp_ibm_cidr,omitempty"`
	PublicIP                 string `json:"public_ip,omitempty"`
	NatPublicIP              string `json:"nat_public_ip,omitempty"`
}
type AwsCloudProvider struct {
	Pop  string `json:"pop,omitempty"`
	Site string `json:"site,omitempty"`
}

type CloudRouterSettings struct {
	AwsRegion                string `json:"aws_region,omitempty"`
	AwsHostedType            string `json:"aws_hosted_type,omitempty"`
	AwsConnectionID          string `json:"aws_connection_id,omitempty"`
	AwsAccountID             string `json:"aws_account_id,omitempty"`
	VlanIDPf                 int    `json:"vlan_id_pf,omitempty"`
	VlanIDCust               int    `json:"vlan_id_cust,omitempty"`
	SvlanIDCust              int    `json:"svlan_id_cust,omitempty"`
	GooglePairingKey         string `json:"google_pairing_key,omitempty"`
	GoogleVlanAttachmentName string `json:"google_vlan_attachment_name,omitempty"`
	VlanPrivate              int    `json:"vlan_private,omitempty"`
	VlanMicrosoft            int    `json:"vlan_microsoft,omitempty"`
	AzureServiceKey          string `json:"azure_service_key,omitempty"`
	AzureServiceTag          int    `json:"azure_service_tag,omitempty"`
	OracleRegion             string `json:"oracle_region,omitempty"`
	VcOcid                   string `json:"vc_ocid,omitempty"`
	PortCrossConnectOcid     string `json:"port_cross_connect_ocid,omitempty"`
	PortCompartmentOcid      string `json:"port_compartment_ocid,omitempty"`
	AccountID                string `json:"account_id,omitempty"`
	GatewayID                string `json:"gateway_id,omitempty"`
	PortID                   string `json:"port_id,omitempty"`
	Name                     string `json:"name,omitempty"`
	BgpAsn                   int    `json:"bgp_asn,omitempty"`
	BgpCerCidr               string `json:"bgp_cer_cidr,omitempty"`
	BgpIbmCidr               string `json:"bgp_ibm_cidr,omitempty"`
}

type AwsBilling struct {
	AccountUUID      string `json:"account_uuid"`
	SubscriptionTerm int    `json:"subscription_term"`
}

type AwsComponents struct {
	IfdPortCircuitIDCust string `json:"ifd_port_circuit_id_cust"`
	IfdPortCircuitIDPf   string `json:"ifd_port_circuit_id_pf"`
}

type DescriptionUpdate struct {
	Description string `json:"description"`
}

type ConnectionDeleteResp struct {
	Message string `json:"message"`
}

type IBMCloudRouterConn struct {
	MaybeNat               bool   `json:"maybe_nat,omitempty"`
	IbmAccountID           string `json:"ibm_account_id,omitempty"`
	IbmBgpAsn              int    `json:"ibm_bgp_asn,omitempty"`
	IbmBgpCerCidr          string `json:"ibm_bgp_cer_cidr,omitempty"`
	IbmBgpIbmCidr          string `json:"ibm_bgp_ibm_cidr,omitempty"`
	Description            string `json:"description,omitempty"`
	AccountUUID            string `json:"account_uuid,omitempty"`
	Pop                    string `json:"pop,omitempty"`
	Zone                   string `json:"zone,omitempty"`
	Speed                  string `json:"speed,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
}

type IPSecRouterConn struct {
	Description                string `json:"description,omitempty"`
	AccountUUID                string `json:"account_uuid,omitempty"`
	Pop                        string `json:"pop,omitempty"`
	Speed                      string `json:"speed,omitempty"`
	IkeVersion                 int    `json:"ike_version,omitempty"`
	Phase1AuthenticationMethod string `json:"phase1_authentication_method,omitempty"`
	Phase1Group                string `json:"phase1_group,omitempty"`
	Phase1EncryptionAlgo       string `json:"phase1_encryption_algo,omitempty"`
	Phase1AuthenticationAlgo   string `json:"phase1_authentication_algo,omitempty"`
	Phase1Lifetime             int    `json:"phase1_lifetime,omitempty"`
	Phase2PfsGroup             string `json:"phase2_pfs_group,omitempty"`
	Phase2EncryptionAlgo       string `json:"phase2_encryption_algo,omitempty"`
	Phase2AuthenticationAlgo   string `json:"phase2_authentication_algo,omitempty"`
	Phase2Lifetime             int    `json:"phase2_lifetime,omitempty"`
	GatewayAddress             string `json:"gateway_address,omitempty"`
	SharedKey                  string `json:"shared_key,omitempty"`
	PublishedQuoteLineUUID     string `json:"published_quote_line_uuid,omitempty"`
}

type OracleCloudRouterConn struct {
	MaybeNat               bool   `json:"maybe_nat,omitempty"`
	VcOcid                 string `json:"vc_ocid,omitempty"`
	Region                 string `json:"region,omitempty"`
	Description            string `json:"description,omitempty"`
	AccountUUID            string `json:"account_uuid,omitempty"`
	Pop                    string `json:"pop,omitempty"`
	Zone                   string `json:"zone,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
}

type IPSecConnUpdate struct {
	CustomerGatewayAddress     string `json:"customer_gateway_address,omitempty"`
	IkeVersion                 int    `json:"ike_version,omitempty"`
	Phase1AuthenticationMethod string `json:"phase1_authentication_method,omitempty"`
	Phase1Group                string `json:"phase1_group,omitempty"`
	Phase1EncryptionAlgo       string `json:"phase1_encryption_algo,omitempty"`
	Phase1AuthenticationAlgo   string `json:"phase1_authentication_algo,omitempty"`
	Phase1Lifetime             int    `json:"phase1_lifetime,omitempty"`
	Phase2PfsGroup             string `json:"phase2_pfs_group,omitempty"`
	Phase2EncryptionAlgo       string `json:"phase2_encryption_algo,omitempty"`
	Phase2AuthenticationAlgo   string `json:"phase2_authentication_algo,omitempty"`
	Phase2Lifetime             int    `json:"phase2_lifetime,omitempty"`
	PreSharedKey               string `json:"pre_shared_key,omitempty"`
}

type IPSecConnUpdateResponse struct {
	CircuitID                  string `json:"circuit_id,omitempty"`
	CustomerGatewayAddress     string `json:"customer_gateway_address,omitempty"`
	LocalGatewayAddress        string `json:"local_gateway_address,omitempty"`
	IkeVersion                 int    `json:"ike_version,omitempty"`
	Phase1AuthenticationMethod string `json:"phase1_authentication_method,omitempty"`
	Phase1Group                string `json:"phase1_group,omitempty"`
	Phase1EncryptionAlgo       string `json:"phase1_encryption_algo,omitempty"`
	Phase1AuthenticationAlgo   string `json:"phase1_authentication_algo,omitempty"`
	Phase1Lifetime             int    `json:"phase1_lifetime,omitempty"`
	Phase2PfsGroup             string `json:"phase2_pfs_group,omitempty"`
	Phase2EncryptionAlgo       string `json:"phase2_encryption_algo,omitempty"`
	Phase2AuthenticationAlgo   string `json:"phase2_authentication_algo,omitempty"`
	Phase2Lifetime             int    `json:"phase2_lifetime,omitempty"`
	PreSharedKey               string `json:"pre_shared_key,omitempty"`
	Deleted                    bool   `json:"deleted,omitempty"`
	TimeCreated                string `json:"time_created,omitempty"`
	TimeUpdated                string `json:"time_updated,omitempty"`
}

type IPSecCloudRouterCreateResp struct {
	VcCircuitID                string `json:"vc_circuit_id,omitempty"`
	CircuitID                  string `json:"circuit_id,omitempty"`
	CustomerGatewayAddress     string `json:"customer_gateway_address,omitempty"`
	LocalGatewayAddress        string `json:"local_gateway_address,omitempty"`
	IkeVersion                 int    `json:"ike_version,omitempty"`
	Phase1AuthenticationMethod string `json:"phase1_authentication_method,omitempty"`
	Phase1Group                string `json:"phase1_group,omitempty"`
	Phase1EncryptionAlgo       string `json:"phase1_encryption_algo,omitempty"`
	Phase1AuthenticationAlgo   string `json:"phase1_authentication_algo,omitempty"`
	Phase1Lifetime             int    `json:"phase1_lifetime,omitempty"`
	Phase2PfsGroup             string `json:"phase2_pfs_group,omitempty"`
	Phase2EncryptionAlgo       string `json:"phase2_encryption_algo,omitempty"`
	Phase2AuthenticationAlgo   string `json:"phase2_authentication_algo,omitempty"`
	Phase2Lifetime             int    `json:"phase2_lifetime,omitempty"`
	PreSharedKey               string `json:"pre_shared_key,omitempty"`
	TimeCreated                string `json:"time_created,omitempty"`
	TimeUpdated                string `json:"time_updated,omitempty"`
	Description                string `json:"description,omitempty"`
	AccountUUID                string `json:"account_uuid,omitempty"`
	Pop                        string `json:"pop,omitempty"`
	Speed                      string `json:"speed,omitempty"`
}

func (c *PFClient) CreateAwsConnection(connection AwsConnection, circuitId string) (*AwsConnectionCreateResponse, error) {
	formatedURI := fmt.Sprintf(awsConnectionURI, circuitId)

	resp := &AwsConnectionCreateResponse{}
	_, err := c.sendRequest(formatedURI, postMethod, connection, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *PFClient) CreateIBMCloudRouteConn(ibmRouter IBMCloudRouterConn, circuitID string) (*CloudRouterConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(ibmCloudRouterConnectionByCidURI, circuitID)

	resp := &CloudRouterConnectionReadResponse{}
	_, err := c.sendRequest(formatedURI, postMethod, ibmRouter, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *PFClient) CreateIPSecCloudRouerConnection(iPSecRouter IPSecRouterConn, circuitID string) (*IPSecCloudRouterCreateResp, error) {
	formatedURI := fmt.Sprintf(ipsecCloudRouterConnectionByCidURI, circuitID)

	resp := &IPSecCloudRouterCreateResp{}
	_, err := c.sendRequest(formatedURI, postMethod, iPSecRouter, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *PFClient) CreateOracleCloudRouerConnection(oracleRouter OracleCloudRouterConn, circuitID string) (*CloudRouterConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(oracleCloudRouterConnectionByCidURI, circuitID)

	resp := &CloudRouterConnectionReadResponse{}
	_, err := c.sendRequest(formatedURI, postMethod, oracleRouter, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *PFClient) ReadAwsConnection(cID, connCid string) (*CloudRouterConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(cloudRouterConnectionByCidURI, cID, connCid)

	resp := &CloudRouterConnectionReadResponse{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) UpdateCloudRouterConnection(cID, connCid string, description DescriptionUpdate) (*CloudRouterConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(cloudRouterConnectionByCidURI, cID, connCid)

	resp := &CloudRouterConnectionReadResponse{}
	_, err := c.sendRequest(formatedURI, patchMethod, description, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) UpdateIPSecConnection(cID string, ipSecUpdate IPSecConnUpdate) (*IPSecConnUpdateResponse, error) {
	formatedURI := fmt.Sprintf(ipsecConnServiceByCidURI, cID)

	resp := &IPSecConnUpdateResponse{}
	_, err := c.sendRequest(formatedURI, patchMethod, ipSecUpdate, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) DeleteCloudRouterConnection(cID, connCid string) (*ConnectionDeleteResp, error) {
	formatedURI := fmt.Sprintf(cloudRouterConnectionByCidURI, cID, connCid)
	if cID == "" {
		return nil, errors.New(errorMsg)
	}

	routerConn, _ := c.ReadAwsConnection(cID, connCid)
	if routerConn == nil {
		return &ConnectionDeleteResp{Message: fmt.Sprintf("No cloud router connection to delete for %s", cID)}, nil
	}

	expectedResp := &ConnectionDeleteResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	// Upon requested on issue #157
	time.Sleep(20 * time.Second)
	return expectedResp, nil
}

func (c *PFClient) GetCloudConnectionStatus(cID, connCID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(awsConnectionStatusURI, cID, connCID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil

}

func (c *PFClient) GetIpsecSpecificConn(cID string) (*IPSecConnUpdateResponse, error) {
	formatedURI := fmt.Sprintf(ipsecConnServiceByCidURI, cID)
	expectedResp := &IPSecConnUpdateResponse{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) ListAwsRouterConnections(cID string) ([]CloudRouterConnectionReadResponse, error) {
	formatedURI := fmt.Sprintf(awsConnectionListURI, cID)
	resp := make([]CloudRouterConnectionReadResponse, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if len(resp) == 0 {
		return resp, nil
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
