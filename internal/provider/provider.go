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
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_HOST", nil),
				Description: "Packet Fabric Cloud API endpoint. " +
					"Example TF files input TF_VAR_pf_api_server shell environment variable. " +
					"Defaults to https://api.packetfabric.com",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_TOKEN", nil),
				Description: "Packet Fabric Cloud API access token. " +
					"Example TF files input TF_VAR_pf_api_key shell environment variable",
				Sensitive: true,
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_USER", nil),
				Sensitive:   true,
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_PASSWORD", nil),
				Sensitive:   true,
			},
		},
		// packetfabric_cloud_router - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_dedicated_connection_post
		// packetfabric_aws_cloud_router_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_dedicated_connection_post
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
			"packetfabric_aws_cloud_router_connection":             resourceRouterConnectionAws(),
			"packetfabric_cloud_router_bgp_session":                resourceBgpSession(),
			"packetfabric_cloud_router_bgp_prefixes":               resourceBgpPrefixes(),
			"packetfabric_backbone_virtual_circuit":                resourceAwsVcBackbone(),
			"packetfabric_cs_aws_hosted_marketplace_connection":    resourceAwsHostedMkt(),
			"packetfabric_cloud_services_aws_hosted_connection":    resourceAwsHostedMktConn(),
			"packetfabric_cs_aws_provision_marketplace":            resourceAwsProvision(),
			"packetfabric_cs_aws_hosted_connection":                resourceAwsRequestHostConn(),
			"packetfabric_cs_aws_dedicated_connection":             resourceAwsReqDedicatedConn(),
			"packetfabric_interface":                               resourceInterfaces(),
			"packetfabric_link_aggregation_group":                  resourceLinkAggregationGroups(),
			"packetfabric_outbound_cross_connect":                  resourceOutboundCrossConnect(),
			"packetfabric_cs_google_hosted_marketplace_connection": resourceGoogleHostedMktConn(),
			"packetfabric_cs_google_provision_marketplace":         resourceGoogleProvision(),
			"packetfabric_cs_google_hosted_connection":             resourceGoogleRequestHostConn(),
			"packetfabric_cs_azure_hosted_marketplace_connection":  resourceAzureHostedMktConn(),
			"packetfabric_cs_azure_provision_marketplace":          resourceAzureProvision(),
			"packetfabric_cs_azure_hosted_connection":              resourceAzureReqExpressConn(),
			"packetfabric_cs_azure_dedicated_connection":           resourceAzureReqExpressDedicatedConn(),
		},
		// packetfabric_cloud_router - https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_list
		// packetfabric_cloud_router_bgp_prefixes - https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_session_settings_list
		// packetfabric_aws_cloud_router_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_connections
		// packetfabric_cloud_router_bgp_session - https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_session_settings_list
		// packetfabric_cs_aws_hosted_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/get_connection
		// packetfabric_cs_aws_dedicated_connection_conn - https://docs.packetfabric.com/api/v2/redoc/#operation/get_connections_dedicated_list
		// packetfabric_billing - https://docs.packetfabric.com/api/v2/redoc/#operation/get_order
		// packetfabric_locations - https://docs.packetfabric.com/api/v2/redoc/#operation/get_location_list
		DataSourcesMap: map[string]*schema.Resource{
			"packetfabric_cloud_router":                           dataSourceCloudRouter(),
			"packetfabric_cloud_router_bgp_prefixes":              dataSourceBgpPrefix(),
			"packetfabric_aws_cloud_router_connection":            dataSourceCloudConn(),
			"packetfabric_cloud_router_bgp_session":               dataSourceBgpSession(),
			"packetfabric_cloud_services_aws_connection_info":     dataSourceCloudServicesConnInfo(),
			"packetfabric_cloud_services_aws_dedicated_conn":      datasourceDedicatedConn(),
			"packetfabric_aws_services_hosted_requested_mkt_conn": datasourceProvisionRequested(),
			"packetfabric_billing":                                dataSourceBilling(),
			"packetfabric_locations":                              dataSourceLocations(),
			"packetfabric_link_aggregation_group":                 datasourceLinkAggregationGroups(),
			"packetfabric_outbound_cross_connect":                 dataSourceOutboundCrossConnect(),
			"packetfabric_gcp_cloud_router_connection":            dataSourceCloudConn(),
			"packetfabric_cloud_services_gcp_connection_info":     dataSourceCloudServicesConnInfo(),
			"packetfabric_cloud_services_gcp_dedicated_conn":      datasourceDedicatedConn(),
			"packetfabric_gcp_services_hosted_requested_mkt_conn": datasourceProvisionRequested(),
			"packetfabric_azr_cloud_router_connection":            dataSourceCloudConn(),
			"packetfabric_cloud_services_azr_connection_info":     dataSourceCloudServicesConnInfo(),
			"packetfabric_cloud_services_azr_dedicated_conn":      datasourceDedicatedConn(),
			"packetfabric_azr_services_hosted_requested_mkt_conn": datasourceProvisionRequested(),
			"packetfabric_activitylog":                            datasourceActivityLog(),
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
