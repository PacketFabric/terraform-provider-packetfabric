package packetfabric

const documentsURI = "/v2/documents"

type DocumentLinks struct {
	Port                  string `json:"port"`
	Service               string `json:"service"`
	Cloud                 string `json:"cloud"`
	CloudRouter           string `json:"cloud_router"`
	CloudRouterConnection string `json:"cloud_router_connection"`
}

type Document struct {
	UUID        string         `json:"uuid"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MimeType    string         `json:"mime_type"`
	Type        string         `json:"type"`
	Size        int            `json:"size"`
	TimeCreated string         `json:"time_created"`
	TimeUpdated string         `json:"time_updated"`
	Links       *DocumentLinks `json:"_links"`
}

type Documents struct {
	Documents []Document `json:"documents"`
}

type DocumentsPayload struct {
	Document      string `json:"document"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	PortCircuitId string `json:"port_circuit_id,omitempty"`
}

func (c *PFClient) CreateDocument(documentsData DocumentsPayload) (*Document, error) {
	resp := &Document{}
	_, err := c.sendRequest(documentsURI, postMethod, documentsData, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) GetDocuments() (*Documents, error) {
	resp := &Documents{}
	_, err := c.sendRequest(documentsURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
