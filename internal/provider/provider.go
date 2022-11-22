package provider

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var once sync.Once

var pfClient *packetfabric.PFClient

func Provider() *schema.Provider {
	once.Do(func() {
		schema.DescriptionKind = schema.StringMarkdown
		schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
			desc := s.Description
			if s.Default != nil {
				desc += fmt.Sprintf("Defaults: %v", s.Default)
			}
			return strings.TrimSpace(desc)
		}
	})
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_HOST", "https://api.packetfabric.com"),
				Description: "PacketFabric API endpoint. " +
					"Can also be set with the PF_HOST environment variable. " +
					"Defaults to https://api.packetfabric.com",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_TOKEN", nil),
				Description: "PacketFabric API access token. " +
					"Can also be set with the PF_TOKEN environment variable.",
				Sensitive: true,
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_USER", nil),
				Description: "PacketFabric username. " +
					"Can also be set with the PF_USER environment variable.",
				Sensitive: true,
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_PASSWORD", nil),
				Description: "PacketFabric username. " +
					"Can also be set with the PF_PASSWORD environment variable.",
				Sensitive: true,
			},
		},
		// packetfabric_cloud_router - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_dedicated_connection_post
		// packetfabric_cloud_router_connection_aws - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_dedicated_connection_post
		// packetfabric_cloud_router_bgp_session - https://docs.packetfabric.com/api/v2/redoc/#tag/BGP-Session-Settings
		// packetfabric_cloud_router_bgp_prefixes - https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_create
		// packetfabric_backbone_virtual_circuit - https://docs.packetfabric.com/api/v2/redoc/#operation/post_service_backbone
		// packetfabric_cs_aws_hosted_marketplace_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/post_aws_marketplace_cloud
		// packetfabric_cloud_services_aws_hosted_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_hosted_connection_post
		// packetfabric_cs_aws_provision_marketplace - https://docs.packetfabric.com/api/v2/redoc/#operation/provision_marketplace_cloud
		// packetfabric_cs_aws_hosted_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_hosted_connection_post
		// packetfabric_cs_aws_dedicated_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_dedicated_connection_post
		ResourcesMap: map[string]*schema.Resource{
			"packetfabric_cloud_router":                            resourceCloudRouter(),
			"packetfabric_cloud_router_connection_aws":             resourceRouterConnectionAws(),
			"packetfabric_cloud_router_connection_oracle":          resourceOracleCloudRouteConn(),
			"packetfabric_cloud_router_bgp_session":                resourceBgpSession(),
			"packetfabric_backbone_virtual_circuit":                resourceVcBackbone(),
			"packetfabric_backbone_virtual_circuit_marketplace":    resourceThirdPartyVirtualCircuit(),
			"packetfabric_backbone_virtual_circuit_speed_burst":    resourceAddSpeedBurst(),
			"packetfabric_cs_aws_hosted_marketplace_connection":    resourceAwsHostedMkt(),
			"packetfabric_cs_aws_hosted_connection":                resourceAwsRequestHostConn(),
			"packetfabric_cs_aws_dedicated_connection":             resourceAwsReqDedicatedConn(),
			"packetfabric_port":                                    resourceInterfaces(),
			"packetfabric_link_aggregation_group":                  resourceLinkAggregationGroups(),
			"packetfabric_outbound_cross_connect":                  resourceOutboundCrossConnect(),
			"packetfabric_cs_google_hosted_marketplace_connection": resourceGoogleHostedMktConn(),
			"packetfabric_cs_google_hosted_connection":             resourceGoogleRequestHostConn(),
			"packetfabric_cs_google_dedicated_connection":          resourceGoogleDedicatedConn(),
			"packetfabric_cloud_router_connection_google":          resourceGoogleCloudRouterConn(),
			"packetfabric_cloud_router_connection_azure":           resourceAzureExpressRouteConn(),
			"packetfabric_cloud_router_connection_ibm":             resourceIBMCloudRouteConn(),
			"packetfabric_cs_azure_hosted_marketplace_connection":  resourceAzureHostedMktConn(),
			"packetfabric_cs_azure_hosted_connection":              resourceAzureReqExpressConn(),
			"packetfabric_cs_azure_dedicated_connection":           resourceAzureReqExpressDedicatedConn(),
			"packetfabric_cs_oracle_hosted_connection":             resourceOracleHostedConn(),
			"packetfabric_cs_oracle_hosted_marketplace_connection": resourceOracleMktCloudConn(),
			"packetfabric_cloud_router_connection_ipsec":           resourceIPSecCloudRouteConn(),
			"packetfabric_ix_virtual_circuit_marketplace":          resourceIxVC(),
			"packetfabric_cloud_router_connection_port":            resourceCustomerOwnedPortConn(),
			"packetfabric_point_to_point":                          resourcePointToPoint(),
			"packetfabric_marketplace_service_accept_request":      resourceProvisionRequestedService(),
			"packetfabric_marketplace_service_reject_request":      resourceRejectRequestedService(),
		},
		// packetfabric_cloud_router - https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_list
		// packetfabric_cloud_router_bgp_prefixes - https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_session_settings_list
		// packetfabric_cloud_router_connections - https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_connections
		// packetfabric_cloud_router_bgp_session - https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_session_settings_list
		// packetfabric_cs_aws_hosted_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/get_connection
		// packetfabric_cs_aws_dedicated_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/get_connections_dedicated_list
		// packetfabric_billing - https://docs.packetfabric.com/api/v2/redoc/#operation/get_order
		// packetfabric_locations - https://docs.packetfabric.com/api/v2/redoc/#operation/get_location_list
		DataSourcesMap: map[string]*schema.Resource{
			"packetfabric_cloud_router":                  dataSourceCloudRouter(),
			"packetfabric_cloud_router_connections":      dataSourceCloudConn(),
			"packetfabric_cloud_router_bgp_session":      dataSourceBgpSession(),
			"packetfabric_cs_aws_hosted_connection":      dataSourceCloudServicesConnInfo(),
			"packetfabric_cs_dedicated_connections":      datasourceDedicatedConn(),
			"packetfabric_billing":                       dataSourceBilling(),
			"packetfabric_port":                          datasourceInterfaces(),
			"packetfabric_locations":                     dataSourceLocations(),
			"packetfabric_link_aggregation_group":        datasourceLinkAggregationGroups(),
			"packetfabric_outbound_cross_connect":        dataSourceOutboundCrossConnect(),
			"packetfabric_cs_google_hosted_connection":   dataSourceCloudServicesConnInfo(),
			"packetfabric_cs_azure_hosted_connection":    dataSourceCloudServicesConnInfo(),
			"packetfabric_cs_oracle_hosted_connection":   dataSourceCloudServicesConnInfo(),
			"packetfabric_cloud_router_connection_ipsec": datasourceIPSec(),
			"packetfabric_activitylog":                   datasourceActivityLog(),
			"packetfabric_marketplace_service_requests":  dataSourceVcRequests(),
			"packetfabric_virtual_circuits":              datasourceBackboneServices(),
			"packetfabric_point_to_point":                datasourcePointToPoint(),
			"packetfabric_locations_market":              dataSourceLocationsMarkets(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if pfClient != nil {
		return pfClient, nil
	}
	token, tOk := d.GetOk("token")
	if !tOk {
		_, uOk := d.GetOk("username")
		_, pOk := d.GetOk("password")
		if !uOk || !pOk {
			return nil, diag.Errorf("please provide a valid Token or Username/Password")
		}
	}
	var host *string
	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}
	var diags diag.Diagnostics
	if token != "" {
		tokenStr := token.(string)
		c, err := packetfabric.NewPFClient(host, &tokenStr)
		c.Ctx = ctx
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create PacketFabric client",
				Detail:   "Unable to authenticate user for authenticated PacketFabric client through token",
			})
			return nil, diags
		}
		return c, diags
	}
	c, err := packetfabric.NewPFClientByUserAndPass(ctx, host, d.Get("username").(string), d.Get("password").(string))
	c.Ctx = ctx
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create PacketFabric Client",
			Detail:   err.Error(),
		})
		return nil, diags
	}
	pfClient = c
	return c, diags
}
