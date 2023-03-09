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
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},
			"autoneg": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Only applicable to 1Gbps ports. Controls whether auto negotiation is on (true) or off (false). The request will fail if specified with 10Gbps. ",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of the port.",
			},
			"media": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Optic media type.\n\n\tEnum: [\"LX\" \"EX\" \"ZX\" \"LR\" \"ER\" \"ER DWDM\" \"ZR\" \"ZR DWDM\" \"LR4\" \"ER4\" \"CWDM4\" \"LR4\" \"ER4 Lite\"]",
			},
			"nni": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Set this to true to provision an ENNI port. ENNI ports will use a nni_svlan_tpid value of 0x8100.\n\n\tBy default, ENNI ports are not available to all users. If you are provisioning your first ENNI port and are unsure if you have permission, contact support@packetfabric.com. ",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Point of presence in which the port should be located.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Speed of the port.\n\n\tEnum: [\"1Gbps\" \"10Gbps\" \"40Gbps\" \"100Gbps\"]",
			},
			"subscription_term": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "Duration of the subscription in months\n\n\tEnum [\"1\" \"12\" \"24\" \"36\"]",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Availability zone of the port.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Change Port Admin Status. Set it to true when port is enabled, false when port is disabled. ",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	interf := extractInterface(d)
	resp, err := c.CreateInterface(interf)
	time.Sleep(30 * time.Second)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		enabled := d.Get("enabled")
		if !enabled.(bool) {
			if toggleErr := _togglePortStatus(c, enabled.(bool), resp.PortCircuitID); toggleErr != nil {
				return diag.FromErr(toggleErr)
			}
		}
		d.SetId(resp.PortCircuitID)
	}
	return diags
}

func resourceReadInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetPortByCID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("autoneg", resp.Autoneg)
		_ = d.Set("media", resp.Media)
		_ = d.Set("nni", resp.IsNni)
		_ = d.Set("pop", resp.Pop)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("subscription_term", resp.SubscriptionTerm)
		_ = d.Set("zone", resp.Zone)
		_ = d.Set("po_number", resp.PONumber)
		if resp.Disabled {
			_ = d.Set("enabled", false)
		} else {
			_ = d.Set("enabled", true)
		}
	}
	return diags
}

func resourceUpdateInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	_, err := _extractUpdateFn(ctx, d, m)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceDeleteInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	_, err := c.DeletePort(d.Id())
	time.Sleep(30 * time.Second)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func extractInterface(d *schema.ResourceData) packetfabric.Interface {
	interf := packetfabric.Interface{
		AccountUUID:      d.Get("account_uuid").(string),
		Description:      d.Get("description").(string),
		Media:            d.Get("media").(string),
		Nni:              d.Get("nni").(bool),
		Pop:              d.Get("pop").(string),
		Speed:            d.Get("speed").(string),
		SubscriptionTerm: d.Get("subscription_term").(int),
		Zone:             d.Get("zone").(string),
		PONumber:         d.Get("po_number").(string),
	}
	if autoneg, ok := d.GetOk("autoneg"); ok {
		interf.Autoneg = autoneg.(bool)
	}
	return interf
}

func _extractUpdateFn(ctx context.Context, d *schema.ResourceData, m interface{}) (resp *packetfabric.InterfaceReadResp, err error) {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	// Update if payload contains Autoneg and Description
	if d.HasChanges([]string{"po_number", "description", "autoneg"}...) {
		portUpdateData := packetfabric.PortUpdate{
			Description: d.Get("description").(string),
			Autoneg:     d.Get("autoneg").(bool),
			PONumber:    d.Get("po_number").(string),
		}
		resp, err = c.UpdatePort(d.Id(), portUpdateData)
	}

	// Update port status
	if enabledHasChanged := d.HasChange("enabled"); enabledHasChanged {
		_, enableChange := d.GetChange("enabled")
		err = _togglePortStatus(c, enableChange.(bool), d.Id())
	}

	// Update port term
	if d.HasChange("subscription_term") {
		if subTerm, ok := d.GetOk("subscription_term"); ok {
			billing := packetfabric.BillingUpgrade{
				SubscriptionTerm: subTerm.(int),
			}
			_, err = c.ModifyBilling(d.Id(), billing)
			_ = d.Set("subscription_term", subTerm.(int))
		}
	}
	return
}

func _togglePortStatus(c *packetfabric.PFClient, enabled bool, portCID string) (err error) {
	if enabled {
		_, err = c.EnablePort(portCID)
	} else {
		_, err = c.DisablePort(portCID)
	}
	return
}
