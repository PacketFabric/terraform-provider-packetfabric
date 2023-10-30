package provider

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePointToPoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePointToPointCreate,
		ReadContext:   resourcePointToPointRead,
		UpdateContext: resourcePointToPointUpdate,
		DeleteContext: resourcePointToPointDelete,
		Timeouts:      schema10MinuteTimeouts(),
		Schema: map[string]*schema.Schema{
			PfId:          schemaStringComputedPlain(),
			PfPtpUuid:     schemaStringComputed(PfPtpUuidDescription2),
			PfDescription: schemaStringRequiredNotEmpty(PfConnectionDescription),
			PfSpeed:       schemaStringRequiredNewValidate(PfSpeedDescription7, validateSpeed1or10or40or100()),
			PfMedia:       schemaStringRequiredNewValidate(PfMediaDescription5, validateMedia()),
			PfEndpoints: {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPop:              schemaStringRequiredNewNotEmpty(PfPopDescription3),
						PfZone:             schemaStringRequiredNewNotEmpty(PfZoneDescription6),
						PfCustomerSiteCode: schemaStringOptionalNewNotEmpty(PfCustomerSiteCodeDescription2),
						PfAutoneg:          schemaBoolRequiredNew(PfAutonegDescription5),
						PfLoa:              schemaStringOptionalNewValidate(PfLoaDescription4, validation.StringIsBase64),
						PfPortCircuitId:    schemaStringComputed(PfPortCircuitIdDescription4),
					},
				},
			},
			PfAccountUuid:            schemaAccountUuid(PfAccountUuidDescription2),
			PfSubscriptionTerm:       schemaIntRequiredValidate(PfSubscriptionTermDescriptionB, validateSubscriptionTerm()),
			PfPublishedQuoteLineUuid: schemaStringOptionalNewValidate(PfPublishedQuoteLineUuidDescription2, validation.IsUUID),
			PfPoNumber:               schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfLabels:                 schemaStringSetOptional(PfLabelsDescription),
			PfEtl:                    schemaFloatComputed(PfEtlDescription),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePointToPointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ptpService := extractPtpService(d)
	resp, err := c.CreatePointToPointService(ptpService)
	if err != nil {
		return diag.FromErr(err)
	}

	host := os.Getenv(PfPfeHost)
	testingInLab := strings.Contains(host, PfDevLab)

	if testingInLab {
		time.Sleep(time.Duration(30) * time.Second)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  MessageDevDelay,
		})
	}

	if err2 := checkPtpStatus(c, resp.PtpCircuitID); err2 != nil {
		return diag.FromErr(err2)
	}
	if resp != nil {
		_ = d.Set(PfPtpUuid, resp.PtpUUID)
		d.SetId(resp.PtpCircuitID)

		if labels, ok := d.GetOk(PfLabels); ok {
			diagnostics, created := createLabels(c, d.Id(), labels)
			if !created {
				return diagnostics
			}
		}
	}
	return diags
}

func resourcePointToPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.ReadPointToPoint(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = setResourceDataKeys(d, resp, PfPtpUuid, PfDescription, PfSpeed, PfMedia, PfPoNumber)
		_ = setResourceDataKeys(d, resp.Billing, PfAccountUuid, PfSubscriptionTerm)

		if len(resp.Interfaces) == 2 {
			interface1 := mapStruct(resp.Interfaces[0], PfPop, PfZone, PfCustomerSiteCode, PfPortCircuitId)
			interface2 := mapStruct(resp.Interfaces[1], PfPop, PfZone, PfCustomerSiteCode, PfPortCircuitId)

			endpoints := []interface{}{interface1, interface2}
			_ = d.Set(PfEndpoints, endpoints)
		}
	}
	// unsetFields: loa, autoneg, published_quote_line_uuid
	if _, ok := d.GetOk(PfLabels); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set(PfLabels, labels)
	}

	etl, err3 := c.GetEarlyTerminationLiability(d.Id())
	if err3 != nil {
		return diag.FromErr(err3)
	}
	if etl > 0 {
		_ = d.Set(PfEtl, etl)
	}
	return diags
}

func resourcePointToPointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	if d.HasChanges([]string{PfPoNumber, PfDescription}...) {
		updatePointToPointData := packetfabric.UpdatePointToPointData{}
		desc, ok := d.GetOk(PfDescription)
		if !ok {
			return diag.FromErr(errors.New(MessageMissingDescription))
		}
		updatePointToPointData.Description = desc.(string)

		poNumber, ok := d.GetOk(PfPoNumber)
		if !ok {
			return diag.FromErr(errors.New(MessageMissingPoNumber))
		}
		updatePointToPointData.PONumber = poNumber.(string)
		ptpUuid := d.Get(PfPtpUuid).(string) // must use the UUID to update the PTP
		if _, err := c.UpdatePointToPoint(ptpUuid, updatePointToPointData); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange(PfSubscriptionTerm) {
		if subTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
			billing := packetfabric.BillingUpgrade{
				SubscriptionTerm: subTerm.(int),
			}
			cID := d.Get(PfPtpCircuitId).(string)
			if _, err := c.ModifyBilling(cID, billing); err != nil {
				return diag.FromErr(err)
			}
			_ = d.Set(PfSubscriptionTerm, subTerm.(int))
		} else {
			return diag.Errorf(MessageMissingSubscriptionTerm)
		}
	}

	if d.HasChange(PfLabels) {
		labels := d.Get(PfLabels)
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}
	return diags
}

func resourcePointToPointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	host := os.Getenv(PfPfeHost)
	testingInLab := strings.Contains(host, PfDevLab)

	if testingInLab {
		endpoints := d.Get(PfEndpoints).([]interface{})
		for _, v := range endpoints {
			endpoint := v.(map[string]interface{})
			portCircuitID := endpoint[PfPortCircuitId].(string)
			if toggleErr := _togglePortStatus(c, false, portCircuitID); toggleErr != nil {
				return diag.FromErr(toggleErr)
			}
		}
		time.Sleep(time.Duration(180) * time.Second)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  MessageDevDisableToDelete,
		})
	}
	etlDiags, err := addETLWarning(c, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	diags = append(diags, etlDiags...)
	ptpUuid := d.Get(PfPtpUuid).(string) // must use the UUID to delete the PTP
	if err := c.DeletePointToPointService(ptpUuid); err != nil {
		return diag.FromErr(err)
	} else {
		if err := checkPtpStatus(c, d.Id()); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(PfEmptyString)
	return diags
}

func extractPtpService(d *schema.ResourceData) packetfabric.PointToPoint {
	ptpService := packetfabric.PointToPoint{}
	if description, ok := d.GetOk(PfDescription); ok {
		ptpService.Description = description.(string)
	}
	if speed, ok := d.GetOk(PfSpeed); ok {
		ptpService.Speed = speed.(string)
	}
	if media, ok := d.GetOk(PfMedia); ok {
		ptpService.Media = media.(string)
	}
	if endpoints, ok := d.GetOk(PfEndpoints); ok {
		edps := make([]packetfabric.Endpoints, 0)
		for _, endpoint := range endpoints.([]interface{}) {
			edps = append(edps, packetfabric.Endpoints{
				Pop:              endpoint.(map[string]interface{})[PfPop].(string),
				Zone:             endpoint.(map[string]interface{})[PfZone].(string),
				CustomerSiteCode: endpoint.(map[string]interface{})[PfCustomerSiteCode].(string),
				Autoneg:          endpoint.(map[string]interface{})[PfAutoneg].(bool),
				Loa:              endpoint.(map[string]interface{})[PfLoa].(string),
				PortCircuitID:    endpoint.(map[string]interface{})[PfPortCircuitId].(string),
			})
		}
		ptpService.Endpoints = edps
	}
	if accountUUID, ok := d.GetOk(PfAccountUuid); ok {
		ptpService.AccountUUID = accountUUID.(string)
	}
	if subTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
		ptpService.SubscriptionTerm = subTerm.(int)
	}
	if quote, ok := d.GetOk(PfPublishedQuoteLineUuid); ok {
		ptpService.PublishedQuoteLineUUID = quote.(string)
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		ptpService.PONumber = poNumber.(string)
	}
	return ptpService
}

func checkPtpStatus(c *packetfabric.PFClient, cid string) error {
	statusOk := make(chan bool)
	defer close(statusOk)

	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetPointToPointStatus(cid)
	}
	go c.CheckServiceStatus(statusOk, fn)
	time.Sleep(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	if !<-statusOk {
		return fmt.Errorf("failed to retrieve the status for %s", cid)
	}
	return nil
}
