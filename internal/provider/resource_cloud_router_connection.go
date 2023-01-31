package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// common function to update or delete cloud router connections (aws, google, azure, oracle, ibm)

func resourceCloudRouterConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	if d.HasChange("aws_account_id") ||
		d.HasChange("google_pairing_key") ||
		d.HasChange("google_vlan_attachment_name") ||
		d.HasChange("azure_service_key") ||
		d.HasChange("ibm_account_id") ||
		d.HasChange("ibm_bgp_asn") ||
		d.HasChange("ibm_bgp_cer_cidr") ||
		d.HasChange("ibm_bgp_ibm_cidr") ||
		d.HasChange("vc_ocid") ||
		d.HasChange("region") ||
		d.HasChange("maybe_nat") ||
		d.HasChange("maybe_dnat") ||
		d.HasChange("pop") ||
		d.HasChange("zone") ||
		d.HasChange("is_public") ||
		d.HasChange("published_quote_line_uuid") {
		return diag.Errorf("only the description or speed field can be updated")
	}

	if cid, ok := d.GetOk("circuit_id"); ok {
		if desc, descOk := d.GetOk("description"); descOk {
			descUpdate := packetfabric.DescriptionUpdate{
				Description: desc.(string),
			}
			if _, err := c.UpdateCloudRouterConnection(cid.(string), d.Id(), descUpdate); err != nil {
				return diag.FromErr(err)
			}
		}
		if speed, ok := d.GetOk("speed"); ok {
			billing := packetfabric.BillingUpgrade{
				Speed: speed.(string),
			}
			if _, err := c.ModifyBilling(cid.(string), billing); err != nil {
				return diag.FromErr(err)
			}
			_ = d.Set("speed", speed.(int))
		}
	}
	return diags
}

func resourceCloudRouterConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		if _, err := c.DeleteCloudRouterConnection(cid.(string), cloudConnCID.(string)); err != nil {
			diags = diag.FromErr(err)
		} else {
			deleteOk := make(chan bool)
			defer close(deleteOk)
			fn := func() (*packetfabric.ServiceState, error) {
				return c.GetCloudConnectionStatus(cid.(string), cloudConnCID.(string))
			}
			go c.CheckServiceStatus(deleteOk, fn)
			if !<-deleteOk {
				return diag.FromErr(err)
			}
			d.SetId("")
		}
	}
	return diags
}
