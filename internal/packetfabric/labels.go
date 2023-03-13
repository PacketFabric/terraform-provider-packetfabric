package packetfabric

import "fmt"

const labelsURI = "/v2/objects/%s/labels"

type LabelsResponse struct {
	Labels []string `json:"labels"`
}

type LabelsPayload struct {
	Labels []string `json:"labels"`
}

func (c *PFClient) CreateLabel(circuitId string, labelsData LabelsPayload) (*LabelsResponse, error) {
	formattedURI := fmt.Sprintf(labelsURI, circuitId)
	resp := &LabelsResponse{}
	_, err := c.sendRequest(formattedURI, putMethod, labelsData, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) UpdateLabel(circuitId string, labelsData LabelsPayload) (*LabelsResponse, error) {
	formattedURI := fmt.Sprintf(labelsURI, circuitId)
	if labelsData.Labels == nil {
		labelsData = LabelsPayload{Labels: []string{}}
	}
	resp := &LabelsResponse{}
	_, err := c.sendRequest(formattedURI, putMethod, labelsData, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) GetLabels(circuitId string) (*LabelsResponse, error) {
	formattedURI := fmt.Sprintf(labelsURI, circuitId)
	resp := &LabelsResponse{}
	_, err := c.sendRequest(formattedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
