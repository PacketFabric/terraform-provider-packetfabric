package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudRouter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudRouterCreate,
		ReadContext:   resourceCloudRouterRead,
		UpdateContext: resourceCloudRouterUpdate,
		DeleteContext: resourceCloudRouterDelete,
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
			"asn": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     4556,
				Description: "The ASN of the cloud router.\n\n\tThis can be the PacketFabric public ASN 4556 (default) or a private ASN from 64512 - 65534. ",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Cloud router name.",
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
			"regions": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The regions in which the Cloud Router connections will be located.\n\t\tUse `[\"US\"]` for North America and `[\"UK\"]` for EMEA. For transatlantic, use `[\"US\",\"UK\"]`.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},
			"capacity": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The cloud router capacity.\n\n\tEnum: \"100Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\" \"20Gbps\" \"30Gbps\" \"40Gbps\" \"50Gbps\" \"60Gbps\" \"80Gbps\" \"100Gbps\" \">100Gbps\"",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudRouterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	router := extractCloudRouter(d)

	resp, err := c.CreateCloudRouter(router)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("asn", resp.Asn)
		_ = d.Set("name", resp.Name)
		_ = d.Set("capacity", resp.Capacity)
		d.SetId(resp.CircuitID)
	}
	return diags
}

func resourceCloudRouterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID := d.Get("id").(string)
	resp, err := c.ReadCloudRouter(cID)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("asn", resp.Asn)
		_ = d.Set("name", resp.Name)
		_ = d.Set("capacity", resp.Capacity)
		var regions []string
		for _, region := range resp.Regions {
			regions = append(regions, region.Code)
		}
		_ = d.Set("regions", regions)
	}
	return diags
}

func resourceCloudRouterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	routerUpdate := packetfabric.CloudRouterUpdate{
		Name:     d.Get("name").(string),
		Regions:  extractRegions(d),
		Capacity: d.Get("capacity").(string),
	}

	cID := d.Get("id").(string)

	resp, err := c.UpdateCloudRouter(routerUpdate, cID)
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("name", resp.Name)
	_ = d.Set("capacity", resp.Capacity)
	return diags
}

func resourceCloudRouterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	cID := d.Get("id").(string)
	_, err := c.DeleteCloudRouter(cID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags

}

func extractCloudRouter(d *schema.ResourceData) packetfabric.CloudRouter {
	router := packetfabric.CloudRouter{}
	if asn, ok := d.GetOk("asn"); ok {
		router.Asn = asn.(int)
	}
	if name, ok := d.GetOk("name"); ok {
		router.Name = name.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		router.AccountUUID = accountUUID.(string)
	}
	if capacity, ok := d.GetOk("capacity"); ok {
		router.Capacity = capacity.(string)
	}
	router.Regions = extractRegions(d)
	return router
}

func extractRegions(d *schema.ResourceData) []string {
	if regions, ok := d.GetOk("regions"); ok {
		regs := make([]string, 0)
		for _, reg := range regions.([]interface{}) {
			regs = append(regs, reg.(string))
		}
		return regs
	}
	return make([]string, 0)
}
