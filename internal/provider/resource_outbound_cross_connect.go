package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
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
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the PacketFabric port to which your are building the cross connect. This starts with \"PF-AP-\".",
			},
			"site": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The site code for the port location.",
			},
			"document_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "Document UUID for the LOA. When you order a cross connect, you must provide an LOA/CFA authorizing PacketFabric access to your equipment.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "PacketFabric outbound cross connect description.",
			},
			"destination_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side company name for the far side of the cross connect.",
			},
			"destination_circuit_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side circuit id for the far side of the cross connect.",
			},
			"panel": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side fiber panel info.",
			},
			"module": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side fiber module info.",
			},
			"position": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Z-side fiber position info.",
			},
			"data_center_cross_connect_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Display ID for the outbound cross connect.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this cross connect should be associated.",
			},
			"user_description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The user description used for update.",
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
	_, err := c.CreateOutboundCrossConnect(crossConn)
	if err != nil {
		return diag.FromErr(err)
	}

	time.Sleep(10 * time.Second)

	crossConns, err := c.ListOutboundCrossConnects()
	if err != nil {
		return diag.FromErr(err)
	}

	matchFound := false
	for _, crossConn := range *crossConns {
		if crossConn.Port == d.Get("port").(string) {
			d.SetId(crossConn.CircuitID)
			matchFound = true
			break
		}
	}
	if !matchFound {
		return diag.Errorf("Failed to find the cross connect with port: %s", d.Get("port").(string))
	}

	return diags
}

func resourceOutboundCrossConnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	resp, err := c.GetOutboundCrossConnect(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	_ = d.Set("port", resp.Port)
	_ = d.Set("site", resp.Site)
	_ = d.Set("document_uuid", resp.DocumentUUID)
	_ = d.Set("outbound_cross_connect_id", resp.OutboundCrossConnectID)
	_ = d.Set("obcc_status", resp.ObccStatus)
	_ = d.Set("description", resp.Description)
	_ = d.Set("user_description", resp.UserDescription)
	_ = d.Set("destination_name", resp.DestinationName)
	_ = d.Set("destination_circuit_id", resp.DestinationCircuitID)
	_ = d.Set("panel", resp.Panel)
	_ = d.Set("module", resp.Module)
	_ = d.Set("position", resp.Position)
	_ = d.Set("data_center_cross_connect_id", resp.DataCenterCrossConnectID)
	_ = d.Set("progress", resp.Progress)
	_ = d.Set("deleted", resp.Deleted)
	_ = d.Set("z_loc_cfa", resp.ZLocCfa)
	_ = d.Set("time_created", resp.TimeCreated)
	_ = d.Set("time_updated", resp.TimeUpdated)

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
	if port, ok := d.GetOk("port"); ok {
		err := c.DeleteOutboundCrossConnect(port.(string))
		if err != nil {
			return diag.FromErr(err)
		}
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
