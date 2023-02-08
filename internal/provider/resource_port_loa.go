package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePortLoa() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePortLoaCreate,
		ReadContext:   resourcePortLoaRead,
		DeleteContext: resourcePortLoaDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port circuit ID.",
			},
			"loa_customer_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The LOA customer name.",
			},
			"destination_email": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The LOA destination e-mail.",
			},
		},
	}
}

func resourcePortLoaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	portCID, ok := d.GetOk("port_circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid port circuit ID")
	}
	loaReq := packetfabric.PortLoa{
		LoaCustomerName:  d.Get("loa_customer_name").(string),
		DestinationEmail: d.Get("destination_email").(string),
	}
	_, err := c.SendPortLoa(portCID.(string), loaReq)
	if err != nil {
		return diag.FromErr(err)
	}
	if d.Id() == "" {
		d.SetId(portCID.(string))
	}
	return diags
}

func resourcePortLoaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourcePortLoaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return diag.Diagnostics{}
}
