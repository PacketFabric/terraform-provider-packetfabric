package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

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
				ForceNew:    true,
				Description: "The region that the new connection will connect to.\n\n\tExample: us-west-1",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A brief description of this connection.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The desired AWS availability zone of the new connection.\n\n\tExample: \"A\"",
			},
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The POP in which the dedicated port should be provisioned (the cloud on-ramp).\n\n\tExample: DAL1",
			},
			"subscription_term": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The billing term, in months, for this connection.\n\n\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
			},
			"service_class": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The service class for the given port, either long haul or metro. Specify metro if the cloud on-ramp (the `pop`) is in the same market as the source ports (the ports to which you will be building out virtual circuits).\n\n\tEnum: [\"longhaul\",\"metro\"]",
			},
			"autoneg": {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Whether the port auto-negotiates or not. This is currently only possible with 1Gbps ports and the request will fail if specified with 10Gbps.",
			},
			"speed": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The desired capacity of the port.\n\n\tEnum: [\"1Gps\", \"10Gbps\"]",
			},
			"should_create_lag": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "Create the dedicated connection as a LAG interface. ",
			},
			"loa": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "A base64 encoded string of a PDF of the LOA that AWS provided.\n\n\tExample: SSBhbSBhIFBERg==",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Label value linked to an object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"etl": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Early Termination Liability (ETL) fees apply when terminating a service before its term ends. ETL is prorated to the remaining contract days.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsReqDedicatedConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	dedicatedConn := extractAwsDedicatedConn(d)
	expectedResp, err := c.CreateDedicadedAWSConn(dedicatedConn)
	if err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
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

	if labels, ok := d.GetOk("labels"); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}
	return diags
}

func resourceAwsReqDedicatedConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("pop", resp.Pop)
		_ = d.Set("subscription_term", resp.SubscriptionTerm)
		_ = d.Set("service_class", resp.ServiceClass)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("aws_region", resp.Settings.AwsRegion)
	}
	resp2, err2 := c.GetPortByCID(d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}

	if resp2 != nil {
		_ = d.Set("autoneg", resp2.Autoneg)
		if _, ok := d.GetOk("zone"); ok {
			_ = d.Set("zone", resp2.Zone)
		}
		if _, ok := d.GetOk("po_number"); ok {
			_ = d.Set("po_number", resp2.PONumber)
		}
		if resp2.IsLag {
			_ = d.Set("should_create_lag", true)
		} else {
			_ = d.Set("should_create_lag", false)
		}
	}
	// unsetFields: loa

	if _, ok := d.GetOk("labels"); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set("labels", labels)
	}
	etl, err3 := c.GetEarlyTerminationLiability(d.Id())
	if err3 != nil {
		return diag.FromErr(err3)
	}
	if etl > 0 {
		_ = d.Set("etl", etl)
	}
	return diags
}

func resourceAwsReqDedicatedConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesDedicatedUpdate(ctx, d, m)
}

func resourceAwsServicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudSourceDelete(ctx, d, m, "AWS Service Delete")
}

func extractAwsDedicatedConn(d *schema.ResourceData) packetfabric.DedicatedAwsConn {
	dedicatedConn := packetfabric.DedicatedAwsConn{}
	if awsRegion, ok := d.GetOk("aws_region"); ok {
		dedicatedConn.AwsRegion = awsRegion.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		dedicatedConn.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		dedicatedConn.Description = description.(string)
	}
	if zone, ok := d.GetOk("zone"); ok {
		dedicatedConn.Zone = zone.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		dedicatedConn.Pop = pop.(string)
	}
	if subTerm, ok := d.GetOk("subscription_term"); ok {
		dedicatedConn.SubscriptionTerm = subTerm.(int)
	}
	if serviceClass, ok := d.GetOk("service_class"); ok {
		dedicatedConn.ServiceClass = serviceClass.(string)
	}
	if autoneg, ok := d.GetOk("autoneg"); ok {
		dedicatedConn.AutoNeg = autoneg.(bool)
	}
	if speed, ok := d.GetOk("speed"); ok {
		dedicatedConn.Speed = speed.(string)
	}
	if shouldCreateLag, ok := d.GetOk("should_create_lag"); ok {
		dedicatedConn.ShouldCreateLag = shouldCreateLag.(bool)
	}
	if loa, ok := d.GetOk("loa"); ok {
		dedicatedConn.Loa = loa.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		dedicatedConn.PONumber = poNumber.(string)
	}
	return dedicatedConn
}
