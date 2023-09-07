package packetfabric

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const attachCustomerPortToCRURI = "/v2/services/cloud-routers/%s/connections/packetfabric"

type CustomerOwnedPort struct {
	AccountUUID            string `json:"account_uuid,omitempty"`
	MaybeNat               bool   `json:"maybe_nat,omitempty"`
	MaybeDNat              bool   `json:"maybe_dnat,omitempty"`
	PortCircuitID          string `json:"port_circuit_id,omitempty"`
	Description            string `json:"description,omitempty"`
	Untagged               bool   `json:"untagged,omitempty"`
	Vlan                   int    `json:"vlan,omitempty"`
	Speed                  string `json:"speed,omitempty"`
	IsPublic               bool   `json:"is_public,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
	PONumber               string `json:"po_number,omitempty"`
	SubscriptionTerm       int    `json:"subscription_term,omitempty" validate:"oneof=1 12 24 36" default:"1"`
}

type CustomerOwnedPortResp struct {
	PortType                  string        `json:"port_type,omitempty"`
	ConnectionType            string        `json:"connection_type,omitempty"`
	PortCircuitID             string        `json:"port_circuit_id,omitempty"`
	PendingDelete             bool          `json:"pending_delete,omitempty"`
	Deleted                   bool          `json:"deleted,omitempty"`
	Speed                     string        `json:"speed,omitempty"`
	State                     string        `json:"state,omitempty"`
	CloudCircuitID            string        `json:"cloud_circuit_id,omitempty"`
	AccountUUID               string        `json:"account_uuid,omitempty"`
	ServiceClass              string        `json:"service_class,omitempty"`
	ServiceProvider           string        `json:"service_provider,omitempty"`
	ServiceType               string        `json:"service_type,omitempty"`
	Description               string        `json:"description,omitempty"`
	UUID                      string        `json:"uuid,omitempty"`
	CloudProviderConnectionID string        `json:"cloud_provider_connection_id,omitempty"`
	CloudSettings             CloudSettings `json:"cloud_settings,omitempty"`
	UserUUID                  string        `json:"user_uuid,omitempty"`
	CustomerUUID              string        `json:"customer_uuid,omitempty"`
	TimeCreated               string        `json:"time_created,omitempty"`
	TimeUpdated               string        `json:"time_updated,omitempty"`
	CloudProvider             CloudProvider `json:"cloud_provider,omitempty"`
	Pop                       string        `json:"pop,omitempty"`
	Site                      string        `json:"site,omitempty"`
	BgpState                  string        `json:"bgp_state,omitempty"`
	CloudRouterCircuitID      string        `json:"cloud_router_circuit_id,omitempty"`
	NatCapable                bool          `json:"nat_capable,omitempty"`
	SubscriptionTerm          int           `json:"subscription_term,omitempty" validate:"oneof=1 12 24 36" default:"1"`
}

func (c *PFClient) AttachCustomerOwnedPortToCR(ownedPort CustomerOwnedPort, cID string) (*CustomerOwnedPortResp, error) {
	formatedURI := fmt.Sprintf(attachCustomerPortToCRURI, cID)

	if err := validator.New().Struct(ownedPort); err != nil {
		return nil, err
	}

	expectedResp := &CustomerOwnedPortResp{}
	_, err := c.sendRequest(formatedURI, postMethod, ownedPort, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
