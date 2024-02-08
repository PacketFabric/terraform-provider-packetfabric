package packetfabric

import (
	"fmt"
)

const HighPerformanceInternetURI = "/v2/services/high-performance-internet"

type HighPerformanceInternet struct {
	AccountUUID          string                                      `json:"account_uuid,omitempty"`          // write
	RoutingConfiguration HighPerformanceInternetRoutingConfiguration `json:"routing_configuration,omitempty"` // write/"read"
	PortCircuitId        string                                      `json:"port_circuit_id,omitempty"`       // read/write
	Speed                string                                      `json:"speed"`                           // read/write
	Vlan                 int                                         `json:"vlan"`                            // read/write
	Description          string                                      `json:"description"`                     // read/write
	CircuitId            string                                      `json:"hpi_circuit_id,omitempty"`        // read
	Market               string                                      `json:"market,omitempty"`                // read
	RoutingType          string                                      `json:"routing_type,omitempty"`          // read
	State                string                                      `json:"state,omitempty"`                 // read
}

type HighPerformanceInternetRoutingConfiguration struct {
	StaticRoutingV4 *HighPerformanceInternetStaticConfiguration `json:"static_routing_v4,omitempty"`
	StaticRoutingV6 *HighPerformanceInternetStaticConfiguration `json:"static_routing_v6,omitempty"`
	BgpV4           *HighPerformanceInternetBgpConfiguration    `json:"bgp_v4,omitempty"`
	BgpV6           *HighPerformanceInternetBgpConfiguration    `json:"bgp_v6,omitempty"`
}

type HighPerformanceInternetStaticConfiguration struct {
	L3Address     string                               `json:"l3_address"`               // "read"/write
	RemoteAddress string                               `json:"remote_address"`           // "read"/write
	Prefixes      []HighPerformanceInternetStaticRoute `json:"prefixes"`                 // "read"/write
	AddressFamily string                               `json:"address_family,omitempty"` // "read"
}

type HighPerformanceInternetBgpConfiguration struct {
	Asn           int                                `json:"asn"`                      // "read"/write
	L3Address     string                             `json:"l3_address"`               // "read"/write
	RemoteAddress string                             `json:"remote_address"`           // "read"/write
	Md5           string                             `json:"md5"`                      // "read"/write
	Prefixes      []HighPerformanceInternetBgpPrefix `json:"prefixes"`                 // "read"/write
	BgpState      string                             `json:"bgp_state,omitempty"`      // "read"
	AddressFamily string                             `json:"address_family,omitempty"` // "read"
}

type HighPerformanceInternetStaticRoute struct {
	Prefix string `json:"prefix"` // "read"/write
}

type HighPerformanceInternetBgpPrefix struct {
	Prefix          string `json:"prefix"`
	LocalPreference int    `json:"local_preference"`
}

type HighPerformanceInternetDeleteResponse struct {
	Message string `json:"message"`
}

func VerifyHighPerformanceInternetStaticOrBGP(highPerformanceInternet *HighPerformanceInternet) (*HighPerformanceInternet, error) {
	var routing = highPerformanceInternet.RoutingConfiguration
	var static = (nil != routing.StaticRoutingV4 || nil != routing.StaticRoutingV6)
	var bgp = (nil != routing.BgpV4 || nil != routing.BgpV6)
	if static && bgp {
		return nil, fmt.Errorf("HighPerformanceInternet should have either static or bgp routes but not both for %s", highPerformanceInternet.CircuitId)
	}
	return highPerformanceInternet, nil
}

// POST   /v2/services/high-performance-internet                     Create a new HPI
// Note: RoutingConfiguration will be empty, see ReadHighPerformanceInternetWithRoutes or AddHighPerformanceInternetRoutes
func (c *PFClient) CreateHighPerformanceInternet(highPerformanceInternet *HighPerformanceInternet) (*HighPerformanceInternet, error) {
	var err error
	highPerformanceInternet, err = VerifyHighPerformanceInternetStaticOrBGP(highPerformanceInternet)
	if err != nil {
		return nil, err
	}
	resp := &HighPerformanceInternet{}
	_, err = c.sendRequest(HighPerformanceInternetURI, postMethod, highPerformanceInternet, &resp)
	if err != nil {
		return nil, err
	}
	resp.AccountUUID = highPerformanceInternet.AccountUUID
	return resp, nil
}

// GET    /v2/services/high-performance-internet/{circuit_id}        Get a HPI by circuit_id
// Note: RoutingConfiguration will be empty, see ReadHighPerformanceInternetWithRoutes or AddHighPerformanceInternetRoutes
func (c *PFClient) ReadHighPerformanceInternet(circuitId string) (*HighPerformanceInternet, error) {
	formatedURI := fmt.Sprintf("%s/%s", HighPerformanceInternetURI, circuitId)
	resp := &HighPerformanceInternet{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PUT    /v2/services/high-performance-internet/{circuit_id}        Update a HPI
// Note: RoutingConfiguration will be empty, see ReadHighPerformanceInternetWithRoutes or AddHighPerformanceInternetRoutes
func (c *PFClient) UpdateHighPerformanceInternet(highPerformanceInternet *HighPerformanceInternet) (*HighPerformanceInternet, error) {
	var err error
	highPerformanceInternet, err = VerifyHighPerformanceInternetStaticOrBGP(highPerformanceInternet)
	if err != nil {
		return nil, err
	}
	formatedURI := fmt.Sprintf("%s/%s", HighPerformanceInternetURI, highPerformanceInternet.CircuitId)
	resp := &HighPerformanceInternet{}
	_, err = c.sendRequest(formatedURI, putMethod, highPerformanceInternet, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DELETE /v2/services/high-performance-internet/{circuit_id}        Delete a HPI by circuit_id
func (c *PFClient) DeleteHighPerformanceInternet(circuitId string) (*HighPerformanceInternetDeleteResponse, error) {
	formatedURI := fmt.Sprintf("%s/%s", HighPerformanceInternetURI, circuitId)
	resp := &HighPerformanceInternetDeleteResponse{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET    /v2/services/high-performance-internet/{circuit_id}/bgp    Get the bgp routing configurations for this HPI
func (c *PFClient) GetHighPerformanceInternetBgpConfiguration(circuitId string) ([]HighPerformanceInternetBgpConfiguration, error) {
	formatedURI := fmt.Sprintf("%s/%s/bgp", HighPerformanceInternetURI, circuitId)
	resp := make([]HighPerformanceInternetBgpConfiguration, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET    /v2/services/high-performance-internet/{circuit_id}/static Get the static routing configurations for this HPI
func (c *PFClient) GetHighPerformanceInternetStaticConfiguration(circuitId string) ([]HighPerformanceInternetStaticConfiguration, error) {
	formatedURI := fmt.Sprintf("%s/%s/static", HighPerformanceInternetURI, circuitId)
	resp := make([]HighPerformanceInternetStaticConfiguration, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET    /v2/services/high-performance-internet                     Get HPIs associated with the current customer
// Note: RoutingConfiguration will be empty, see ReadHighPerformanceInternetWithRoutes or AddHighPerformanceInternetRoutes
func (c *PFClient) ReadHighPerformanceInternets() ([]HighPerformanceInternet, error) {
	resp := make([]HighPerformanceInternet, 0)
	_, err := c.sendRequest(HighPerformanceInternetURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// GET    /v2/services/high-performance-internet/{circuit_id}        Get a HPI by circuit_id
// Note: RoutingConfiguration will be empty and must be filled with calls to GET bgp and static
func (c *PFClient) ReadHighPerformanceInternetWithRoutes(circuitId string) (*HighPerformanceInternet, error) {
	formatedURI := fmt.Sprintf("%s/%s", HighPerformanceInternetURI, circuitId)
	resp := &HighPerformanceInternet{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return c.AddHighPerformanceInternetRoutes(resp)
}

func (c *PFClient) AddHighPerformanceInternetRoutes(hpi *HighPerformanceInternet) (*HighPerformanceInternet, error) {
	var err error
	if "bgp" == hpi.RoutingType {
		hpi, err = c.AddHighPerformanceInternetBGPRoutes(hpi)
		if err != nil {
			return nil, err
		}
	} else {
		if "static" == hpi.RoutingType {
			hpi, err = c.AddHighPerformanceInternetStaticRoutes(hpi)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("Unexpected HighPerformanceInternet RoutingType \"%s\", expected one of \"bgp\" or \"static\"", hpi.RoutingType)
		}
	}
	return hpi, nil
}

func (c *PFClient) AddHighPerformanceInternetBGPRoutes(hpi *HighPerformanceInternet) (*HighPerformanceInternet, error) {
	bgpRoutes, err := c.GetHighPerformanceInternetBgpConfiguration(hpi.CircuitId)
	if err != nil {
		return nil, err
	}
	for _, bgpRoute := range bgpRoutes {
		if "v4" == bgpRoute.AddressFamily {
			if nil == hpi.RoutingConfiguration.BgpV4 {
				hpi.RoutingConfiguration.BgpV4 = &bgpRoute
			} else {
				return nil, fmt.Errorf("Expected only 1 v4 BGP route for %s", hpi.CircuitId)
			}
		} else {
			if "v6" == bgpRoute.AddressFamily {
				if nil == hpi.RoutingConfiguration.BgpV6 {
					hpi.RoutingConfiguration.BgpV6 = &bgpRoute
				} else {
					return nil, fmt.Errorf("Expected only 1 v6 BGP route for %s", hpi.CircuitId)
				}
			} else {
				return nil, fmt.Errorf("Unexpected HighPerformanceInternet BgpConfiguration Address family \"%s\", expected on of \"v4\" or \"v6\"", bgpRoute.AddressFamily)
			}
		}
	}
	return hpi, nil
}

func (c *PFClient) AddHighPerformanceInternetStaticRoutes(hpi *HighPerformanceInternet) (*HighPerformanceInternet, error) {
	staticRoutes, err := c.GetHighPerformanceInternetStaticConfiguration(hpi.CircuitId)
	if err != nil {
		return nil, err
	}
	for _, staticRoute := range staticRoutes {
		if "v4" == staticRoute.AddressFamily {
			if nil == hpi.RoutingConfiguration.StaticRoutingV4 {
				hpi.RoutingConfiguration.StaticRoutingV4 = &staticRoute
			} else {
				return nil, fmt.Errorf("Expected only 1 v4 static route for %s", hpi.CircuitId)
			}
		} else {
			if "v6" == staticRoute.AddressFamily {
				if nil == hpi.RoutingConfiguration.StaticRoutingV6 {
					hpi.RoutingConfiguration.StaticRoutingV6 = &staticRoute
				} else {
					return nil, fmt.Errorf("Expected only 1 v6 static route for %s", hpi.CircuitId)
				}
			} else {
				return nil, fmt.Errorf("Unexpected HighPerformanceInternet StaticConfiguration Address family \"%s\", expected on of \"v4\" or \"v6\"", staticRoute.AddressFamily)
			}
		}

	}
	return hpi, nil
}
