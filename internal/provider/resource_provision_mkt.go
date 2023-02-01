package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const azureProvider = "azure"
const googleProvider = "google"
const oracleProvider = "oracle"
const awsProvider = "aws"

func resourceProvision() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"vc_request_uuid": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "UUID of the service request you received from the marketplace customer.",
		},
		"port_circuit_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The circuit ID of the port on which you want to provision the request. This starts with \"PF-AP-\".",
		},
		"vlan": {
			Type:         schema.TypeInt,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IntBetween(4, 4094),
			Description:  "Valid VLAN range is from 4-4094, inclusive.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "A brief description of this connection.",
		},
	}
}

func resourceProvisionAzure() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"vc_request_uuid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "UUID of the service request you received from the marketplace customer.",
		},
		"port_circuit_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The circuit ID of the port on which you want to provision the request. This starts with \"PF-AP-\".",
		},
		"vlan_private": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(4, 4094),
			Description:  "Valid VLAN range is from 4-4094, inclusive.",
		},
		"vlan_microsoft": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(4, 4094),
			Description:  "Valid VLAN range is from 4-4094, inclusive.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A brief description of this connection.",
		},
	}
}

func resourceProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(packetfabric.ServiceAwsMktConn, string, string) (*packetfabric.MktConnProvisionResp, error), provider string) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	requestUUID, ok := d.GetOk("vc_request_uuid")
	if !ok {
		return diag.Errorf("please provide a valid VC Request UUID")
	}
	provision := extractProvision(d, provider)
	_, err := fn(provision, requestUUID.(string), provider)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func extractProvision(d *schema.ResourceData, provider string) packetfabric.ServiceAwsMktConn {
	mktConn := packetfabric.ServiceAwsMktConn{Provider: provider}
	interf := packetfabric.ServiceAwsInterf{}
	if portCid, ok := d.GetOk("port_circuit_id"); ok {
		interf.PortCircuitID = portCid.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		interf.Vlan = vlan.(int)
	}
	if vlanMicrosoft, ok := d.GetOk("vlan_microsoft"); ok {
		interf.VlanMicrosoft = vlanMicrosoft.(int)
	}
	if vlanPriv, ok := d.GetOk("vlan_private"); ok {
		interf.VlanPrivate = vlanPriv.(int)
	}
	if desc, ok := d.GetOk("description"); ok {
		mktConn.Description = desc.(string)
	}
	mktConn.Interface = interf
	return mktConn
}
