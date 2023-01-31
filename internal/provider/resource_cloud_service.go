package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// used for all Hosted Clouds
func resourceServicesHostedUpdate(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(string, string) (*packetfabric.CloudServiceConnCreateResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid cloud service id")
	}
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
		d.HasChange("pop") ||
		d.HasChange("zone") {
		return diag.Errorf("only the description or speed field can be updated")
	}
	if desc, ok := d.GetOk("description"); !ok {
		return diag.Errorf("please provide a valid description for Cloud Service")
	} else {
		resp, err := fn(desc.(string), cloudCID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		_ = d.Set("description", resp.Description)
	}
	// speed only for hosted cloud
	if d.HasChange("speed") {
		billing := packetfabric.BillingUpgrade{}
		if speed, ok := d.GetOk("speed"); ok {
			billing.Speed = speed.(string)
		}
		if _, err := c.ModifyBilling(cloudCID.(string), billing); err != nil {
			return diag.FromErr(err)
		}
		if speed, ok := d.GetOk("speed"); ok {
			_ = d.Set("speed", speed.(string))
		}
	}
	return diags
}

// used for all Dedicated Clouds
func resourceServicesDedicatedUpdate(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(string, string) (*packetfabric.CloudServiceConnCreateResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid cloud service id")
	}
	if d.HasChange("description") ||
		d.HasChange("autoneg") ||
		d.HasChange("aws_region") ||
		d.HasChange("pop") ||
		d.HasChange("speed") ||
		d.HasChange("loa") ||
		d.HasChange("should_create_lag") ||
		d.HasChange("zone") {
		return diag.Errorf("only the subscription_term or service_class field can be updated")
	}
	// Updating the description for Dedicated Cloud is not supported a this time
	// if desc, ok := d.GetOk("description"); !ok {
	// 	return diag.Errorf("please provide a valid description for Cloud Service")
	// } else {
	// 	resp, err := fn(desc.(string), cloudCID.(string))
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	_ = d.Set("description", resp.Description)
	// }
	// subscription_term and service_class only for dedicated cloud
	if d.HasChange("subscription_term") || d.HasChange("service_class") {
		billing := packetfabric.BillingUpgrade{}
		if subTerm, ok := d.GetOk("subscription_term"); ok {
			billing.SubscriptionTerm = subTerm.(int)
		}
		if serviceClass, ok := d.GetOk("service_class"); ok {
			billing.ServiceClass = serviceClass.(string)
		}
		if _, err := c.ModifyBilling(cloudCID.(string), billing); err != nil {
			return diag.FromErr(err)
		}
		if subTerm, ok := d.GetOk("subscription_term"); ok {
			_ = d.Set("subscription_term", subTerm.(int))
		}
		if serviceClass, ok := d.GetOk("service_class"); ok {
			_ = d.Set("service_class", serviceClass.(string))
		}
	}
	return diags
}

func resourceCloudSourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}, diagSummary string) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCID, ok := d.GetOk("id")
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  diagSummary,
			Detail:   cloudCidNotFoundDetailsMsg,
		})
		return diags
	}
	err := c.DeleteCloudService(cloudCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	deleteOkCh := make(chan bool)
	defer close(deleteOkCh)
	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetCloudServiceStatus(cloudCID.(string))
	}
	go c.CheckServiceStatus(deleteOkCh, fn)
	if !<-deleteOkCh {
		return diag.FromErr(err)
	}
	return diags
}
