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
		DeleteContext: resourceIpamPrefixDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"length": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"family": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				Default:      "ipv4",
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),
			},
			"market": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"admin_ipam_contact_uuid": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
			},
			"tech_ipam_contact_uuid": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"circuit_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"address": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Required for ARIN if org_id was not provided, otherwise optional.",
			},
			"city": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Required for ARIN if org_id was not provided, otherwise optional.",
			},
			"postal_code": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Required for ARIN if org_id was not provided, otherwise optional.",
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipj_details": {
				Type:     schema.TypeSet,
				ForceNew: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rejection_reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_prefixes": {
							Type:     schema.TypeSet,
							Optional: true,
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
						"planned_prefix": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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

	ipamPrefix := extractIpamPrefix(d)
	resp, err := c.CreateIpamPrefix(ipamPrefix)
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.CircuitId)
	return resourceIpamPrefixRead(ctx, d, m)
}

func resourceIpamPrefixRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	ipamPrefix, err := c.ReadIpamPrefix(d.Id())
	if err != nil || ipamPrefix == nil {
		return diag.FromErr(err)
	}

	_ = d.Set("length", ipamPrefix.Length)
	_ = d.Set("market", ipamPrefix.Market)
	_ = d.Set("family", ipamPrefix.Family)
	_ = d.Set("ip_address", ipamPrefix.IpAddress)
	_ = d.Set("circuit_id", ipamPrefix.CircuitId)
	_ = d.Set("type", ipamPrefix.Type)
	_ = d.Set("org_id", ipamPrefix.OrgId)
	_ = d.Set("address", ipamPrefix.Address)
	_ = d.Set("city", ipamPrefix.City)
	_ = d.Set("postal_code", ipamPrefix.PostalCode)
	_ = d.Set("admin_ipam_contact_uuid", ipamPrefix.AdminIpamContactUuid)
	_ = d.Set("tech_ipam_contact_uuid", ipamPrefix.TechIpamContactUuid)
	_ = d.Set("state", ipamPrefix.State)

	if nil != ipamPrefix.IpjDetails {
		ipjDetails := flattenIpjDetails(ipamPrefix.IpjDetails)
		if err := d.Set("ipj_details", ipjDetails); err != nil {
			return diag.Errorf("error setting 'ipj_details': %s", err)
		}
	}
	return diags
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
	if length, ok := d.GetOk("length"); ok {
		ipamPrefix.Length = length.(int)
	}
	if market, ok := d.GetOk("market"); ok {
		ipamPrefix.Market = market.(string)
	}
	if family, ok := d.GetOk("family"); ok {
		ipamPrefix.Family = family.(string)
	}
	if ip_address, ok := d.GetOk("ip_address"); ok {
		ipamPrefix.IpAddress = ip_address.(string)
	}
	if circuit_id, ok := d.GetOk("circuit_id"); ok {
		ipamPrefix.CircuitId = circuit_id.(string)
	}
	if type_, ok := d.GetOk("type"); ok {
		ipamPrefix.Type = type_.(string)
	}
	if org_id, ok := d.GetOk("org_id"); ok {
		ipamPrefix.OrgId = org_id.(string)
	}
	if address, ok := d.GetOk("address"); ok {
		ipamPrefix.Address = address.(string)
	}
	if city, ok := d.GetOk("city"); ok {
		ipamPrefix.City = city.(string)
	}
	if postal_code, ok := d.GetOk("postal_code"); ok {
		ipamPrefix.PostalCode = postal_code.(string)
	}
	if admin_ipam_contact_uuid, ok := d.GetOk("admin_ipam_contact_uuid"); ok {
		ipamPrefix.AdminIpamContactUuid = admin_ipam_contact_uuid.(string)
	}
	if tech_ipam_contact_uuid, ok := d.GetOk("tech_ipam_contact_uuid"); ok {
		ipamPrefix.TechIpamContactUuid = tech_ipam_contact_uuid.(string)
	}
	if state, ok := d.GetOk("state"); ok {
		ipamPrefix.State = state.(string)
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
		CurrentPrefixes: extractIpamCurrentPrefixes(ipj_details_map["current_prefixes"].(*schema.Set)),
		PlannedPrefix:   extractIpamPlannedPrefix(ipj_details_map["planned_prefix"].(*schema.Set)),
		RejectionReason: ipj_details_map["rejection_reason"].(string),
	}
}

func extractIpamCurrentPrefixes(current_prefixes *schema.Set) []packetfabric.IpamCurrentPrefixes {
	ipamCurrentPrefixes := make([]packetfabric.IpamCurrentPrefixes, 0)
	for _, ipamCurrentPrefix := range current_prefixes.List() {
		ipamCurrentPrefixMap := ipamCurrentPrefix.(map[string]interface{})
		ipamCurrentPrefixes = append(ipamCurrentPrefixes, packetfabric.IpamCurrentPrefixes{
			Prefix:       ipamCurrentPrefixMap["prefix"].(string),
			IpsInUse:     ipamCurrentPrefixMap["ips_in_use"].(int),
			Description:  ipamCurrentPrefixMap["description"].(string),
			IspName:      ipamCurrentPrefixMap["isp_name"].(string),
			WillRenumber: ipamCurrentPrefixMap["will_renumber"].(bool),
		})
	}
	return ipamCurrentPrefixes
}

func extractIpamPlannedPrefix(planned_prefix *schema.Set) *packetfabric.IpamPlannedPrefix {
	for _, ipamPlannedPrefix := range planned_prefix.List() {
		ipamPlannedPrefixMap := ipamPlannedPrefix.(map[string]interface{})
		return &packetfabric.IpamPlannedPrefix{
			Description: ipamPlannedPrefixMap["description"].(string),
			Location:    ipamPlannedPrefixMap["location"].(string),
			Usage30d:    ipamPlannedPrefixMap["usage_30d"].(int),
			Usage3m:     ipamPlannedPrefixMap["usage_3m"].(int),
			Usage6m:     ipamPlannedPrefixMap["usage_6m"].(int),
			Usage1y:     ipamPlannedPrefixMap["usage_1y"].(int),
		}
	}
	return nil
}

func flattenIpjDetails(ipjDetails *packetfabric.IpjDetails) []interface{} {
	result := make([]interface{}, 0)
	data := make(map[string]interface{})
	data["current_prefixes"] = flattenCurrentPrefixes(ipjDetails.CurrentPrefixes)
	data["planned_prefix"] = flattenPlannedPrefix(ipjDetails.PlannedPrefix)
	result = append(result, data)
	return result
}

func flattenCurrentPrefixes(currentPrefixes []packetfabric.IpamCurrentPrefixes) []interface{} {
	result := make([]interface{}, len(currentPrefixes))
	for i, prefix := range currentPrefixes {
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

func flattenPlannedPrefix(plannedPrefix *packetfabric.IpamPlannedPrefix) map[string]interface{} {
	data := make(map[string]interface{})
	data["description"] = plannedPrefix.Description
	data["location"] = plannedPrefix.Location
	data["usage_30d"] = plannedPrefix.Usage30d
	data["usage_3m"] = plannedPrefix.Usage3m
	data["usage_6m"] = plannedPrefix.Usage6m
	data["usage_1y"] = plannedPrefix.Usage1y
	return data
}
