package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudServicesConnInfoAWS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudServicesAwsRead,
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
		},
	}
}

// https://docs.packetfabric.com/api/v2/redoc/#operation/get_connection
func dataSourceCloudServicesAwsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("cloud_circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid cloud_circuit_id")
	}
	service, err := c.GetCloudConnInfo(cID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	_ = d.Set("customer_uuid", service.CustomerUUID)
	_ = d.Set("user_uuid", service.UserUUID)
	_ = d.Set("state", service.State)
	_ = d.Set("service_provider", service.ServiceProvider)
	_ = d.Set("service_class", service.ServiceClass)
	_ = d.Set("port_type", service.PortType)
	_ = d.Set("speed", service.Speed)
	_ = d.Set("description", service.Description)
	_ = d.Set("cloud_provider_pop", service.CloudProvider.Pop)
	_ = d.Set("cloud_provider_region", service.CloudProvider.Region)
	_ = d.Set("time_created", service.TimeCreated)
	_ = d.Set("time_updated", service.TimeUpdated)
	_ = d.Set("pop", service.Pop)
	_ = d.Set("site", service.Site)
	d.SetId(service.CloudCircuitID)
	return diags
}
