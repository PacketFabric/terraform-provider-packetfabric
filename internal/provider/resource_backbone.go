package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackbone() *schema.Resource {
	return &schema.Resource{
		Timeouts:      schema10MinuteTimeouts(),
		CreateContext: resourceBackboneCreate,
		UpdateContext: resourceBackboneUpdate,
		ReadContext:   resourceBackboneRead,
		DeleteContext: resourceBackboneDelete,
		Schema: map[string]*schema.Schema{
			PfId:          schemaStringComputedPlain(),
			PfDescription: schemaStringRequired(PfConnectionDescription),
			PfBandwidth: {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfAccountUuid:      schemaAccountUuid(PfAccountUuidDescription2),
						PfSpeed:            schemaStringOptionalValidate(PfSpeedDescriptionB, validateSpeed()),
						PfSubscriptionTerm: schemaIntOptionalValidate(PfSubscriptionTermDescriptionA, validateSubscriptionTerm()),
						PfLonghaulType:     schemaStringOptionalValidate(PfLonghaulTypeDescription, validateLongHaul()),
					},
				},
			},
			PfInterfaceA: resourceBackboneInterface(),
			PfInterfaceZ: resourceBackboneInterface(),
			PfRateLimitIn:     schemaIntOptional(PfRateLimitInDescription2),
			PfRateLimitOut:    schemaIntOptional(PfRateLimitOutDescription2),
			PfEpl:             schemaBoolOptionalNewDefault(PfEplDescription, false),
			PfFlexBandwidthId: schemaStringOptionalNew(PfFlexBandwidthIdDescription),
			PfPoNumber:        schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfLabels:          schemaStringSetOptional(PfLabelsDescription),
			PfEtl:             schemaFloatComputed(PfEtlDescription),
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, d *schema.ResourceDiff, m interface{}) error {
				if d.Id() == PfEmptyString {
					return nil
				}

				interfaces := []string{PfInterfaceA, PfInterfaceZ}

				for _, iface := range interfaces {
					oldRaw, newRaw := d.GetChange(iface)
					oldSet := oldRaw.(*schema.Set)
					newSet := newRaw.(*schema.Set)

					for _, oldElem := range oldSet.List() {
						for _, newElem := range newSet.List() {
							oldResource := oldElem.(map[string]interface{})
							newResource := newElem.(map[string]interface{})

							if oldResource[PfPortCircuitId] != newResource[PfPortCircuitId] {
								return fmt.Errorf("updating %s port_circuit_id in-place is not supported, delete and recreate the resource with the updated values", iface)
							}
						}
					}
				}

				return nil
			},
		),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// TODO: consider moving this to common_schema.go
func resourceBackboneInterface() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				PfPortCircuitId: schemaStringRequired(PfPortCircuitIdDescription4),
				PfVlan:          schemaIntOptionalValidateDefault(PfVlanDescription2, validateVlan(), 0),
				PfSvlan:         schemaIntOptionalValidateDefault(PfSvlanDescription2, validateVlan(), 0),
				PfUntagged:      schemaBoolOptionalDefault(PfUntaggedDescription4, false),
			},
		},
	}
}

func resourceBackboneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	backboneVC := extractBack(d)
	resp, err := c.CreateBackbone(backboneVC)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		checkBackboneComplete(c, resp.VcCircuitID)
		d.SetId(resp.VcCircuitID)

		if labels, ok := d.GetOk(PfLabels); ok {
			diagnostics, created := createLabels(c, d.Id(), labels)
			if !created {
				return diagnostics
			}
		}
	}
	return diags
}

func resourceBackboneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetBackboneByVcCID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set(PfDescription, resp.Description)
		if resp.Mode == PfEpl {
			_ = d.Set(PfEpl, true)
		} else {
			_ = d.Set(PfEpl, false)
		}
		// Create a new schema set for the bandwidth attribute
		bandwidthSet := schema.NewSet(
			func(i interface{}) int { return 0 },
			[]interface{}{},
		)
		// Add the bandwidth values to the set
		if resp.Bandwidth.LonghaulType == PfDedicated {
			bandwidthSet.Add(mapStruct(resp.Bandwidth, PfAccountUuid, PfLonghaulType, PfSubscriptionTerm, PfSpeed))
		}
		// metro dedicated doesn't need longhaul_type
		if resp.Bandwidth.LonghaulType == PfEmptyString {
			bandwidthSet.Add(mapStruct(resp.Bandwidth, PfAccountUuid, PfSubscriptionTerm, PfSpeed))
		}
		if resp.Bandwidth.LonghaulType == PfUsage {
			bandwidthSet.Add(mapStruct(resp.Bandwidth, PfAccountUuid, PfLonghaulType))
		}
		if resp.Bandwidth.LonghaulType == PfHourly {
			bandwidthSet.Add(mapStruct(resp.Bandwidth, PfAccountUuid, PfLonghaulType, PfSpeed))
		}
		// Set the bandwidth attribute to the schema set
		_ = d.Set(PfBandwidth, bandwidthSet)

		if len(resp.Interfaces) == 2 {
			interfaceFields := stringsToMap(PfPortCircuitId, PfVlan, PfSvlan, PfUntagged)
			interfaceA := structToMap(resp.Interfaces[0], interfaceFields)
			interfaceZ := structToMap(resp.Interfaces[1], interfaceFields)
			_ = d.Set(PfInterfaceA, []interface{}{interfaceA})
			_ = d.Set(PfInterfaceZ, []interface{}{interfaceZ})
		}
		if _, ok := d.GetOk(PfRateLimitIn); ok {
			_ = d.Set(PfRateLimitIn, resp.RateLimitIn)
		}
		if _, ok := d.GetOk(PfRateLimitOut); ok {
			_ = d.Set(PfRateLimitOut, resp.RateLimitOut)
		}
		if _, ok := d.GetOk(PfFlexBandwidthId); ok {
			_ = d.Set(PfFlexBandwidthId, resp.AggregateCapacityID)
		} else {
			_ = d.Set(PfFlexBandwidthId, nil)
		}
		_ = d.Set(PfPoNumber, resp.PONumber)
	}

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

// used for Backbone VC
func resourceBackboneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	settings := extractServiceSettings(d)
	backboneVC := extractBack(d)

	if d.HasChange(PfBandwidth) {
		billing := packetfabric.BillingUpgrade{
			Speed:            backboneVC.Bandwidth.Speed,
			SubscriptionTerm: backboneVC.Bandwidth.SubscriptionTerm,
		}
		if _, err := c.ModifyBilling(d.Id(), billing); err != nil {
			return diag.FromErr(err)
		}
		checkBackboneComplete(c, d.Id())
	}

	if _, err := c.UpdateServiceSettings(d.Id(), settings); err != nil {
		return diag.FromErr(err)
	}
	checkBackboneComplete(c, d.Id())

	if d.HasChange(PfLabels) {
		labels := d.Get(PfLabels)
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}
	return diags
}

func checkBackboneComplete(c *packetfabric.PFClient, id string) {
	done := make(chan bool)
	defer close(done)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			if ok := c.IsBackboneComplete(id); ok {
				ticker.Stop()
				done <- true
			}
		}
	}()
	<-done
}

func resourceBackboneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCircuitID, ok := d.GetOk(PfId); ok {
		etlDiags, err2 := addETLWarning(c, vcCircuitID.(string))
		if err2 != nil {
			return diag.FromErr(err2)
		}
		diags = append(diags, etlDiags...)
		_, err := c.DeleteBackbone(vcCircuitID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(PfEmptyString)
		return diags
	}
	return diag.Errorf(MesssageVCCIdRequiredForDeletion)
}

func extractBack(d *schema.ResourceData) packetfabric.Backbone {
	backboneVC := packetfabric.Backbone{
		Description: d.Get(PfDescription).(string),
		Epl:         d.Get(PfEpl).(bool),
	}
	for _, interfA := range d.Get(PfInterfaceA).(*schema.Set).List() {
		backboneVC.Interfaces = append(backboneVC.Interfaces, extractBackboneInterface(interfA.(map[string]interface{})))
	}
	for _, interfZ := range d.Get(PfInterfaceZ).(*schema.Set).List() {
		backboneVC.Interfaces = append(backboneVC.Interfaces, extractBackboneInterface(interfZ.(map[string]interface{})))
	}
	for _, bw := range d.Get(PfBandwidth).(*schema.Set).List() {
		backboneVC.Bandwidth = extractBandwidth(bw.(map[string]interface{}))
	}
	if rateLimitIn, ok := d.GetOk(PfRateLimitIn); ok {
		backboneVC.RateLimitIn = rateLimitIn.(int)
	}
	if rateLimitOut, ok := d.GetOk(PfRateLimitOut); ok {
		backboneVC.RateLimitOut = rateLimitOut.(int)
	}
	if flexBandID, ok := d.GetOk(PfFlexBandwidthId); ok {
		backboneVC.FlexBandwidthID = flexBandID.(string)
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		backboneVC.PONumber = poNumber.(string)
	}
	return backboneVC
}

func extractServiceSettings(d *schema.ResourceData) packetfabric.ServiceSettingsUpdate {
	settUpdate := packetfabric.ServiceSettingsUpdate{}

	if rateLimitIn, ok := d.GetOk(PfRateLimitIn); ok {
		settUpdate.RateLimitIn = rateLimitIn.(int)
	}
	if rateLimitOut, ok := d.GetOk(PfRateLimitOut); ok {
		settUpdate.RateLimitOut = rateLimitOut.(int)
	}
	if description, ok := d.GetOk(PfDescription); ok {
		settUpdate.Description = description.(string)
	}
	if _, ok := d.GetOk(PfInterface); ok {
		for _, interf := range d.Get(PfInterface).(*schema.Set).List() {
			settUpdate.Interfaces = append(settUpdate.Interfaces, extractBackboneInterface(interf.(map[string]interface{})))
		}
	}
	if _, ok := d.GetOk(PfInterfaceA); ok {
		for _, interf := range d.Get(PfInterfaceA).(*schema.Set).List() {
			settUpdate.Interfaces = append(settUpdate.Interfaces, extractBackboneInterface(interf.(map[string]interface{})))
		}
	}
	if _, ok := d.GetOk(PfInterfaceZ); ok {
		for _, interf := range d.Get(PfInterfaceZ).(*schema.Set).List() {
			// Only include interface_z if it was modified
			if d.HasChange(PfInterfaceZ) {
				settUpdate.Interfaces = append(settUpdate.Interfaces, extractBackboneInterface(interf.(map[string]interface{})))
			}
		}
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		settUpdate.PONumber = poNumber.(string)
	}

	return settUpdate
}

func extractBandwidth(bw map[string]interface{}) packetfabric.Bandwidth {
	bandwidth := packetfabric.Bandwidth{}
	bandwidth.AccountUUID = bw[PfAccountUuid].(string)
	longhaulType := bw[PfLonghaulType]
	if longhaulType != nil {
		bandwidth.LonghaulType = longhaulType.(string)
	}
	if subsTerm := bw[PfSubscriptionTerm]; subsTerm != nil {
		bandwidth.SubscriptionTerm = subsTerm.(int)
	}
	if speed := bw[PfSpeed]; speed != nil {
		bandwidth.Speed = speed.(string)
	}
	return bandwidth
}

func extractBackboneInterface(interf map[string]interface{}) packetfabric.Interfaces {
	backboneInter := packetfabric.Interfaces{}
	if portCID := interf[PfPortCircuitId]; portCID != nil {
		backboneInter.PortCircuitID = portCID.(string)
	}
	if vlan := interf[PfVlan]; vlan != nil {
		backboneInter.Vlan = vlan.(int)
	}
	if untagged := interf[PfUntagged]; untagged != nil {
		backboneInter.Untagged = untagged.(bool)
	}
	if svlan := interf[PfSvlan]; svlan != nil {
		backboneInter.Svlan = svlan.(int)
	}
	return backboneInter
}
