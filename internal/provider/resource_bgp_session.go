package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceBgpSession() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBgpSessionCreate,
		ReadContext:   resourceBgpSessionRead,
		UpdateContext: resourceBgpSessionUpdate,
		DeleteContext: resourceBgpSessionDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Circuit ID of the target cloud router.\n\t\tExample: \"PF-L3-CUST-2\"",
			},
			"connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The circuit ID of the connection to update.\n\t\tExample: \"PF-AE-1234\"",
			},
			"md5": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The MD5 value of the authenticated BGP sessions.",
			},
			"l3_address": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The L3 Address of this instance. Not used for Azure connections.",
			},
			"primary_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Currently for Azure use only, provide this as the primary subnet when creating an Azure cloud router connection.",
			},
			"secondary_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Currently for Azure use only, provide this as the secondary subnet when creating an Azure cloud router connection.",
			},
			"address_family": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Whether this instance is IPv4 or IPv6.\n\t\tEnum: \"v4\" \"v6\"",
			},
			"remote_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cloud-side address of the instance.",
			},
			"remote_asn": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The cloud-side ASN of the instance.",
			},
			"multihop_ttl": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The TTL of this session.\n\t\tDefaults to 1.",
			},
			"local_preference": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The preference for this instance.",
			},
			"med": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Multi-Exit Discriminator of this instance.",
			},
			"community": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The BGP community for this instance.",
			},
			"as_prepend": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The BGP prepend value for this instance.",
			},
			"orlonger": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to use exact match or longer for all prefixes.",
			},
			"bfd_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Minimum interval, in microseconds, for transmitting BFD Control packets.\n\t\tAvailable range is 3 through 30000.",
			},
			"bfd_multiplier": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of BFD Control packets not received by a neighbor that causes the session to be declared down.\n\t\tAvailable range is 2 through 16.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether this BGP session is disabled.\n\t\tDefault \"false\"",
			},
			"pre_nat_sources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "The source IP address + mask of the host before NAT translation.\n\t\tExample: 10.0.0.0/24",
				},
			},
			"pool_prefixes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "The source IP address + mask of the NAT pool prefix.\n\t\tExample: 10.0.0.0/32",
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceBgpSessionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}
	connCID, ok := d.GetOk("connection_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Cloud Router Connection ID"))
	}
	var diags diag.Diagnostics
	session := extractBgpSession(d)
	resp, err := c.CreateBgpSession(session, cID.(string), connCID.(string))
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.BgpSettingsUUID)
	return diags
}

func resourceBgpSessionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	const warningSummary = "Retrieve BGP Session after create"
	var diags diag.Diagnostics
	var cID, connCID, bgpSettingsUUID string
	if circuitID, ok := d.GetOk("circuit_id"); !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  warningSummary,
			Detail:   "Cloud not extract Circuit ID from Resource Data",
		})
	} else {
		cID = circuitID.(string)
	}
	if connectionCID, ok := d.GetOk("connection_id"); !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  warningSummary,
			Detail:   "Cloud not extract Connection Circuit ID from Resource Data",
		})
	} else {
		connCID = connectionCID.(string)
	}
	if settingsUUID, ok := d.GetOk("id"); !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  warningSummary,
			Detail:   "Cloud not extract BGP Settings UUID from Resource Data",
		})
	} else {
		bgpSettingsUUID = settingsUUID.(string)
	}
	if diags != nil || len(diags) > 0 {
		return diags
	}
	_, err := c.GetBgpSessionBy(cID, connCID, bgpSettingsUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  warningSummary,
			Detail:   "Cloud not retrieve BGP session after create",
		})
		return diags
	}
	return diags
}

func resourceBgpSessionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}
	connCID, ok := d.GetOk("connection_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Cloud Router Connection ID"))
	}
	session := extractBgpSession(d)
	_, _, err := c.UpdateBgpSession(session, cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceBgpSessionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	bgpSettingsUUID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid BGP Settings UUID")
	}
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid Circuit ID")
	}
	connCID, ok := d.GetOk("connection_id")
	if !ok {
		return diag.Errorf("please provide a valid Cloud Router Connection ID")
	}
	session, err := c.GetBgpSessionBy(cID.(string), connCID.(string), bgpSettingsUUID.(string))
	if err != nil {
		return diag.Errorf("could not find BGP session associated with the provided Cloud Router ID: %v", err)
	}
	sessionToDisable := session.BuildNewBgpSessionInstance()
	sessionPrefixes, err := c.ReadBgpSessionPrefixes(bgpSettingsUUID.(string))
	if err != nil || len(sessionPrefixes) <= 0 {
		resp, err := c.DeleteBgpSession(cID.(string), connCID.(string), bgpSettingsUUID.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "BGP Settings to be deleted",
				Detail: fmt.Sprintf("BGP with Settings UUID (%s) "+
					"might be associated with an active Cloud Router Connection "+
					"and will be deleted together with current Cloud Router Connection.", bgpSettingsUUID.(string)),
			})
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "BGP Settings deleted",
				Detail:   resp.Message,
			})
		}
		d.SetId("")
		return diags
	}
	if l3Address, ok := d.GetOk("l3_address"); !ok {
		return diag.Errorf("please provide a valid l3_address")
	} else {
		sessionToDisable.L3Address = l3Address.(string)
	}
	sessionToDisable.BgpSettingsUUID = bgpSettingsUUID.(string)
	sessionToDisable.Disabled = true
	sessionToDisable.Prefixes = make([]packetfabric.BgpSessionResponse, 0)
	for _, prefix := range sessionPrefixes {
		sessionToDisable.Prefixes = append(sessionToDisable.Prefixes, packetfabric.BgpSessionResponse{
			BgpPrefixUUID: prefix.BgpPrefixUUID,
			Prefix:        prefix.Prefix,
			Type:          prefix.Type,
			Order:         prefix.Order,
		})
	}
	err = c.DisableBgpSession(sessionToDisable, cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func extractBgpSession(d *schema.ResourceData) packetfabric.BgpSession {
	bgpSession := packetfabric.BgpSession{}
	if l3Address, ok := d.GetOk("l3_address"); ok {
		bgpSession.L3Address = l3Address.(string)
	}
	if addressFamily, ok := d.GetOk("address_family"); ok {
		bgpSession.AddressFamily = addressFamily.(string)
	}
	if remoteAddress, ok := d.GetOk("remote_address"); ok {
		bgpSession.RemoteAddress = remoteAddress.(string)
	}
	if remoteAsn, ok := d.GetOk("remote_asn"); ok {
		bgpSession.RemoteAsn = remoteAsn.(int)
	}
	if multihopTTL, ok := d.GetOk("multihop_ttl"); ok {
		bgpSession.MultihopTTL = multihopTTL.(int)
	}
	if localPreference, ok := d.GetOk("local_preference"); ok {
		bgpSession.LocalPreference = localPreference.(int)
	}
	if med, ok := d.GetOk("med"); ok {
		bgpSession.Med = med.(int)
	}
	if community, ok := d.GetOk("community"); ok {
		bgpSession.Community = community.(int)
	}
	if asPrepend, ok := d.GetOk("as_prepend"); ok {
		bgpSession.AsPrepend = asPrepend.(int)
	}
	if orlonger, ok := d.GetOk("orlonger"); ok {
		bgpSession.Orlonger = orlonger.(bool)
	}
	if bfdInterval, ok := d.GetOk("bfd_interval"); ok {
		bgpSession.BfdInterval = bfdInterval.(int)
	}
	if bfdMultiplier, ok := d.GetOk("bfd_multiplier"); ok {
		bgpSession.BfdMultiplier = bfdMultiplier.(int)
	}
	if md5, ok := d.GetOk("md5"); ok {
		bgpSession.Md5 = md5.(string)
	}
	return bgpSession
}
