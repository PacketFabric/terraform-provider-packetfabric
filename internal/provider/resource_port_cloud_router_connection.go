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
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The UUID of the contact that will be billed.",
			},
			"circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not this connection is intended for NAT later.",
			},
			"port_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the port to connect to the cloud router.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of this connection.",
			},
			"vlan": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The desired vlan to use on the customer-owned port.",
			},
			"untagged": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not the customer port is untagged. Do not specify a VLAN if this is to be an untagged connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The desired speed of the new connection.",
			},
			"is_public": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not PacketFabric should allocate an IP address for the user.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
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
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceCustomerOwnedPortConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudSourceDelete(ctx, d, m, "Customer Owned Port attach Service Delete")
}

func extractOwnedPortConn(d *schema.ResourceData) packetfabric.CustomerOwnedPort {
	ownedPort := packetfabric.CustomerOwnedPort{}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		ownedPort.AccountUUID = accountUUID.(string)
	}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		ownedPort.MaybeNat = maybeNat.(bool)
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
