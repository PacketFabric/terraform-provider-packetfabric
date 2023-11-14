package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpamContact() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamContactCreate,
		ReadContext:   resourceIpamContactRead,
		UpdateContext: resourceIpamContactUpdate,
		DeleteContext: resourceIpamContactDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"contact_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "IPAM Contact Name.",
			},
			"org_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "IPAM Contact Organization Name.",
			},
			"address": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "IPAM Contact Address.",
			},
			"phone": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[0-9 ()+.-]+(\s?(x|ex|ext|ete|extn)?(\.|\.\s|\s)?[\d]{1,9})?$`), "Phone number must match the pattern ^[0-9 ()+.-]+(\\s?(x|ex|ext|ete|extn)?(\\.|\\.\\s|\\s)?[\\d]{1,9})?$"),
				Description:  "IPAM Contact phone number.",
			},
			"email": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact e-mail. Please note that this email address can only be updated by the IPAM contact themselves after creation.",
			},
			"arin_org_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact ARIN Organization ID.",
			},
			"apnic_org_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact APNIC Organization ID.",
			},
			"ripe_org_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact RIPE Organization ID",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIpamContactCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ipamContact := extractIpamContact(d)
	resp, err := c.CreateIpamContact(ipamContact)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		d.SetId(resp.UUID)
	}
	return diags
}

func resourceIpamContactRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ipamContactID := d.Id()
	resp, err := c.ReadIpamContact(ipamContactID)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("contact_name", resp.ContactName)
		_ = d.Set("org_name", resp.OrgName)
		_ = d.Set("address", resp.Address)
		_ = d.Set("phone", resp.Phone)
		_ = d.Set("email", resp.Email)
		_ = d.Set("arin_org_id", resp.ArinOrgId)
		_ = d.Set("apnic_org_id", resp.ApnicOrgId)
		_ = d.Set("ripe_org_id", resp.RipeOrgId)
	}
	return diags
}

func resourceIpamContactUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	// Not sure we want this check
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update IPAM contact's email",
			Detail:   "IPAM Contact's email can only be updated by the IPAM contact himself. Please ask the IPAM contact to update their email if needed.",
		})
	}
	if !d.HasChange("email") {
		ipamContactUpdate := packetfabric.IpamContact{
			UUID:        d.Id(),
			ContactName: d.Get("contact_name").(string),
			OrgName:     d.Get("org_name").(string),
			Address:     d.Get("address").(string),
			Phone:       d.Get("phone").(string),
			Email:       d.Get("email").(string),
			ArinOrgId:   d.Get("arin_org_id").(string),
			ApnicOrgId:  d.Get("apnic_org_id").(string),
			RipeOrgId:   d.Get("ripe_org_id").(string),
		}
		_, err := c.UpdateIpamContact(ipamContactUpdate)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(d.Get("uuid").(string))
	}
	return diags
}

func resourceIpamContactDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ipamContactID := d.Id()
	_, err := c.DeleteIpamContact(ipamContactID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func extractIpamContact(d *schema.ResourceData) packetfabric.IpamContact {
	ipamContact := packetfabric.IpamContact{}
	ipamContact.UUID = d.Id()
	if contact_name, ok := d.GetOk("contact_name"); ok {
		ipamContact.ContactName = contact_name.(string)
	}
	if org_name, ok := d.GetOk("org_name"); ok {
		ipamContact.OrgName = org_name.(string)
	}
	if address, ok := d.GetOk("address"); ok {
		ipamContact.Address = address.(string)
	}
	if phone, ok := d.GetOk("phone"); ok {
		ipamContact.Phone = phone.(string)
	}
	if email, ok := d.GetOk("email"); ok {
		ipamContact.Email = email.(string)
	}
	if arin_org_id, ok := d.GetOk("arin_org_id"); ok {
		ipamContact.ArinOrgId = arin_org_id.(string)
	}
	if apnic_org_id, ok := d.GetOk("apnic_org_id"); ok {
		ipamContact.ApnicOrgId = apnic_org_id.(string)
	}
	if ripe_org_id, ok := d.GetOk("ripe_org_id"); ok {
		ipamContact.RipeOrgId = ripe_org_id.(string)
	}
	return ipamContact
}
