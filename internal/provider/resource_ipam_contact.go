package provider

import (
	"context"
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
		DeleteContext: resourceIpamContactDelete,
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
			"name": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "IPAM Contact Name.",
			},
			"address": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "IPAM Contact Address.",
			},
			"country_code": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(2, 2),
				Description:  "IPAM Contact Country Code.",
			},
			"phone": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[0-9 ()+.-]+(\s?(x|ex|ext|ete|extn)?(\.|\.\s|\s)?[\d]{1,9})?$`), "Phone number must match the pattern ^[0-9 ()+.-]+(\\s?(x|ex|ext|ete|extn)?(\\.|\\.\\s|\\s)?[\\d]{1,9})?$"),
				Description:  "IPAM Contact phone number.",
			},
			"email": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact e-mail. Please note that this email address can only be updated by the IPAM contact themselves after creation.",
			},
			"apnic_org_id": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact APNIC Organization ID.",
			},
			"apnic_ref": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact APNIC Reference.",
			},
			"ripe_org_id": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact RIPE Organization ID.",
			},
			"ripe_ref": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "IPAM Contact RIPE Reference.",
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
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
		_ = d.Set("name", resp.Name)
		_ = d.Set("address", resp.Address)
		_ = d.Set("country_code", resp.CountryCode)
		_ = d.Set("phone", resp.Phone)
		_ = d.Set("email", resp.Email)
		_ = d.Set("apnic_org_id", resp.ApnicOrgId)
		_ = d.Set("ripe_org_id", resp.RipeOrgId)
		_ = d.Set("apnic_ref", resp.ApnicRef)
		_ = d.Set("ripe_ref", resp.RipeRef)
		_ = d.Set("time_updated", resp.TimeUpdated)
		_ = d.Set("time_created", resp.TimeCreated)
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
	if name, ok := d.GetOk("name"); ok {
		ipamContact.Name = name.(string)
	}
	if address, ok := d.GetOk("address"); ok {
		ipamContact.Address = address.(string)
	}
	if country_code, ok := d.GetOk("country_code"); ok {
		ipamContact.CountryCode = country_code.(string)
	}
	if phone, ok := d.GetOk("phone"); ok {
		ipamContact.Phone = phone.(string)
	}
	if email, ok := d.GetOk("email"); ok {
		ipamContact.Email = email.(string)
	}
	if apnic_org_id, ok := d.GetOk("apnic_org_id"); ok {
		ipamContact.ApnicOrgId = apnic_org_id.(string)
	}
	if ripe_org_id, ok := d.GetOk("ripe_org_id"); ok {
		ipamContact.RipeOrgId = ripe_org_id.(string)
	}
	if apnic_ref, ok := d.GetOk("apnic_ref"); ok {
		ipamContact.ApnicRef = apnic_ref.(string)
	}
	if ripe_ref, ok := d.GetOk("ripe_ref"); ok {
		ipamContact.RipeRef = ripe_ref.(string)
	}
	if time_created, ok := d.GetOk("time_created"); ok {
		ipamContact.TimeCreated = time_created.(string)
	}
	if time_updated, ok := d.GetOk("time_updated"); ok {
		ipamContact.TimeUpdated = time_updated.(string)
	}
	return ipamContact
}
