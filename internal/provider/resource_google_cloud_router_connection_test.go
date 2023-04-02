package provider

import (
	"log"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclCloudRouterConnectionGoogleRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	cloudRouterConnectionAwsResult := testutil.RHclCloudRouterConnectionGoogle()
	log.Fatal(cloudRouterConnectionAwsResult.Hcl)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_CRC_GOOGLE_PAIRING_KEY,
				testutil.PF_CRC_GOOGLE_VLAN_ATTACHMENT_NAME_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterConnectionAwsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionAwsResult.ResourceName, "description", cloudRouterConnectionAwsResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionAwsResult.ResourceName, "google_pairing_key", cloudRouterConnectionAwsResult.GooglePairingKey),
					resource.TestCheckResourceAttr(cloudRouterConnectionAwsResult.ResourceName, "google_vlan_attachment_name", cloudRouterConnectionAwsResult.GoogleVlanAttachmentName),
					resource.TestCheckResourceAttr(cloudRouterConnectionAwsResult.ResourceName, "pop", cloudRouterConnectionAwsResult.Pop),
					resource.TestCheckResourceAttr(cloudRouterConnectionAwsResult.ResourceName, "speed", cloudRouterConnectionAwsResult.Speed),
				),
			},
		},
	})
}
