package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOutboundCrossConnect() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		CreateContext: resourceOutboundCrossConnectCreate,
		ReadContext:   resourceOutboundCrossConnectRead,
		UpdateContext: resourceOutboundCrossConnectUpdate,
		DeleteContext: resourceOutboundCrossConnectDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "port_circuit_id to use for the cross connect.",
			},
			"site": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "site_code for the port location.",
			},
			"document_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "Document UUID for the LOAD",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "PacketFabric outbound cross connect description.",
			},
			"destination_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side company name for the far side of the cross connect.",
			},
			"destination_circuit_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side circuit id for the far side of the cross connect.",
			},
			"panel": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side fiber panel info.",
			},
			"module": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side fiber module info.",
			},
			"position": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side fiber position info.",
			},
			"data_center_cross_connect_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Display ID for the OBCC.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this cross connect should be associated.",
			},
			"user_description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The user desctiption used for update.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceOutboundCrossConnectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	crossConn := extractCrossConnect(d)
	resp, err := c.CreateOutboundCrossConnect(crossConn)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Outbound Cross Connect Create",
		Detail:   resp.Message,
	})
	return diags
}

func resourceOutboundCrossConnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if crossConnID, ok := d.GetOk("data_center_cross_connect_id"); ok {
		resp, err := c.GetOutboundCrossConnect(crossConnID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Existing outbound cross connect ID",
			Detail:   resp.OutboundCrossConnectID,
		})
	}
	return diags
}

func resourceOutboundCrossConnectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if userDesc, ok := d.GetOk("user_description"); ok {
		err := c.UpdateOutboundCrossConnect(d.Id(), userDesc.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Outbound Cross Connect update",
			Detail:   "Please provide a valid User Description for update.",
		})
	}
	return diags
}

func resourceOutboundCrossConnectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	err := c.DeleteOutboundCrossConnect(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func extractCrossConnect(d *schema.ResourceData) packetfabric.OutboundCrossConnect {
	crossConn := packetfabric.OutboundCrossConnect{}
	crossConn.Port = d.Get("port").(string)
	crossConn.Site = d.Get("site").(string)
	crossConn.DocumentUUID = d.Get("document_uuid").(string)
	if desc, ok := d.GetOk("description"); ok {
		crossConn.Description = desc.(string)
	}
	if destinationName, ok := d.GetOk("destination_name"); ok {
		crossConn.DestinationName = destinationName.(string)
	}
	if destinationCID, ok := d.GetOk("destination_circuit_id"); ok {
		crossConn.DestinationCircuitID = destinationCID.(string)
	}
	if panel, ok := d.GetOk("panel"); ok {
		crossConn.Panel = panel.(string)
	}
	if module, ok := d.GetOk("module"); ok {
		crossConn.Module = module.(string)
	}
	if position, ok := d.GetOk("position"); ok {
		crossConn.Position = position.(string)
	}
	if dataCenterCrossConnID, ok := d.GetOk("data_center_cross_connect_id"); ok {
		crossConn.DataCenterCrossConnectID = dataCenterCrossConnID.(string)
	}
	if publishedQuoteLineUUID, ok := d.GetOk("published_quote_line_uuid"); ok {
		crossConn.PublishedQuoteLineUUID = publishedQuoteLineUUID.(string)
	}
	return crossConn
}
