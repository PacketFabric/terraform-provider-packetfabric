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
			"scope": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Whether the cloud router is private or public.\n\t\tValid Values: " + "\"private\" , \"public\"",
			},
			"asn": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ASN of the instance.\n\t\tValues must be within 64512 - 65534, or 4556.\n\t\tDefaults to 4556 if unspecified.",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The name of this particular CloudRouter.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "PacketFabric account UUID. The contact that will be billed.",
			},
			"regions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of PacketFabric Reigions.\n\t\tExample: \"[LAX1,SEA]\"\n\t\tPacketFabric Locations: https://packetfabric.com/locations",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},
			"capacity": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The total capacity of this particular Cloud Router.\n\t\tExample: 1Gbps or 10Gbps",
			},
			"circuit_id": {
				Type:     schema.TypeString,
				Optional: true,
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
		_ = d.Set("scope", resp.Scope)
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
		_ = d.Set("scope", resp.Scope)
		_ = d.Set("asn", resp.Asn)
		_ = d.Set("name", resp.Name)
		_ = d.Set("capacity", resp.Capacity)
	}
	return diags
}

func resourceCloudRouterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	routerUpdate := packetfabric.CloudRouterUpdate{
		Name:     d.Get("name").(string),
		Regions:  d.Get("regions").([]interface{}),
		Capacity: d.Get("capacity").(string),
	}

	cID := d.Get("id").(string)

	resp, err := c.UpdateCloudRouter(routerUpdate, cID)
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("scope", resp.Scope)
	_ = d.Set("asn", resp.Asn)
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
	router := packetfabric.CloudRouter{Regions: make([]packetfabric.Region, 0)}
	if scope, ok := d.GetOk("scope"); ok {
		router.Scope = scope.(string)
	}
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
	if regions, ok := d.GetOk("regions"); ok {
		for _, region := range regions.([]interface{}) {
			router.Regions = append(router.Regions, region.(packetfabric.Region))
		}
	}
	return router
}
