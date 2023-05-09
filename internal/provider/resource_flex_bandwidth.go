package provider

import (
	"context"
	"strconv"
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
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Description of the flex bandwidth container.",
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
			"subscription_term": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
				Description:  "The billing term, in months, of the flex bandwidth container.\n\n\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
			},
			"capacity": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedFlexOptions(), true),
				Description:  "Capacity of the flex bandwidth container. Must be in the format XXGbps.\n\n\tEnum: [\"100Gbps\" \"150Gbps\" \"200Gbps\" \"250Gbps\" \"300Gbps\" \"350Gbps\" \"400Gbps\" \"450Gbps\" \"500Gbps\"]",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"used_capacity_mbps": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Used capacity in Mbps of the flex bandwidth container.",
			},
			"available_capacity_mbps": {
				Type:        schema.TypeInt,
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
		_ = d.Set("used_capacity_mbps", resp.UsedCapacityMbps)
		_ = d.Set("available_capacity_mbps", resp.AvailableCapacityMbps)
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
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("subscription_term", resp.SubscriptionTerm)
		// convert int to string and append "Gbps"
		capacityGbps := resp.CapacityMbps / 1000
		capacityMbps_string := strconv.Itoa(capacityGbps) + "Gbps"
		_ = d.Set("capacity", capacityMbps_string)
		_ = d.Set("used_capacity_mbps", resp.UsedCapacityMbps)
		_ = d.Set("available_capacity_mbps", resp.AvailableCapacityMbps)
		if _, ok := d.GetOk("po_number"); ok {
			_ = d.Set("po_number", resp.PONumber)
		}
	}
	return diags
}

func resourceFlexBandwidthUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	flexID := d.Get("id").(string)
	capacity, ok := d.GetOk("capacity")
	if !ok {
		return diag.Errorf("please provide a valid capacity")
	}
	billing := packetfabric.BillingUpgrade{
		Capacity: capacity.(string),
	}

	if resp, err := c.ModifyBilling(flexID, billing); err != nil {
		return diag.FromErr(err)
	} else {
		diags = make(diag.Diagnostics, 0)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Flex Bandwidth ID billing upgrade.",
			Detail:   resp.Message,
		})
	}
	_ = d.Set("capacity", capacity.(string))

	return diags
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
		flex.PONumber = poNumber.(string)
	}
	return flex
}

func speedFlexOptions() []string {
	return []string{
		"50Gbps", "100Gbps", "150Gbps", "200Gbps",
		"250Gbps", "300Gbps", "350Gbps", "400Gbps",
		"450Gbps", "500Gbps"}
}
