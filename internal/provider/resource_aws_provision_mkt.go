package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsProvision() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsProvisionCreate,
		UpdateContext: resourceAwsProvisionUpdate,
		ReadContext:   resourceAwsProvisionRead,
		DeleteContext: resourceAwsProvisionDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vc_request_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "UUID of the service request you received from the marketplace customer.",
			},
			"port_circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The circuit ID of the port on which you want to provision the request. This starts with \"PF-AP-\".",
			},
			"vlan": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Valid VLAN range is from 4-4094, inclusive.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A brief description of this connection.",
			},
		},
	}
}

func resourceAwsProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	requestUUID, ok := d.GetOk("vc_request_uuid")
	if !ok {
		return diag.Errorf("please provide a valid VC Request UUID")
	}
	awsProvision := extractAwsProvision(d)
	_, err := c.CreateAwsProvisionReq(awsProvision, requestUUID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func resourceAwsProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceAwsProvisionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceAwsProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Connection provision requests",
		Detail:   "Connection provision requests must be deleted by deleting a connection request.",
	})
	d.SetId("")
	return diags
}

func extractAwsProvision(d *schema.ResourceData) packetfabric.ServiceAwsMktConn {
	return packetfabric.ServiceAwsMktConn{
		Provider: "aws",
		Interface: packetfabric.ServiceAwsInterf{
			PortCircuitID: d.Get("port_circuit_id").(string),
			Vlan:          d.Get("vlan").(int),
		},
		Description: d.Get("description").(string),
	}
}
