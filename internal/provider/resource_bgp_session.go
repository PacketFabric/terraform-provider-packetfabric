package provider

import (
	"context"
	"errors"

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
				ForceNew:    true,
				Description: "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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
				Description: "Currently for Azure use only. Provide this as the primary subnet when creating the primary Azure cloud router connection.",
			},
			"secondary_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Currently for Azure use only. Provide this as the secondary subnet when creating the secondary Azure cloud router connection.",
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
				Default:     1,
				Description: "The TTL of this session. The default is `1`. For Google Cloud connections, see [the PacketFabric doc](https://docs.packetfabric.com/cr/bgp/bgp_google/#ttl). ",
			},
			"local_preference": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The local preference for this instance. When the same route is received in multiple locations, those with a higher local preference value are preferred by the cloud router. It is used when type = in.",
			},
			"med": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Multi-Exit Discriminator of this instance. When the same route is advertised in multiple locations, those with a lower MED are preferred by the peer AS. It is used when type = out.",
			},
			"community": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The BGP community for this instance.",
			},
			"as_prepend": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The BGP prepend value for this instance. It is used when type = out.",
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
			"nat": {
				Type:        schema.TypeSet,
				MaxItems:    1,
				Optional:    true,
				Description: "Translate the source or destination IP address.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pre_nat_sources": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "If using NAT overload, this is the prefixes from the cloud that you want to associate with the NAT pool.\n\n\tExample: 10.0.0.0/24",
							Elem: &schema.Schema{
								Type:        schema.TypeString,
								Description: "IP prefix using CIDR format.",
							},
						},
						"pool_prefixes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "If using NAT overload, all prefixes that are NATed on this connection will be translated to the pool prefix address.\n\n\tExample: 10.0.0.0/32",
							Elem: &schema.Schema{
								Type:        schema.TypeString,
								Description: "IP prefix using CIDR format.",
							},
						},
						"direction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "If using NAT overload, the direction of the NAT connection. Output is the default.\n\t\tEnum: output, input.",
						},
						"nat_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The NAT type of the NAT connection, source NAT (overload) or destination NAT (inline_dnat). Overload is the default.\n\t\tEnum: overload, inline_dnat.",
						},
						"dnat_mappings": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Translate the destination IP address.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"private_prefix": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The private prefix of this DNAT mapping.",
									},
									"public_prefix": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The public prefix of this DNAT mapping.",
									},
									"conditional_prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The conditional prefix prefix of this DNAT mapping.",
									},
								},
							},
						},
					},
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
							Description:  "The match type of this prefix.\n\n\tEnum: `\"exact\"` `\"orlonger\"` `\"longer\"`",
						},
						"as_prepend": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The BGP prepend value of this prefix. It is used when type = out.",
						},
						"med": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The MED of this prefix. It is used when type = out.",
						},
						"local_preference": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The local_preference of this prefix. It is used when type = in.",
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"in", "out"}, true),
							Description:  "Whether this prefix is in (Allowed Prefixes from Cloud) or out (Allowed Prefixes to Cloud).\n\t\tEnum: in, out.",
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
	session := extractBgpSessionCreate(d)
	resp, err := c.CreateBgpSession(session, cID.(string), connCID.(string))
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}
	// check Cloud Router Connection status
	createOkCh := make(chan bool)
	defer close(createOkCh)
	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetCloudConnectionStatus(cID.(string), connCID.(string))
	}
	go c.CheckServiceStatus(createOkCh, fn)
	if !<-createOkCh {
		return diag.FromErr(err)
	}
	d.SetId(resp.BgpSettingsUUID)
	return diags
}

func resourceBgpSessionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	var cID, connCID, bgpSettingsUUID string
	if circuitID, ok := d.GetOk("circuit_id"); !ok {
		return diag.FromErr(errors.New("could not extract cloud router circuit id from resource data"))
	} else {
		cID = circuitID.(string)
	}
	if connectionCID, ok := d.GetOk("connection_id"); !ok {
		return diag.FromErr(errors.New("could not extract cloud router connection circuit id from resource data"))
	} else {
		connCID = connectionCID.(string)
	}
	if settingsUUID, ok := d.GetOk("id"); !ok {
		return diag.FromErr(errors.New("could not extract bgp settings uuid from resource data"))
	} else {
		bgpSettingsUUID = settingsUUID.(string)
	}
	if diags != nil || len(diags) > 0 {
		return diags
	}
	_, err := c.GetBgpSessionBy(cID, connCID, bgpSettingsUUID)
	if err != nil {
		return diag.FromErr(errors.New("could not retrieve bgp session"))
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
	if d.HasChange("primary_subnet") && d.HasChange("secondary_subnet") {
		return diag.FromErr(errors.New("cannot modify both primary_subnet and secondary_subnet at the same time"))
	}
	session := extractBgpSessionUpdate(d)
	_, resp, err := c.UpdateBgpSession(session, cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	// check Cloud Router Connection status
	updateOkCh := make(chan bool)
	defer close(updateOkCh)
	fn := func() (*packetfabric.ServiceState, error) {
		return c.GetCloudConnectionStatus(cID.(string), connCID.(string))
	}
	go c.CheckServiceStatus(updateOkCh, fn)
	if !<-updateOkCh {
		return diag.FromErr(err)
	}
	d.SetId(resp.BgpSettingsUUID)
	return diags
}

func resourceBgpSessionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "BGP session cannot be deleted.",
		Detail:   "It will be deleted together with the Cloud Router Connection.",
	})
	d.SetId("")
	return diags
}

func extractBgpSessionCreate(d *schema.ResourceData) packetfabric.BgpSession {
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
	if nat, ok := d.GetOk("nat"); ok {
		for _, nat := range nat.(*schema.Set).List() {
			bgpSession.Nat = extractConnBgpSessionNat(nat.(map[string]interface{}))
		}
	} else {
		bgpSession.Nat = nil
	}
	bgpSession.Prefixes = extractConnBgpSessionPrefixes(d)
	return bgpSession
}

func extractBgpSessionUpdate(d *schema.ResourceData) packetfabric.BgpSession {
	bgpSession := packetfabric.BgpSession{}
	if l3Address, ok := d.GetOk("l3_address"); ok {
		bgpSession.L3Address = l3Address.(string)
	}
	// https://docs.packetfabric.com/api/v2/swagger/#/Cloud%20Router%20BGP%20Session%20Settings/cloud_routers_bgp_update
	// Azure BGP session Update: set l3Address based on the values of primarySubnet and secondarySubnet when modified
	// This is a temporary solution until the BGP API is refactored.
	if d.HasChange("primary_subnet") {
		if primarySubnet, ok := d.GetOk("primary_subnet"); ok {
			bgpSession.L3Address = primarySubnet.(string)
		}
	}
	if d.HasChange("secondary_subnet") {
		if secondarySubnet, ok := d.GetOk("secondary_subnet"); ok {
			bgpSession.L3Address = secondarySubnet.(string)
		}
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
	if nat, ok := d.GetOk("nat"); ok {
		for _, nat := range nat.(*schema.Set).List() {
			bgpSession.Nat = extractConnBgpSessionNat(nat.(map[string]interface{}))
		}
	} else {
		bgpSession.Nat = nil
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

func extractConnBgpSessionNat(n map[string]interface{}) *packetfabric.BgpNat {
	nat := packetfabric.BgpNat{}
	if direction := n["direction"]; direction != nil {
		nat.Direction = direction.(string)
	}
	if natType := n["nat_type"]; natType != nil {
		nat.NatType = natType.(string)
	}
	nat.PreNatSources = extractPreNatSources(n["pre_nat_sources"])
	nat.PoolPrefixes = extractPoolPrefixes(n["pool_prefixes"])
	nat.DnatMappings = extractConnBgpSessionDnat(n["dnat_mappings"].(*schema.Set))
	return &nat
}

func extractPreNatSources(d interface{}) []interface{} {
	if PreNatSources, ok := d.([]interface{}); ok {
		regs := make([]interface{}, 0)
		for _, reg := range PreNatSources {
			regs = append(regs, reg.(string))
		}
		return regs
	}
	return make([]interface{}, 0)
}

func extractPoolPrefixes(d interface{}) []interface{} {
	if PoolPrefixes, ok := d.([]interface{}); ok {
		regs := make([]interface{}, 0)
		for _, reg := range PoolPrefixes {
			regs = append(regs, reg.(string))
		}
		return regs
	}
	return make([]interface{}, 0)
}

func extractConnBgpSessionDnat(d *schema.Set) []packetfabric.BgpDnatMapping {
	sessionDnat := make([]packetfabric.BgpDnatMapping, 0)
	for _, dnat := range d.List() {
		sessionDnat = append(sessionDnat, packetfabric.BgpDnatMapping{
			PrivateIP:         dnat.(map[string]interface{})["private_prefix"].(string),
			PublicIP:          dnat.(map[string]interface{})["public_prefix"].(string),
			ConditionalPrefix: dnat.(map[string]interface{})["conditional_prefix"].(string),
		})
	}
	return sessionDnat
}
