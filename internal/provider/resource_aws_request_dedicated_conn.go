package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const cloudCidNotFoundSummaryMsg = "cloud_circuit_id not created yet"
const cloudCidNotFoundDetailsMsg = "Please wait few minutes then run: terraform refresh"

func resourceAwsReqDedicatedConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		CreateContext: resourceAwsReqDedicatedConnCreate,
		UpdateContext: resourceAwsReqDedicatedConnUpdate,
		ReadContext:   resourceAwsReqDedicatedConnRead,
		DeleteContext: resourceAwsServicesDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The region that the new connection will connect to.\n\t\tExample: us-west-1",
			},
			"account_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The UUID of the contact that will be billed.\n\t\tExample: a2115890-ed02-4795-a6dd-c485bec3529c",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of this connection.\n\t\tExample: AWS Hosted connection for Foo Corp",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The desired AWS Availability zone of the new connection.\n\t\tExample: \"A\"",
			},
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The desired location for the new AWS Hosted Connection.\n\t\tExample: DAL1",
			},
			"subscription_term": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The billing term, in months, for this connection.\n\t\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
			},
			"service_class": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The service class for the given port, either long haul or metro.\n\t\tEnum: [\"longhaul\",\"metro\"]",
			},
			"autoneg": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether the port auto-negotiates or not, this is currently only possible with 1Gbps ports and the request will fail if specified with 10Gbps.",
			},
			"speed": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The desired speed of the new connection.\n\t\tEnum: []\"1gps\", \"10gbps\"]",
			},
			"should_create_lag": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Create the dedicated connection as a LAG interface.",
			},
			"loa": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A base64 encoded string of a PDF of a LOA\n\t\tExample: SSBhbSBhIFBERg==",
			},
		},
	}
}

func resourceAwsReqDedicatedConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	dedicatedConn := extractDedicatedConn(d)
	expectedResp, err := c.CreateDedicadedAWSConn(dedicatedConn)
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			dedicatedConns, err := c.GetCurrentCustomersDedicated()
			if dedicatedConns != nil && err == nil && len(dedicatedConns) > 0 {
				for _, conn := range dedicatedConns {
					if expectedResp.UUID == conn.UUID && conn.State == "active" {
						expectedResp.CloudCircuitID = conn.CloudCircuitID
						ticker.Stop()
						createOk <- true
					}
				}
			}
		}
	}()
	<-createOk
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(expectedResp.CloudCircuitID)
	return diags
}

func resourceAwsReqDedicatedConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceAwsReqDedicatedConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceAwsServicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCID, ok := d.GetOk("id")
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Aws Services Delete",
			Detail:   cloudCidNotFoundDetailsMsg,
		})
		return diags
	}
	err := c.DeleteCloudService(cloudCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	deleteOkCh := make(chan bool)
	defer close(deleteOkCh)
	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetCloudServiceStatus(cloudCID.(string))
	}
	go c.CheckServiceStatus(deleteOkCh, err, fn)
	if !<-deleteOkCh {
		return diag.FromErr(err)
	}
	return diags
}

func extractDedicatedConn(d *schema.ResourceData) packetfabric.DedicatedAwsConn {
	return packetfabric.DedicatedAwsConn{
		AwsRegion:        d.Get("aws_region").(string),
		AccountUUID:      d.Get("account_uuid").(string),
		Description:      d.Get("description").(string),
		Zone:             d.Get("zone").(string),
		Pop:              d.Get("pop").(string),
		SubscriptionTerm: d.Get("subscription_term").(int),
		ServiceClass:     d.Get("service_class").(string),
		AutoNeg:          d.Get("autoneg").(bool),
		Speed:            d.Get("speed").(string),
		ShouldCreateLag:  d.Get("should_create_lag").(bool),
		Loa:              d.Get("loa").(interface{}),
	}
}
