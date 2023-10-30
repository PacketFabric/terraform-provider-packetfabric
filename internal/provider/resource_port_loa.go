package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePortLoa() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePortLoaCreate,
		ReadContext:   resourcePortLoaRead,
		DeleteContext: resourcePortLoaDelete,
		Timeouts:      schemaTimeoutsCRD(10, 10, 10),
		Schema: map[string]*schema.Schema{
			PfId:               schemaStringComputedPlain(),
			PfPortCircuitId:    schemaStringRequiredNewNotEmpty(PfPortCircuitIdDescription2),
			PfLoaCustomerName:  schemaStringRequiredNewNotEmpty(PfLoaCustomerNameDescription),
			PfDestinationEmail: schemaStringRequiredNewNotEmpty(PfDestinationEmailDescription),
		},
	}
}

func resourcePortLoaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	portCID, ok := d.GetOk(PfPortCircuitId)
	if !ok {
		return diag.Errorf(MessageMissingPortCiruitId)
	}
	loaReq := packetfabric.PortLoa{
		LoaCustomerName:  d.Get(PfLoaCustomerName).(string),
		DestinationEmail: d.Get(PfDestinationEmail).(string),
	}
	_, err := c.SendPortLoa(portCID.(string), loaReq)
	if err != nil {
		return diag.FromErr(err)
	}
	if d.Id() == PfEmptyString {
		d.SetId(portCID.(string))
	}
	return diags
}

func resourcePortLoaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourcePortLoaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId(PfEmptyString)
	return diag.Diagnostics{}
}
