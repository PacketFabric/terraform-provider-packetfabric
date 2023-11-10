package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpamPrefixConfirmation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamPrefixConfirmationCreate,
		ReadContext:   resourceIpamPrefixConfirmationRead,
		UpdateContext: resourceIpamPrefixConfirmationUpdate,
		DeleteContext: resourceIpamPrefixConfirmationDelete,
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

func resourceIpamPrefixConfirmationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	ipamPrefixConfirmation := extractIpamPrefixConfirmation(d)
	resp, err := c.CreateIpamPrefixConfirmation(ipamPrefixConfirmation)
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PrefixUuid)
	return diags
}

func resourceIpamPrefixConfirmationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	ipamPrefixConfirmation, err := c.ReadIpamPrefixConfirmation(d.Id())
	if err != nil || ipamPrefixConfirmation == nil {
		return diag.FromErr(err)
	}

	_ = d.Set("admin_contact_uuid", ipamPrefixConfirmation.AdminContactUuid)
	_ = d.Set("tech_contact_uuid", ipamPrefixConfirmation.TechContactUuid)

	if nil != ipamPrefixConfirmation.IpjDetails {
		ipjDetails := flattenIpjDetails(ipamPrefixConfirmation.IpjDetails)
		if err := d.Set("ipj_details", ipjDetails); err != nil {
			return diag.Errorf("error setting 'ipj_details': %s", err)
		}
	}
	return diags
}

func resourceIpamPrefixConfirmationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	updateIpamPrefixConfirmation := extractIpamPrefixConfirmation(d)
	ipamPrefixConfirmation, err := c.UpdateIpamPrefixConfirmation(updateIpamPrefixConfirmation)
	if err != nil || ipamPrefixConfirmation == nil {
		return diag.FromErr(err)
	}
	return resourceIpamPrefixConfirmationRead(ctx, d, m)
}

func resourceIpamPrefixConfirmationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	response, err := c.DeleteIpamPrefixConfirmation(d.Id())
	if err != nil || response == nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func extractIpamPrefixConfirmation(d *schema.ResourceData) packetfabric.IpamPrefixConfirmation {
	ipamPrefixConfirmation := packetfabric.IpamPrefixConfirmation{}
	if admin_contact_uuid, ok := d.GetOk("admin_contact_uuid"); ok {
		ipamPrefixConfirmation.AdminContactUuid = admin_contact_uuid.(string)
	}
	if tech_contact_uuid, ok := d.GetOk("tech_contact_uuid"); ok {
		ipamPrefixConfirmation.TechContactUuid = tech_contact_uuid.(string)
	}
	if ipj_details, ok := d.GetOk("ipj_details"); ok {
		ipamPrefixConfirmation.IpjDetails = extractIpjDetails(ipj_details.(*schema.Set))
	} else {
		ipamPrefixConfirmation.IpjDetails = nil
	}
	return ipamPrefixConfirmation
}
