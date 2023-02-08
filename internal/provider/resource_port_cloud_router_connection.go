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
				Default:     false,
				Description: "Set this to true if you intend to use NAT on this connection. ",
			},
			"maybe_dnat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Set this to true if you intend to use DNAT on this connection. ",
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
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Valid VLAN range is from 4-4094, inclusive.",
			},
			"untagged": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the interface should be untagged. Do not specify a VLAN if this is to be an untagged connection.",
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
		d.SetId(resp.CloudCircuitID)
	}
	return diags
}

func resourceCustomerOwnedPortConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
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
	return ownedPort
}
