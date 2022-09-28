package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIPSecUpdate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpSecConnCreate,
		ReadContext:   resourceIpSecConnRead,
		UpdateContext: resourceIpSecConnUpdate,
		DeleteContext: resourceIpSecConnDelete,
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
				Description:  "IPSec circuit ID or its associated VC.",
			},
			"ike_version": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 2}),
				Description:  "The IKE version.",
			},
			"phase1_authentication_method": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The phase 1 authentication method.",
			},
			"phase1_group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"group5", "group14"}, true),
				Description:  "The phase 1 group.",
			},
			"phase1_encryption_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"aes-256-cbc", "aes-128-cbc"}, true),
				Description:  "Phase 1 encryption algo.",
			},
			"phase1_authentication_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"sha-384", "sha1"}, true),
				Description:  "The Phase 1 authentication algorithm.",
			},
			"phase1_lifetime": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "The phase 1 lifetime.",
			},
			"phase2_pfs_group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"group5", "group14"}, true),
				Description:  "Phase 2 PFS Group.",
			},
			"phase2_encryption_algo": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"aes-256-cbc", "aes-256-gcm", "aes-128-cbc"}, true),
				Description:  "The phase 2 encryption algorithm.",
			},
			"phase2_authentication_algo": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"sha-384", "sha1"}, true),
				Description:  "This field cannot be null if phase2_encryption_algo is a CBC algorithm.",
			},
			"pahse2_lifetime": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "The phase 2 lifetime.",
			},
			"pre_shared_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The pre-shared key for this connection.",
			},
		},
	}
}

func resourceIpSecConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		ipsecUpdated := extractIPSecUpdate(d)
		resp, err := c.UpdateIPSecConnection(cid.(string), ipsecUpdated)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(resp.CircuitID)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Circuit ID not present",
			Detail:   "Please provide a valid Circuit ID.",
		})
	}
	return diags
}

func resourceIpSecConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		_, err := c.GetIpsecSpecificConn(cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpSecConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	// The IPSec Resource doesn't have an update operation
	return
}

func resourceIpSecConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	// The IPSec Resource doesn't have a delete operation
	d.SetId("")
	return
}

func extractIPSecUpdate(d *schema.ResourceData) packetfabric.IPSecConnUpdate {
	ipsec := packetfabric.IPSecConnUpdate{}
	if custGatewayAdd, ok := d.GetOk("customer_gateway_address"); ok {
		ipsec.CustomerGatewayAddress = custGatewayAdd.(string)
	}
	if ikeVersion, ok := d.GetOk("ike_version"); ok {
		ipsec.IkeVersion = ikeVersion.(int)
	}
	if phaseOneAuthMethod, ok := d.GetOk("phase1_authentication_method"); ok {
		ipsec.Phase1AuthenticationMethod = phaseOneAuthMethod.(string)
	}
	if phaseOneGroup, ok := d.GetOk("pahse1_group"); ok {
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
	if preSharedKey, ok := d.GetOk("pre_shared_key"); ok {
		ipsec.PreSharedKey = preSharedKey.(string)
	}
	return ipsec
}
