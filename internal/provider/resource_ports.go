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
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},

			"autoneg": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Only applicable to 1Gbps ports. Controls whether auto negotiation is on (true) or off (false). The request will fail if specified with 10Gbps.",
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
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Optic media type.\n\n\tEnum: [\"LX\" \"EX\" \"ZX\" \"LR\" \"ER\" \"ER DWDM\" \"ZR\" \"ZR DWDM\" \"LR4\" \"ER4\" \"CWDM4\" \"LR4\" \"ER4 Lite\"]",
			},
			"nni": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set this to true to provision an ENNI port. ENNI ports will use a nni_svlan_tpid value of 0x8100.\n\n\tBy default, ENNI ports are not available to all users. If you are provisioning your first ENNI port and are unsure if you have permission, contact support@packetfabric.com.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Point of presence in which the port should be located.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
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
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Availability zone of the port.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "True when port is enabled.",
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
	enabled := d.Get("enabled")
	if !enabled.(bool) {
		if toggleErr := _togglePortStatus(c, enabled.(bool), resp.PortCircuitID); toggleErr != nil {
			return diag.FromErr(toggleErr)
		}
	}
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
	_, err := c.GetPortByCID(d.Id())
	if err != nil {
		return diag.FromErr(err)
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
	resp, err := c.DeletePort(d.Id())
	time.Sleep(30 * time.Second)
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
	interf := packetfabric.Interface{
		AccountUUID:      d.Get("account_uuid").(string),
		Description:      d.Get("description").(string),
		Media:            d.Get("media").(string),
		Nni:              d.Get("nni").(bool),
		Pop:              d.Get("pop").(string),
		Speed:            d.Get("speed").(string),
		SubscriptionTerm: d.Get("subscription_term").(int),
		Zone:             d.Get("zone").(string),
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
	if autoneg, autoNegOk := d.GetOk("autoneg"); autoNegOk {
		if desc, descOk := d.GetOk("description"); descOk {
			resp, err = c.UpdatePort(autoneg.(bool), d.Id(), desc.(string))
		} else {
			// Update if payload contains Autoneg only
			resp, err = c.UpdatePortAutoNegOnly(autoneg.(bool), d.Id())
		}
	}
	// Update if payload contains Description only
	if desc, descOk := d.GetOk("description"); descOk {
		if _, autoNegOk := d.GetOk("autoneg"); !autoNegOk {
			resp, err = c.UpdatePortDescriptionOnly(d.Id(), desc.(string))
		}
	}
	// Update port status
	if enabledHasChanged := d.HasChange("enabled"); enabledHasChanged {
		_, enableChange := d.GetChange("enabled")
		err = _togglePortStatus(c, enableChange.(bool), d.Id())
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
