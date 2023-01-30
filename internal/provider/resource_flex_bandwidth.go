package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFlexBandwidth() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFlexBandwidthCreate,
		ReadContext:   resourceFlexBandwidthRead,
		UpdateContext: resourceFlexBandwidthUpdate,
		DeleteContext: resourceFlexBandwidthDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Description of the flex bandwidth container.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},
			"subscription_term": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The billing term, in months, of the flex bandwidth container.\n\n\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
			},
			"capacity": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Capacity of the flex bandwidth container. Must be in the format XXGbps or XXMbps.\n\n\tExample: [\"100Gbps\" \"150Gbps\" \"200Gbps\" \"250Gbps\" \"300Gbps\" \"350Gbps\" \"400Gbps\" \"450Gbps\" \"500Gbps\"]",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"capacity_mbps": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Capacity in Mbps of the flex bandwidth container.",
			},
			"used_capacity_mbps": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Used capacity in Mbps of the flex bandwidth container.",
			},
			"available_capacity_mbps": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Available capacity in Mbps of the flex bandwidth container.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceFlexBandwidthCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	flex := extractFlexBandwidth(d)

	resp, err := c.CreateFlexBandwidth(flex)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("description", resp.Description)
		_ = d.Set("subscription_term", resp.SubscriptionTerm)
		_ = d.Set("capacity_mbps", resp.CapacityMbps)
		_ = d.Set("used_capacity_mbps", resp.UsedCapacityMbps)
		_ = d.Set("available_capacity_mbps", resp.AvailableCapacityMbps)
		_ = d.Set("po_number", resp.PoNumber)
		d.SetId(resp.FlexBandwidthID)
	}
	return diags
}

func resourceFlexBandwidthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	flexID := d.Get("id").(string)
	resp, err := c.ReadFlexBandwidth(flexID)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("description", resp.Description)
		_ = d.Set("subscription_term", resp.SubscriptionTerm)
		_ = d.Set("capacity_mbps", resp.CapacityMbps)
		_ = d.Set("used_capacity_mbps", resp.UsedCapacityMbps)
		_ = d.Set("available_capacity_mbps", resp.AvailableCapacityMbps)
		_ = d.Set("po_number", resp.PoNumber)
	}
	return diags
}

func resourceFlexBandwidthUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Flex Bandwidth Update.",
		Detail:   "Warning: use packetfabric_billing_modify_order resource to modify the capacity.",
	}}
}

func resourceFlexBandwidthDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	flexID := d.Get("id").(string)
	_, err := c.DeleteFlexBandwidth(flexID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags

}

func extractFlexBandwidth(d *schema.ResourceData) packetfabric.FlexBandwidth {
	flex := packetfabric.FlexBandwidth{}
	if description, ok := d.GetOk("description"); ok {
		flex.Description = description.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		flex.AccountUUID = accountUUID.(string)
	}
	if subscriptionTerm, ok := d.GetOk("subscription_term"); ok {
		flex.SubscriptionTerm = subscriptionTerm.(int)
	}
	if capacity, ok := d.GetOk("capacity"); ok {
		flex.Capacity = capacity.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		flex.PoNumber = poNumber.(string)
	}
	return flex
}
