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
				ValidateFunc: validation.StringIsNotEmpty,
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
		// cloud_router - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_dedicated_connection_post
		// aws_cloud_router_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_dedicated_connection_post
		// cloud_router_bgp_session - https://docs.packetfabric.com/api/v2/redoc/#tag/BGP-Session-Settings
		// cloud_router_bgp_prefixes - https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_create
		// cloud_services_aws_create_backbone_dedicated_cr - https://docs.packetfabric.com/api/v2/redoc/#operation/post_service_backbone
		// cloud_services_aws_hosted_marketplace - https://docs.packetfabric.com/api/v2/redoc/#operation/post_aws_marketplace_cloud
		// cloud_services_aws_hosted_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_hosted_connection_post
		// cloud_services_aws_provision_requested_mkt_conn - https://docs.packetfabric.com/api/v2/redoc/#operation/provision_marketplace_cloud
		// cloud_services_aws_req_hosted_conn - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_hosted_connection_post
		// cloud_services_aws_dedicated - https://docs.packetfabric.com/api/v2/redoc/#operation/aws_dedicated_connection_post
		ResourcesMap: map[string]*schema.Resource{
			"cloud_router":                                    resourceCloudRouter(),
			"aws_cloud_router_connection":                     resourceRouterConnectionAws(),
			"cloud_router_bgp_session":                        resourceBgpSession(),
			"cloud_router_bgp_prefixes":                       resourceBgpPrefixes(),
			"cloud_services_aws_create_backbone_dedicated_cr": resourceAwsBackbone(),
			"cloud_services_aws_hosted_marketplace":           resourceAwsHostedMkt(),
			"cloud_services_aws_hosted_connection":            resourceAwsHostedMktConn(),
			"cloud_services_aws_provision_requested_mkt_conn": resourceAwsProvision(),
			"cloud_services_aws_req_hosted_conn":              resourceAwsRequestHostConn(),
			"cloud_services_aws_dedicated":                    resourceAwsReqDedicatedConn(),
			"interface":                                       resourceInterfaces(),
		},
		// cloud_router - https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_list
		// cloud_router_bgp_prefixes - https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_session_settings_list
		// aws_cloud_router_connection - https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_connections
		// cloud_router_bgp_session - https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_session_settings_list
		// cloud_services_aws_connection_info - https://docs.packetfabric.com/api/v2/redoc/#operation/get_connection
		// cloud_services_aws_dedicated_conn - https://docs.packetfabric.com/api/v2/redoc/#operation/get_connections_dedicated_list
		// billing - https://docs.packetfabric.com/api/v2/redoc/#operation/get_order
		// locations - https://docs.packetfabric.com/api/v2/redoc/#operation/get_location_list
		DataSourcesMap: map[string]*schema.Resource{
			"cloud_router":                           dataSourceCloudRouter(),
			"cloud_router_bgp_prefixes":              dataSourceBgpPrefix(),
			"aws_cloud_router_connection":            dataSourceCloudConnAws(),
			"cloud_router_bgp_session":               dataSourceBgpSession(),
			"cloud_services_aws_connection_info":     dataSourceCloudServicesConnInfoAWS(),
			"cloud_services_aws_dedicated_conn":      datasourceAwsDedicatedConn(),
			"aws_services_hosted_requested_mkt_conn": datasourceAwsProvisionRequested(),
			"billing":                                dataSourceBilling(),
			"locations":                              dataSourceLocations(),
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
