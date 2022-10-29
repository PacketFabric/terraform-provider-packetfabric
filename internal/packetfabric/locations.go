package packetfabric

import (
	"fmt"
	"net/http"
)

const locationsURI = "/v2/locations"
const portAvailabilityURI = "/v2/locations/%s/port-availability"

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

type PortAvailability struct {
	Zone    string `json:"zone"`
	Speed   string `json:"speed"`
	Media   string `json:"media"`
	Count   int    `json:"count"`
	Partial bool   `json:"partial"`
	Enni    bool   `json:"enni"`
}

func (c *PFClient) GetLocationPortAvailability(pop string) ([]PortAvailability, error) {
	resp := make([]PortAvailability, 0)
	_, err := c.sendRequest(fmt.Sprintf(portAvailabilityURI, pop), http.MethodGet, nil, &resp)
	if len(resp) == 0 {
		return resp, nil
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
