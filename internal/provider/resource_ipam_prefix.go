package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpamPrefix() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamPrefixCreate,
		ReadContext:   resourceIpamPrefixRead,
		UpdateContext: resourceIpamPrefixUpdate,
		DeleteContext: resourceIpamPrefixDelete,
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
			"prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"length": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"version": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      4,
				ValidateFunc: validation.IntInSlice([]int{4, 6}),
			},
			"bgp_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_contact_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
			},
			"tech_contact_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
			},
			"ipj_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"currently_used_prefixes": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validateIPAddressWithPrefix,
									},
									"ips_in_use": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"isp_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"will_renumber": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"planned_prefixes": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validateIPAddressWithPrefix,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"location": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"usage_30d": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"usage_3m": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"usage_6m": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"usage_1y": {
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceIpamPrefixCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	ipamPrefix := extractIpamPrefix(d)
	resp, err := c.CreateIpamPrefix(ipamPrefix)
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PrefixUuid)
	return diags
}

func resourceIpamPrefixRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	ipamPrefix, err := c.ReadIpamPrefix(d.Id())
	if err != nil || ipamPrefix == nil {
		return diag.FromErr(err)
	}

	_ = d.Set("prefix", ipamPrefix.Prefix)
	_ = d.Set("length", ipamPrefix.Length)
	_ = d.Set("version", ipamPrefix.Version)
	_ = d.Set("bgp_region", ipamPrefix.BgpRegion)
	_ = d.Set("admin_contact_uuid", ipamPrefix.AdminContactUuid)
	_ = d.Set("tech_contact_uuid", ipamPrefix.TechContactUuid)

	if nil != ipamPrefix.IpjDetails {
		ipjDetails := flattenIpjDetails(ipamPrefix.IpjDetails)
		if err := d.Set("ipj_details", ipjDetails); err != nil {
			return diag.Errorf("error setting 'ipj_details': %s", err)
		}
	}
	return diags
}

func resourceIpamPrefixUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	updateIpamPrefix := extractIpamPrefix(d)
	ipamPrefix, err := c.UpdateIpamPrefix(updateIpamPrefix)
	if err != nil || ipamPrefix == nil {
		return diag.FromErr(err)
	}
	return resourceIpamPrefixRead(ctx, d, m)
}

func resourceIpamPrefixDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	response, err := c.DeleteIpamPrefix(d.Id())
	if err != nil || response == nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func extractIpamPrefix(d *schema.ResourceData) packetfabric.IpamPrefix {
	ipamPrefix := packetfabric.IpamPrefix{}
	if prefix, ok := d.GetOk("prefix"); ok {
		ipamPrefix.Prefix = prefix.(string)
	}
	if length, ok := d.GetOk("length"); ok {
		ipamPrefix.Length = length.(int)
	}
	if version, ok := d.GetOk("version"); ok {
		ipamPrefix.Version = version.(int)
	}
	if bgp_region, ok := d.GetOk("bgp_region"); ok {
		ipamPrefix.BgpRegion = bgp_region.(string)
	}
	if admin_contact_uuid, ok := d.GetOk("admin_contact_uuid"); ok {
		ipamPrefix.AdminContactUuid = admin_contact_uuid.(string)
	}
	if tech_contact_uuid, ok := d.GetOk("tech_contact_uuid"); ok {
		ipamPrefix.TechContactUuid = tech_contact_uuid.(string)
	}
	if ipj_details, ok := d.GetOk("ipj_details"); ok {
		ipamPrefix.IpjDetails = extractIpjDetails(ipj_details.(*schema.Set))
	} else {
		ipamPrefix.IpjDetails = nil
	}
	return ipamPrefix
}

func extractIpjDetails(ipj_details *schema.Set) *packetfabric.IpjDetails {
	ipj_details_map := ipj_details.List()[0].(map[string]interface{})
	return &packetfabric.IpjDetails{
		CurrentlyUsedPrefixes: extractIpamCurrentlyUsedPrefixes(ipj_details_map["currently_used_prefixes"].(*schema.Set)),
		PlannedPrefixes:       extractIpamPlannedPrefixes(ipj_details_map["planned_prefixes"].(*schema.Set)),
	}
}

func extractIpamCurrentlyUsedPrefixes(currently_used_prefixes *schema.Set) []packetfabric.IpamCurrentlyUsedPrefixes {
	ipamCurrentlyUsedPrefixes := make([]packetfabric.IpamCurrentlyUsedPrefixes, 0)
	for _, ipamCurrentlyUsedPrefix := range currently_used_prefixes.List() {
		ipamCurrentlyUsedPrefixMap := ipamCurrentlyUsedPrefix.(map[string]interface{})
		ipamCurrentlyUsedPrefixes = append(ipamCurrentlyUsedPrefixes, packetfabric.IpamCurrentlyUsedPrefixes{
			Prefix:       ipamCurrentlyUsedPrefixMap["prefix"].(string),
			IpsInUse:     ipamCurrentlyUsedPrefixMap["ips_in_use"].(int),
			Description:  ipamCurrentlyUsedPrefixMap["description"].(string),
			IspName:      ipamCurrentlyUsedPrefixMap["isp_name"].(string),
			WillRenumber: ipamCurrentlyUsedPrefixMap["will_renumber"].(bool),
		})
	}
	return ipamCurrentlyUsedPrefixes
}

func extractIpamPlannedPrefixes(planned_prefixes *schema.Set) []packetfabric.IpamPlannedPrefixes {
	ipamPlannedPrefixes := make([]packetfabric.IpamPlannedPrefixes, 0)
	for _, ipamPlannedPrefix := range planned_prefixes.List() {
		ipamPlannedPrefixMap := ipamPlannedPrefix.(map[string]interface{})
		ipamPlannedPrefixes = append(ipamPlannedPrefixes, packetfabric.IpamPlannedPrefixes{
			Prefix:      ipamPlannedPrefixMap["prefix"].(string),
			Description: ipamPlannedPrefixMap["description"].(string),
			Location:    ipamPlannedPrefixMap["location"].(string),
			Usage30d:    ipamPlannedPrefixMap["usage_30d"].(int),
			Usage3m:     ipamPlannedPrefixMap["usage_3m"].(int),
			Usage6m:     ipamPlannedPrefixMap["usage_6m"].(int),
			Usage1y:     ipamPlannedPrefixMap["usage_1y"].(int),
		})
	}
	return ipamPlannedPrefixes
}

func flattenIpjDetails(ipjDetails *packetfabric.IpjDetails) map[string]interface{} {
	data := make(map[string]interface{})
	data["currently_used_prefixes"] = flattenCurrentlyUsedPrefixes(ipjDetails.CurrentlyUsedPrefixes)
	data["planned_prefixes"] = flattenPlannedPrefixes(ipjDetails.PlannedPrefixes)
	return data
}

func flattenCurrentlyUsedPrefixes(currentlyUsedPrefixes []packetfabric.IpamCurrentlyUsedPrefixes) []interface{} {
	result := make([]interface{}, len(currentlyUsedPrefixes))
	for i, prefix := range currentlyUsedPrefixes {
		data := make(map[string]interface{})
		data["prefix"] = prefix.Prefix
		data["ips_in_use"] = prefix.IpsInUse
		data["description"] = prefix.Description
		data["isp_name"] = prefix.IspName
		data["will_renumber"] = prefix.WillRenumber
		result[i] = data
	}
	return result
}

func flattenPlannedPrefixes(plannedPrefixes []packetfabric.IpamPlannedPrefixes) []interface{} {
	result := make([]interface{}, len(plannedPrefixes))
	for i, prefix := range plannedPrefixes {
		data := make(map[string]interface{})
		data["prefix"] = prefix.Prefix
		data["description"] = prefix.Description
		data["location"] = prefix.Location
		data["usage_30d"] = prefix.Usage30d
		data["usage_3m"] = prefix.Usage3m
		data["usage_6m"] = prefix.Usage6m
		data["usage_1y"] = prefix.Usage1y
		result[i] = data
	}
	return result
}
