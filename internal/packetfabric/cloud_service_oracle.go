package packetfabric

const cloudServiceOracleURI = "/v2/services/third-party/hosted/oracle"
const cloudServiceHostedOracleURI = "/v2/services/cloud/hosted/oracle"

type CloudServiceOracle struct {
	RoutingID   string `json:"routing_id,omitempty"`
	Market      string `json:"market,omitempty"`
	Description string `json:"description,omitempty"`
	AccountUUID string `json:"account_uuid,omitempty"`
	VcOcid      string `json:"vc_ocid,omitempty"`
	Region      string `json:"region,omitempty"`
	Pop         string `json:"pop,omitempty"`
	ServiceUUID string `json:"service_uuid,omitempty"`
}

type CloudServiceOracleConn struct {
	VcOcid                 string `json:"vc_ocid,omitempty"`
	Region                 string `json:"region,omitempty"`
	Description            string `json:"description,omitempty"`
	AccountUUID            string `json:"account_uuid,omitempty"`
	Pop                    string `json:"pop,omitempty"`
	Port                   string `json:"port,omitempty"`
	Zone                   string `json:"zone,omitempty"`
	Vlan                   int    `json:"vlan,omitempty"`
	SrcSvlan               int    `json:"src_svlan,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
}

type CloudServiceOracleConnResp struct {
	UUID                    string           `json:"uuid,omitempty"`
	CustomerUUID            string           `json:"customer_uuid,omitempty"`
	UserUUID                string           `json:"user_uuid,omitempty"`
	ServiceProvider         string           `json:"service_provider,omitempty"`
	PortType                string           `json:"port_type,omitempty"`
	Deleted                 bool             `json:"deleted,omitempty"`
	CloudCircuitID          string           `json:"cloud_circuit_id,omitempty"`
	AccountUUID             string           `json:"account_uuid,omitempty"`
	CustomerSiteName        string           `json:"customer_site_name,omitempty"`
	CustomerSiteCode        string           `json:"customer_site_code,omitempty"`
	ServiceClass            string           `json:"service_class,omitempty"`
	Description             string           `json:"description,omitempty"`
	State                   string           `json:"state,omitempty"`
	Settings                Settings         `json:"settings,omitempty"`
	Billing                 Billing          `json:"billing,omitempty"`
	Components              OracleComponents `json:"components,omitempty"`
	IsCloudRouterConnection bool             `json:"is_cloud_router_connection,omitempty"`
	Speed                   string           `json:"speed,omitempty"`
}

type CrossConnectMappings struct {
	BgpMd5AuthKey                     string `json:"bgp_md5_auth_key,omitempty"`
	CrossConnectOrCrossConnectGroupID string `json:"cross_connect_or_cross_connect_group_id,omitempty"`
	CustomerBgpPeeringIP              string `json:"customer_bgp_peering_ip,omitempty"`
	CustomerBgpPeeringIpv6            string `json:"customer_bgp_peering_ipv6,omitempty"`
	OracleBgpPeeringIP                string `json:"oracle_bgp_peering_ip,omitempty"`
	OracleBgpPeeringIpv6              string `json:"oracle_bgp_peering_ipv6,omitempty"`
	Vlan                              int    `json:"vlan,omitempty"`
}

type CloudProviderProvisioningResponse struct {
	VcOcid               string                 `json:"vc_ocid,omitempty"`
	Bandwidth            string                 `json:"bandwidth,omitempty"`
	BgpManagement        string                 `json:"bgp_management,omitempty"`
	BgpSessionState      string                 `json:"bgp_session_state,omitempty"`
	CompartmentID        string                 `json:"compartment_id,omitempty"`
	CrossConnectMappings []CrossConnectMappings `json:"cross_connect_mappings,omitempty"`
	CustomerAsn          int                    `json:"customer_asn,omitempty"`
	GatewayID            string                 `json:"gateway_id,omitempty"`
	LifecycleState       string                 `json:"lifecycle_state,omitempty"`
	OracleBgpAsn         int                    `json:"oracle_bgp_asn,omitempty"`
	ProviderServiceID    string                 `json:"provider_service_id,omitempty"`
	ProviderState        string                 `json:"provider_state,omitempty"`
	ReferenceComment     string                 `json:"reference_comment,omitempty"`
	ServiceType          string                 `json:"service_type,omitempty"`
	Type                 string                 `json:"type,omitempty"`
}

type OracleComponents struct {
	IfdPortCircuitIDCust              string                            `json:"ifd_port_circuit_id_cust,omitempty"`
	CloudProviderProvisioningResponse CloudProviderProvisioningResponse `json:"cloud_provider_provisioning_response,omitempty"`
}

func (c *PFClient) RequestHostedOracleMktConn(oracleService CloudServiceOracle) (*VcRequest, error) {
	expectedResp := &VcRequest{}
	_, err := c.sendRequest(cloudServiceOracleURI, postMethod, oracleService, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) RequestNewHostedOracleConn(oracleHosted CloudServiceOracleConn) (*CloudServiceOracleConnResp, error) {
	expectedResp := &CloudServiceOracleConnResp{}
	_, err := c.sendRequest(cloudServiceHostedOracleURI, postMethod, oracleHosted, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
