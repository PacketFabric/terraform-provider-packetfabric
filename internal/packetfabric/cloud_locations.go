package packetfabric

import "fmt"

const cloudLocationsURI = "/v2/locations/cloud?cloud_provider=%s&cloud_connection_type=%s&nat_capable=%v&has_cloud_router=%v&any_type=%v&pop=%s&city=%s&state=%s&market=%s&region=%s"

type CloudLocation struct {
	Pop                    string                 `json:"pop,omitempty"`
	Region                 string                 `json:"region,omitempty"`
	Market                 string                 `json:"market,omitempty"`
	MarketDescription      string                 `json:"market_description,omitempty"`
	Zones                  []string               `json:"zones,omitempty"`
	Vendor                 string                 `json:"vendor,omitempty"`
	Site                   string                 `json:"site,omitempty"`
	SiteCode               string                 `json:"site_code,omitempty"`
	Type                   string                 `json:"type,omitempty"`
	Status                 string                 `json:"status,omitempty"`
	Latitude               string                 `json:"latitude,omitempty"`
	Longitude              string                 `json:"longitude,omitempty"`
	Timezone               interface{}            `json:"timezone,omitempty"`
	Notes                  interface{}            `json:"notes,omitempty"`
	Pcode                  interface{}            `json:"pcode,omitempty"`
	LeadTime               string                 `json:"lead_time,omitempty"`
	SingleArmed            bool                   `json:"single_armed,omitempty"`
	Address1               string                 `json:"address1,omitempty"`
	Address2               interface{}            `json:"address2,omitempty"`
	City                   string                 `json:"city,omitempty"`
	State                  string                 `json:"state,omitempty"`
	Postal                 string                 `json:"postal,omitempty"`
	Country                string                 `json:"country,omitempty"`
	CloudProvider          string                 `json:"cloud_provider,omitempty"`
	CloudConnectionDetails CloudConnectionDetails `json:"cloud_connection_details,omitempty"`
	NetworkProvider        string                 `json:"network_provider,omitempty"`
	TimeCreated            string                 `json:"time_created,omitempty"`
	EnniSupported          bool                   `json:"enni_supported,omitempty"`
}
type CloudConnectionDetails struct {
	Region            string `json:"region,omitempty"`
	HostedType        string `json:"hosted_type,omitempty"`
	RegionDescription string `json:"region_description,omitempty"`
}

func (c *PFClient) GetCloudLocations(
	cloudProvider, cloudConnType string,
	natCapable, hasCloudRouter, anyType bool,
	pop, city, state, market, region string) ([]CloudLocation, error) {
	formatedURI := fmt.Sprintf(cloudLocationsURI, cloudProvider, cloudConnType,
		natCapable, hasCloudRouter, anyType,
		pop, city, state, market, region)
	expectedResp := make([]CloudLocation, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
