package packetfabric

const locationRegionsURI = "/v2/locations/regions"

type LocationRegion struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

func (c *PFClient) GetLocationRegions() ([]LocationRegion, error) {
	expectedResp := make([]LocationRegion, 0)
	_, err := c.sendRequest(locationRegionsURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
