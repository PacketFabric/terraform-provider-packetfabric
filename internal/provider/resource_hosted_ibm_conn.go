package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceHostedIbmConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourceHostedIbmConnCreate,
		ReadContext:   resourceHostedIbmConnRead,
		UpdateContext: resourceHostedIbmConnUpdate,
		DeleteContext: resourceHostedIbmConnDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ibm_account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Your IBM account ID.",
			},
			"ibm_bgp_asn": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Enter an ASN to use with your BGP session. This should be the same ASN you used for your Cloud Router.",
			},
			"ibm_bgp_cer_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The IP address in CIDR format for the PacketFabric-side router in the BGP session. If you do not specify an address, IBM will assign one on your behalf.",
			},
			"ibm_bgp_ibm_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The IP address in CIDR format for the IBM-side router in the BGP session. If you do not specify an address, IBM will assign one on your behalf. See the documentation for information on which IP ranges are allowed.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of this connection. This will appear as the connection name from the IBM side.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID of the contact that will be billed.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which you want to provision the connection (the on-ramp).",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port to connect to IBM.",
			},
			"vlan": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid VLAN range is from 4-4094, inclusive.",
			},
			"src_svlan": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Valid S-VLAN range is from 4-4094, inclusive.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired availability zone of the connection.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceHostedIbmConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ibmConn := extractHostedIBMConn(d)
	expectedResp, err := c.CreateHostedIBMConn(ibmConn)
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			dedicatedConns, err := c.GetCurrentCustomersHosted()
			if dedicatedConns != nil && err == nil && len(dedicatedConns) > 0 {
				for _, conn := range dedicatedConns {
					if expectedResp.UUID == conn.UUID && conn.State == "active" {
						expectedResp.CloudCircuitID = conn.CloudCircuitID
						ticker.Stop()
						createOk <- true
					}
				}
			}
		}
	}()
	<-createOk
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(expectedResp.CloudCircuitID)
	return diags

}

func resourceHostedIbmConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceHostedIbmConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceHostedIbmConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudSourceDelete(ctx, d, m, "IBM Service Delete")
}

func extractHostedIBMConn(d *schema.ResourceData) packetfabric.HostedIBMConn {
	hostedConn := packetfabric.HostedIBMConn{}
	if accountID, ok := d.GetOk("ibm_account_id"); ok {
		hostedConn.IbmAccountID = accountID.(string)
	}
	if ibmBgpAsn, ok := d.GetOk("ibm_bgp_asn"); ok {
		hostedConn.IbmBgpAsn = ibmBgpAsn.(int)
	}
	if cerCidr, ok := d.GetOk("ibm_bgp_cer_cidr"); ok {
		hostedConn.IbmBgpCerCidr = cerCidr.(string)
	}
	if bgpCidr, ok := d.GetOk("ibm_bgp_ibm_cidr"); ok {
		hostedConn.IbmBgpIbmCidr = bgpCidr.(string)
	}
	if desc, ok := d.GetOk("description"); ok {
		hostedConn.Description = desc.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedConn.AccountUUID = accountUUID.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		hostedConn.Pop = pop.(string)
	}
	if port, ok := d.GetOk("port"); ok {
		hostedConn.Port = port.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		hostedConn.Vlan = vlan.(int)
	}
	if srcVlan, ok := d.GetOk("src_vlan"); ok {
		hostedConn.SrcSvlan = srcVlan.(int)
	}
	if zone, ok := d.GetOk("zone"); ok {
		hostedConn.Zone = zone.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedConn.Speed = speed.(string)
	}
	return hostedConn
}
