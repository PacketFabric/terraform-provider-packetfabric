package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// used for all Hosted Clouds
func resourceServicesHostedUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid cloud service id")
	}

	updateServiceConnData := packetfabric.UpdateServiceConn{}
	changed := false

	if d.HasChanges([]string{"po_number", "description"}...) {
		if desc, ok := d.GetOk("description"); ok {
			updateServiceConnData.Description = desc.(string)
		}
		if poNumber, ok := d.GetOk("po_number"); ok {
			updateServiceConnData.PONumber = poNumber.(string)
		}
		changed = true
	}

	if d.HasChange("cloud_settings") {
		if cloudSettings, ok := d.GetOk("cloud_settings"); ok {
			cs := cloudSettings.(map[string]interface{})
			if cs["aws_vif_type"] != nil {
				updateServiceConnData.CloudSettings = &packetfabric.CloudSettingsHosted{
					CredentialsUUID: cs["credentials_uuid"].(string),
					AWSRegion:       cs["aws_region"].(string),
					MTU:             cs["mtu"].(int),
					AWSVIFType:      cs["aws_vif_type"].(string),
					BGPSettings: &packetfabric.BGPSettings{
						CustomerASN:   cs["customer_asn"].(int),
						AddressFamily: cs["address_family"].(string),
					},
				}
				changed = true
			}
		}
	}

	if changed {
		if _, err := c.UpdateServiceHostedConn(cloudCID.(string), updateServiceConnData); err != nil {
			return diag.FromErr(err)
		}
	}

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

	if d.HasChange("labels") {
		labels := d.Get("labels")
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}
	return diags
}

// used for all Dedicated Clouds
func resourceServicesDedicatedUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid cloud service id")
	}

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

	if d.HasChanges([]string{"po_number", "description"}...) {
		portUpdateData := packetfabric.PortUpdate{
			Description: d.Get("description").(string),
			PONumber:    d.Get("po_number").(string),
		}
		if _, err := c.UpdatePort(d.Id(), portUpdateData); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("labels") {
		labels := d.Get("labels")
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}
	return diags
}

// used for all Hosted and Dedicated Clouds
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
