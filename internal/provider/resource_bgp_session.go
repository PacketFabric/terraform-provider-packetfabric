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
				Description: "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The circuit ID of the connection associated with the BGP session. This starts with \"PF-L3-CON-\".",
			},
			"md5": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The MD5 value of the authenticated BGP sessions. Required for AWS.",
			},
			"l3_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The L3 address of this instance. Not used for Azure connections. Required for all other CSP.",
			},
			"primary_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Currently for Azure use only. Provide this as the primary subnet when creating an Azure cloud router connection.",
			},
			"secondary_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Currently for Azure use only. Provide this as the secondary subnet when creating an Azure cloud router connection.",
			},
			"address_family": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Whether this instance is IPv4 or IPv6. At this time, only IPv4 is supported.\n\n\tEnum: \"v4\" \"v6\"",
			},
			"remote_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cloud-side router peer IP. Not used for Azure connections. Required for all other CSP.",
			},
			"remote_asn": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The cloud-side ASN.",
			},
			"multihop_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The TTL of this session. The default is `1`. For Google Cloud connections, see [the PacketFabric doc](https://docs.packetfabric.com/cr/bgp/bgp_google/#ttl).",
			},
			"local_preference": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The local preference for this instance. When the same route is received in multiple locations, those with a higher local preference value are preferred by the cloud router. Deprecated.",
			},
			"med": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Multi-Exit Discriminator of this instance. When the same route is advertised in multiple locations, those with a lower MED are preferred by the peer AS. Deprecated.",
			},
			"community": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The BGP community for this instance. Deprecated.",
			},
			"as_prepend": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The BGP prepend value for this instance. Deprecated.",
			},
			"orlonger": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to use exact match or longer for all prefixes.",
			},
			"bfd_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "If you are using BFD, this is the interval (in milliseconds) at which to send test packets to peers.\n\n\tAvailable range is 3 through 30000.",
			},
			"bfd_multiplier": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "If you are using BFD, this is the number of consecutive packets that can be lost before BFD considers a peer down and shuts down BGP.\n\n\tAvailable range is 2 through 16.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether this BGP session is disabled. Default is false.",
			},
			"pre_nat_sources": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "If using NAT, this is the prefixes from the cloud that you want to associate with the NAT pool.\n\n\tExample: 10.0.0.0/24",
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "IP prefix using CIDR format.",
				},
			},
			"pool_prefixes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "If using NAT, all prefixes that are NATed on this connection will be translated to the pool prefix address.\n\n\tExample: 10.0.0.0/32",
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "IP prefix using CIDR format.",
				},
			},
			"prefixes": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The list of BGP prefixes",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The actual IP Prefix of this instance.",
						},
						"match_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger", "longer"}, true),
							Description:  "The match type of this prefix.",
						},
						"as_prepend": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The BGP prepend value of this prefix.",
						},
						"med": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The MED of this prefix.",
						},
						"local_preference": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The local_preference of this prefix.",
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"in", "out"}, true),
							Description:  "Whether this prefix is in or out.",
						},
						"order": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The order of this prefix against the others.",
						},
					},
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
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Can't find the BGP Settings.",
			Detail: fmt.Sprintf("BGP with Settings UUID (%s) "+
				"is associated with a Cloud Router Connection "+
				"and will be (or has been) deleted together with the Cloud Router Connection.", bgpSettingsUUID.(string)),
		})
	}
	sessionToDisable := session.BuildNewBgpSessionInstance()
	sessionPrefixes, err := c.ReadBgpSessionPrefixes(bgpSettingsUUID.(string))
	if err != nil || len(sessionPrefixes) <= 0 {
		resp, err := c.DeleteBgpSession(cID.(string), connCID.(string), bgpSettingsUUID.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Can't find the BGP Settings.",
				Detail: fmt.Sprintf("BGP with Settings UUID (%s) "+
					"is associated with a Cloud Router Connection "+
					"and will be (or has been) deleted together with the Cloud Router Connection.", bgpSettingsUUID.(string)),
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

	l3Address, ok := d.GetOk("l3_address")
	if !ok {
		return diag.Errorf("please provide a valid l3_address")
	}
	sessionToDisable.L3Address = l3Address.(string)
	sessionToDisable.BgpSettingsUUID = bgpSettingsUUID.(string)
	sessionToDisable.Disabled = true
	sessionToDisable.Prefixes = make([]packetfabric.BgpPrefix, 0)
	for _, prefix := range sessionPrefixes {
		sessionToDisable.Prefixes = append(sessionToDisable.Prefixes, packetfabric.BgpPrefix{
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
	if primarySubnet, ok := d.GetOk("primary_subnet"); ok {
		bgpSession.PrimarySubnet = primarySubnet.(string)
	}
	if secondarySubnet, ok := d.GetOk("secondary_subnet"); ok {
		bgpSession.SecondarySubnet = secondarySubnet.(string)
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
	bgpSession.Prefixes = extractConnBgpSessionPrefixes(d)
	return bgpSession
}

func extractConnBgpSessionPrefixes(d *schema.ResourceData) []packetfabric.BgpPrefix {
	if prefixes, ok := d.GetOk("prefixes"); ok {
		sessionPrefixes := make([]packetfabric.BgpPrefix, 0)
		for _, pref := range prefixes.(*schema.Set).List() {
			sessionPrefixes = append(sessionPrefixes, packetfabric.BgpPrefix{
				Prefix:          pref.(map[string]interface{})["prefix"].(string),
				MatchType:       pref.(map[string]interface{})["match_type"].(string),
				AsPrepend:       pref.(map[string]interface{})["as_prepend"].(int),
				Med:             pref.(map[string]interface{})["med"].(int),
				LocalPreference: pref.(map[string]interface{})["local_preference"].(int),
				Type:            pref.(map[string]interface{})["type"].(string),
				Order:           pref.(map[string]interface{})["order"].(int),
			})
		}
		return sessionPrefixes
	}
	return make([]packetfabric.BgpPrefix, 0)
}
