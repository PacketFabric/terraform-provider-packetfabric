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
			PfCircuitId:                  schemaStringRequired(PfCircuitIdDescription3),
			PfCustomerGatewayAddress:     schemaStringComputed(PfCustomerGatewayAddressDescription),
			PfLocalGatewayAddress:        schemaStringComputed(PfLocalGatewayAddressDescription),
			PfIkeVersion:                 schemaIntComputed(PfIkeVersionDescription),
			PfPhase1AuthenticationMethod: schemaStringComputed(PfPhase1AuthenticationMethodDescription),
			PfPhase1Group:                schemaStringComputed(PfPhase1GroupDescription),
			PfPhase1EncryptionAlgo:       schemaStringComputed(PfPhase1EncryptionAlgoDescription),
			PfPhase1AuthenticationAlgo:   schemaStringComputed(PfPhase1AuthenticationAlgoDescription),
			PfPhase1Lifetime:             schemaIntComputed(PfPhase1LifetimeDescription),
			PfPhase2PfsGroup:             schemaStringComputed(PfPhase2PfsGroupDescription),
			PfPhase2EncryptionAlgo:       schemaStringComputed(PfPhase2EncryptionAlgoDescription),
			PfPhase2AuthenticationAlgo:   schemaStringComputed(PfPhase2AuthenticationAlgoDescription),
			PfPhase2Lifetime:             schemaIntComputed(PfPhase2LifetimeDescription),
			PfPreSharedKey:               schemaStringComputed(PfPreSharedKeyDescription),
			PfTimeCreated:                schemaStringComputed(PfTimeCreatedDescription8),
			PfTimeUpdated:                schemaStringComputed(PfTimeUpdatedDescription6),
		},
	}
}

func datasourceIPSecRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cID, ok := d.GetOk(PfCircuitId); ok {
		ipsec, err := c.GetIpsecSpecificConn(cID.(string))
		if err != nil {
			return diag.FromErr(err)
		}

		_ = setResourceDataKeys(d, ipsec, PfCustomerGatewayAddress, PfLocalGatewayAddress, PfIkeVersion, PfPhase1AuthenticationMethod, PfPhase1Group, PfPhase1EncryptionAlgo, PfPhase1AuthenticationAlgo, PfPhase1Lifetime, PfPhase2PfsGroup, PfPhase2EncryptionAlgo, PfPhase2AuthenticationAlgo, PfPhase2Lifetime, PfPreSharedKey, PfTimeCreated, PfTimeUpdated)

		d.SetId(cID.(string) + "-data")
	}
	return diags
}
