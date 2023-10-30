package provider

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaces() *schema.Resource {
	return &schema.Resource{
		Timeouts:      schemaTimeouts(30, 10, 10, 30),
		CreateContext: resourceCreateInterface,
		ReadContext:   resourceReadInterface,
		UpdateContext: resourceUpdateInterface,
		DeleteContext: resourceDeleteInterface,
		Schema: map[string]*schema.Schema{
			PfId:               schemaStringComputedPlain(),
			PfAccountUuid:      schemaAccountUuid(PfAccountUuidDescription2),
			PfAutoneg:          schemaBoolOptional(PfAutonegDescription6),
			PfDescription:      schemaStringRequiredNotEmpty(PfPortDescription9),
			PfMedia:            schemaStringRequiredNewNotEmpty(PfMediaDescription5),
			PfNni:              schemaBoolOptionalNewDefault(PfNniDescription, false),
			PfPop:              schemaStringRequiredNewNotEmpty(PfPopDescription3),
			PfSpeed:            schemaStringRequiredNewNotEmpty(PfSpeedDescription8),
			PfSubscriptionTerm: schemaIntRequiredValidate(PfSubscriptionTermDescriptionB, validateSubscriptionTerm()),
			PfZone:             schemaStringRequiredNewNotEmpty(PfZoneDescription6),
			PfEnabled:          schemaBoolOptionalDefault(PfEnabledDescription2, true),
			PfPoNumber:         schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfLabels:           schemaStringSetOptional(PfLabelsDescription),
			PfEtl:              schemaFloatComputed(PfEtlDescription),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	_, ok := d.GetOk(PfAutoneg)
	if ok && d.Get(PfSpeed).(string) != Pf1Gbps {
		return diag.Errorf(MessageAutonegOnly1Gbps)
	}
	interf := extractInterface(d)
	resp, err := c.CreateInterface(interf)
	time.Sleep(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		enabled := d.Get(PfEnabled)
		if !enabled.(bool) {
			if toggleErr := _togglePortStatus(c, enabled.(bool), resp.PortCircuitID); toggleErr != nil {
				return diag.FromErr(toggleErr)
			}
		}
		autoneg := d.Get(PfAutoneg)
		// if autoneg = false and speed 1Gbps
		if !autoneg.(bool) && d.Get(PfSpeed).(string) == Pf1Gbps {
			if togglePortAutonegErr := _togglePortAutoneg(c, autoneg.(bool), resp.PortCircuitID); togglePortAutonegErr != nil {
				return diag.FromErr(togglePortAutonegErr)
			}
		}
		d.SetId(resp.PortCircuitID)

		if labels, ok := d.GetOk(PfLabels); ok {
			diagnostics, created := createLabels(c, d.Id(), labels)
			if !created {
				return diagnostics
			}
		}
	}
	return diags
}

func resourceReadInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetPortByCID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = setResourceDataKeys(d, resp, PfAccountUuid, PfDescription, PfMedia, PfNni, PfPop, PfSpeed, PfSubscriptionTerm, PfZone, PfPoNumber)

		if _, ok := d.GetOk(PfAutoneg); ok {
			_ = d.Set(PfAutoneg, resp.Autoneg)
		}

		if resp.Disabled {
			_ = d.Set(PfEnabled, false)
		} else {
			_ = d.Set(PfEnabled, true)
		}
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

func resourceUpdateInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	_, ok := d.GetOk(PfAutoneg)
	if ok && d.Get(PfSpeed).(string) != Pf1Gbps {
		return diag.Errorf(MessageAutonegOnly1Gbps)
	}
	_, err := _extractUpdateFn(ctx, d, m)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange(PfLabels) {
		c := m.(*packetfabric.PFClient)
		labels := d.Get(PfLabels)
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}
	return diags
}

func resourceDeleteInterface(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	etlDiags, err2 := addETLWarning(c, d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	diags = append(diags, etlDiags...)

	host := os.Getenv(PfPfeHost)
	testingInLab := strings.Contains(host, PfDevLab)

	if testingInLab {
		enabled := d.Get(PfEnabled)
		if enabled.(bool) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  MessageDevDisableToDelete,
			})

			if toggleErr := _togglePortStatus(c, false, d.Id()); toggleErr != nil {
				return diag.FromErr(toggleErr)
			}
		}
		// allow time for port to be disabled
		time.Sleep(time.Duration(120) * time.Second)
	}

	_, err := c.DeletePort(d.Id())
	time.Sleep(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(PfEmptyString)
	return diags
}

func extractInterface(d *schema.ResourceData) packetfabric.Interface {
	interf := packetfabric.Interface{
		AccountUUID:      d.Get(PfAccountUuid).(string),
		Description:      d.Get(PfDescription).(string),
		Media:            d.Get(PfMedia).(string),
		Nni:              d.Get(PfNni).(bool),
		Pop:              d.Get(PfPop).(string),
		Speed:            d.Get(PfSpeed).(string),
		SubscriptionTerm: d.Get(PfSubscriptionTerm).(int),
		Zone:             d.Get(PfZone).(string),
		PONumber:         d.Get(PfPoNumber).(string),
	}
	if autoneg, ok := d.GetOk(PfAutoneg); ok {
		interf.Autoneg = autoneg.(bool)
	}
	return interf
}

func _extractUpdateFn(ctx context.Context, d *schema.ResourceData, m interface{}) (resp *packetfabric.InterfaceReadResp, err error) {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	// Update if payload contains description and po_number
	if d.HasChanges([]string{PfPoNumber, PfDescription}...) {
		portUpdateData := packetfabric.PortUpdate{
			Description: d.Get(PfDescription).(string),
			PONumber:    d.Get(PfPoNumber).(string),
		}
		resp, err = c.UpdatePort(d.Id(), portUpdateData)
	}

	// Update autoneg only if speed == 1Gbps
	if d.HasChange(PfAutoneg) && d.Get(PfSpeed).(string) == Pf1Gbps {
		_, autonegChange := d.GetChange(PfAutoneg)
		err = _togglePortAutoneg(c, autonegChange.(bool), d.Id())
	}

	// Update port status
	if enabledHasChanged := d.HasChange(PfEnabled); enabledHasChanged {
		_, enableChange := d.GetChange(PfEnabled)
		err = _togglePortStatus(c, enableChange.(bool), d.Id())
	}

	// Update port term
	if d.HasChange(PfSubscriptionTerm) {
		if subTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
			billing := packetfabric.BillingUpgrade{
				SubscriptionTerm: subTerm.(int),
			}
			_, err = c.ModifyBilling(d.Id(), billing)
			_ = d.Set(PfSubscriptionTerm, subTerm.(int))
		}
	}
	return
}

func _togglePortStatus(c *packetfabric.PFClient, enabled bool, portCID string) (err error) {
	if enabled {
		_, err = c.EnablePort(portCID)
	} else {
		_, err = c.DisablePort(portCID)
	}
	return
}

func _togglePortAutoneg(c *packetfabric.PFClient, enabled bool, portCID string) (err error) {
	if enabled {
		_, err = c.EnablePortAutoneg(portCID)
	} else {
		_, err = c.DisablePortAutoneg(portCID)
	}
	return
}
