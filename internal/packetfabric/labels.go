package packetfabric

import "fmt"

const labelsURI = "/v2/objects/%s/labels"
const labelsUpdateURI = "/v2/labels/%s"

type LabelsResponse struct {
	Labels []string `json:"labels"`
}

type LabelsCreatePayload struct {
	Labels []string `json:"labels"`
}

type LabelsUpdateResponse struct {
	Objects []string `json:"objects"`
}

type LabelsUpdatePayload struct {
	Objects []string `json:"objects"`
}

func (c *PFClient) CreateLabel(circuitId string, labelsData LabelsCreatePayload) (*LabelsResponse, error) {
	formattedURI := fmt.Sprintf(labelsURI, circuitId)
	resp := &LabelsResponse{}
	_, err := c.sendRequest(formattedURI, putMethod, labelsData, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) UpdateLabel(labelValue string, labelsData LabelsUpdatePayload) (*LabelsUpdateResponse, error) {
	formattedURI := fmt.Sprintf(labelsUpdateURI, labelValue)
	resp := &LabelsUpdateResponse{}
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
