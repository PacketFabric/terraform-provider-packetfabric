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

func showWarningForUnsetFields(unsetFields []string, diags *diag.Diagnostics) diag.Diagnostics {
	if len(unsetFields) > 0 {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Field(s) not set.",
			Detail:   fmt.Sprintf("The following fields: %s cannot be set. Update the Terraform state file manually if needed.", unsetFields),
		})
	}
	return *diags
}
