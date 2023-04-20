package packetfabric

import (
	"fmt"
)

const locationsURI = "/v2/locations"
const portAvailabilityURI = "/v2/locations/%s/port-availability"
const locationsZonesURI = "/v2/locations/%s/zones"
const mktServicesURI = "/v2/marketplace/services?market=%s"

type Location struct {
	Pop               string `json:"pop"`
	Region            string `json:"region"`
	Market            string `json:"market"`
	MarketDescription string `json:"market_description"`
	Vendor            string `json:"vendor"`
	Site              string `json:"site"`
	SiteCode          string `json:"site_code"`
	Type              string `json:"type"`
	Status            string `json:"status"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	Timezone          string `json:"timezone,omitempty"`
	Notes             string `json:"notes,omitempty"`
	Pcode             int    `json:"pcode"`
	LeadTime          string `json:"lead_time"`
	SingleArmed       bool   `json:"single_armed"`
	Address1          string `json:"address1"`
	Address2          string `json:"address2,omitempty"`
	City              string `json:"city"`
	State             string `json:"state"`
	Postal            string `json:"postal"`
	Country           string `json:"country"`
	NetworkProvider   string `json:"network_provider"`
	TimeCreated       string `json:"time_created"`
	EnniSupported     bool   `json:"enni_supported"`
}

type PortAvailability struct {
	Zone    string `json:"zone,omitempty"`
	Speed   string `json:"speed,omitempty"`
	Media   string `json:"media,omitempty"`
	Count   int    `json:"count,omitempty"`
	Partial bool   `json:"partial,omitempty"`
	Enni    bool   `json:"enni,omitempty"`
}

func (c *PFClient) ListLocations() ([]Location, error) {
	resp := make([]Location, 0)
	_, err := c.sendRequest(locationsURI, getMethod, nil, &resp)
	if len(resp) == 0 {
		return resp, nil
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) GetLocationPortAvailability(pop string) ([]PortAvailability, error) {
	resp := make([]PortAvailability, 0)
	_, err := c.sendRequest(fmt.Sprintf(portAvailabilityURI, pop), getMethod, nil, &resp)
	if len(resp) == 0 {
		return resp, nil
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) GetLocationsZones(pop string) ([]string, error) {
	formatedURI := fmt.Sprintf(locationsZonesURI, pop)
	expectedResp := make([]string, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetMarketplacePopByMarket(market string) (string, error) {
	type MktServices []struct {
		Services []struct {
			Locations []string `json:"locations"`
		} `json:"services"`
		Links struct {
			CustomerImage string `json:"customer_image"`
		} `json:"_links"`
	}
	mktServices := make(MktServices, 0)
	_, err := c.sendRequest(fmt.Sprintf(mktServicesURI, market), getMethod, nil, &mktServices)
	if err != nil {
		return "", err
	}
	if len(mktServices) == 0 {
		return "", fmt.Errorf("empty marketplace services")
	}
	services := mktServices
	if len(services) == 0 {
		return "", fmt.Errorf("empty marketplace services locations")
	}
	locations := services[0].Services
	return locations[0].Locations[0], nil
}
