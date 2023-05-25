package provider

import (
	"context"
	"errors"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceMarketplaceService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMarketplaceServiceCreate,
		ReadContext:   resourceMarketplaceServiceRead,
		UpdateContext: resourceMarketplaceServiceUpdate,
		DeleteContext: resourceMarketplaceServiceDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Name of the service.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Brief description of what the service does.",
			},
			"sku": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A SKU identifier for the service. This is not shown to the A side user (the requestor).",
			},
			"locations": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Locations in which the service will operate (port service only). The location should be a POP, e.g. `NYC5`.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},
			"categories": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Categories in which the service will fit.\n\n\tEnum: `\"cloud-computing\"`, `\"content-delivery-network\"`, `\"edge-computing\"`, `\"sd-wan\"`, `\"data-storage\"`, `\"developer-platform\"`, `\"internet-service-provider\"`, `\"security\"`, `\"video-conferencing\"`, `\"voice-and-messaging\"`, `\"web-hosting\"`, `\"internet-of-things\"`, `\"private-connectivity\"`, `\"bare-metal-hosting\"`",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"cloud-computing", "content-delivery-network", "edge-computing", "sd-wan", "data-storage", "developer-platform", "internet-service-provider", "security", "video-conferencing", "voice-and-messaging", "web-hosting", "internet-of-things", "private-connectivity", "bare-metal-hosting"}, true),
				},
			},
			"published": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "If published, the service appears in your marketplace listing.",
			},
			"service_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"port-service", "quick-connect-service"}, true),
				Default:      "port-service",
				Description:  "The service type of this service. Enum: `\"port-service\"`, `\"quick-connect-service\"`\n\n\t",
			},
			"cloud_router_circuit_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the Cloud Router this service is associated with (Quick Connect service only).",
			},
			"route_set": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The route set's description.",
						},
						"is_private": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "In a private route set, the return traffic is private. In other words, in a public route set, anyone who imports this route set can also see other clients who are importing the route based on return traffic. ",
						},
						"prefixes": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "A prefix, in CIDR format, to include in this route set.",
									},
									"match_type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger"}, true),
										Description:  "The match type for this prefix. Options are: `\"exact\"` and `\"orlonger\"`.",
									},
								},
							},
						},
					},
				},
				Description: "The Cloud Router route set to export (Quick Connect service only).",
			},
			"connection_circuit_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The circuit IDs of the Cloud Router connections that will be included in this service. (Quick Connect service only).",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},
			"route_set_circuit_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The route set circuit ID.",
			},
			"service_uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The marketplace service UUID",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The marketplace service state.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceMarketplaceServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if validationDiags := _validate(d); validationDiags != nil {
		return validationDiags
	}
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var resp *packetfabric.MarketplaceServiceResp
	var err error
	if d.Get("service_type").(string) == "quick-connect-service" {
		resp, err = c.CreateMarketplaceServiceWithRouteSet(extractMarketplace(d), extractMarketplaceRouteSet(d))
	} else {
		resp, err = c.CreateMarketplaceService(extractMarketplace(d))
	}
	var diags diag.Diagnostics
	if err != nil {
		return diag.FromErr(err)
	}
	_ = d.Set("service_uuid", resp.UUID)
	_ = d.Set("route_set_circuit_id", resp.RouteSetCircuitID)
	d.SetId(resp.UUID)
	return diags
}

func _validate(d *schema.ResourceData) (diags diag.Diagnostics) {
	if serviceType := d.Get("service_type"); serviceType == "quick-connect-service" {
		if _, ok := d.GetOk("route_set"); !ok {
			return diag.Errorf("Route Sets are required when service_type is cloud-router-service.")
		}
		if ciIDs := d.Get("connection_circuit_ids"); len(ciIDs.([]interface{})) <= 0 {
			return diag.Errorf("Connnection circuit IDs cannot be empty")
		}
	}
	return
}

func resourceMarketplaceServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	serviceUUID := d.Get("service_uuid")
	var diags diag.Diagnostics
	if resp, err := c.GetMarketPlaceService(serviceUUID.(string)); err == nil {
		_ = d.Set("state", resp.State)
	} else {
		return diag.FromErr(err)
	}
	return diags
}

func resourceMarketplaceServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	serviceUUID := d.Get("service_uuid").(string)
	if err := c.UpdateMarketPlaceService(serviceUUID, extractMarketplace(d)); err != nil {
		return diag.FromErr(err)
	}
	serviceType := d.Get("service_type").(string)
	if mktService, err := extractMarketplaceUpdate(d); err == nil && serviceType == "quick-connect-service" {
		if err := c.UpdateMarketPlaceServiceRouteSet(mktService.CloudRouterCircuitID, mktService.RouteSetCircuitID, mktService); err != nil {
			return diag.FromErr(err)
		} else {
			if err := c.UpdateMarketPlaceConnection(mktService.CloudRouterCircuitID, mktService.RouteSetCircuitID, _extractMapOfConnectionCircuiIDs(d, "connection_circuit_ids")); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return diag.Diagnostics{}
}

func resourceMarketplaceServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if serviceType := d.Get("service_type"); serviceType != "quick-connect-service" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Marketplace delete",
			Detail:   "Marketplace cannot be deleted when service_type is not cloud-router-service.",
		})
		return diags
	}
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	err := c.DeleteMarketPlaceService(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func extractMarketplace(d *schema.ResourceData) packetfabric.MarketplaceService {
	mkt := packetfabric.MarketplaceService{}
	if name, ok := d.GetOk("name"); ok {
		mkt.Name = name.(string)
	}
	if desc, ok := d.GetOk("description"); ok {
		mkt.Description = desc.(string)
	}
	if state, ok := d.GetOk("state"); ok {
		mkt.State = state.(string)
	}
	if sku, ok := d.GetOk("sku"); ok {
		mkt.Sku = sku.(string)
	}
	if published, ok := d.GetOk("published"); ok {
		mkt.Published = published.(bool)
	}

	mkt.Locations = _extractList(d, "locations")
	mkt.Categories = _extractList(d, "categories")
	if serviceType, ok := d.GetOk("service_type"); ok {
		mkt.ServiceType = serviceType.(string)
	}
	return mkt
}

func extractMarketplaceRouteSet(d *schema.ResourceData) packetfabric.MarketplaceServiceRouteSet {
	mkt := packetfabric.MarketplaceServiceRouteSet{}
	mkt.ConnectionCircuitIDs = _extractMapOfConnectionCircuiIDs(d, "connection_circuit_ids")
	mkt.RouteSet = extractRouteSet(d)
	if cloudRouterCID, ok := d.GetOk("cloud_router_circuit_id"); ok {
		mkt.CloudRouterCircuitID = cloudRouterCID.(string)
	}
	return mkt
}

func extractMarketplaceUpdate(d *schema.ResourceData) (packetfabric.MarketplaceServiceRouteSet, error) {
	mkt := packetfabric.MarketplaceServiceRouteSet{}
	if serviceType, ok := d.GetOk("service_type"); ok {
		if serviceType == "quick-connect-service" {
			mkt.RouteSet = extractRouteSet(d)
		}
	}
	if cloudRouterCID, ok := d.GetOk("cloud_router_circuit_id"); ok {
		mkt.CloudRouterCircuitID = cloudRouterCID.(string)
	} else {
		return mkt, errors.New("cloud_router_circuit_id cannot be empty")
	}
	if routeSetCircuitID, ok := d.GetOk("route_set_circuit_id"); ok {
		mkt.RouteSetCircuitID = routeSetCircuitID.(string)
	}
	return mkt, nil

}

func _extractList(d *schema.ResourceData, key string) []string {
	if locations, ok := d.GetOk(key); ok {
		locs := make([]string, 0)
		for _, loc := range locations.([]interface{}) {
			locs = append(locs, loc.(string))
		}
		return locs
	}
	return make([]string, 0)
}

func _extractMapOfConnectionCircuiIDs(d *schema.ResourceData, key string) packetfabric.ConnectionCircuitIDs {
	if locations, ok := d.GetOk(key); ok {
		locs := make(packetfabric.ConnectionCircuitIDs)
		ids := make([]string, 0)
		for _, loc := range locations.([]interface{}) {
			ids = append(ids, loc.(string))
		}
		locs[key] = ids
		return locs
	}
	return packetfabric.ConnectionCircuitIDs{}
}

func extractRouteSet(d *schema.ResourceData) packetfabric.RouteSet {
	var routeSet packetfabric.RouteSet
	if sets, ok := d.GetOk("route_set"); ok {
		for _, set := range sets.(*schema.Set).List() {
			desc := ""
			isPrivate := true
			if descSet := set.(map[string]interface{})["description"]; descSet != nil {
				desc = descSet.(string)
			}
			if isPrivateSet := set.(map[string]interface{})["is_private"]; isPrivateSet != nil {
				isPrivate = isPrivateSet.(bool)
			}
			routeSet = packetfabric.RouteSet{
				Description: desc,
				IsPrivate:   isPrivate,
				Prefixes:    extractyRouteSetPrefixes(set.(map[string]interface{})["prefixes"]),
			}

		}
	}
	return routeSet
}

func extractyRouteSetPrefixes(prefixes interface{}) []packetfabric.QuickConnectImportFilters {
	routePrefixes := make([]packetfabric.QuickConnectImportFilters, 0)
	for _, prefix := range prefixes.(*schema.Set).List() {
		routePrefix := packetfabric.QuickConnectImportFilters{}
		if p := prefix.(map[string]interface{})["prefix"]; p != nil {
			routePrefix.Prefix = p.(string)
		}
		if m := prefix.(map[string]interface{})["match_type"]; m != nil {
			routePrefix.MatchType = m.(string)
		}
		routePrefixes = append(routePrefixes, routePrefix)
	}
	return routePrefixes
}
