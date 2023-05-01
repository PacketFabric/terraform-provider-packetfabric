package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceDocuments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDocumentsRead,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the document",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Document name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Document description",
			},
			"mime_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mime type of the document",
			},
			"size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Document size",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Document type. One of [\"loa\", \"msa\"]",
			},
			"time_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Document creation time",
			},
			"time_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Document's last update time",
			},
			"_links": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A link to the port this document refers to",
						},
						"service": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A link to the service this document refers to",
						},
						"cloud": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A link to the cloud service this document refers to",
						},
						"cloud_router": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A link to the cloud service this document refers to",
						},
						"cloud_router_connection": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A link to the cloud router connection this document refers to",
						},
					},
				},
			},
		},
	}
}

func dataSourceDocumentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	_, err := c.GetDocuments()
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
