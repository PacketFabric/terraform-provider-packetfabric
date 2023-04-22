package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclGoogleReqHostedConnectRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	csGoogleReqHostedConnectResult := testutil.RHclCsGoogleReqHostedConnect()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_CS_GOOGLE_HOS_CONN_SPEED1_KEY,
				testutil.PF_CS_GOOGLE_HOS_CONN_PAIRING_KEY,
				testutil.PF_CS_GOOGLE_HOS_CONN_VLAN_ATTACHMENT_NAME_KEY,
				testutil.PF_CS_GOOGLE_HOS_CONN_VLAN1_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: csGoogleReqHostedConnectResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csGoogleReqHostedConnectResult.ResourceName, "description", csGoogleReqHostedConnectResult.Desc),
					resource.TestCheckResourceAttr(csGoogleReqHostedConnectResult.ResourceName, "google_pairing_key", csGoogleReqHostedConnectResult.GooglePairingKey),
					resource.TestCheckResourceAttr(csGoogleReqHostedConnectResult.ResourceName, "google_vlan_attachment_name", csGoogleReqHostedConnectResult.GoogleVlan),
					resource.TestCheckResourceAttr(csGoogleReqHostedConnectResult.ResourceName, "pop", csGoogleReqHostedConnectResult.Pop),
					resource.TestCheckResourceAttr(csGoogleReqHostedConnectResult.ResourceName, "speed", csGoogleReqHostedConnectResult.Speed),
					resource.TestCheckResourceAttr(csGoogleReqHostedConnectResult.ResourceName, "vlan", strconv.Itoa(csGoogleReqHostedConnectResult.Vlan)),
				),
			},
			{
				ResourceName:      csGoogleReqHostedConnectResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
