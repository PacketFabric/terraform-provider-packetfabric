package provider

import (
	"context"
	"errors"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				ForceNew:    true,
				Description: "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"aws_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_AWS_ACCOUNT_ID", nil),
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
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Set this to true if you intend to use NAT on this connection. Default: false.",
			},
			"maybe_dnat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Set this to true if you intend to use DNAT on this connection. Default: false.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A brief description of this connection.",
			},
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The POP in which you want to provision the connection.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The desired AWS availability zone of the new connection.",
			},
			"is_public": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether PacketFabric should allocate a public IP address for this connection. Set this to true if you intend to use a public VIF on the AWS side. ",
			},
			"speed": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The desired speed of the new connection.\n\n\t Available: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line which this connection should be associated.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
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
	go c.CheckServiceStatus(createOkCh, fn)
	if !<-createOkCh {
		return diag.FromErr(err)
	}
	if conn != nil {
		_ = d.Set("speed", conn.Speed)
		_ = d.Set("account_uuid", conn.AccountUUID)
		d.SetId(conn.CloudCircuitID)
	}
	return diags
}

func resourceRouterConnectionAwsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	circuitID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}

	cloudConnCID := d.Get("id")
	resp, err := c.ReadCloudRouterConnection(circuitID.(string), cloudConnCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("account_uuid", resp.AccountUUID)
	_ = d.Set("circuit_id", resp.CloudRouterCircuitID)
	_ = d.Set("maybe_nat", resp.NatCapable)
	_ = d.Set("maybe_dnat", resp.DNatCapable)
	_ = d.Set("description", resp.Description)
	_ = d.Set("speed", resp.Speed)
	_ = d.Set("pop", resp.Pop)
	_ = d.Set("zone", resp.Zone)
	_ = d.Set("aws_account_id", resp.CloudSettings.AwsAccountID)

	if resp.CloudSettings.PublicIP != "" {
		_ = d.Set("is_public", true)
	} else {
		_ = d.Set("is_public", false)
	}
	// unsetFields: published_quote_line_uuid
	return diags
}

func resourceRouterConnectionAwsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceRouterConnectionAwsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractAwsConnection(d *schema.ResourceData) packetfabric.AwsConnection {
	return packetfabric.AwsConnection{
		AwsAccountID:           d.Get("aws_account_id").(string),
		AccountUUID:            d.Get("account_uuid").(string),
		MaybeNat:               d.Get("maybe_nat").(bool),
		MaybeDNat:              d.Get("maybe_dnat").(bool),
		Description:            d.Get("description").(string),
		Pop:                    d.Get("pop").(string),
		Zone:                   d.Get("zone").(string),
		IsPublic:               d.Get("is_public").(bool),
		Speed:                  d.Get("speed").(string),
		PublishedQuoteLineUUID: d.Get("published_quote_line_uuid").(string),
	}
}
