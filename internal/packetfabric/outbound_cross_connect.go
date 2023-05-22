package packetfabric

import "fmt"

const outboundCrossConnectURI = "/v2/outbound-cross-connects"
const outboundCrossConnectWithIDURI = "/v2/outbound-cross-connects/%s"

type OutboundCrossConnect struct {
	Port                     string `json:"port,omitempty"`
	Site                     string `json:"site,omitempty"`
	DocumentUUID             string `json:"document_uuid,omitempty"`
	Description              string `json:"description,omitempty"`
	DestinationName          string `json:"destination_name,omitempty"`
	DestinationCircuitID     string `json:"destination_circuit_id,omitempty"`
	Panel                    string `json:"panel,omitempty"`
	Module                   string `json:"module,omitempty"`
	Position                 string `json:"position,omitempty"`
	DataCenterCrossConnectID string `json:"data_center_cross_connect_id,omitempty"`
	PublishedQuoteLineUUID   string `json:"published_quote_line_uuid,omitempty"`
}

type OutboundCrossConnectResp struct {
	Port                     string `json:"port,omitempty"`
	Site                     string `json:"site,omitempty"`
	DocumentUUID             string `json:"document_uuid,omitempty"`
	OutboundCrossConnectID   string `json:"outbound_cross_connect_id,omitempty"`
	ObccStatus               string `json:"obcc_status,omitempty"`
	Description              string `json:"description,omitempty"`
	UserDescription          string `json:"user_description,omitempty"`
	DestinationName          string `json:"destination_name,omitempty"`
	DestinationCircuitID     string `json:"destination_circuit_id,omitempty"`
	Panel                    string `json:"panel,omitempty"`
	Module                   string `json:"module,omitempty"`
	Position                 string `json:"position,omitempty"`
	DataCenterCrossConnectID string `json:"data_center_cross_connect_id,omitempty"`
	Progress                 int    `json:"progress,omitempty"`
	Deleted                  bool   `json:"deleted,omitempty"`
	ZLocCfa                  string `json:"z_loc_cfa,omitempty"`
	TimeCreated              string `json:"time_created,omitempty"`
	TimeUpdated              string `json:"time_updated,omitempty"`
}

type OutboundCrossConnectMessageResp struct {
	Message string `json:"message"`
}

func (c *PFClient) CreateOutboundCrossConnect(crossConn OutboundCrossConnect) (*OutboundCrossConnectMessageResp, error) {
	expectedResp := &OutboundCrossConnectMessageResp{}
	_, err := c.sendRequest(outboundCrossConnectURI, postMethod, crossConn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetOutboundCrossConnect(outboundCrossConnID string) (*OutboundCrossConnectResp, error) {
	formatedURI := fmt.Sprintf(outboundCrossConnectWithIDURI, outboundCrossConnID)
	expectedResp := &OutboundCrossConnectResp{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) ListOutboundCrossConnects() (*[]OutboundCrossConnectResp, error) {
	expectedResp := make([]OutboundCrossConnectResp, 0)
	_, err := c.sendRequest(outboundCrossConnectURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return &expectedResp, nil
}

func (c *PFClient) UpdateOutboundCrossConnect(outboundCrossConnID, userDesc string) error {
	formatedURI := fmt.Sprintf(outboundCrossConnectWithIDURI, outboundCrossConnID)
	type UpdateCrossConn struct {
		UserDescription string `json:"user_description"`
	}
	updatePayload := UpdateCrossConn{
		UserDescription: userDesc,
	}
	_, err := c.sendRequest(formatedURI, patchMethod, updatePayload, nil)
	return err
}

func (c *PFClient) DeleteOutboundCrossConnect(port string) error {
	crossConns, err := c.ListOutboundCrossConnects()
	if err != nil {
		return err
	}
	for _, crossConn := range *crossConns {
		if crossConn.Port == port {
			formatedURI := fmt.Sprintf(outboundCrossConnectWithIDURI, crossConn.DataCenterCrossConnectID)
			_, err = c.sendRequest(formatedURI, deleteMethod, nil, nil)
		}
	}
	return err
}
