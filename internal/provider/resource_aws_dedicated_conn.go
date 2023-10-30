package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const cloudCidNotFoundDetailsMsg = MessagePleaseWaitAndRefresh

func resourceAwsReqDedicatedConn() *schema.Resource {
	return &schema.Resource{
		Timeouts:      schema10MinuteTimeouts(),
		CreateContext: resourceAwsReqDedicatedConnCreate,
		UpdateContext: resourceAwsReqDedicatedConnUpdate,
		ReadContext:   resourceAwsReqDedicatedConnRead,
		DeleteContext: resourceAwsServicesDelete,
		Schema: map[string]*schema.Schema{
			PfId:               schemaStringComputedPlain(),
			PfAwsRegion:        schemaStringRequiredNew(PfAwsRegionDescription3),
			PfAccountUuid:      schemaAccountUuid(PfAccountUuidDescription2),
			PfDescription:      schemaStringRequiredNew(PfConnectionDescription),
			PfZone:             schemaStringRequiredNew(PfZoneDescription),
			PfPop:              schemaStringRequiredNew(PfPopDescriptionD),
			PfSubscriptionTerm: schemaIntRequired(PfSubscriptionTermDescription4),
			PfServiceClass:     schemaStringRequired(PfServiceClassDescription6),
			PfAutoneg:          schemaBoolRequiredNew(PfAutonegDescription3),
			PfSpeed:            schemaStringRequiredNew(PfSpeedDescriptionF),
			PfShouldCreateLag:  schemaBoolOptionalNewDefault(PfShouldCreateLagDescription, true),
			PfLoa:              schemaStringOptionalNew(PfLoaDescription),
			PfPoNumber:         schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfLabels:           schemaStringSetOptional(PfLabelsDescription),
			PfEtl:              schemaFloatComputed(PfEtlDescription),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsReqDedicatedConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	dedicatedConn := extractAwsDedicatedConn(d)
	expectedResp, err := c.CreateDedicadedAWSConn(dedicatedConn)
	if err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			dedicatedConns, err := c.GetCurrentCustomersDedicated()
			if dedicatedConns != nil && err == nil && len(dedicatedConns) > 0 {
				for _, conn := range dedicatedConns {
					if expectedResp.UUID == conn.UUID && conn.State == "active" {
						expectedResp.CloudCircuitID = conn.CloudCircuitID
						ticker.Stop()
						createOk <- true
					}
				}
			}
		}
	}()
	<-createOk
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(expectedResp.CloudCircuitID)

	if labels, ok := d.GetOk("labels"); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}
	return diags
}

func resourceAwsReqDedicatedConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = setResourceDataKeys(d, resp, PfAccountUuid, PfDescription, PfPop, PfSubscriptionTerm, PfServiceClass, PfSpeed)
		_ = d.Set(PfAwsRegion, resp.Settings.AwsRegion)
	}
	resp2, err2 := c.GetPortByCID(d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	if resp2 != nil {
		_ = setResourceDataKeys(d, resp2, PfAutoneg, PfZone, PfPoNumber)
		if resp2.IsLag {
			_ = d.Set(PfShouldCreateLag, true)
		} else {
			_ = d.Set(PfShouldCreateLag, false)
		}
	}
	// unsetFields: loa

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

func resourceAwsReqDedicatedConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesDedicatedUpdate(ctx, d, m)
}

func resourceAwsServicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudSourceDelete(ctx, d, m, PfAwsServiceDelete)
}

func extractAwsDedicatedConn(d *schema.ResourceData) packetfabric.DedicatedAwsConn {
	dedicatedConn := packetfabric.DedicatedAwsConn{}
	if awsRegion, ok := d.GetOk(PfAwsRegion); ok {
		dedicatedConn.AwsRegion = awsRegion.(string)
	}
	if accountUUID, ok := d.GetOk(PfAccountUuid); ok {
		dedicatedConn.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk(PfDescription); ok {
		dedicatedConn.Description = description.(string)
	}
	if zone, ok := d.GetOk(PfZone); ok {
		dedicatedConn.Zone = zone.(string)
	}
	if pop, ok := d.GetOk(PfPop); ok {
		dedicatedConn.Pop = pop.(string)
	}
	if subTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
		dedicatedConn.SubscriptionTerm = subTerm.(int)
	}
	if serviceClass, ok := d.GetOk(PfServiceClass); ok {
		dedicatedConn.ServiceClass = serviceClass.(string)
	}
	if autoneg, ok := d.GetOk(PfAutoneg); ok {
		dedicatedConn.AutoNeg = autoneg.(bool)
	}
	if speed, ok := d.GetOk(PfSpeed); ok {
		dedicatedConn.Speed = speed.(string)
	}
	if shouldCreateLag, ok := d.GetOk(PfShouldCreateLag); ok {
		dedicatedConn.ShouldCreateLag = shouldCreateLag.(bool)
	}
	if loa, ok := d.GetOk(PfLoa); ok {
		dedicatedConn.Loa = loa.(string)
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		dedicatedConn.PONumber = poNumber.(string)
	}
	return dedicatedConn
}
