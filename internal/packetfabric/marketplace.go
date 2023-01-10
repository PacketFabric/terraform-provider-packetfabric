package packetfabric

import (
	"fmt"

	"github.com/google/uuid"
)

const marketPlaceSeviceURI = "/v2/marketplace/services"
const marketPlaceRouteSetsURI = "/v2/services/cloud-routers/%s/route-sets/%s/connections"
const marketPlaceByUUIDURI = "/v2/marketplace/services/%s"
const marketPlaceUpdateURI = "/v2/services/cloud-routers/%s/route-sets/%s"

type MarketplaceService struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	State       string   `json:"state,omitempty"`
	Sku         string   `json:"sku,omitempty"`
	Locations   []string `json:"locations,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Published   bool     `json:"published"`
	ServiceType string   `json:"service_type,omitempty"`
}

type MarketplaceServiceRouteSet struct {
	CloudRouterCircuitID string               `json:"cloud_router_circuit_id,omitempty"`
	RouteSet             RouteSet             `json:"route_set,omitempty"`
	ConnectionCircuitIDs ConnectionCircuitIDs `json:",omitempty"`
	RouteSetCircuitID    string               `json:",omitempty"`
}
type MktPrefix struct {
	Prefix    string `json:"prefix,omitempty"`
	MatchType string `json:"match_type,omitempty"`
}
type RouteSet struct {
	Description string      `json:"description,omitempty"`
	IsPrivate   bool        `json:"is_private"`
	Prefixes    []MktPrefix `json:"prefixes,omitempty"`
}

type MarketplaceServiceResp struct {
	UUID                 string          `json:"uuid,omitempty"`
	Locations            []string        `json:"locations,omitempty"`
	Categories           []MktCategories `json:"categories,omitempty"`
	Name                 string          `json:"name,omitempty"`
	ServiceType          string          `json:"service_type,omitempty"`
	Description          string          `json:"description,omitempty"`
	Published            bool            `json:"published,omitempty"`
	State                string          `json:"state,omitempty"`
	Sku                  string          `json:"sku,omitempty"`
	CloudRouterCircuitID string          `json:"cloud_router_circuit_id,omitempty"`
	BgpRouteSetCircuitID string          `json:"bgp_route_set_circuit_id,omitempty"`
	Links                MktLinks        `json:"_links,omitempty"`
}
type MktCategories struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}
type MktLinks struct {
	ServiceImage string `json:"service_image,omitempty"`
	RouteSets    string `json:"route_sets,omitempty"`
}

type ConnectionCircuitIDs map[string]interface{}

func (c *PFClient) CreateMarketplaceService(service MarketplaceService) (*MarketplaceServiceResp, error) {
	expectedResp := &MarketplaceServiceResp{}
	_, err := c.sendRequest(marketPlaceSeviceURI, postMethod, service, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, err
}

func (c *PFClient) CreateMarketplaceServiceWithRouteSet(service MarketplaceService, mktServiceRouteSet MarketplaceServiceRouteSet) (*MarketplaceServiceResp, error) {
	expectedResp := &MarketplaceServiceResp{}
	if service.ServiceType == "quick-connect-service" {
		// Please see this Github issue for more details: #248
		service.ServiceType = "cloud-router-service"
	}
	var err error
	type MarketplaceWithRouteSet struct {
		Name                 string   `json:"name,omitempty"`
		Description          string   `json:"description,omitempty"`
		Sku                  string   `json:"sku,omitempty"`
		Locations            []string `json:"locations,omitempty"`
		Categories           []string `json:"categories,omitempty"`
		Published            bool     `json:"published"`
		ServiceType          string   `json:"service_type,omitempty"`
		CloudRouterCircuitID string   `json:"cloud_router_circuit_id,omitempty"`
		RouteSet             RouteSet `json:"route_set,omitempty"`
	}
	mktWithRouteSet := MarketplaceWithRouteSet{
		Name:                 service.Name,
		Description:          service.Description,
		Sku:                  service.Sku,
		Locations:            service.Locations,
		Categories:           service.Categories,
		Published:            service.Published,
		ServiceType:          service.ServiceType,
		CloudRouterCircuitID: mktServiceRouteSet.CloudRouterCircuitID,
		RouteSet:             mktServiceRouteSet.RouteSet,
	}
	_, err = c.sendRequest(marketPlaceSeviceURI, postMethod, mktWithRouteSet, expectedResp)
	if err != nil {
		return nil, err
	}
	if service.ServiceType == "cloud-router-service" && len(mktServiceRouteSet.ConnectionCircuitIDs) > 0 {
		err = c.UpdateMarketPlaceConnection(mktServiceRouteSet.CloudRouterCircuitID, expectedResp.BgpRouteSetCircuitID, mktServiceRouteSet.ConnectionCircuitIDs)
		if err != nil {
			return nil, err
		}
	}
	return expectedResp, err
}

func (c *PFClient) GetMarketPlaceService(serviceUUID string) (*MarketplaceServiceResp, error) {
	formatedURI := fmt.Sprintf(marketPlaceByUUIDURI, serviceUUID)
	expectedResp := &MarketplaceServiceResp{}
	if _, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp); err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) UpdateMarketPlaceConnection(crCircuitID, routeSetCircuitID string, circuitIDs ConnectionCircuitIDs) error {
	formatedURI := fmt.Sprintf(marketPlaceRouteSetsURI, crCircuitID, routeSetCircuitID)
	_, err := c.sendRequest(formatedURI, putMethod, circuitIDs, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *PFClient) UpdateMarketPlaceServiceRouteSet(crCircuitID, routeSetCircuitID string, service MarketplaceServiceRouteSet) error {
	formatedURI := fmt.Sprintf(marketPlaceUpdateURI, crCircuitID, routeSetCircuitID)
	type MarketPlaceRouteSetUpdate struct {
		Description string      `json:"description"`
		Prefixes    []MktPrefix `json:"prefixes"`
	}
	mktUpdate := MarketPlaceRouteSetUpdate{
		Description: service.RouteSet.Description,
		Prefixes:    service.RouteSet.Prefixes,
	}
	_, err := c.sendRequest(formatedURI, patchMethod, mktUpdate, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *PFClient) UpdateMarketPlaceService(serviceUUID string, service MarketplaceService) error {
	formatedURI := fmt.Sprintf(marketPlaceByUUIDURI, serviceUUID)
	type MarketPlaceServiceUpdate struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Sku         string   `json:"sku"`
		Categories  []string `json:"categories"`
		Published   bool     `json:"published"`
	}
	update := MarketPlaceServiceUpdate{
		Name:        service.Name,
		Description: service.Description,
		Sku:         service.Sku,
		Categories:  service.Categories,
		Published:   service.Published,
	}
	if _, err := c.sendRequest(formatedURI, patchMethod, update, nil); err != nil {
		return err
	}
	return nil
}

func (c *PFClient) DeleteMarketPlaceService(mktUUID string) error {
	if _, uuidErr := uuid.Parse(mktUUID); uuidErr != nil {
		return uuidErr
	}
	_, err := c.sendRequest(fmt.Sprintf(marketPlaceByUUIDURI, mktUUID), deleteMethod, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
