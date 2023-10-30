package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBilling() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBillingRead,
		Schema: map[string]*schema.Schema{
			PfCircuitId: schemaStringRequired(PfCircuitIdDescription2),
			PfBillings: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfOrderId:   schemaIntComputed(PfOrderIdDescription),
						PfAccountId: schemaStringComputed(PfAccountIdDescription),
						PfOrderType: schemaStringComputed(PfOrderTypeDescription),
						PfTerm: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfStartDate:         schemaStringComputed(PfStartDateDescription),
									PfMonths:            schemaIntComputed(PfMonthsDescription),
									PfTerminationDate:   schemaStringComputed(PfTerminationDateDescription),
									PfCommitmentEndDate: schemaStringComputed(PfCommitmentEndDateDescription),
								},
							},
						},
						PfBillables: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfAccountId:     schemaStringComputed(PfAccountIdDescription),
									PfBillableId:    schemaIntComputed(PfBillableIdDescription),
									PfOrderId:       schemaIntComputed(PfOrderIdDescription),
									PfPriceType:     schemaStringComputed(PfPriceTypeDescription),
									PfCurrencyCode:  schemaStringComputed(PfCurrencyCodeDescription),
									PfPrice:         schemaIntComputed(PfPriceDescription),
									PfAdjustedPrice: schemaIntComputed(PfAdjustedPriceDescription),
								},
							},
						},
						PfProductDetails: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfProductId:      schemaIntComputed(PfProductIdDescription),
									PfName:           schemaStringComputed(PfOrderTypeDescription),
									PfVcType:         schemaStringComputed(PfVcTypeDescription),
									PfVcServiceClass: schemaStringComputed(PfVcServiceClassDescription),
									PfBundleType:     schemaStringComputed(PfBundleTypeDescription),
									PfActiveDate:     schemaStringComputed(PfActiveDateDescription),
									PfEndDate:        schemaStringComputed(PfEndDateDescription),
									PfTranslationId:  schemaIntComputed(PfTranslationIdDescription),
								},
							},
						},
						PfParentOrder: schemaIntComputed(PfParentOrderDescription),
					},
				},
			},
		},
	}
}

func dataSourceBillingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk(PfCircuitId)
	if !ok {
		return diag.Errorf(MessageMissingCircuitIdDetail)
	}
	billings, err := c.ReadBilling(cID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfBillings, flattenBillings(&billings))
	d.SetId(uuid.New().String())
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func flattenBillings(billings *[]packetfabric.BillingResponse) []interface{} {
	fields := stringsToMap(PfOrderId, PfAccountId, PfOrderType, PfParentOrder)
	if billings != nil {
		flattens := make([]interface{}, len(*billings), len(*billings))

		for i, billing := range *billings {
			flatten := structToMap(&billing, fields)
			flatten[PfTerm] = flattenTerm(&billing.Term)
			flatten[PfBillables] = flattenBillables(&billing.Billables)
			flatten[PfProductDetails] = flattenProductDetails(&billing.ProductDetails)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenBillables(billables *[]packetfabric.Billables) []interface{} {
	fields := stringsToMap(PfAccountId, PfBillableId, PfOrderId, PfPriceType, PfCurrencyCode, PfPrice, PfAdjustedPrice)
	if billables != nil {
		flattens := make([]interface{}, len(*billables), len(*billables))

		for i, billable := range *billables {
			flattens[i] = structToMap(&billable, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenProductDetails(details *packetfabric.ProductDetails) []interface{} {
	fields := stringsToMap(PfProductId, PfName, PfVcType, PfVcServiceClass, PfBundleType, PfActiveDate, PfEndDate, PfTranslationId)
	flattens := make([]interface{}, 0)
	if details != nil {
		flattens = append(flattens, structToMap(details, fields))
	}
	return flattens
}

func flattenTerm(term *packetfabric.Term) []interface{} {
	fields := stringsToMap(PfStartDate, PfMonths, PfTerminationDate, PfCommitmentEndDate)
	flattens := make([]interface{}, 0)
	if term != nil {
		flattens = append(flattens, structToMap(term, fields))
	}
	return flattens
}
