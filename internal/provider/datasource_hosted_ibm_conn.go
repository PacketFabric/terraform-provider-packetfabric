package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceHostedIBMConn() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHostedIbmConnRead,
		Schema: map[string]*schema.Schema{
			"cloud_circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique PF circuit ID for this connection\n\t\tExample: PF-AP-LAX1-1002",
			},
			"customer_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The UUID for the customer this connection belongs to.",
			},
			"user_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The UUID for the user this connection belongs to.",
			},
			"state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The state of the connection.\n\t\tEnum: [ \"active\", \"deleting\", \"inactive\", \"pending\", \"requested\" ]",
			},
			"service_provider": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The service provider of the connection\n\t\tEnum: [ \"aws\", \"azure\", \"packet\", \"google\", \"ibm\", \"salesforce\", \"webex\" ]",
			},
			"service_class": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The service class for the given port, either long haul or metro.\n\t\tEnum: [ \"longhaul\", \"metro\" ]",
			},
			"port_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port type for the given port.\n\t\tEnum: [ \"hosted\", \"dedicated\" ]",
			},
			"speed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The desired speed of the connection.\n\t\tEnum: [ \"50Mbps\", \"100Mbps\", \"200Mbps\", \"300Mbps\", \"400Mbps\", \"500Mbps\", \"1Gbps\", \"2Gbps\", \"5Gbps\", \"10Gbps\" ]",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this connection.\n\t\tExample: AWS connection for Foo Corp.",
			},
			"cloud_provider_pop": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Point of Presence for the cloud provider location.\n\t\tExample: DAL1",
			},
			"cloud_provider_region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Region short name.\n\t\tExample: us-west-1",
			},
			"time_created": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Date and time of connection creation",
			},
			"time_updated": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Date and time connection was last updated",
			},
			"pop": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Point of Presence for the connection.\n\t\tExample: LAS1",
			},
			"site": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Site name\n\t\tExample: SwitchNAP Las Vegas 7",
			},
			"customer_site_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The customer site name",
			},
			"customer_site_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The customer site code",
			},
			"is_awaiting_onramp": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if this connection is waiting on RAMP",
			},
		},
	}
}

func dataSourceHostedIbmConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cCID, ok := d.GetOk("cloud_circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid cloud circuit ID")
	}
	ibmConn, err := c.GetCloudConnInfo(cCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	if ibmConn != nil {
		_ = d.Set("customer_uuid", ibmConn.CustomerUUID)
		_ = d.Set("user_uuid", ibmConn.UserUUID)
		_ = d.Set("state", ibmConn.State)
		_ = d.Set("service_provider", ibmConn.ServiceProvider)
		_ = d.Set("service_class", ibmConn.ServiceClass)
		_ = d.Set("port_type", ibmConn.PortType)
		_ = d.Set("speed", ibmConn.Speed)
		_ = d.Set("description", ibmConn.Description)
		_ = d.Set("cloud_provider_pop", ibmConn.CloudProvider.Pop)
		_ = d.Set("cloud_provider_region", ibmConn.CloudProvider.Region)
		_ = d.Set("time_created", ibmConn.TimeCreated)
		_ = d.Set("time_updated", ibmConn.TimeUpdated)
		_ = d.Set("pop", ibmConn.Pop)
		_ = d.Set("site", ibmConn.Site)
		_ = d.Set("customer_site_name", ibmConn.CustomerSiteName)
		_ = d.Set("customer_site_code", ibmConn.CustomerSiteCode)
		_ = d.Set("is_awaiting_onramp", ibmConn.IsAwaitingOnramp)
		d.SetId(cCID.(string))
	}
	return diags
}
