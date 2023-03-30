package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceHostedCloudRouterConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceHostedCloudRouterConfigRead,
		Schema: map[string]*schema.Schema{
			"cloud_circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique PF circuit ID for this connection\n\t\tExample: PF-AP-LAX1-1002",
			},
			"router_type": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Router Type Identifier.\n\n\tEnum: [\"CiscoSystemsInc-2900SeriesRouters-IOS124\", \"CiscoSystemsInc-3700SeriesRouters-IOS124\", \"CiscoSystemsInc-7200SeriesRouters-IOS124\", \"CiscoSystemsInc-Nexus7000SeriesSwitches-NXOS51\", \"CiscoSystemsInc-Nexus9KSeriesSwitches-NXOS93\", \"JuniperNetworksInc-MMXSeriesRouters-JunOS95\", \"JuniperNetworksInc-SRXSeriesRouters-JunOS95\", \"JuniperNetworksInc-TSeriesRouters-JunOS95\", \"PaloAltoNetworks-PA3000and5000series-PANOS803\"]",
				ValidateFunc: validation.StringInSlice(getValidRouterTypes(), false),
			},
			"router_config": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The router configuration returned by the API.",
			},
		},
	}
}

func datasourceHostedCloudRouterConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCircuitID, ok := d.GetOk("cloud_circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid Cloud Circuit ID")
	}
	routerType, ok := d.GetOk("router_type")
	if !ok {
		return diag.Errorf("please provide a valid Router Type")
	}

	routerConfig, err := c.GetRouterConfiguration(cloudCircuitID.(string), routerType.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("router_config", routerConfig.RouterConfig)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cloudCircuitID.(string))

	return diags
}

func getValidRouterTypes() []string {
	return []string{
		"CiscoSystemsInc-2900SeriesRouters-IOS124",
		"CiscoSystemsInc-3700SeriesRouters-IOS124",
		"CiscoSystemsInc-7200SeriesRouters-IOS124",
		"CiscoSystemsInc-Nexus7000SeriesSwitches-NXOS51",
		"CiscoSystemsInc-Nexus9KSeriesSwitches-NXOS93",
		"JuniperNetworksInc-MMXSeriesRouters-JunOS95",
		"JuniperNetworksInc-SRXSeriesRouters-JunOS95",
		"JuniperNetworksInc-TSeriesRouters-JunOS95",
		"PaloAltoNetworks-PA3000and5000series-PANOS803",
	}
}
