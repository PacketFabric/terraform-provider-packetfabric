package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudRouter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudRouterCreate,
		ReadContext:   resourceCloudRouterRead,
		UpdateContext: resourceCloudRouterUpdate,
		DeleteContext: resourceCloudRouterDelete,
		Timeouts:      schema10MinuteTimeouts(),
		Schema: map[string]*schema.Schema{
			PfId:               schemaStringComputedPlain(),
			PfAsn:              schemaIntOptionalNewDefault(PfAsnDescription3, 4556),
			PfName:             schemaStringRequiredNotEmpty(PfNameDescription8),
			PfAccountUuid:      schemaAccountUuid(PfAccountUuidDescription2),
			PfRegions:          schemaStringListRequiredNotEmpty(PfRegionsDescription2),
			PfCapacity:         schemaStringRequiredNotEmpty(PfCapacityDescription2),
			PfPoNumber:         schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfLabels:           schemaStringSetOptional(PfLabelsDescription),
			PfEtl:              schemaFloatComputed(PfEtlDescription),
			PfSubscriptionTerm: schemaIntOptionalValidateDefault(PfSubscriptionTermDescription2, validateSubscriptionTerm(), 1),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudRouterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	router := extractCloudRouter(d)

	resp, err := c.CreateCloudRouter(router)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = setResourceDataKeys(d, resp, PfAsn, PfName, PfCapacity, PfSubscriptionTerm)
		d.SetId(resp.CircuitID)

		if labels, ok := d.GetOk(PfLabels); ok {
			diagnostics, created := createLabels(c, d.Id(), labels)
			if !created {
				return diagnostics
			}
		}
	}
	return diags
}

func resourceCloudRouterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID := d.Get(PfId).(string)
	resp, err := c.ReadCloudRouter(cID)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = setResourceDataKeys(d, resp, PfAccountUuid, PfAsn, PfName, PfCapacity, PfSubscriptionTerm, PfPoNumber)
		var regions []string
		for _, region := range resp.Regions {
			regions = append(regions, region.Code)
		}
		_ = d.Set(PfRegions, regions)
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

func resourceCloudRouterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	routerUpdate := packetfabric.CloudRouterUpdate{
		Name:     d.Get(PfName).(string),
		Regions:  extractRegions(d),
		Capacity: d.Get(PfCapacity).(string),
	}

	cID := d.Get(PfId).(string)

	resp, err := c.UpdateCloudRouter(routerUpdate, cID)
	if err != nil {
		return diag.FromErr(err)
	}

	_ = setResourceDataKeys(d, resp, PfName, PfCapacity, PfSubscriptionTerm)
	if d.HasChange(PfPoNumber) {
		_ = d.Set(PfPoNumber, resp.PONumber)
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

func resourceCloudRouterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics

	cID := d.Get(PfId).(string)

	etlDiags, err := addETLWarning(c, cID)
	if err != nil {
		return diag.FromErr(err)
	}
	diags = append(diags, etlDiags...)

	_, err2 := c.DeleteCloudRouter(cID)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(PfEmptyString)
	return diags

}

func extractCloudRouter(d *schema.ResourceData) packetfabric.CloudRouter {
	router := packetfabric.CloudRouter{}
	if asn, ok := d.GetOk(PfAsn); ok {
		router.Asn = asn.(int)
	}
	if name, ok := d.GetOk(PfName); ok {
		router.Name = name.(string)
	}
	if accountUUID, ok := d.GetOk(PfAccountUuid); ok {
		router.AccountUUID = accountUUID.(string)
	}
	if capacity, ok := d.GetOk(PfCapacity); ok {
		router.Capacity = capacity.(string)
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		router.PONumber = poNumber.(string)
	}
	if subscriptionTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
		router.SubscriptionTerm = subscriptionTerm.(int)
	}
	router.Regions = extractRegions(d)
	return router
}

func extractRegions(d *schema.ResourceData) []string {
	if regions, ok := d.GetOk(PfRegions); ok {
		regs := make([]string, 0)
		for _, reg := range regions.([]interface{}) {
			regs = append(regs, reg.(string))
		}
		return regs
	}
	return make([]string, 0)
}
