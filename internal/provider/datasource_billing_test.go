package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceBillingComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceBillingResult := testutil.DHclDatasourceBilling()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceBillingResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "circuit_id"),

					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.order_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.account_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.order_type"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.parent_order"),

					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.term.0.start_date"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.term.0.months"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.term.0.termination_date"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.term.0.commitment_end_date"),

					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.account_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.billable_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.order_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.price_type"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.currency_code"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.adjusted_price"),

					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.product_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.name"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.vc_type"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.vc_service_class"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.bundle_type"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.active_date"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.end_date"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.translation_id"),
				),
			},
		},
	})

}
