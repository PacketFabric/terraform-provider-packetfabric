package packetfabric

import (
	"fmt"
)

const portsURI = "/v2/ports"
const portStatusURI = "/v2.1/ports/%s/status"
const portByCIDURI = "/v2/ports/%s"
const portEnableURI = "/v2/ports/%s/enable"
const portDisableURI = "/v2/ports/%s/disable"
const portLoaURI = "/v2/ports/%s/letter-of-authorization"
const portVlanSummaryURI = "/v2/ports/%s/vlan-summary"
const portDeviceInfoURI = "/v2/ports/%s/device-info"
const portRouterLogsURI = "/v2/ports/%s/router-logs?time_from=%s&time_to=%s"

type Interface struct {
	Autoneg          bool   `json:"autoneg,omitempty"`
	Nni              bool   `json:"nni,omitempty"`
	SubscriptionTerm int    `json:"subscription_term,omitempty"`
	AccountUUID      string `json:"account_uuid,omitempty"`
	Pop              string `json:"pop,omitempty"`
	Speed            string `json:"speed,omitempty"`
	Media            string `json:"media,omitempty"`
	Zone             string `json:"zone,omitempty"`
	Description      string `json:"description,omitempty"`
	PortCircuitID    string `json:"port_circuit_id,omitempty"`
	Vlan             int    `json:"vlan,omitempty"`
	Svlan            int    `json:"svlan,omitempty"`
	VlanMicrosoft    int    `json:"vlan_microsoft,omitempty"`
	VlanPrivate      int    `json:"vlan_private,omitempty"`
	Untagged         bool   `json:"untagged,omitempty"`
	PONumber         string `json:"po_number,omitempty"`
}

type InterfaceCreateResp struct {
	Autoneg             bool        `json:"autoneg,omitempty"`
	PortCircuitID       string      `json:"port_circuit_id,omitempty"`
	State               string      `json:"state,omitempty"`
	Pop                 string      `json:"pop,omitempty"`
	Speed               string      `json:"speed,omitempty"`
	Media               string      `json:"media,omitempty"`
	Zone                string      `json:"zone,omitempty"`
	Mtu                 int         `json:"mtu,omitempty"`
	Description         string      `json:"description,omitempty"`
	VcMode              interface{} `json:"vc_mode,omitempty"`
	IsLag               bool        `json:"is_lag,omitempty"`
	IsLagMember         bool        `json:"is_lag_member,omitempty"`
	IsCloud             bool        `json:"is_cloud,omitempty"`
	IsPtp               bool        `json:"is_ptp,omitempty"`
	LagInterval         interface{} `json:"lag_interval,omitempty"`
	MemberCount         interface{} `json:"member_count,omitempty"`
	ParentLagCircuitID  interface{} `json:"parent_lag_circuit_id,omitempty"`
	Disabled            bool        `json:"disabled,omitempty"`
	Status              string      `json:"status,omitempty"`
	TimeCreated         string      `json:"time_created,omitempty"`
	TimeUpdated         string      `json:"time_updated,omitempty"`
	IsCloudRouter       interface{} `json:"is_cloud_router,omitempty"`
	IsNatCapable        interface{} `json:"is_nat_capable,omitempty"`
	IsIpsecCapable      interface{} `json:"is_ipsec_capable,omitempty"`
	Provider            string      `json:"provider,omitempty"`
	Region              string      `json:"region,omitempty"`
	Market              string      `json:"market,omitempty"`
	MarketDescription   string      `json:"market_description,omitempty"`
	Site                string      `json:"site,omitempty"`
	SiteCode            string      `json:"site_code,omitempty"`
	OperationalStatus   interface{} `json:"operational_status,omitempty"`
	AdminStatus         interface{} `json:"admin_status,omitempty"`
	AccountUUID         string      `json:"account_uuid,omitempty"`
	SubscriptionTerm    int         `json:"subscription_term,omitempty"`
	IsNni               bool        `json:"is_nni,omitempty"`
	CustomerName        string      `json:"customer_name,omitempty"`
	CustomerUUID        string      `json:"customer_uuid,omitempty"`
	MaxCloudRouterSpeed string      `json:"max_cloud_router_speed,omitempty"`
}

type InterfaceReadResp struct {
	Autoneg             bool   `json:"autoneg,omitempty"`
	PortCircuitID       string `json:"port_circuit_id,omitempty"`
	State               string `json:"state,omitempty"`
	Pop                 string `json:"pop,omitempty"`
	Speed               string `json:"speed,omitempty"`
	Media               string `json:"media,omitempty"`
	Zone                string `json:"zone,omitempty"`
	Mtu                 int    `json:"mtu,omitempty"`
	Description         string `json:"description,omitempty"`
	VcMode              string `json:"vc_mode,omitempty"`
	IsLag               bool   `json:"is_lag,omitempty"`
	IsLagMember         bool   `json:"is_lag_member,omitempty"`
	IsCloud             bool   `json:"is_cloud,omitempty"`
	IsPtp               bool   `json:"is_ptp,omitempty"`
	LagInterval         string `json:"lag_interval,omitempty"`
	MemberCount         int    `json:"member_count,omitempty"`
	ParentLagCircuitID  string `json:"parent_lag_circuit_id,omitempty"`
	Disabled            bool   `json:"disabled,omitempty"`
	Status              string `json:"status,omitempty"`
	TimeCreated         string `json:"time_created,omitempty"`
	TimeUpdated         string `json:"time_updated,omitempty"`
	IsCloudRouter       bool   `json:"is_cloud_router,omitempty"`
	IsNatCapable        bool   `json:"is_nat_capable,omitempty"`
	IsIpsecCapable      bool   `json:"is_ipsec_capable,omitempty"`
	Provider            string `json:"provider,omitempty"`
	Region              string `json:"region,omitempty"`
	Market              string `json:"market,omitempty"`
	MarketDescription   string `json:"market_description,omitempty"`
	Site                string `json:"site,omitempty"`
	SiteCode            string `json:"site_code,omitempty"`
	OperationalStatus   string `json:"operational_status,omitempty"`
	AdminStatus         string `json:"admin_status,omitempty"`
	AccountUUID         string `json:"account_uuid,omitempty"`
	SubscriptionTerm    int    `json:"subscription_term,omitempty"`
	IsNni               bool   `json:"is_nni,omitempty"`
	CustomerName        string `json:"customer_name,omitempty"`
	CustomerUUID        string `json:"customer_uuid,omitempty"`
	MaxCloudRouterSpeed string `json:"max_cloud_router_speed,omitempty"`
	Links               Links  `json:"_links,omitempty"`
	PONumber            string `json:"po_number,omitempty"`
}

type Links struct {
	DeviceInfo string `json:"device_info,omitempty"`
	Location   string `json:"location,omitempty"`
}

type PortMessageResp struct {
	Message string `json:"message"`
}

type PortLoa struct {
	LoaCustomerName  string `json:"loa_customer_name,omitempty"`
	DestinationEmail string `json:"destination_email,omitempty"`
}

type PortLoaResp struct {
	UUID        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mime_type,omitempty"`
	Size        int    `json:"size,omitempty"`
	TimeCreated string `json:"time_created,omitempty"`
	TimeUpdated string `json:"time_updated,omitempty"`
}

type PortVlanSummary struct {
	LowestAvailableVlan int `json:"lowest_available_vlan,omitempty"`
	MaxVlan             int `json:"max_vlan,omitempty"`
}

type PortDeviceInfo struct {
	AdjacentRouter              interface{}                   `json:"adjacent_router,omitempty"`
	DeviceName                  string                        `json:"device_name,omitempty"`
	DeviceMake                  string                        `json:"device_make,omitempty"`
	AdminStatus                 string                        `json:"admin_status,omitempty"`
	OperStatus                  string                        `json:"oper_status,omitempty"`
	AutoNegotiation             bool                          `json:"auto_negotiation,omitempty"`
	IfaceName                   string                        `json:"iface_name,omitempty"`
	Speed                       string                        `json:"speed,omitempty"`
	OpticsDiagnosticsLaneValues []OpticsDiagnosticsLaneValues `json:"optics_diagnostics_lane_values,omitempty"`
	Polltime                    interface{}                   `json:"polltime,omitempty"`
	TimeFlapped                 string                        `json:"time_flapped,omitempty"`
	TrafficRxBps                int                           `json:"traffic_rx_bps,omitempty"`
	TrafficRxBytes              int                           `json:"traffic_rx_bytes,omitempty"`
	TrafficRxIpv6Bytes          int                           `json:"traffic_rx_ipv6_bytes,omitempty"`
	TrafficRxIpv6Packets        int                           `json:"traffic_rx_ipv6_packets,omitempty"`
	TrafficRxPackets            int                           `json:"traffic_rx_packets,omitempty"`
	TrafficRxPps                int                           `json:"traffic_rx_pps,omitempty"`
	TrafficTxBps                int                           `json:"traffic_tx_bps,omitempty"`
	TrafficTxBytes              int                           `json:"traffic_tx_bytes,omitempty"`
	TrafficTxIpv6Bytes          int                           `json:"traffic_tx_ipv6_bytes,omitempty"`
	TrafficTxIpv6Packets        int                           `json:"traffic_tx_ipv6_packets,omitempty"`
	TrafficTxPackets            int                           `json:"traffic_tx_packets,omitempty"`
	TrafficTxPps                int                           `json:"traffic_tx_pps,omitempty"`
	WiringMedia                 string                        `json:"wiring_media,omitempty"`
	WiringModule                string                        `json:"wiring_module,omitempty"`
	WiringPanel                 string                        `json:"wiring_panel,omitempty"`
	WiringPosition              string                        `json:"wiring_position,omitempty"`
	WiringReach                 string                        `json:"wiring_reach,omitempty"`
	WiringType                  string                        `json:"wiring_type,omitempty"`
	LagSpeed                    int                           `json:"lag_speed,omitempty"`
	DeviceCanLag                bool                          `json:"device_can_lag,omitempty"`
}

type OpticsDiagnosticsLaneValues struct {
	TxPowerDbm  float64 `json:"tx_power_dbm,omitempty"`
	TxPower     float64 `json:"tx_power,omitempty"`
	LaneIndex   string  `json:"lane_index,omitempty"`
	RxPower     float64 `json:"rx_power,omitempty"`
	RxPowerDbm  float64 `json:"rx_power_dbm,omitempty"`
	BiasCurrent float64 `json:"bias_current,omitempty"`
	TxStatus    string  `json:"tx_status,omitempty"`
	RxStatus    string  `json:"rx_status,omitempty"`
}

type PortRouterLogs struct {
	DeviceName   string `json:"device_name,omitempty"`
	IfaceName    string `json:"iface_name,omitempty"`
	Message      string `json:"message,omitempty"`
	Severity     int    `json:"severity,omitempty"`
	SeverityName string `json:"severity_name,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
}

type PortUpdate struct {
	Description string `json:"description"`
	PONumber    string `json:"po_number,omitempty"`
}

func (c *PFClient) CreateInterface(interf Interface) (*InterfaceCreateResp, error) {
	expectedResp := &InterfaceCreateResp{}

	_, err := c.sendRequest(portsURI, postMethod, interf, expectedResp)
	if err != nil {
		return nil, err
	}
	checked := checkPortStatus(c, expectedResp.PortCircuitID)
	if checked {
		return expectedResp, nil
	} else {
		return nil, fmt.Errorf("could not determine port status")
	}
}

func (c *PFClient) SendPortLoa(portCID string, portLoa PortLoa) (*PortLoaResp, error) {
	formatedURI := fmt.Sprintf(portLoaURI, portCID)
	expectedResp := &PortLoaResp{}
	_, err := c.sendRequest(formatedURI, postMethod, portLoa, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) EnablePort(portCID string) (*PortMessageResp, error) {
	return c.changePortState(portCID, true)
}

func (c *PFClient) DisablePort(portCID string) (*PortMessageResp, error) {
	return c.changePortState(portCID, false)
}

func (c *PFClient) changePortState(portCID string, enable bool) (*PortMessageResp, error) {
	expectedResp := &PortMessageResp{}
	var formatedURI string
	if enable {
		formatedURI = fmt.Sprintf(portEnableURI, portCID)
	} else {
		formatedURI = fmt.Sprintf(portDisableURI, portCID)
	}
	_, err := c.sendRequest(formatedURI, postMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	statusChangeOk := checkPortStatus(c, portCID)
	if statusChangeOk {
		return expectedResp, nil
	} else {
		return nil, fmt.Errorf("could not determine port status")
	}
}

func (c *PFClient) EnablePortAutoneg(portCID string) (*InterfaceReadResp, error) {
	return c.changePortStateAutoneg(portCID, true)
}

func (c *PFClient) DisablePortAutoneg(portCID string) (*InterfaceReadResp, error) {
	return c.changePortStateAutoneg(portCID, false)
}

func (c *PFClient) changePortStateAutoneg(portCID string, enable bool) (*InterfaceReadResp, error) {
	formatedURI := fmt.Sprintf(portByCIDURI, portCID)
	expectedResp := &InterfaceReadResp{}
	type PortUpdateAutoNegOnly struct {
		Autoneg bool `json:"autoneg"`
	}
	portUpdateAutoneg := PortUpdateAutoNegOnly{
		Autoneg: enable,
	}
	_, err := c.sendRequest(formatedURI, patchMethod, portUpdateAutoneg, expectedResp)
	if err != nil {
		return nil, err
	}
	updateAutonegChangeOk := checkPortStatus(c, portCID)
	if updateAutonegChangeOk {
		return expectedResp, nil
	} else {
		return nil, fmt.Errorf("could not determine port status")
	}
}

func (c *PFClient) GetPortByCID(portCID string) (*InterfaceReadResp, error) {
	formatedURI := fmt.Sprintf(portByCIDURI, portCID)
	expectedResp := &InterfaceReadResp{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) ListPorts() (*[]InterfaceReadResp, error) {
	expectedResp := make([]InterfaceReadResp, 0)
	_, err := c.sendRequest(portsURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return &expectedResp, nil
}

func (c *PFClient) GetPortStatus(portCID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(portStatusURI, portCID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetPortVlanSummary(portCID string) (*PortVlanSummary, error) {
	formatedURI := fmt.Sprintf(portVlanSummaryURI, portCID)
	expectedResp := &PortVlanSummary{}

	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetPortDeviceInfo(portCID string) (*PortDeviceInfo, error) {
	formatedURI := fmt.Sprintf(portDeviceInfoURI, portCID)
	expectedResp := &PortDeviceInfo{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetPortRouterLogs(portCID, timeFrom, timeTo string) ([]PortRouterLogs, error) {
	formatedURI := fmt.Sprintf(portRouterLogsURI, portCID, timeFrom, timeTo)
	expectedResp := make([]PortRouterLogs, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) UpdatePort(portCID string, portUpdateData PortUpdate) (*InterfaceReadResp, error) {
	formatedURI := fmt.Sprintf(portByCIDURI, portCID)
	expectedResp := &InterfaceReadResp{}
	_, err := c.sendRequest(formatedURI, patchMethod, portUpdateData, expectedResp)
	if err != nil {
		return nil, err
	}
	updated := checkPortStatus(c, expectedResp.PortCircuitID)
	if updated {
		return expectedResp, nil
	} else {
		return nil, fmt.Errorf("could not determine port status")
	}
}

func (c *PFClient) DeletePort(portCID string) (*PortMessageResp, error) {
	formatedURI := fmt.Sprintf(portByCIDURI, portCID)
	expectedResp := &PortMessageResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	deleted := checkPortStatus(c, portCID)
	if deleted {
		return expectedResp, nil
	} else {
		return nil, fmt.Errorf("could not determine port status")
	}
}

func checkPortStatus(c *PFClient, portId string) bool {
	done := make(chan bool)
	defer close(done)
	fn := func() (*ServiceState, error) {
		return c.GetPortStatus(portId)
	}
	go c.CheckServiceStatus(done, fn)
	return <-done
}
