package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceBilling() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBillingCreate,
		ReadContext:   resourceBillingRead,
		UpdateContext: resourceBillingUpdate,
		DeleteContext: resourceBillingDelete,
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
			"circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the service to modify.",
			},
			"subscription_term": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
				Description:  "The subscription term in months. Options are: 1, 12, 24, 36.",
			},
			"speed": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "New speed of the connection.",
			},
			"billing_product_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Billing type for virtual circuits. The only option is `longhaul_dedicated`. This is applicable when upgrading an hourly or usage-based circuit.",
			},
			"service_class": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"longhaul", "metro"}, true),
				Description:  "Only applicable to dedicated cloud connections. This can be `longhaul` or `metro`.",
			},
			"capacity": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "New capacity if modifying a flex bandwidth container.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceBillingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid circuit_id")
	}
	if resp, err := c.ModifyBilling(cID.(string), _extractBillingUpgrade(d)); err != nil {
		return diag.FromErr(err)
	} else {
		diags = make(diag.Diagnostics, 0)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Circuit ID billing upgrade.",
			Detail:   resp.Message,
		})
	}
	d.SetId(cID.(string))
	return diags
}

func resourceBillingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceBillingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceBillingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return diag.Diagnostics{}
}

func _extractBillingUpgrade(d *schema.ResourceData) (billing packetfabric.BillingUpgrade) {
	billing = packetfabric.BillingUpgrade{}
	if term, ok := d.GetOk("subscription_term"); ok {
		billing.SubscriptionTerm = term.(int)
	}
	if speed, ok := d.GetOk("speed"); ok {
		billing.Speed = speed.(string)
	}
	if productType, ok := d.GetOk("billing_product_type"); ok {
		billing.BillingProductType = productType.(string)
	}
	if serviceClass, ok := d.GetOk("service_class"); ok {
		billing.ServiceClass = serviceClass.(string)
	}
	if capacity, ok := d.GetOk("capacity"); ok {
		billing.Capacity = capacity.(string)
	}
	return
}
