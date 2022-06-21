package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceInterfaces() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		CreateContext: resourceCreateInterface,
		ReadContext:   resourceReadInterface,
		UpdateContext: resourceUpdateInterface,
		DeleteContext: resourceDeleteInterface,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "PacketFabric account UUID. The contact that will be billed.",
			},
			"autoneg": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether the port auto-negotiates or not, this is currently only possible with 1Gbps ports and the request will fail if specified with 10Gbps.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Description for this interface/port",
			},
			"media": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Optic media type.\n\t\tEnum: [\"LX\" \"EX\" \"ZX\" \"LR\" \"ER\" \"ER DWDM\" \"ZR\" \"ZR DWDM\" \"LR4\" \"ER4\" \"CWDM4\" \"LR4\" \"ER4 Lite\"]",
			},
			"nni": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Flag to enable NNI on this port. NNI ports will use a nni_svlan_tpid value of 0x8100.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Point of Presence for the newly created port",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Speed of the port\n\t\tEnum: [\"1Gbps\" \"10Gbps\" \"40Gbps\" \"100Gbps\"]",
			},
			"subscription_term": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "Duration of the subscription in months\n\t\tEnum [\"1\" \"12\" \"24\" \"36\"]",
			},
			"zone": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Availability zone of the port",
			},
		},
	}
}

func resourceCreateInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	interf := extractInterface(d)
	resp, err := c.CreateInterface(interf)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		d.SetId(resp.PortCircuitID)
	}
	return diags
}

func resourceReadInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	portCID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid Port Circuit ID")
	}
	_, err := c.GetPortByCID(portCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUpdateInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	var autoneg bool
	var portCID, description string
	if autonegData, ok := d.GetOk("autoneg"); !ok {
		return diag.Errorf("autoneg is a required field")
	} else {
		autoneg = autonegData.(bool)
	}
	if portCIDData, ok := d.GetOk("id"); !ok {
		return diag.Errorf("port circuit ID is a required field")
	} else {
		portCID = portCIDData.(string)
	}
	if descriptionData, ok := d.GetOk("description"); !ok {
		return diag.Errorf("description is a required field")
	} else {
		description = descriptionData.(string)
	}
	_, err := c.UpdatePort(autoneg, portCID, description)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceDeleteInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	var portCID string
	if portCIDData, ok := d.GetOk("id"); !ok {
		return diag.Errorf("please provide a valid Port Circuit ID")
	} else {
		portCID = portCIDData.(string)
	}
	resp, err := c.DeletePort(portCID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Port Delete",
		Detail:   resp.Message,
	})
}

func extractInterface(d *schema.ResourceData) packetfabric.Interface {
	return packetfabric.Interface{
		AccountUUID:      d.Get("account_uuid").(string),
		Autoneg:          d.Get("autoneg").(bool),
		Description:      d.Get("description").(string),
		Media:            d.Get("media").(string),
		Nni:              d.Get("nni").(bool),
		Pop:              d.Get("pop").(string),
		Speed:            d.Get("speed").(string),
		SubscriptionTerm: d.Get("subscription_term").(int),
		Zone:             d.Get("zone").(string),
	}
}
