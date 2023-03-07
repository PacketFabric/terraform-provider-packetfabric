package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAwsRequestHostConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsReqHostConnCreate,
		UpdateContext: resourceAwsReqHostConnUpdate,
		ReadContext:   resourceAwsReqHostConnRead,
		DeleteContext: resourceAwsServicesDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_AWS_ACCOUNT_ID", nil),
				ValidateFunc: validation.StringIsNotEmpty,
				Description: "The AWS account ID to connect with. Must be 12 characters long. " +
					"Can also be set with the PF_AWS_ACCOUNT_ID environment variable.",
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the hosted connection should be provisioned (the cloud on-ramp).",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the PacketFabric port you want to connect to AWS. This starts with \"PF-AP-\".",
			},
			"vlan": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid VLAN range is from 4-4094, inclusive.",
			},
			"src_svlan": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Valid S-VLAN range is from 4-4094, inclusive.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The desired zone of the new connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The speed of the new connection.\n\n\tAvailable: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
			},
			"po_number": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Purchase order number or identifier of a service.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsReqHostConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	reqConn := extractReqConn(d)
	expectedResp, err := c.CreateAwsHostedConn(reqConn)
	if err != nil {
		return diag.FromErr(err)
	}
	// Cloud Everywhere: if cloud_circuit_id is null display error
	if expectedResp.CloudCircuitID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Hosted location Requested",
			Detail: "On-ramp location does not have a Hosted port currently available. " +
				"Check in the Portal when your hosted cloud is provisioned and import the resource into your Terraform state file.",
		})
		return diags
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			dedicatedConns, err := c.GetCurrentCustomersHosted()
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

func resourceAwsReqHostConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("cloud_circuit_id", resp.CloudCircuitID)
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("vlan", resp.Settings.VlanIDCust)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("pop", resp.CloudProvider.Pop)
		_ = d.Set("aws_account_id", resp.Settings.AwsAccountID)
	}
	resp2, err2 := c.GetBackboneByVcCID(d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	if resp2 != nil {
		_ = d.Set("port", resp2.Interfaces[0].PortCircuitID) // Port A
		if resp2.Interfaces[0].Svlan != 0 {
			_ = d.Set("src_svlan", resp2.Interfaces[0].Svlan) // Port A if ENNI
		}
		_ = d.Set("zone", resp2.Interfaces[1].Zone) // Port Z
		_ = d.Set("po_number", resp2.PONumber)
	}
	return diags
}

func resourceAwsReqHostConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesHostedUpdate(ctx, d, m)
}

func extractReqConn(d *schema.ResourceData) packetfabric.HostedAwsConnection {
	hostedAwsConn := packetfabric.HostedAwsConnection{}
	if awsAccountID, ok := d.GetOk("aws_account_id"); ok {
		hostedAwsConn.AwsAccountID = awsAccountID.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedAwsConn.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		hostedAwsConn.Description = description.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		hostedAwsConn.Pop = pop.(string)
	}
	if port, ok := d.GetOk("port"); ok {
		hostedAwsConn.Port = port.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		hostedAwsConn.Vlan = vlan.(int)
	}
	if srcSvlan, ok := d.GetOk("src_svlan"); ok {
		hostedAwsConn.SrcSvlan = srcSvlan.(int)
	}
	if zone, ok := d.GetOk("zone"); ok {
		hostedAwsConn.Zone = zone.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedAwsConn.Speed = speed.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		hostedAwsConn.PONumber = poNumber.(string)
	}
	return hostedAwsConn
}
