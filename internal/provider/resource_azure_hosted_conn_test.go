//go:build resource || hosted_cloud || all

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAzureHostedConnectionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"ARM_SUBSCRIPTION_ID", "ARM_CLIENT_ID", "ARM_CLIENT_SECRET", "ARM_TENANT_ID"})
	azureHostedConnectionResult := testutil.RHclCsAzureHostedConnection()
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config:             azureHostedConnectionResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(azureHostedConnectionResult.ResourceName, "description", azureHostedConnectionResult.Desc),
					resource.TestCheckResourceAttr(azureHostedConnectionResult.ResourceName, "speed", azureHostedConnectionResult.Speed),
					resource.TestCheckResourceAttr(azureHostedConnectionResult.ResourceName, "vlan_private", strconv.Itoa(azureHostedConnectionResult.VlanPrivate)),
				),
			},
			{
				ResourceName:            googleHostedConnectionResult.ResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"azure_service_key"},
			},
		},
	})
}
