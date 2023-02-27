package provider

import (
	"context"
	"errors"
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
		if d.HasChange("description") {
			if desc, descOk := d.GetOk("description"); descOk {
				descUpdate := packetfabric.DescriptionUpdate{
					Description: desc.(string),
				}
				if _, err := c.UpdateCloudRouterConnection(cid.(string), d.Id(), descUpdate); err != nil {
					return diag.FromErr(err)
				}
			}
		}
		if d.HasChange("speed") {
			if speed, ok := d.GetOk("speed"); ok {
				billing := packetfabric.BillingUpgrade{
					Speed: speed.(string),
				}
				if _, err := c.ModifyBilling(d.Id(), billing); err != nil {
					return diag.FromErr(err)
				}
				_ = d.Set("speed", speed.(string))
			}
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
