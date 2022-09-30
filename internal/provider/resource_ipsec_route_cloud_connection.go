package provider

import (
	"context"
	"fmt"
	"time"

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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID for the billing account that should be billed.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which you want to provision the connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecSpeedOptions(), true),
				Description:  "The desired speed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\"]",
			},
			"ike_version": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 2}),
				Description:  "The Internet Key Exchange (IKE) version supported by your device. In most cases, this is v2 (v1 is deprecated).\n\n\tEnum: 1, 2.",
			},
			"phase1_authentication_method": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The authentication method to use during phase 1. For example, \"pre-shared-key\".",
			},
			"phase1_group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPahsesGroupOptions(), false),
				Description:  "Phase 1 is when the VPN peers are authenticated and we establish security associations (SAs) to protect IKE messaging between the two endpoints (which in this case is PacketFabric and your VPN device). This is also known as the IKE phase.\n\n\tThe Phase 1 group is the Diffie-Hellman (DH) algorithm used to create a shared secret between the endpoints.\n\n\tEnum: \"group1\" \"group14\" \"group15\" \"group16\" \"group19\" \"group2\" \"group20\" \"group24\" \"group5\" ",
			},
			"phase1_encryption_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPhase1EncryptionAlgoOptions(), false),
				Description:  "The encryption algorithm to use during phase 1.\n\n\tEnum: \"3des-cbc\" \"aes-128-cbc\" \"aes-192-cbc\" \"aes-256-cbc\" \"des-cbc\" ",
			},
			"phase1_authentication_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPhase1AuthenticationAlgoOptions(), false),
				Description:  "The authentication algorithm to use during phase 1.\n\n\tEnum: \"md5\" \"sha-256\" \"sha-384\" \"sha1\" ",
			},
			"phase1_lifetime": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(180, 86400),
				Description:  "The time in seconds before a tunnel will need to re-authenticate. The phase 1 lifetime should be equal to or longer than phase 2. This can be between 180 and 86400.",
			},
			"phase2_pfs_group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPahsesGroupOptions(), false),
				Description:  "Phase 2 is when SAs are further established to protect and encrypt IP traffic within the tunnel. This is also known as the IPsec phase.\n\n\tThe PFS group is the Perfect Forward Secrecy group. This means that rather than using the keys from phase 1, new keys are generated per the selected Diffie-Hellman algorithm (same as those listed above).\n\n\tEnum: \"group1\" \"group14\" \"group15\" \"group16\" \"group19\" \"group2\" \"group20\" \"group24\" \"group5\" ",
			},
			"phase2_encryption_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPhase2EncryptionAlgoOptions(), false),
				Description:  "The encryption algorithm to use during phase 2.\n\n\tEnum: \"3des-cbc\" \"aes-128-cbc\" \"aes-128-gcm\" \"aes-192-cbc\" \"aes-192-gcm\" \"aes-256-cbc\" \"aes-256-gcm\" \"des-cbc\" ",
			},
			"phase2_authentication_algo": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(iPSecPhase2AuthenticationAlgoOptions(), false),
				Description:  "The authentication algorithm to use during phase 2.\n\n\tEnum: \"hmac-md5-96\" \"hmac-sha-256-128\" \"hmac-sha1-96\" ",
			},
			"phase2_lifetime": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(180, 86400),
				Description:  "The time in seconds before phase 2 expires and needs to reauthenticate. We recommend that the phase 2 lifetime is equal to or shorter than phase 1. This can be between 180 and 86400.",
			},
			"gateway_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPv4Address,
				Description:  "The gateway address of your VPN device. Because VPNs traverse the public internet, this must be a public IP address owned by you.",
			},
			"shared_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The pre-shared-key to use for authentication.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
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
	if cid, ok := d.GetOk("circuit_id"); ok {
		resp, err := c.CreateIPSecCloudRouerConnection(ipSecRouter, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		createOkCh := make(chan bool)
		defer close(createOkCh)
		fn := func() (*packetfabric.ServiceState, error) {
			return c.GetCloudConnectionStatus(cid.(string), resp.VcCircuitID)
		}
		go c.CheckIPSecStatus(createOkCh, err, fn)
		if !<-createOkCh {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.VcCircuitID)
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Circuit ID not present",
			Detail:   "Please provide a valid Circuit ID.",
		})
	}
	return diags
}

func resourceIPSecCloudRouteConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		_, err := c.ReadAwsConnection(cid.(string), cloudConnCID.(string))
		if err != nil {
			diags = diag.FromErr(err)
		}
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
	return diags
}

func resourceIPSecCloudRouteConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		if _, err := c.DeleteAwsConnection(cid.(string), cloudConnCID.(string)); err != nil {
			diags = diag.FromErr(err)
		} else {
			deleteOk := make(chan bool)
			defer close(deleteOk)
			fn := func() (*packetfabric.ServiceState, error) {
				return c.GetCloudConnectionStatus(cid.(string), cloudConnCID.(string))
			}
			go c.CheckServiceStatus(deleteOk, err, fn)
			if !<-deleteOk {
				return diag.FromErr(err)
			}
			d.SetId("")
		}
	}
	return diags
}

func extractIPSecRouteConn(d *schema.ResourceData) (packetfabric.IPSecRouterConn, error) {
	iPSecRouter := packetfabric.IPSecRouterConn{}
	if desc, ok := d.GetOk("description"); ok {
		iPSecRouter.Description = desc.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		iPSecRouter.AccountUUID = accountUUID.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		iPSecRouter.Pop = pop.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		iPSecRouter.Speed = speed.(string)
	}
	if ikeVersion, ok := d.GetOk("ike_version"); ok {
		iPSecRouter.IkeVersion = ikeVersion.(int)
	}
	if phaseOneAuthMethod, ok := d.GetOk("phase1_authentication_method"); ok {
		iPSecRouter.Phase1AuthenticationMethod = phaseOneAuthMethod.(string)
	}
	if phaseOneGroup, ok := d.GetOk("phase1_group"); ok {
		iPSecRouter.Phase1Group = phaseOneGroup.(string)
	}
	if phaseOneEncryptionAlgo, ok := d.GetOk("phase1_encryption_algo"); ok {
		iPSecRouter.Phase1EncryptionAlgo = phaseOneEncryptionAlgo.(string)
	}
	if phaseOneAuthAlgo, ok := d.GetOk("phase1_authentication_algo"); ok {
		iPSecRouter.Phase1AuthenticationAlgo = phaseOneAuthAlgo.(string)
	}
	if phaseOneLifetime, ok := d.GetOk("phase1_lifetime"); ok {
		iPSecRouter.Phase1Lifetime = phaseOneLifetime.(int)
	}
	if phaseTwoPfsGroup, ok := d.GetOk("phase2_pfs_group"); ok {
		iPSecRouter.Phase2PfsGroup = phaseTwoPfsGroup.(string)
	}
	if phaseTwoEncryptionAlgo, ok := d.GetOk("phase2_encryption_algo"); ok {
		iPSecRouter.Phase2EncryptionAlgo = phaseTwoEncryptionAlgo.(string)
	}
	if phaseTwoAuthAlgo, ok := d.GetOk("phase2_authentication_algo"); ok {
		iPSecRouter.Phase2AuthenticationAlgo = phaseTwoAuthAlgo.(string)
	} else {
		return iPSecRouter, fmt.Errorf("The field 'phase2_authentication_algo' is required.")
	}
	if phaseTwoLifetime, ok := d.GetOk("phase2_lifetime"); ok {
		iPSecRouter.Phase2Lifetime = phaseTwoLifetime.(int)
	}
	if gatewayAddress, ok := d.GetOk("gateway_address"); ok {
		iPSecRouter.GatewayAddress = gatewayAddress.(string)
	}
	if sharedKey, ok := d.GetOk("shared_key"); ok {
		iPSecRouter.SharedKey = sharedKey.(string)
	}
	if publishedQuote, ok := d.GetOk("published_quote_line_uuid"); ok {
		iPSecRouter.PublishedQuoteLineUUID = publishedQuote.(string)
	}
	return iPSecRouter, nil
}

func extractIPSecUpdate(d *schema.ResourceData) packetfabric.IPSecConnUpdate {
	ipsec := packetfabric.IPSecConnUpdate{}
	if custGatewayAdd, ok := d.GetOk("gateway_address"); ok {
		ipsec.CustomerGatewayAddress = custGatewayAdd.(string)
	}
	if ikeVersion, ok := d.GetOk("ike_version"); ok {
		ipsec.IkeVersion = ikeVersion.(int)
	}
	if phaseOneAuthMethod, ok := d.GetOk("phase1_authentication_method"); ok {
		ipsec.Phase1AuthenticationMethod = phaseOneAuthMethod.(string)
	}
	if phaseOneGroup, ok := d.GetOk("phase1_group"); ok {
		ipsec.Phase1Group = phaseOneGroup.(string)
	}
	if phaseOneEncAlgo, ok := d.GetOk("phase1_encryption_algo"); ok {
		ipsec.Phase1EncryptionAlgo = phaseOneEncAlgo.(string)
	}
	if phaseOneAuthAlgo, ok := d.GetOk("phase1_authentication_algo"); ok {
		ipsec.Phase1AuthenticationAlgo = phaseOneAuthAlgo.(string)
	}
	if phaseOneLifetime, ok := d.GetOk("phase1_lifetime"); ok {
		ipsec.Phase1Lifetime = phaseOneLifetime.(int)
	}
	if phaseTwoPfsGroup, ok := d.GetOk("phase2_pfs_group"); ok {
		ipsec.Phase2PfsGroup = phaseTwoPfsGroup.(string)
	}
	if phaseTwoEncryptationAlgo, ok := d.GetOk("phase2_encryption_algo"); ok {
		ipsec.Phase2EncryptionAlgo = phaseTwoEncryptationAlgo.(string)
	}
	if phaseTwoAuthAlgo, ok := d.GetOk("phase2_authentication_algo"); ok {
		ipsec.Phase2AuthenticationAlgo = phaseTwoAuthAlgo.(string)
	}
	if phaseTwoLifetime, ok := d.GetOk("phase2_lifetime"); ok {
		ipsec.Phase2Lifetime = phaseTwoLifetime.(int)
	}
	if preSharedKey, ok := d.GetOk("shared_key"); ok {
		ipsec.PreSharedKey = preSharedKey.(string)
	}
	return ipsec
}

func iPSecSpeedOptions() []string {
	return []string{
		"50Mbps", "100Mbps", "200Mbps", "300Mbps",
		"400Mbps", "500Mbps", "1Gbps", "2Gbps",
		"5Gbps", "10Gbps"}
}

func iPSecPahsesGroupOptions() []string {
	return []string{
		"group1", "group14", "group15", "group16",
		"group19", "group2", "group20", "group24",
		"group5"}
}

func iPSecPhase1EncryptionAlgoOptions() []string {
	return []string{
		"3des-cbc", "aes-128-cbc", "aes-192-cbc", "aes-256-cbc",
		"des-cbc"}
}

func iPSecPhase1AuthenticationAlgoOptions() []string {
	return []string{
		"md5", "sha-256", "sha-384", "sha1"}
}

func iPSecPhase2PfsGroup() []string {
	return []string{
		"md5", "sha-256", "sha-384", "sha1"}
}

func iPSecPhase2EncryptionAlgoOptions() []string {
	return []string{
		"3des-cbc", "aes-128-cbc", "aes-128-gcm", "aes-192-cbc", "aes-256-cbc", "aes-192-gcm",
		"aes-256-cbc", "aes-256-gcm", "des-cbc"}
}

func iPSecPhase2AuthenticationAlgoOptions() []string {
	return []string{
		"hmac-md5-96", "hmac-sha-256-128", "hmac-sha1-96"}
}
