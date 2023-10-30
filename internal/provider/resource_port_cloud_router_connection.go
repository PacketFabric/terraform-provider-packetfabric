package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCustomerOwnedPortConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCustomerOwnedPortConnCreate,
		ReadContext:   resourceCustomerOwnedPortConnRead,
		UpdateContext: resourceCustomerOwnedPortConnUpdate,
		DeleteContext: resourceCustomerOwnedPortConnDelete,
		Timeouts:      schema10MinuteTimeouts(),
		Schema: map[string]*schema.Schema{
			PfId:                     schemaStringComputedPlain(),
			PfAccountUuid:            schemaAccountUuid(PfAccountUuidDescription2),
			PfCircuitId:              schemaStringRequiredNewNotEmpty(PfCircuitIdDescription),
			PfMaybeNat:               schemaBoolOptionalNew(PfMaybeNatDescription),
			PfMaybeDnat:              schemaBoolOptionalNew(PfMaybeDnatDescription),
			PfPortCircuitId:          schemaStringRequiredNewNotEmpty(PfPortCircuitIdDescription5),
			PfDescription:            schemaStringRequiredNotEmpty(PfConnectionDescription),
			PfVlan:                   schemaIntOptionalNewValidateDefault(PfVlanDescription2, validateVlan(), 0),
			PfUntagged:               schemaBoolOptionalNewDefault(PfUntaggedDescription, false),
			PfSpeed:                  schemaStringRequiredValidate(PfSpeedDescription2, validateSpeed()),
			PfIsPublic:               schemaBoolOptionalNewDefault(PfIsPublicDescription3, false),
			PfPublishedQuoteLineUuid: schemaStringOptionalNewValidate(PfPublishedQuoteLineUuidDescription2, validation.IsUUID),
			PfPoNumber:               schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfSubscriptionTerm:       schemaIntOptionalValidateDefault(PfSubscriptionTermDescription2, validateSubscriptionTerm(), 1),
			PfLabels:                 schemaStringSetOptional(PfLabelsDescription),
			PfEtl:                    schemaFloatComputed(PfEtlDescription),
		},
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceCustomerOwnedPortConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ownedPort := extractOwnedPortConn(d)
	if cID, ok := d.GetOk(PfCircuitId); ok {
		resp, err := c.AttachCustomerOwnedPortToCR(ownedPort, cID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		createOkCh := make(chan bool)
		defer close(createOkCh)
		fn := func() (*packetfabric.ServiceState, error) {
			return c.GetCloudConnectionStatus(cID.(string), resp.CloudCircuitID)
		}
		go c.CheckServiceStatus(createOkCh, fn)
		if !<-createOkCh {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.CloudCircuitID)

			if labels, ok := d.GetOk(PfLabels); ok {
				diagnostics, created := createLabels(c, d.Id(), labels)
				if !created {
					return diagnostics
				}
			}
		}
	}
	return diags
}

func resourceCustomerOwnedPortConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		_ = setResourceDataKeys(d, resp, PfAccountUuid, PfCircuitId, PfPortCircuitId, PfDescription, PfVlan, PfSpeed, PfPoNumber, PfSubscriptionTerm)
		if resp.CloudSettings.PublicIP != PfEmptyString {
			_ = d.Set(PfIsPublic, true)
		} else {
			_ = d.Set(PfIsPublic, false)
		}
		if resp.Vlan == 0 {
			_ = d.Set(PfUntagged, true)
		} else {
			_ = d.Set(PfUntagged, false)
		}
		// unsetFields: published_quote_line_uuid
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

func resourceCustomerOwnedPortConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceCustomerOwnedPortConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractOwnedPortConn(d *schema.ResourceData) packetfabric.CustomerOwnedPort {
	ownedPort := packetfabric.CustomerOwnedPort{}
	if accountUUID, ok := d.GetOk(PfAccountUuid); ok {
		ownedPort.AccountUUID = accountUUID.(string)
	}
	if maybeNat, ok := d.GetOk(PfMaybeNat); ok {
		ownedPort.MaybeNat = maybeNat.(bool)
	}
	if maybeDNat, ok := d.GetOk(PfMaybeDnat); ok {
		ownedPort.MaybeDNat = maybeDNat.(bool)
	}
	if portCircuitID, ok := d.GetOk(PfPortCircuitId); ok {
		ownedPort.PortCircuitID = portCircuitID.(string)
	}
	if description, ok := d.GetOk(PfDescription); ok {
		ownedPort.Description = description.(string)
	}
	if untagged, ok := d.GetOk(PfUntagged); ok {
		ownedPort.Untagged = untagged.(bool)
	}
	if vlan, ok := d.GetOk(PfVlan); ok {
		ownedPort.Vlan = vlan.(int)
	}
	if speed, ok := d.GetOk(PfSpeed); ok {
		ownedPort.Speed = speed.(string)
	}
	if isPublic, ok := d.GetOk(PfIsPublic); ok {
		ownedPort.IsPublic = isPublic.(bool)
	}
	if publishedQuote, ok := d.GetOk(PfPublishedQuoteLineUuid); ok {
		ownedPort.PublishedQuoteLineUUID = publishedQuote.(string)
	}
	if poNumber, ok := d.GetOk(PfPoNumber); ok {
		ownedPort.PONumber = poNumber.(string)
	}
	if subscriptionTerm, ok := d.GetOk(PfSubscriptionTerm); ok {
		ownedPort.SubscriptionTerm = subscriptionTerm.(int)
	}
	return ownedPort
}
