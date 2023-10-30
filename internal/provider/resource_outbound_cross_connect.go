package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOutboundCrossConnect() *schema.Resource {
	return &schema.Resource{
		Timeouts:      schema10MinuteTimeouts(),
		CreateContext: resourceOutboundCrossConnectCreate,
		ReadContext:   resourceOutboundCrossConnectRead,
		UpdateContext: resourceOutboundCrossConnectUpdate,
		DeleteContext: resourceOutboundCrossConnectDelete,
		Schema: map[string]*schema.Schema{
			PfId:                       schemaStringComputedPlain(),
			PfPort:                     schemaStringRequiredNewNotEmpty(PfPortDescription8),
			PfSite:                     schemaStringRequiredNewNotEmpty(PfSiteDescription8),
			PfDocumentUuid:             schemaStringRequiredNewValidate(PfDocumentUuidDescription2, validation.IsUUID),
			PfDescription:              schemaStringRequiredNewNotEmpty(PfOutboundCrossConnectsDescription3),
			PfDestinationName:          schemaStringOptionalNewNotEmpty(PfDestinationNameDescription2),
			PfDestinationCircuitId:     schemaStringOptionalNewNotEmpty(PfDestinationCircuitIdDescription2),
			PfPanel:                    schemaStringOptionalNewNotEmpty(PfPanelDescription2),
			PfModule:                   schemaStringOptionalNewNotEmpty(PfModuleDescription2),
			PfPosition:                 schemaStringOptionalNewNotEmpty(PfPositionDescription2),
			PfDataCenterCrossConnectId: schemaStringOptionalNewNotEmpty(PfDataCenterCrossConnectIdDescription2),
			PfPublishedQuoteLineUuid:   schemaStringOptionalNewValidate(PfPublishedQuoteLineUuidDescription4, validation.IsUUID),
			PfUserDescription:          schemaStringOptionalNotEmpty(PfUserDescriptionDescription2),
			PfOutboundCrossConnectId:   schemaStringComputed(PfOutboundCrossConnectIdDescription),
			PfObccStatus:               schemaStringComputed(PfObccStatusDescription),
			PfProgress:                 schemaIntComputed(PfProgressDescription),
			PfDeleted:                  schemaBoolComputed(PfDeletedDescription4),
			PfZLocCfa:                  schemaStringComputed(PfZLocCfaDescription),
			PfTimeCreated:              schemaStringComputed(PfTimeCreatedDescription7),
			PfTimeUpdated:              schemaStringComputed(PfTimeUpdatedDescription7),
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

	circuitID, err := getCrossConnectCircuitID(c, d.Get(PfPort).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if circuitID == "" {
		return diag.Errorf("Failed to find the cross connect with port: %s", d.Get(PfPort).(string))
	}

	d.SetId(circuitID)

	return diags
}

func getCrossConnectCircuitID(c *packetfabric.PFClient, port string) (string, error) {
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			crossConns, err := c.ListOutboundCrossConnects()
			if err != nil {
				return "", err
			}

			for _, crossConn := range *crossConns {
				if crossConn.Port == port {
					return crossConn.CircuitID, nil
				}
			}
		case <-timeout:
			return "", fmt.Errorf("timed out waiting for cross connect with port: %s", port)
		}
	}
}

func resourceOutboundCrossConnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	resp, err := c.GetOutboundCrossConnect(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	_ = setResourceDataKeys(d, resp, PfPort, PfSite, PfDocumentUuid, PfOutboundCrossConnectId, PfObccStatus, PfDescription, PfUserDescription, PfDestinationName, PfDestinationCircuitId, PfPanel, PfModule, PfPosition, PfDataCenterCrossConnectId, PfProgress, PfDeleted, PfZLocCfa, PfTimeCreated, PfTimeUpdated)

	return diags
}

func resourceOutboundCrossConnectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if userDesc, ok := d.GetOk(PfUserDescription); ok {
		err := c.UpdateOutboundCrossConnect(d.Id(), userDesc.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  MessageObccUpdate,
			Detail:   MessageObccUpdateDescription,
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
	crossConn.Port = d.Get(PfPort).(string)
	crossConn.Site = d.Get(PfSite).(string)
	crossConn.DocumentUUID = d.Get(PfDocumentUuid).(string)
	if desc, ok := d.GetOk(PfDescription); ok {
		crossConn.Description = desc.(string)
	}
	if destinationName, ok := d.GetOk(PfDestinationName); ok {
		crossConn.DestinationName = destinationName.(string)
	}
	if destinationCID, ok := d.GetOk(PfDestinationCircuitId); ok {
		crossConn.DestinationCircuitID = destinationCID.(string)
	}
	if panel, ok := d.GetOk(PfPanel); ok {
		crossConn.Panel = panel.(string)
	}
	if module, ok := d.GetOk(PfModule); ok {
		crossConn.Module = module.(string)
	}
	if position, ok := d.GetOk(PfPosition); ok {
		crossConn.Position = position.(string)
	}
	if dataCenterCrossConnID, ok := d.GetOk(PfDataCenterCrossConnectId); ok {
		crossConn.DataCenterCrossConnectID = dataCenterCrossConnID.(string)
	}
	if publishedQuoteLineUUID, ok := d.GetOk(PfPublishedQuoteLineUuid); ok {
		crossConn.PublishedQuoteLineUUID = publishedQuoteLineUUID.(string)
	}
	return crossConn
}
