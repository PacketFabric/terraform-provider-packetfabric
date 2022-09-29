package provider

import (
	"context"
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
				Description:  "The circuit ID.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of this connection.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID of the contact that will be billed",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired location for the new connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecSpeedOptions(), true),
				Description:  "The peed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\"]",
			},
			"ike_version": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 2}),
				Description:  "Enum: 1, 2.",
			},
			"phase1_authentication_method": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The phase 1 auth method.",
			},
			"phase1_group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPahsesGroupOptions(), false),
				Description:  "The phase 1 group.",
			},
			"phase1_encryption_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPhase1EncryptionAlgoOptions(), false),
				Description:  "The phase 1 encryption algo.",
			},
			"phase1_authentication_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPhase1AuthenticationAlgoOptions(), false),
				Description:  "The phase 1 authentication algo.",
			},
			"phase1_lifetime": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(180, 86400),
				Description:  "The phase 1 lifetime.",
			},
			"phase2_pfs_group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPahsesGroupOptions(), false),
				Description:  "The phase 1 authentication algo.",
			},
			"phase2_encryption_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPhase2EncryptionAlgoOptions(), false),
				Description:  "The phase 2 encryption algo.",
			},
			"phase2_authentication_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(iPSecPhase2AuthenticationAlgoOptions(), false),
				Description:  "The phase 2 authentication algo.",
			},
			"phase2_lifetime": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(180, 86400),
				Description:  "The phase 2 lifetime.",
			},
			"gateway_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPv4Address,
				Description:  "The customer-side (remote) gateway address.",
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
	ipSecRouter := extractIPSecRouteConn(d)
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
		go c.CheckServiceStatus(createOkCh, err, fn)
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
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		desc := d.Get("description")
		descUpdate := packetfabric.DescriptionUpdate{
			Description: desc.(string),
		}
		if _, err := c.UpdateAwsConnection(cid.(string), cloudConnCID.(string), descUpdate); err != nil {
			diags = diag.FromErr(err)
		}
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

func extractIPSecRouteConn(d *schema.ResourceData) packetfabric.IPSecRouterConn {
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
	return iPSecRouter
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
