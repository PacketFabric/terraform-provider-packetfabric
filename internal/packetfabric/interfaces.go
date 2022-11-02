package packetfabric

import (
	"fmt"
)

const portsURI = "/v2/ports"
const portStatusURI = "/v2.1/ports/%s/status"
const portByCIDURI = "/v2/ports/%s"
const portEnableURI = "/v2/ports/%s/enable"
const portDisableURI = "/v2/ports/%s/disable"

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
}

type Links struct {
	DeviceInfo string `json:"device_info,omitempty"`
	Location   string `json:"location,omitempty"`
}

type PortMessageResp struct {
	Message string `json:"message"`
}

func (c *PFClient) CreateInterface(interf Interface) (*InterfaceCreateResp, error) {
	expectedResp := &InterfaceCreateResp{}
	_, err := c.sendRequest(portsURI, postMethod, interf, expectedResp)
	if err != nil {
		return nil, err
	}
	createOk := make(chan bool)
	defer close(createOk)
	fn := func() (*ServiceState, error) {
		return c.GetPortStatus(expectedResp.PortCircuitID)
	}
	go c.CheckServiceStatus(createOk, fn)
	if !<-createOk {
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
	return expectedResp, nil
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

func (c *PFClient) UpdatePort(autoNeg bool, portCID, description string) (*InterfaceReadResp, error) {
	formatedURI := fmt.Sprintf(portByCIDURI, portCID)
	expectedResp := &InterfaceReadResp{}
	type PortUpdate struct {
		Autoneg     bool   `json:"autoneg"`
		Description string `json:"description"`
	}
	portUpdate := PortUpdate{
		Autoneg:     autoNeg,
		Description: description,
	}
	_, err := c.sendRequest(formatedURI, patchMethod, portUpdate, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) UpdatePortAutoNegOnly(autoNeg bool, portCID string) (*InterfaceReadResp, error) {
	formatedURI := fmt.Sprintf(portByCIDURI, portCID)
	expectedResp := &InterfaceReadResp{}
	type PortUpdate struct {
		Autoneg bool `json:"autoneg"`
	}
	portUpdate := PortUpdate{
		Autoneg: autoNeg,
	}
	_, err := c.sendRequest(formatedURI, patchMethod, portUpdate, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) UpdatePortDescriptionOnly(portCID, description string) (*InterfaceReadResp, error) {
	formatedURI := fmt.Sprintf(portByCIDURI, portCID)
	expectedResp := &InterfaceReadResp{}
	type PortUpdate struct {
		Description string `json:"description"`
	}
	portUpdate := PortUpdate{
		Description: description,
	}
	_, err := c.sendRequest(formatedURI, patchMethod, portUpdate, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) DeletePort(portCID string) (*PortMessageResp, error) {
	formatedURI := fmt.Sprintf(portByCIDURI, portCID)
	expectedResp := &PortMessageResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	deleteOk := make(chan bool)
	fn := func() (*ServiceState, error) {
		return c.GetPortStatus(portCID)
	}
	go c.CheckServiceStatus(deleteOk, fn)
	if !<-deleteOk {
		return nil, err
	}
	return expectedResp, nil
}
