package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIPSec() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIPSecRead,
		Schema: map[string]*schema.Schema{
			"circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IPSec circuit ID or its associated VC.",
			},
			"customer_gateway_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Customer Gateway Address.",
			},
			"local_gateway_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Local gateway address.",
			},
			"ike_version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The IKE version.",
			},
			"phase1_authentication_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Phase 1 authentication method.",
			},
			"phase1_group": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phase 1 group.",
			},
			"phase1_encryption_algo": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phase 1 encryption algorithm.",
			},
			"phase1_authentication_algo": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phase 1 authentication algorithm.",
			},
			"phase1_lifetime": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The phase 1 lifetime.",
			},
			"phase2_pfs_group": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phase 2 PFS group.",
			},
			"phase2_encryption_algo": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phase 2 encryption algorithm.",
			},
			"phase2_authentication_algo": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phase 2 authentication algorithm.",
			},
			"phase2_lifetime": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The phase 2 lifetime.",
			},
			"pre_shared_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The pre shared key.",
			},
			"time_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time created.",
			},

			"time_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time updated.",
			},
		},
	}
}

func datasourceIPSecRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cID, ok := d.GetOk("circuit_id"); ok {
		ipsec, err := c.GetIpsecSpecificConn(cID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		_ = d.Set("customer_gateway_address", ipsec.CustomerGatewayAddress)
		_ = d.Set("local_gateway_address", ipsec.LocalGatewayAddress)
		_ = d.Set("ike_version", ipsec.IkeVersion)
		_ = d.Set("phase1_authentication_method", ipsec.Phase1AuthenticationMethod)
		_ = d.Set("phase1_group", ipsec.Phase1Group)
		_ = d.Set("phase1_encryption_algo", ipsec.Phase1EncryptionAlgo)
		_ = d.Set("phase1_authentication_algo", ipsec.Phase1AuthenticationAlgo)
		_ = d.Set("phase1_lifetime", ipsec.Phase1Lifetime)
		_ = d.Set("phase2_pfs_group", ipsec.Phase2PfsGroup)
		_ = d.Set("phase2_encryption_algo", ipsec.Phase2EncryptionAlgo)
		_ = d.Set("phase2_authentication_algo", ipsec.Phase2AuthenticationAlgo)
		_ = d.Set("phase2_lifetime", ipsec.Phase2Lifetime)
		_ = d.Set("pre_shared_key", ipsec.PreSharedKey)
		_ = d.Set("time_created", ipsec.TimeCreated)
		_ = d.Set("time_updated", ipsec.TimeUpdated)
		d.SetId(cID.(string) + "-data")
	}
	return diags
}
