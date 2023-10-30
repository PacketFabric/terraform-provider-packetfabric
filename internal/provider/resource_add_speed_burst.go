package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAddSpeedBurst() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAddSpeedBurstCreate,
		ReadContext:   resourceAddSpeedBurstRead,
		DeleteContext: resourceAddSpeedBurstDelete,
		Timeouts:      schemaTimeoutsCRD(10, 10, 10),
		Schema: map[string]*schema.Schema{
			PfId:          schemaStringComputedPlain(),
			PfVcCircuitId: schemaStringRequiredNewNotEmpty(PfVcCircuitIdDescription2),
			PfSpeed:       schemaStringRequiredNewNotEmpty(PfSpeedDescriptionD),
		},
	}
}

func resourceAddSpeedBurstCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCID, ok := d.GetOk(PfVcCircuitId); ok {
		if speed, ok := d.GetOk(PfSpeed); ok {
			if _, err := c.AddSpeedBurstToCircuit(vcCID.(string), speed.(string)); err != nil {
				return diag.FromErr(err)
			}
			createOk := make(chan bool)
			defer close(createOk)
			ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
			go func() {
				for range ticker.C {
					if ok := c.IsBackboneComplete(vcCID.(string)); ok {
						ticker.Stop()
						createOk <- true
					}
				}
			}()
			<-createOk
			d.SetId(vcCID.(string))
		}
	}
	return diags
}

func resourceAddSpeedBurstRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCID, ok := d.GetOk(PfVcCircuitId); ok {
		if _, err := c.GetBackboneByVcCID(vcCID.(string)); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAddSpeedBurstDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if _, err := c.DeleteSpeedBurst(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			if ok := c.IsBackboneComplete(d.Id()); ok {
				ticker.Stop()
				createOk <- true
			}
		}
	}()
	<-createOk
	d.SetId(PfEmptyString)
	return diags
}
