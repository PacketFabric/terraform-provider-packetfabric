package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCustomerOwnedPortConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCustomerOwnedPortConnCreate,
		ReadContext:   resourceCustomerOwnedPortConnRead,
		UpdateContext: resourceCustomerOwnedPortConnUpdate,
		DeleteContext: resourceCustomerOwnedPortConnDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
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
			"circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set this to true if you intend to use NAT on this connection. Defaults: false",
			},
			"maybe_dnat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set this to true if you intend to use DNAT on this connection. Defaults: false",
			},
			"port_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the port to connect to the cloud router. This starts with \"PF-AP-\".",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"vlan": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid VLAN range is from 4-4094, inclusive. ",
			},
			"untagged": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether the interface should be untagged. Do not specify a VLAN if this is to be an untagged connection. ",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
			"is_public": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether PacketFabric should allocate a public IP address for this connection. ",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"labels": {
				Type:        schema.TypeList,
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
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceCustomerOwnedPortConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ownedPort := extractOwnedPortConn(d)
	if cID, ok := d.GetOk("circuit_id"); ok {
		resp, err := c.AttachCustomerOwnedPortToCR(ownedPort, cID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		createOkCh := make(chan bool)
		defer close(createOkCh)
		fn := func() (*packetfabric.ServiceState, error) {
			return c.GetCloudConnectionStatus(cID.(string), resp.CloudCircuitID)
		}
		go c.CheckServiceStatus(createOkCh, fn)
		if !<-createOkCh {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.CloudCircuitID)

			if labels, ok := d.GetOk("labels"); ok {
				diagnostics, created := createLabels(c, d.Id(), labels)
				if !created {
					return diagnostics
				}
			}
		}
	}
	return diags
}

func resourceCustomerOwnedPortConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		resp, err := c.ReadCloudRouterConnection(cid.(string), cloudConnCID.(string))
		if err != nil {
			diags = diag.FromErr(err)
			return diags
		}

		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("circuit_id", resp.CloudRouterCircuitID)
		_ = d.Set("port_circuit_id", resp.PortCircuitID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("vlan", resp.Vlan)
		_ = d.Set("speed", resp.Speed)
		if _, ok := d.GetOk("po_number"); ok {
			_ = d.Set("po_number", resp.PONumber)
		}
		if resp.CloudSettings.PublicIP != "" {
			_ = d.Set("is_public", true)
		} else {
			_ = d.Set("is_public", false)
		}
		if resp.Vlan == 0 {
			_ = d.Set("untagged", true)
		} else {
			_ = d.Set("untagged", false)
		}
		// unsetFields: published_quote_line_uuid
	}

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

func resourceCustomerOwnedPortConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceCustomerOwnedPortConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractOwnedPortConn(d *schema.ResourceData) packetfabric.CustomerOwnedPort {
	ownedPort := packetfabric.CustomerOwnedPort{}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		ownedPort.AccountUUID = accountUUID.(string)
	}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		ownedPort.MaybeNat = maybeNat.(bool)
	}
	if maybeDNat, ok := d.GetOk("maybe_dnat"); ok {
		ownedPort.MaybeDNat = maybeDNat.(bool)
	}
	if portCircuitID, ok := d.GetOk("port_circuit_id"); ok {
		ownedPort.PortCircuitID = portCircuitID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		ownedPort.Description = description.(string)
	}
	if untagged, ok := d.GetOk("untagged"); ok {
		ownedPort.Untagged = untagged.(bool)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		ownedPort.Vlan = vlan.(int)
	}
	if speed, ok := d.GetOk("speed"); ok {
		ownedPort.Speed = speed.(string)
	}
	if isPublic, ok := d.GetOk("is_public"); ok {
		ownedPort.IsPublic = isPublic.(bool)
	}
	if publishedQuote, ok := d.GetOk("published_quote_line_uuid"); ok {
		ownedPort.PublishedQuoteLineUUID = publishedQuote.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		ownedPort.PONumber = poNumber.(string)
	}
	return ownedPort
}
