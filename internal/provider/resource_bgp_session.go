package provider

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

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
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
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
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The MD5 value of the authenticated BGP sessions. Required for AWS.",
			},
			"l3_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIPAddressWithPrefix,
				Description:  "The L3 address of this instance. Not used for Azure connections. Required for all other CSP.",
			},
			"primary_subnet": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIPAddressWithPrefix,
				Description:  "Currently for Azure use only. Provide this as the primary subnet when creating the primary Azure cloud router connection.",
			},
			"secondary_subnet": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIPAddressWithPrefix,
				Description:  "Currently for Azure use only. Provide this as the secondary subnet when creating the secondary Azure cloud router connection.",
			},
			"address_family": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "v4",
				ValidateFunc: validation.StringInSlice([]string{"v4", "v6"}, true),
				Description:  "Whether this instance is IPv4 or IPv6. At this time, only IPv4 is supported.\n\n\tEnum: \"v4\" \"v6\" ",
			},
			"remote_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIPAddressWithPrefix,
				Description:  "The cloud-side router peer IP. Not used for Azure connections. Required for all other CSP.",
			},
			"remote_asn": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The cloud-side ASN.",
			},
			"multihop_ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validation.IntBetween(1, 4),
				Description:  "The TTL of this session. For Google Cloud connections, see [the PacketFabric doc](https://docs.packetfabric.com/cr/bgp/bgp_google/#ttl).\n\n\tAvailable range is 1 through 4. ",
			},
			"local_preference": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The local preference for this instance. When the same route is received in multiple locations, those with a higher local preference value are preferred by the cloud router. It is used when type = in.\n\n\tAvailable range is 1 through 4294967295. ",
			},
			"med": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The Multi-Exit Discriminator of this instance. When the same route is advertised in multiple locations, those with a lower MED are preferred by the peer AS. It is used when type = out.\n\n\tAvailable range is 1 through 4294967295. ",
			},
			"as_prepend": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(1, 5),
				Description:  "The BGP prepend value for this instance. It is used when type = out.\n\n\tAvailable range is 1 through 5. ",
			},
			"orlonger": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to use exact match or longer for all prefixes. ",
			},
			"bfd_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(3, 30000),
				Description:  "If you are using BFD, this is the interval (in milliseconds) at which to send test packets to peers.\n\n\tAvailable range is 3 through 30000. ",
			},
			"bfd_multiplier": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(2, 16),
				Description:  "If you are using BFD, this is the number of consecutive packets that can be lost before BFD considers a peer down and shuts down BGP.\n\n\tAvailable range is 2 through 16. ",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this BGP session is disabled. ",
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
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "output",
							ValidateFunc: validation.StringInSlice([]string{"output", "input"}, true),
							Description:  "If using NAT overload, the direction of the NAT connection (input=ingress, output=egress). \n\t\tEnum: output, input. ",
						},
						"nat_type": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "overload",
							ValidateFunc: validation.StringInSlice([]string{"overload", "inline_dnat"}, true),
							Description:  "The NAT type of the NAT connection, source NAT (overload) or destination NAT (inline_dnat). \n\t\tEnum: overload, inline_dnat. ",
						},
						"dnat_mappings": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Translate the destination IP address.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"private_prefix": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validateIPAddressWithPrefix,
										Description:  "Post-translation IP prefix.",
									},
									"public_prefix": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validateIPAddressWithPrefix,
										Description:  "Pre-translation IP prefix.",
									},
									"conditional_prefix": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validateIPAddressWithPrefix,
										Description:  "Post-translation prefix must be equal to or included within the conditional IP prefix.",
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
							ValidateFunc: validateIPAddressWithPrefix,
							Description:  "The actual IP Prefix of this instance.",
						},
						"match_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger"}, true),
							Description:  "The match type of this prefix.\n\n\tEnum: `\"exact\"` `\"orlonger\"` ",
						},
						"as_prepend": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntBetween(1, 5),
							Description:  "The BGP prepend value of this prefix. It is used when type = out.\n\n\tAvailable range is 1 through 5. ",
						},
						"med": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The MED of this prefix. It is used when type = out.\n\n\tAvailable range is 1 through 4294967295. ",
						},
						"local_preference": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The local_preference of this prefix. It is used when type = in.\n\n\tAvailable range is 1 through 4294967295. ",
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"in", "out"}, true),
							Description:  "Whether this prefix is in (Allowed Prefixes from Cloud) or out (Allowed Prefixes to Cloud).\n\t\tEnum: in, out.",
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: BgpImportStatePassthroughContext,
		},
	}
}

func resourceBgpSessionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	prefixesSet := d.Get("prefixes").(*schema.Set)
	prefixesList := prefixesSet.List()
	if err := validatePrefixes(prefixesList); err != nil {
		return diag.FromErr(err)
	}
	session := extractBgpSessionCreate(d)
	resp, err := c.CreateBgpSession(session, cID.(string), connCID.(string))
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}
	if err := checkCloudRouterConnectionStatus(c, cID.(string), connCID.(string)); err != nil {
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
	bgp, err := c.GetBgpSessionBy(cID, connCID, bgpSettingsUUID)
	if err != nil {
		return diag.FromErr(errors.New("could not retrieve bgp session"))
	}

	_ = d.Set("remote_asn", bgp.RemoteAsn)
	_ = d.Set("disabled", bgp.Disabled)
	_ = d.Set("orlonger", bgp.Orlonger)
	_ = d.Set("address_family", bgp.AddressFamily)
	_ = d.Set("multihop_ttl", bgp.MultihopTTL)

	// If not Azure (Subnet empty)
	if bgp.Subnet == "" {
		_ = d.Set("l3_address", bgp.L3Address)
		_ = d.Set("remote_address", bgp.RemoteAddress)
	} else {
		// If Azure will unset l3_address remote_address as those aren't in the BGP resource definition for Azure
		_ = d.Set("l3_address", nil)
		_ = d.Set("remote_address", nil)
		// There is no way to know which Subnet is the primary or the secondary one
		// Display warning in case none of the is set in the state file
		if _, ok := d.GetOk("primary_subnet"); !ok {
			if _, ok := d.GetOk("secondary_subnet"); !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  fmt.Sprintf("Manually set primary_subnet or secondary_subnet in Terraform state file using %s", bgp.Subnet),
				})
			}
		}
	}
	if bgp.Md5 != "" {
		_ = d.Set("md5", bgp.Md5)
	}
	_ = d.Set("med", bgp.Med)
	_ = d.Set("as_prepend", bgp.AsPrepend)
	_ = d.Set("local_preference", bgp.LocalPreference)
	_ = d.Set("bfd_interval", bgp.BfdInterval)
	_ = d.Set("bfd_multiplier", bgp.BfdMultiplier)

	if bgp.Nat != nil {
		nat := flattenNatConfiguration(bgp.Nat)
		if err := d.Set("nat", nat); err != nil {
			return diag.Errorf("error setting 'nat': %s", err)
		}
	}
	prefixes := flattenPrefixConfiguration(bgp.Prefixes)
	if err := d.Set("prefixes", prefixes); err != nil {
		return diag.Errorf("error setting 'prefixes': %s", err)
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
	prefixesSet := d.Get("prefixes").(*schema.Set)
	prefixesList := prefixesSet.List()
	if err := validatePrefixes(prefixesList); err != nil {
		return diag.FromErr(err)
	}
	session, err := extractBgpSessionUpdate(d, c, cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	_, resp, err := c.UpdateBgpSession(session, cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	if err := checkCloudRouterConnectionStatus(c, cID.(string), connCID.(string)); err != nil {
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
		Summary:  "BGP session will be deleted together with the Cloud Router Connection.",
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

func extractBgpSessionUpdate(d *schema.ResourceData, c *packetfabric.PFClient, cID string, connCID string) (packetfabric.BgpSession, error) {
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
	// For Azure, the L3Address is not set so determine which subnet to use
	if "" == bgpSession.L3Address {
		l3Address, err := extractL3Address(&bgpSession, c, cID, connCID)
		if err != nil {
			return packetfabric.BgpSession{}, err
		}
		bgpSession.L3Address = l3Address
	}
	if d.HasChange("disabled") {
		if disabled, ok := d.GetOk("disabled"); ok {
			bgpSession.Disabled = disabled.(bool)
		}
	}
	if addressFamily, ok := d.GetOk("address_family"); ok {
		bgpSession.AddressFamily = addressFamily.(string)
	}
	//remote_address not used for Azure
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
	return bgpSession, nil
}

func extractL3Address(bgpSession *packetfabric.BgpSession, c *packetfabric.PFClient, cID string, connCID string) (string, error) {
	value := ""
	crConnection, err := c.ReadCloudRouterConnection(cID, connCID)
	if err != nil {
		return value, err
	}
	serviceProvider := crConnection.ServiceProvider
	if "azure" != serviceProvider {
		return value, fmt.Errorf("The l3_address is a required field for %s", serviceProvider)
	}
	connectionType := crConnection.CloudSettings.AzureConnectionType
	switch connectionType {
	case "primary":
		value = bgpSession.PrimarySubnet
	case "secondary":
		value = bgpSession.SecondarySubnet
	default:
		err = fmt.Errorf("Invalid value for subnet: \"%s\"", connectionType)
	}
	if "" == value && nil == err {
		noValueMessageFormat := "The l3_address should use \"%s\" subnet but it has no value"
		err = fmt.Errorf(noValueMessageFormat, connectionType)
	}
	return value, err
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

func flattenNatConfiguration(nat *packetfabric.BgpNat) []interface{} {
	if nat == nil {
		return nil
	}

	data := make(map[string]interface{})
	data["pre_nat_sources"] = nat.PreNatSources
	data["pool_prefixes"] = nat.PoolPrefixes
	data["direction"] = nat.Direction
	data["nat_type"] = nat.NatType

	if nat.DnatMappings != nil {
		data["dnat_mappings"] = flattenDnatMappings(nat.DnatMappings)
	}

	return []interface{}{data}
}

func flattenDnatMappings(dnatMappings []packetfabric.BgpDnatMapping) []interface{} {
	result := make([]interface{}, len(dnatMappings))
	for i, dnat := range dnatMappings {
		data := make(map[string]interface{})
		data["private_prefix"] = dnat.PrivateIP
		data["public_prefix"] = dnat.PublicIP
		data["conditional_prefix"] = dnat.ConditionalPrefix
		result[i] = data
	}
	return result
}

func flattenPrefixConfiguration(prefixes []packetfabric.BgpPrefix) []interface{} {
	result := make([]interface{}, len(prefixes))
	for i, prefix := range prefixes {
		data := make(map[string]interface{})
		data["prefix"] = prefix.Prefix
		data["match_type"] = prefix.MatchType
		data["as_prepend"] = prefix.AsPrepend
		data["med"] = prefix.Med
		data["local_preference"] = prefix.LocalPreference
		data["type"] = prefix.Type
		result[i] = data
	}
	return result
}

func validatePrefixes(prefixesList []interface{}) error {
	inCount, outCount := 0, 0
	for _, prefix := range prefixesList {
		prefixMap := prefix.(map[string]interface{})
		prefixType := prefixMap["type"].(string)

		if prefixType == "in" {
			inCount++
		} else if prefixType == "out" {
			outCount++
		}
	}
	if inCount == 0 || outCount == 0 {
		return fmt.Errorf("at least 1 'in' and 1 'out' prefix must be provided")
	}
	return nil
}

func validateIPAddressWithPrefix(val interface{}, key string) (warns []string, errs []error) {
	value := val.(string)
	_, _, err := net.ParseCIDR(value)
	if err != nil {
		errs = append(errs, fmt.Errorf("%q is not a valid IP address with prefix: %s", key, value))
	}
	return
}
