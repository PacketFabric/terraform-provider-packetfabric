package packetfabric

const locationsMarketsURI = "/v2/locations/markets"

type LocationMarket struct {
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Country string `json:"country,omitempty"`
}

func (c *PFClient) GetLocationsMarkets() ([]LocationMarket, error) {
	expectedResp := make([]LocationMarket, 0)
	_, err := c.sendRequest(locationsMarketsURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
