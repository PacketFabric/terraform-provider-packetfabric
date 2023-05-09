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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The connection UUID.",
			},
			"account_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The account UUID.",
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
			"deleted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if the service connection is deleted.",
			},
			"cloud_provider": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pop": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Point of Presence for the cloud provider location\n\t\tExample: LAX1",
						},
						"site": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Region short name\n\t\tExample: us-west-1",
						},
					},
				},
			},
			"vlan_id_pf": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_id_cust": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"svlan_id_cust": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bgp_asn": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bgp_cer_cidr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bgp_ibm_cidr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_term": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The subscription term in months.",
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
			"is_cloud_router_connection": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if this is a cloud router connection",
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
		_ = d.Set("cloud_circuit_id", ibmConn.CloudCircuitID)
		_ = d.Set("uuid", ibmConn.UUID)
		_ = d.Set("customer_uuid", ibmConn.CustomerUUID)
		_ = d.Set("user_uuid", ibmConn.UserUUID)
		_ = d.Set("account_uuid", ibmConn.AccountUUID)
		_ = d.Set("state", ibmConn.State)
		_ = d.Set("service_provider", ibmConn.ServiceProvider)
		_ = d.Set("service_class", ibmConn.ServiceClass)
		_ = d.Set("port_type", ibmConn.PortType)
		_ = d.Set("deleted", ibmConn.Deleted)
		_ = d.Set("speed", ibmConn.Speed)
		_ = d.Set("description", ibmConn.Description)
		_ = d.Set("cloud_provider_pop", ibmConn.CloudProvider.Pop)
		_ = d.Set("time_created", ibmConn.TimeCreated)
		_ = d.Set("time_updated", ibmConn.TimeUpdated)
		_ = d.Set("pop", ibmConn.Pop)
		_ = d.Set("site", ibmConn.Site)
		_ = d.Set("subscription_term", ibmConn.SubscriptionTerm)
		_ = d.Set("customer_site_name", ibmConn.CustomerSiteName)
		_ = d.Set("customer_site_code", ibmConn.CustomerSiteCode)
		_ = d.Set("is_awaiting_onramp", ibmConn.IsAwaitingOnramp)
		_ = d.Set("is_cloud_router_connection", ibmConn.IsCloudRouterConnection)
		_ = d.Set("cloud_provider", flattenCloudProvider(&ibmConn.CloudProvider))
		_ = d.Set("vlan_id_pf", ibmConn.Settings.VlanIDPf)
		_ = d.Set("vlan_id_cust", ibmConn.Settings.VlanIDCust)
		_ = d.Set("svlan_id_cust", ibmConn.Settings.SvlanIDCust)
		_ = d.Set("account_id", ibmConn.Settings.AccountID)
		_ = d.Set("gateway_id", ibmConn.Settings.GatewayID)
		_ = d.Set("port_id", ibmConn.Settings.PortID)
		_ = d.Set("name", ibmConn.Settings.Name)
		_ = d.Set("bgp_asn", ibmConn.Settings.BgpAsn)
		_ = d.Set("bgp_cer_cidr", ibmConn.Settings.BgpCerCidr)
		_ = d.Set("bgp_ibm_cidr", ibmConn.Settings.BgpIbmCidr)
		d.SetId(cCID.(string) + "-data")
	}
	return diags
}
