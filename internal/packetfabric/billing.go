package packetfabric

import (
	"fmt"
	"time"
)

const billingURI = "/v2/billing/services/%s"
const billingModifyURI = "/v2/billing/services/%s/modify"
const etlURI = "/v2/billing/services/early-termination-liability/%s"

type BillingUpgrade struct {
	SubscriptionTerm   int    `json:"subscription_term,omitempty"`
	Speed              string `json:"speed,omitempty"`
	BillingProductType string `json:"billing_product_type,omitempty"`
	ServiceClass       string `json:"service_class,omitempty"`
	Capacity           string `json:"capacity,omitempty"`
}

type BillingResponse struct {
	OrderID        int            `json:"order_id"`
	AccountID      string         `json:"account_id"`
	CircuitID      string         `json:"circuit_id"`
	OrderType      string         `json:"order_type"`
	Term           Term           `json:"term"`
	Billables      []Billables    `json:"billables"`
	ProductDetails ProductDetails `json:"product_details"`
	ParentOrder    interface{}    `json:"parent_order"`
}
type Term struct {
	StartDate         string      `json:"start_date"`
	Months            int         `json:"months"`
	TerminationDate   interface{} `json:"termination_date"`
	CommitmentEndDate string      `json:"commitment_end_date"`
}
type Billables struct {
	AccountID     string  `json:"account_id"`
	BillableID    int     `json:"billable_id"`
	OrderID       int     `json:"order_id"`
	PriceType     string  `json:"price_type"`
	CurrencyCode  string  `json:"currency_code"`
	Price         float64 `json:"price"`
	AdjustedPrice float64 `json:"adjusted_price"`
}
type ProductDetails struct {
	ProductID      int         `json:"product_id"`
	Name           string      `json:"name"`
	VcType         string      `json:"vc_type"`
	VcServiceClass string      `json:"vc_service_class"`
	BundleType     string      `json:"bundle_type"`
	ActiveDate     string      `json:"active_date"`
	EndDate        interface{} `json:"end_date"`
	TranslationID  int         `json:"translation_id"`
}

type BillingUpgradeResp struct {
	Message string `json:"message,omitempty"`
}

// https://docs.packetfabric.com/api/v2/redoc/#operation/get_order
func (c *PFClient) ReadBilling(cID string) ([]BillingResponse, error) {
	formatedURI := fmt.Sprintf(billingURI, cID)

	expectedResp := make([]BillingResponse, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) ModifyBilling(cID string, billing BillingUpgrade) (*BillingUpgradeResp, error) {
	formatedURI := fmt.Sprintf(billingModifyURI, cID)

	expectedResp := &BillingUpgradeResp{}
	if _, err := c.sendRequest(formatedURI, postMethod, billing, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetEarlyTerminationLiability(circuitID string) (float64, error) {
	formattedURI := fmt.Sprintf(etlURI, circuitID)
	var resp float64
	// Add a delay of 5 seconds to allow the billing system to catch up
	time.Sleep(5 * time.Second)
	_, err := c.sendRequest(formattedURI, getMethod, nil, &resp)
	if err != nil {
		return 0, err
	}
	return resp, nil
}
