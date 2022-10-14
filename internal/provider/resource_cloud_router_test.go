package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclCloudRouter(name, account_uuid, region, capacity, asn string) (hcl string, resourceName string) {
	hcl = fmt.Sprintf(`
	resource "packetfabric_cloud_router" "cloud_router" {
		name         = "%s"
		account_uuid = "%s"
		regions      = ["%s"]
		capacity     = "%s"
		asn          = "%s"
	}`, name, account_uuid, region, capacity, asn)
	resourceName = "packetfabric_cloud_router.cloud_router"
	return
}

func TestAccCloudRouter(t *testing.T) {
	hcl, resourceName := hclCloudRouter("tf-acc-test", testutil.GetAccountUUID(), "US", "100Mbps", "12345")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "tf-acc-test"),
					resource.TestCheckResourceAttr(resourceName, "account_uuid", testutil.GetAccountUUID()),
					resource.TestCheckResourceAttr(resourceName, "regions.0", "US"),
					resource.TestCheckResourceAttr(resourceName, "capacity", "100Mbps"),
					resource.TestCheckResourceAttr(resourceName, "asn", "12345"),
				),
			},
		},
	})
}
