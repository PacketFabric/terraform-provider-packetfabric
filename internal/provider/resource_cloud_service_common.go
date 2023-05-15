package provider

import (
	"context"
	"os"
	"strings"
	"time"

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
			updateServiceConnData.CloudSettings = &packetfabric.CloudSettings{}
			if credentialsUUID, ok := cs["credentials_uuid"].(string); ok {
				updateServiceConnData.CloudSettings.CredentialsUUID = credentialsUUID
			}
			if awsRegion, ok := cs["aws_region"].(string); ok {
				updateServiceConnData.CloudSettings.AwsRegion = awsRegion
			}
			if googleRegion, ok := cs["google_region"].(string); ok {
				updateServiceConnData.CloudSettings.GoogleRegion = googleRegion
			}
			if mtu, ok := cs["mtu"].(int); ok {
				updateServiceConnData.CloudSettings.Mtu = mtu
			}
			if bgpSettings, ok := cs["bgp_settings"].([]interface{}); ok && len(bgpSettings) > 0 {
				bgp := bgpSettings[0].(map[string]interface{})
				updateServiceConnData.CloudSettings.BgpSettings = &packetfabric.BgpSettings{}
				if advertisedPrefixes, ok := bgp["advertised_prefixes"].([]interface{}); ok {
					ap_aws := make([]string, len(advertisedPrefixes))
					for i, elem := range advertisedPrefixes {
						ap_aws[i] = elem.(string)
					}
					updateServiceConnData.CloudSettings.BgpSettings.AdvertisedPrefixes = ap_aws
				}
				if customerAsn, ok := bgp["customer_asn"].(int); ok {
					updateServiceConnData.CloudSettings.BgpSettings.CustomerAsn = customerAsn
				}
				if remoteAsn, ok := bgp["remote_asn"].(int); ok {
					updateServiceConnData.CloudSettings.BgpSettings.RemoteAsn = remoteAsn
				}
				if md5, ok := bgp["md5"].(string); ok {
					updateServiceConnData.CloudSettings.BgpSettings.Md5 = md5
				}
				if googleKeepaliveInterval, ok := bgp["google_keepalive_interval"].(int); ok {
					updateServiceConnData.CloudSettings.BgpSettings.GoogleKeepaliveInterval = googleKeepaliveInterval
				}

				if advertisedPrefixes, ok := bgp["google_advertised_ip_ranges"].([]interface{}); ok {
					ap_google := make([]string, len(advertisedPrefixes))
					for i, elem := range advertisedPrefixes {
						ap_google[i] = elem.(string)
					}
					updateServiceConnData.CloudSettings.BgpSettings.GoogleAdvertisedIPRanges = ap_google
				}
			}
			changed = true
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
	host := os.Getenv("PF_HOST")
	testingInLab := strings.Contains(host, "api.dev")

	if testingInLab {
		resp, err3 := c.GetCloudConnInfo(d.Id())
		if err3 != nil {
			return diag.FromErr(err3)
		}
		if resp.PortType == "dedicated" {
			if resp.ServiceProvider == "aws" { // LAG is not enabled in the ACC in dev environment
				if toggleErr := _togglePortStatus(c, false, cloudCID.(string)); toggleErr != nil {
					return diag.FromErr(toggleErr)
				}
			}
			if resp.ServiceProvider == "google" {
				if toggleErr := DisableLinkAggregationGroup(cloudCID.(string)); toggleErr != nil {
					return diag.FromErr(toggleErr)
				}
			}
			time.Sleep(time.Duration(180) * time.Second)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "In the dev environment, ports are disabled prior to deletion.",
			})
		}
	}
	etlDiags, err2 := addETLWarning(c, cloudCID.(string))
	if err2 != nil {
		return diag.FromErr(err2)
	}
	diags = append(diags, etlDiags...)
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
