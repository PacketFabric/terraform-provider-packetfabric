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
			"circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Circuit ID of the target cloud router.\n\t\tExample: \"PF-L3-CUST-2\"",
			},
			"billings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"order_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "PacketFabric Order Number computed for this request.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "PacketFabric account UUID. The contact that will be billed.",
						},
						"order_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "PacketFabric product name",
						},
						"term": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Billing date-time start",
									},
									"months": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: "Contract term in months",
									},
									"termination_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Date-time for service/order pre-mature termination",
									},
									"commitment_end_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Date-time for service/order contracted/planned termination",
									},
								},
							},
						},
						"billables": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "PacketFabric account UUID. The contact that will be billed.",
									},
									"billable_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: "PacketFabric billing item tracking ID",
									},
									"order_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: "PacketFabric Order Number computed for this request.",
									},
									"price_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Pricing interval\n\t\tEnum: [\"monthly\"]",
									},
									"currency_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Currency short name.\n\t\tEnum [\"USD\"]",
									},
									"price": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: "Catalog currency unit quantity",
									},
									"adjusted_price": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: "Contractual currency unit quantity",
									},
								},
							},
						},
						"product_details": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"product_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: "PacketFabric product catalog product code",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "PacketFabric product name",
									},
									"vc_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "PacketFabric connectivity type",
									},
									"vc_service_class": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "The service class for the given port, either long haul or metro.\n\t\tEnum: [\"longhaul\",\"metro\"]",
									},
									"bundle_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "PacketFabric product package name",
									},
									"active_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Service date-time start",
									},
									"end_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Service date-time planned end",
									},
									"translation_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: "Internal translation identifier",
									},
								},
							},
						},
						"parent_order": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "Service order number dependancy.",
						},
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
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid circuit_id")
	}
	billings, err := c.ReadBilling(cID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("billings", flattenBillings(&billings))
	d.SetId(uuid.New().String())
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func flattenBillings(billings *[]packetfabric.BillingResponse) []interface{} {
	if billings != nil {
		flattens := make([]interface{}, len(*billings), len(*billings))

		for i, billing := range *billings {
			flatten := make(map[string]interface{})
			flatten["order_id"] = billing.OrderID
			flatten["account_id"] = billing.AccountID
			flatten["order_type"] = billing.OrderType
			flatten["parent_order"] = billing.ParentOrder
			flatten["billables"] = flattenBillables(&billing.Billables)
			flatten["product_details"] = flattenProductDetails(&billing.ProductDetails)
			flatten["term"] = flattenTerm(&billing.Term)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenBillables(billables *[]packetfabric.Billables) []interface{} {
	if billables != nil {
		flattens := make([]interface{}, len(*billables), len(*billables))

		for i, billable := range *billables {
			flatten := make(map[string]interface{})
			flatten["account_id"] = billable.AccountID
			flatten["billable_id"] = billable.BillableID
			flatten["order_id"] = billable.OrderID
			flatten["price_type"] = billable.PriceType
			flatten["currency_code"] = billable.CurrencyCode
			flatten["price"] = billable.Price
			flatten["adjusted_price"] = billable.AdjustedPrice
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenProductDetails(details *packetfabric.ProductDetails) []interface{} {
	flattens := make([]interface{}, 0)
	if details != nil {
		flatten := make(map[string]interface{})
		flatten["product_id"] = details.ProductID
		flatten["name"] = details.Name
		flatten["vc_type"] = details.VcType
		flatten["vc_service_class"] = details.VcServiceClass
		flatten["bundle_type"] = details.BundleType
		flatten["active_date"] = details.ActiveDate
		flatten["end_date"] = details.EndDate
		flatten["translation_id"] = details.TranslationID
		flattens = append(flattens, flatten)
	}
	return flattens
}

func flattenTerm(term *packetfabric.Term) []interface{} {
	flattens := make([]interface{}, 0)
	if term != nil {
		flatten := make(map[string]interface{})
		flatten["start_date"] = term.StartDate
		flatten["months"] = term.Months
		flatten["termination_date"] = term.TerminationDate
		flatten["commitment_end_date"] = term.CommitmentEndDate
		flattens = append(flattens, flatten)
	}
	return flattens
}
