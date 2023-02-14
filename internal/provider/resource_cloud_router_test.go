package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclCloudRouter(name, accountUUID, region, capacity, asn string) (hcl string, resourceName string) {
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_cloud_router." + hclName
	hcl = fmt.Sprintf(`
	resource "packetfabric_cloud_router" "%s" {
		name         = "%s"
		account_uuid = "%s"
		regions      = ["%s"]
		capacity     = "%s"
		asn          = "%s"
	}
	
	resource "time_sleep" "delay_cloud_router_destroy" {
		depends_on = [packetfabric_cloud_router.%s]
	  
		destroy_duration = "10s"
	}`, hclName, name, accountUUID, region, capacity, asn, hclName)
	return
}

func TestAccCloudRouter(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	name := testutil.GenerateUniqueName(testPrefix)
	hcl, resourceName := hclCloudRouter(name, testutil.GetAccountUUID(), "US", "100Mbps", "4556")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutil.PreCheck(t, nil) },
		ProviderFactories: testAccProviderFactories,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_uuid", testutil.GetAccountUUID()),
					resource.TestCheckResourceAttr(resourceName, "regions.0", "US"),
					resource.TestCheckResourceAttr(resourceName, "capacity", "100Mbps"),
					resource.TestCheckResourceAttr(resourceName, "asn", "4556"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
