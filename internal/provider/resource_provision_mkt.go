package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const awsProvider = "aws"
const azureProvider = "azure"
const googleProvider = "google"

func resourceProvision() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"vc_request_uuid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "UUID of the service request",
		},
		"port_circuit_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The circuit ID of the customer's port.",
		},
		"vlan": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Valid VLAN range is from 4-4094, inclusive.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Connection description",
		},
	}
}

func resourceProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(packetfabric.ServiceAwsMktConn, string) (*packetfabric.MktConnProvisionResp, error), provider string) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	requestUUID, ok := d.GetOk("vc_request_uuid")
	if !ok {
		return diag.Errorf("please provide a valid VC Request UUID")
	}
	awsProvision := extractProvision(d, provider)
	_, err := fn(awsProvision, requestUUID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func resourceProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Connection provision requests",
		Detail:   "Connection provision requests must be deleted by deleting a connection request.",
	})
	d.SetId("")
	return diags
}

func extractProvision(d *schema.ResourceData, provider string) packetfabric.ServiceAwsMktConn {
	return packetfabric.ServiceAwsMktConn{
		Provider: provider,
		Interface: packetfabric.ServiceAwsInterf{
			PortCircuitID: d.Get("port_circuit_id").(string),
			Vlan:          d.Get("vlan").(int),
		},
		Description: d.Get("description").(string),
	}
}
