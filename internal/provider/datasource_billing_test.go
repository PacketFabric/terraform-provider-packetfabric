//go:build datasource || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceBillingComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourceBillingResult := testutil.DHclDatasourceBilling()

	resource.ParallelTest(t, resource.TestCase{
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

					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.account_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.billable_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.order_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.price_type"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.currency_code"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.billables.0.adjusted_price"),

					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.product_id"),
					resource.TestCheckResourceAttrSet(datasourceBillingResult.ResourceName, "billings.0.product_details.0.name"),
				),
			},
		},
	})

}
