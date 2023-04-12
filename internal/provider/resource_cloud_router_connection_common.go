package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const StringSeparator = ":"

type CloudRouterCircuitIdData struct {
	cloudRouterCircuitId           string
	cloudRouterConnectionCircuitId string
}

type CloudRouterCircuitBgpIdData struct {
	cloudRouterCircuitId           string
	cloudRouterConnectionCircuitId string
	bgpSessionUUID                 string
}

// common function to update or delete cloud router connections (aws, google, azure, oracle, ibm)
func resourceCloudRouterConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	if cid, ok := d.GetOk("circuit_id"); ok {
		sectionsToUpdate := []struct {
			condition bool
			action    func() diag.Diagnostics
		}{
			{
				condition: d.HasChange("description") ||
					d.HasChange("po_number") ||
					d.HasChange("cloud_settings.0.credentials_uuid") ||
					d.HasChange("cloud_settings.0.mtu") ||
					d.HasChange("cloud_settings.0.bgp_settings.0.google_keepalive_interval"),
				action: func() diag.Diagnostics {
					updateData := packetfabric.CloudRouterUpdateData{}
					if desc, descOk := d.GetOk("description"); descOk {
						updateData.Description = desc.(string)
					}
					if poNumber, ok := d.GetOk("po_number"); ok {
						updateData.PONumber = poNumber.(string)
					}
					if credentialsUUID, ok := d.GetOk("cloud_settings.0.credentials_uuid"); ok {
						updateData.CloudSettings.CredentialsUUID = credentialsUUID.(string)
					}
					if mtu, ok := d.GetOk("cloud_settings.0.mtu"); ok {
						updateData.CloudSettings.Mtu = mtu.(int)
					}
					if googleKeepaliveInterval, ok := d.GetOk("cloud_settings.0.bgp_settings.0.google_keepalive_interval"); ok {
						updateData.CloudSettings.BgpSettings = &packetfabric.BgpSettings{}
						updateData.CloudSettings.BgpSettings.GoogleKeepaliveInterval = googleKeepaliveInterval.(int)
					}
					if _, err := c.UpdateCloudRouterConnection(cid.(string), d.Id(), updateData); err != nil {
						return diag.FromErr(err)
					}
					return diag.FromErr(checkCloudRouterConnectionStatus(c, cid.(string), d.Id()))
				},
			},
			{
				condition: d.HasChange("speed"),
				action: func() diag.Diagnostics {
					speed, _ := d.GetOk("speed")
					billing := packetfabric.BillingUpgrade{Speed: speed.(string)}
					if _, err := c.ModifyBilling(d.Id(), billing); err != nil {
						return diag.FromErr(err)
					}
					if err := checkCloudRouterConnectionStatus(c, cid.(string), d.Id()); err != nil {
						return diag.FromErr(err)
					}
					_ = d.Set("speed", speed.(string))
					return nil
				},
			},
			{
				condition: d.HasChange("cloud_settings.0.bgp_settings"),
				action: func() diag.Diagnostics {
					prefixesSet := d.Get("cloud_settings.0.bgp_settings.0.prefixes").(*schema.Set)
					prefixesList := prefixesSet.List()
					if err := validatePrefixes(prefixesList); err != nil {
						return diag.FromErr(err)
					}
					session := extractBgpSessionUpdate(d)
					_, resp, err := c.UpdateBgpSession(session, cid.(string), d.Id())
					if err != nil {
						return diag.FromErr(err)
					}
					if err := checkCloudRouterConnectionStatus(c, cid.(string), d.Id()); err != nil {
						return diag.FromErr(err)
					}
					// Update the bgp_settings_id within the bgp_settings block
					if err := d.Set("cloud_settings.0.bgp_settings.0.bgp_settings_id", resp.BgpSettingsUUID); err != nil {
						return diag.FromErr(err)
					}
					return nil
				},
			},
		}

		for _, section := range sectionsToUpdate {
			if section.condition {
				if diags := section.action(); diags.HasError() {
					return diags
				}
			}
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

func resourceCloudRouterConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		if _, err := c.DeleteCloudRouterConnection(cid.(string), cloudConnCID.(string)); err != nil {
			diags = diag.FromErr(err)
		} else {
			if err := checkCloudRouterConnectionStatus(c, cid.(string), cloudConnCID.(string)); err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
		}
	}
	return diags
}

// Used to import Cloud Router Connection part of a Cloud Router
func splitCloudRouterCircuitIdString(data string) (CloudRouterCircuitIdData, error) {
	stringArr := strings.Split(data, StringSeparator)
	if len(stringArr) != 2 {
		return CloudRouterCircuitIdData{}, errors.New("to import a cloud router connection, use the format {cloud_router_circuit_id}:{cloud_router_connection_circuit_id}")
	}
	return CloudRouterCircuitIdData{cloudRouterCircuitId: stringArr[0], cloudRouterConnectionCircuitId: stringArr[1]}, nil
}

func CloudRouterImportStatePassthroughContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	cloudRouterCircuitIdData, err := splitCloudRouterCircuitIdString(d.Id())
	if err != nil {
		return []*schema.ResourceData{}, err
	}
	_ = d.Set("circuit_id", cloudRouterCircuitIdData.cloudRouterCircuitId)
	d.SetId(cloudRouterCircuitIdData.cloudRouterConnectionCircuitId)
	return []*schema.ResourceData{d}, nil
}

// Used to import BGP session part of a Cloud Router Connection
func splitCloudRouterCircuitBgpIdString(data string) (CloudRouterCircuitBgpIdData, error) {
	stringArr := strings.Split(data, StringSeparator)
	if len(stringArr) != 3 {
		return CloudRouterCircuitBgpIdData{}, errors.New("to import a BGP session, use the format {cloud_router_circuit_id}:{cloud_router_connection_circuit_id}:{bgp_session_id}")
	}
	return CloudRouterCircuitBgpIdData{cloudRouterCircuitId: stringArr[0], cloudRouterConnectionCircuitId: stringArr[1], bgpSessionUUID: stringArr[2]}, nil
}

func BgpImportStatePassthroughContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	CloudRouterCircuitBgpIdData, err := splitCloudRouterCircuitBgpIdString(d.Id())
	if err != nil {
		return []*schema.ResourceData{}, err
	}
	_ = d.Set("circuit_id", CloudRouterCircuitBgpIdData.cloudRouterCircuitId)
	_ = d.Set("connection_id", CloudRouterCircuitBgpIdData.cloudRouterConnectionCircuitId)
	d.SetId(CloudRouterCircuitBgpIdData.bgpSessionUUID)

	return []*schema.ResourceData{d}, nil
}

func checkCloudRouterConnectionStatus(c *packetfabric.PFClient, cid string, id string) error {
	updateOk := make(chan bool)
	defer close(updateOk)

	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetCloudConnectionStatus(cid, id)
	}
	go c.CheckServiceStatus(updateOk, fn)

	if !<-updateOk {
		return fmt.Errorf("Failed to update connection status")
	}
	return nil
}
