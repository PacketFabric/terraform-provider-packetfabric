package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDocuments() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDocumentsCreate,
		ReadContext:   resourceDocumentsRead,
		DeleteContext: resourceDocumentsDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID for document",
			},
			"document": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Document file name",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Document type. One of [\"loa\", \"msa\"]",
				ValidateFunc: validation.StringInSlice([]string{"loa", "msa"}, false),
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Document description",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"port_circuit_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "Port circuit id. This field is required only for \"loa\" document type",
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func resourceDocumentsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	payload := packetfabric.DocumentsPayload{
		Document:    d.Get("document").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
	}

	if port_circuit_id, ok := d.GetOk("port_circuit_id"); ok {
		payload.PortCircuitId = port_circuit_id.(string)
	}

	resp, err := c.CreateDocument(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.UUID)
	return diags
}

func resourceDocumentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceDocumentsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Document cannot be deleted.",
	})
	d.SetId("")
	return diags
}
