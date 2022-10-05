package packetfabric

import "fmt"

const pointToPointURI = "/v2/services/point-to-point"
const pointToPointByUUIDURI = "/v2/services/point-to-point/%s"
const pointToPointStatus = "/v2.1/services/point-to-point/%s/status"

type PointToPoint struct {
	Description            string      `json:"description,omitempty"`
	Speed                  string      `json:"speed,omitempty"`
	Media                  string      `json:"media,omitempty"`
	Endpoints              []Endpoints `json:"endpoints,omitempty"`
	AccountUUID            string      `json:"account_uuid,omitempty"`
	SubscriptionTerm       int         `json:"subscription_term,omitempty"`
	PublishedQuoteLineUUID string      `json:"published_quote_line_uuid,omitempty"`
}
type Endpoints struct {
	Pop              string `json:"pop,omitempty"`
	Zone             string `json:"zone,omitempty"`
	CustomerSiteCode string `json:"customer_site_code,omitempty"`
	Autoneg          bool   `json:"autoneg,omitempty"`
	Loa              string `json:"loa,omitempty"`
}

type PointToPointResp struct {
	PtpUUID      string       `json:"ptp_uuid,omitempty"`
	PtpCircuitID string       `json:"ptp_circuit_id,omitempty"`
	Description  string       `json:"description,omitempty"`
	Speed        string       `json:"speed,omitempty"`
	Media        string       `json:"media,omitempty"`
	State        string       `json:"state,omitempty"`
	Billing      Billing      `json:"billing,omitempty"`
	TimeCreated  string       `json:"time_created,omitempty"`
	TimeUpdated  string       `json:"time_updated,omitempty"`
	Deleted      bool         `json:"deleted,omitempty"`
	ServiceClass string       `json:"service_class,omitempty"`
	Interfaces   []Interfaces `json:"interfaces,omitempty"`
}

func (c *PFClient) CreatePointToPointService(ptp PointToPoint) (*PointToPointResp, error) {
	expectedResp := &PointToPointResp{}
	_, err := c.sendRequest(pointToPointURI, postMethod, ptp, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetPointToPointInfo(ptpUUID string) (*PointToPointResp, error) {
	formatedURI := fmt.Sprintf(pointToPointByUUIDURI, ptpUUID)
	expectedResp := &PointToPointResp{}
	if _, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetPointToPointInfos() ([]PointToPointResp, error) {
	expectedResp := make([]PointToPointResp, 0)
	if _, err := c.sendRequest(pointToPointURI, getMethod, nil, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetPointToPointStatus(ptpCircuitID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(pointToPointStatus, ptpCircuitID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) UpdatePointToPoint(ptpUUID, description string) (*PointToPointResp, error) {
	formatedURI := fmt.Sprintf(pointToPointByUUIDURI, ptpUUID)
	expectedResp := &PointToPointResp{}
	type DescUpdate struct {
		Description string `json:"description"`
	}
	if _, err := c.sendRequest(formatedURI, patchMethod, &DescUpdate{Description: description}, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) DeletePointToPointService(ptpUUID string) error {
	formatedURI := fmt.Sprintf(pointToPointByUUIDURI, ptpUUID)
	if _, err := c.sendRequest(formatedURI, deleteMethod, nil, nil); err != nil {
		return err
	}
	return nil
}
