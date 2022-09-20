package provider

import (
	"context"
	"errors"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterConnectionAws() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterConnectionAwsCreate,
		ReadContext:   resourceRouterConnectionAwsRead,
		UpdateContext: resourceRouterConnectionAwsUpdate,
		DeleteContext: resourceRouterConnectionAwsDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"aws_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The AWS account ID to connect with. Must be 12 characters long.",
			},
			"account_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The UUID for the billing account that should be billed.",
			},
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set this to true if you intend to use NAT on this connection.\n\n\tDefaults to false if unspecified.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A brief description of this connection.",
			},
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The POP in which you want to provision the connection.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The desired AWS availability zone of the new connection.",
			},
			"is_public": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether PacketFabric should allocate a public IP address for this connection. Set this to true if you intend to use a public VIF on the AWS side.\n\n\tDefaults to false if unspecified.",
			},
			"speed": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The desired speed of the new connection.\n\n\t Available: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceRouterConnectionAwsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	awsConn := extractAwsConnection(d)
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}
	conn, err := c.CreateAwsConnection(awsConn, cID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	createOkCh := make(chan bool)
	defer close(createOkCh)
	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetCloudConnectionStatus(cID.(string), conn.CloudCircuitID)
	}
	go c.CheckServiceStatus(createOkCh, err, fn)
	if !<-createOkCh {
		return diag.FromErr(err)
	}
	if conn != nil {
		_ = d.Set("port_type", conn.PortType)
		_ = d.Set("speed", conn.Speed)
		_ = d.Set("state", conn.State)
		_ = d.Set("cloud_circuit_id", conn.CloudCircuitID)
		_ = d.Set("account_uuid", conn.AccountUUID)
		_ = d.Set("service_class", conn.ServiceClass)
		_ = d.Set("service_provider", conn.ServiceProvider)
		_ = d.Set("description", conn.Description)
		_ = d.Set("user_uuid", conn.UserUUID)
		_ = d.Set("customer_uuid", conn.CustomerUUID)
		_ = d.Set("uuid", conn.UUID)
		d.SetId(conn.CloudCircuitID)
	}
	return diags
}

func resourceRouterConnectionAwsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}
	conns, err := c.ListAwsRouterConnections(cID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	for _, conn := range conns {
		_ = d.Set("port_type", conn.PortType)
		_ = d.Set("connection_type", conn.ConnectionType)
		_ = d.Set("port_circuit_id", conn.PortCircuitID)
		_ = d.Set("pending_delete", conn.PendingDelete)
		_ = d.Set("deleted", conn.Deleted)
		_ = d.Set("speed", conn.Speed)
		_ = d.Set("state", conn.State)
		_ = d.Set("cloud_circuit_id", conn.CloudCircuitID)
		_ = d.Set("account_uuid", conn.AccountUUID)
		_ = d.Set("service_class", conn.ServiceClass)
		_ = d.Set("service_provider", conn.ServiceProvider)
		_ = d.Set("service_type", conn.ServiceType)
		_ = d.Set("description", conn.Description)
		_ = d.Set("cloud_provider_connection_id", conn.CloudProviderConnectionID)
		_ = d.Set("user_uuid", conn.UserUUID)
		_ = d.Set("customer_uuid", conn.CustomerUUID)
		_ = d.Set("time_created", conn.TimeCreated)
		_ = d.Set("time_updated", conn.TimeUpdated)
		_ = d.Set("pop", conn.Pop)
		_ = d.Set("site", conn.Site)
		_ = d.Set("cloud_router_circuit_id", conn.CloudRouterCircuitID)
		_ = d.Set("nat_capable", conn.NatCapable)
		_ = d.Set("uuid", conn.UUID)
	}
	return diags
}

func resourceRouterConnectionAwsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}
	connCid := d.Get("id").(string)
	description := d.Get("description").(string)
	_, err := c.UpdateAwsConnection(cID.(string), connCid, packetfabric.DescriptionUpdate{Description: description})
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceRouterConnectionAwsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}
	connCID := d.Get("id").(string)
	resp, err := c.DeleteAwsConnection(cID.(string), connCID)
	if err != nil {
		return diag.FromErr(err)
	}
	deleteOkCh := make(chan bool)
	defer close(deleteOkCh)
	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetCloudConnectionStatus(cID.(string), connCID)
	}
	go c.CheckServiceStatus(deleteOkCh, err, fn)
	if !<-deleteOkCh {
		return diag.FromErr(err)
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "AWS Router connection result",
		Detail:   resp.Message,
	})
	d.SetId("")
	return diags
}

func extractAwsConnection(d *schema.ResourceData) packetfabric.AwsConnection {
	return packetfabric.AwsConnection{
		AwsAccountID: d.Get("aws_account_id").(string),
		AccountUUID:  d.Get("account_uuid").(string),
		MaybeNat:     d.Get("maybe_nat").(bool),
		Description:  d.Get("description").(string),
		Pop:          d.Get("pop").(string),
		Zone:         d.Get("zone").(string),
		IsPublic:     d.Get("is_public").(bool),
		Speed:        d.Get("speed").(string),
	}
}
