package packetfabric

const locationsURI = "/v2/locations"

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
