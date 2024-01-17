package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpamAsn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamAsnCreate,
		ReadContext:   resourceIpamAsnRead,
		UpdateContext: resourceIpamAsnUpdate,
		DeleteContext: resourceIpamAsnDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"asn_byte_type": {
				Type:     schema.TypeInt,
				Required: true,
				Optional: false,
			},
			"asn": {
				Type:     schema.TypeInt,
				Computed: true,
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

func resourceIpamAsnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ipamAsn := extractIpamAsn(d)
	resp, err := c.CreateIpamAsn(ipamAsn)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		d.SetId(fmt.Sprintf("%d", resp.Asn))
	}
	return diags
}

func resourceIpamAsnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	asn, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	resp, err := c.ReadIpamAsn(asn)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("asn_byte_type", resp.AsnByteType)
		_ = d.Set("asn", resp.Asn)
		_ = d.Set("time_created", resp.TimeCreated)
		_ = d.Set("time_updated", resp.TimeUpdated)
	}
	return diags
}

func resourceIpamAsnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	asn, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	ipamAsnUpdate := packetfabric.IpamAsn{
		AsnByteType: d.Get("asn_byte_type").(int),
		Asn: asn,
		TimeCreated: d.Get("time_created").(string),
		TimeUpdated: d.Get("time_updated").(string),
	}
	_, err = c.UpdateIpamAsn(ipamAsnUpdate)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(d.Get("id").(string))
	return diags
}

func resourceIpamAsnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	asn, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = c.DeleteIpamAsn(asn)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func extractIpamAsn(d *schema.ResourceData) packetfabric.IpamAsn {
	ipamAsn := packetfabric.IpamAsn{}
	asn, err := strconv.Atoi(d.Id())
	if err == nil {
		ipamAsn.Asn = asn
	}
	if asn_byte_type, ok := d.GetOk("asn_byte_type"); ok {
		ipamAsn.AsnByteType = asn_byte_type.(int)
	}
	if time_created, ok := d.GetOk("time_created"); ok {
		ipamAsn.TimeCreated = time_created.(string)
	}
	if time_updated, ok := d.GetOk("time_updated"); ok {
		ipamAsn.TimeUpdated = time_updated.(string)
	}
	return ipamAsn
}
