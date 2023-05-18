//go:build resource || hosted_cloud || all

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccOracleHostedConnectionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"TF_VAR_tenancy_ocid", "TF_VAR_user_ocid", "TF_VAR_fingerprint", "TF_VAR_private_key", "TF_VAR_parent_compartment_id", "TF_VAR_pf_cs_oracle_drg_ocid"})
	csOracleHostedConnectionResult := testutil.RHclCsOracleHostedConnection()
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: csOracleHostedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "account_uuid", csOracleHostedConnectionResult.AccountUuid),
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "pop", csOracleHostedConnectionResult.Pop),
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "zone", csOracleHostedConnectionResult.Zone),
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "vlan", strconv.Itoa(csOracleHostedConnectionResult.Vlan)),
				),
			},
			{
				ResourceName:            csOracleHostedConnectionResult.ResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vc_ocid", "region"},
			},
		},
	})

}
