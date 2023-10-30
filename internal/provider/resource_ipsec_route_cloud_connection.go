package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIPSecCloudRouteConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIPSecCloudRouteConnCreate,
		ReadContext:   resourceIPSecCloudRouteConnRead,
		UpdateContext: resourceIPSecCloudRouteConnUpdate,
		DeleteContext: resourceIPSecCloudRouteConnDelete,
		Timeouts:      schema10MinuteTimeouts(),
		Schema: map[string]*schema.Schema{
			PfId:                         schemaStringComputedPlain(),
			PfCircuitId:                  schemaStringRequiredNewNotEmpty(PfCircuitIdDescription),
			PfDescription:                schemaStringRequiredNotEmpty(PfConnectionDescription),
			PfAccountUuid:                schemaAccountUuid(PfAccountUuidDescription2),
			PfPop:                        schemaStringRequiredNewNotEmpty(PfPopDescription5),
			PfSpeed:                      schemaStringRequiredValidate(PfSpeedDescription5, validateIpSecSpeed()),
			PfIkeVersion:                 schemaIntRequiredValidate(PfIkeVersionDescription2, validateIkeVersion()),
			PfPhase1AuthenticationMethod: schemaStringRequiredNotEmpty(PfPhase1AuthenticationMethodDescription2),
			PfPhase1Group:                schemaStringRequiredValidate(PfPhase1GroupDescription2, validatePhase1Group()),
			PfPhase1EncryptionAlgo:       schemaStringRequiredValidate(PfPhase1EncryptionAlgoDescription2, validatePhase1EncryptionAlgo()),
			PfPhase1AuthenticationAlgo:   schemaStringRequiredValidate(PfPhase1AuthenticationAlgoDescription2, validatePhase1AuthenticationAlgo()),
			PfPhase1Lifetime:             schemaIntRequiredValidate(PfPhase1LifetimeDescription2, validatePhase1Lifetime()),
			PfPhase2PfsGroup:             schemaStringRequiredValidate(PfPhase2PfsGroupDescription2, validatePhase2PfsGroup()),
			PfPhase2EncryptionAlgo:       schemaStringRequiredValidate(PfPhase2EncryptionAlgoDescription2, validatePhase2EncryptionAlgo()),
			PfPhase2AuthenticationAlgo:   schemaStringOptionalValidate(PfPhase2AuthenticationAlgoDescription2, validatePhase2AuthenticationAlgo()),
			PfPhase2Lifetime:             schemaIntRequiredValidate(PfPhase2LifetimeDescription2, validatePhase2Lifetime()),
			PfGatewayAddress:             schemaStringRequiredValidate(PfGatewayAddressDescription, validation.IsIPv4Address),
			PfSharedKey:                  schemaStringRequiredNotEmpty(PfSharedKeyDescription),
			PfPublishedQuoteLineUuid:     schemaStringOptionalNewValidate(PfPublishedQuoteLineUuidDescription2, validation.IsUUID),
			PfPoNumber:                   schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfSubscriptionTerm:           schemaIntOptionalValidateDefault(PfSubscriptionTermDescription2, validateSubscriptionTerm(), 1),
			PfLabels:                     schemaStringSetOptional(PfLabelsDescription),
			PfEtl:                        schemaFloatComputed(PfEtlDescription),
		},
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceIPSecCloudRouteConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	ipSecRouter, err := extractIPSecRouteConn(d)
	if err != nil {
		return diag.FromErr(err)
	}
	if cid, ok := d.GetOk(PfCircuitId); ok {
		resp, err := c.CreateIPSecCloudRouerConnection(ipSecRouter, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		createOkCh := make(chan bool)
		defer close(createOkCh)
		fn := func() (*packetfabric.ServiceState, error) {
			return c.GetCloudConnectionStatus(cid.(string), resp.VcCircuitID)
		}
		go c.CheckIPSecStatus(createOkCh, fn)
		if !<-createOkCh {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.VcCircuitID)

			if labels, ok := d.GetOk(PfLabels); ok {
				diagnostics, created := createLabels(c, d.Id(), labels)
				if !created {
					return diagnostics
				}
			}
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  MessageMissingCircuitId,
			Detail:   MessageMissingCircuitIdDetail,
		})
	}
	return diags
}

func resourceIPSecCloudRouteConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk(PfCircuitId); ok {
		cloudConnCID := d.Get(PfId)
		resp, err := c.ReadCloudRouterConnection(cid.(string), cloudConnCID.(string))
		if err != nil {
			diags = diag.FromErr(err)
			return diags
		}
		_ = setResourceDataKeys(d, resp, PfAccountUuid, PfCircuitId, PfDescription, PfPop, PfSpeed, PfSubscriptionTerm, PfPoNumber)

		resp2, err2 := c.GetIpsecSpecificConn(cloudConnCID.(string))
		if err2 != nil {
			diags = diag.FromErr(err2)
		}
		errD := setResourceDataKeys(d, resp2, PfIkeVersion, PfPhase1AuthenticationMethod, PfPhase1Group, PfPhase1EncryptionAlgo, PfPhase1AuthenticationAlgo, PfPhase1Lifetime, PfPhase2PfsGroup, PfPhase2EncryptionAlgo, PfPhase2AuthenticationAlgo, PfPhase2Lifetime, PfGatewayAddress, PfSharedKey, PfPoNumber, PfSubscriptionTerm)
		if nil != errD {
			return diag.FromErr(errD)
		}
		_ = d.Set(PfSharedKey, resp2.PreSharedKey)
		_ = d.Set(PfGatewayAddress, resp2.CustomerGatewayAddress)

		// unsetFields: published_quote_line_uuid
	}
	if _, ok := d.GetOk(PfLabels); ok {
		labels, err3 := getLabels(c, d.Id())
		if err3 != nil {
			return diag.FromErr(err3)
		}
		_ = d.Set(PfLabels, labels)
	}

	etl, err4 := c.GetEarlyTerminationLiability(d.Id())
	if err4 != nil {
		return diag.FromErr(err4)
	}
	if etl > 0 {
		_ = d.Set(PfEtl, etl)
	}

	return diags
}

func resourceIPSecCloudRouteConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	ipsecUpdated := extractIPSecUpdate(d)
	_, err := c.UpdateIPSecConnection(d.Id(), ipsecUpdated)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges([]string{PfDescription, PfSpeed, PfPoNumber, PfLabels}...) {
		return resourceCloudRouterConnUpdate(ctx, d, m)
	}

	return diags
}

func resourceIPSecCloudRouteConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractIPSecRouteConn(d *schema.ResourceData) (packetfabric.IPSecRouterConn, error) {
	iPSecRouter := packetfabric.IPSecRouterConn{}
	if desc, ok := d.GetOk(PfDescription); ok {
		iPSecRouter.Description = desc.(string)
	}
	if accountUUID, ok := d.GetOk(PfAccountUuid); ok {
		iPSecRouter.AccountUUID = accountUUID.(string)
	}
	if pop, ok := d.GetOk(PfPop); ok {
		iPSecRouter.Pop = pop.(string)
	}
	if speed, ok := d.GetOk(PfSpeed); ok {
		iPSecRouter.Speed = speed.(string)
	}
	if ikeVersion, ok := d.GetOk(PfIkeVersion); ok {
		iPSecRouter.IkeVersion = ikeVersion.(int)
	}
	if phaseOneAuthMethod, ok := d.GetOk(PfPhase1AuthenticationMethod); ok {
		iPSecRouter.Phase1AuthenticationMethod = phaseOneAuthMethod.(string)
	}
	if phaseOneGroup, ok := d.GetOk(PfPhase1Group); ok {
		iPSecRouter.Phase1Group = phaseOneGroup.(string)
	}
	if phaseOneEncryptionAlgo, ok := d.GetOk(PfPhase1EncryptionAlgo); ok {
		iPSecRouter.Phase1EncryptionAlgo = phaseOneEncryptionAlgo.(string)
	}
	if phaseOneAuthAlgo, ok := d.GetOk(PfPhase1AuthenticationAlgo); ok {
		iPSecRouter.Phase1AuthenticationAlgo = phaseOneAuthAlgo.(string)
	}
	if phaseOneLifetime, ok := d.GetOk(PfPhase1Lifetime); ok {
		iPSecRouter.Phase1Lifetime = phaseOneLifetime.(int)
	}
	if phaseTwoPfsGroup, ok := d.GetOk(PfPhase2PfsGroup); ok {
		iPSecRouter.Phase2PfsGroup = phaseTwoPfsGroup.(string)
	}
	if phaseTwoEncryptionAlgo, ok := d.GetOk(PfPhase2EncryptionAlgo); ok {
		iPSecRouter.Phase2EncryptionAlgo = phaseTwoEncryptionAlgo.(string)
	}
	if phaseTwoAuthAlgo, ok := d.GetOk(PfPhase2AuthenticationAlgo); ok {
		iPSecRouter.Phase2AuthenticationAlgo = phaseTwoAuthAlgo.(string)
	}
	if phaseTwoLifetime, ok := d.GetOk(PfPhase2Lifetime); ok {
		iPSecRouter.Phase2Lifetime = phaseTwoLifetime.(int)
	}
	if gatewayAddress, ok := d.GetOk(PfGatewayAddress); ok {
		iPSecRouter.GatewayAddress = gatewayAddress.(string)
	}
	if sharedKey, ok := d.GetOk(PfSharedKey); ok {
		iPSecRouter.SharedKey = sharedKey.(string)
	}
	if publishedQuote, ok := d.GetOk(PfPublishedQuoteLineUuid); ok {
		iPSecRouter.PublishedQuoteLineUUID = publishedQuote.(string)
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		iPSecRouter.PONumber = poNumber.(string)
	}
	if subscriptionTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
		iPSecRouter.SubscriptionTerm = subscriptionTerm.(int)
	}
	return iPSecRouter, nil
}

func extractIPSecUpdate(d *schema.ResourceData) packetfabric.IPSecConnUpdate {
	ipsec := packetfabric.IPSecConnUpdate{}
	if custGatewayAdd, ok := d.GetOk(PfGatewayAddress); ok {
		ipsec.CustomerGatewayAddress = custGatewayAdd.(string)
	}
	if ikeVersion, ok := d.GetOk(PfIkeVersion); ok {
		ipsec.IkeVersion = ikeVersion.(int)
	}
	if phaseOneAuthMethod, ok := d.GetOk(PfPhase1AuthenticationMethod); ok {
		ipsec.Phase1AuthenticationMethod = phaseOneAuthMethod.(string)
	}
	if phaseOneGroup, ok := d.GetOk(PfPhase1Group); ok {
		ipsec.Phase1Group = phaseOneGroup.(string)
	}
	if phaseOneEncAlgo, ok := d.GetOk(PfPhase1EncryptionAlgo); ok {
		ipsec.Phase1EncryptionAlgo = phaseOneEncAlgo.(string)
	}
	if phaseOneAuthAlgo, ok := d.GetOk(PfPhase1AuthenticationAlgo); ok {
		ipsec.Phase1AuthenticationAlgo = phaseOneAuthAlgo.(string)
	}
	if phaseOneLifetime, ok := d.GetOk(PfPhase1Lifetime); ok {
		ipsec.Phase1Lifetime = phaseOneLifetime.(int)
	}
	if phaseTwoPfsGroup, ok := d.GetOk(PfPhase2PfsGroup); ok {
		ipsec.Phase2PfsGroup = phaseTwoPfsGroup.(string)
	}
	if phaseTwoEncryptationAlgo, ok := d.GetOk(PfPhase2EncryptionAlgo); ok {
		ipsec.Phase2EncryptionAlgo = phaseTwoEncryptationAlgo.(string)
	}
	if phaseTwoAuthAlgo, ok := d.GetOk(PfPhase2AuthenticationAlgo); ok {
		ipsec.Phase2AuthenticationAlgo = phaseTwoAuthAlgo.(string)
	}
	if phaseTwoLifetime, ok := d.GetOk(PfPhase2Lifetime); ok {
		ipsec.Phase2Lifetime = phaseTwoLifetime.(int)
	}
	if preSharedKey, ok := d.GetOk(PfSharedKey); ok {
		ipsec.PreSharedKey = preSharedKey.(string)
	}
	return ipsec
}
