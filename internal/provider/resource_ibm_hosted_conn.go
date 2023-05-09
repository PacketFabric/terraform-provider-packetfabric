package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceHostedIbmConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_IBM_ACCOUNT_ID", nil),
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description: "Your IBM account ID. " +
					"Can also be set with the PF_IBM_ACCOUNT_ID environment variable.",
			},
			"ibm_bgp_asn": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Enter an ASN to use with your BGP session. This should be the same ASN you used for your Cloud Router.",
			},
			"ibm_bgp_cer_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The IP address in CIDR format for the PacketFabric-side router in the BGP session. If you do not specify an address, IBM will assign one on your behalf.",
			},
			"ibm_bgp_ibm_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The IP address in CIDR format for the IBM-side router in the BGP session. If you do not specify an address, IBM will assign one on your behalf. See the documentation for information on which IP ranges are allowed.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of this connection. This will appear as the connection name from the IBM side. Allows only numbers, letters, underscores and dashes.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which you want to provision the connection (the on-ramp).",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port to connect to IBM.",
			},
			"vlan": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid VLAN range is from 4-4094, inclusive.",
			},
			"src_svlan": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid S-VLAN range is from 4-4094, inclusive.",
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
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired availability zone of the connection.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"labels": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Label value linked to an object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gateway_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IBM Gateway ID.",
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
	if err != nil {
		return diag.FromErr(err)
	}
	// Cloud Everywhere: if cloud_circuit_id is null display error
	if expectedResp.CloudCircuitID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Hosted location Requested",
			Detail: "On-ramp location does not have a Hosted port currently available. " +
				"Check in the Portal when your hosted cloud is provisioned and import the resource into your Terraform state file.",
		})
		return diags
	}
	b := make(map[string]interface{})
	b["ibm"] = expectedResp
	tflog.Debug(ctx, "\n#### CREATED IBM CONN", b)
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			dedicatedConns, err := c.GetCurrentCustomersHosted()
			if dedicatedConns != nil && err == nil && len(dedicatedConns) > 0 {
				for _, conn := range dedicatedConns {
					// we stop when the state = requested
					if expectedResp.UUID == conn.UUID && conn.State == "requested" {
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

	time.Sleep(90 * time.Second) // wait for the connection to show in IBM
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp.Settings.GatewayID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Incomplete Cloud Information",
			Detail:   "The gateway_id is currently unavailable.",
		})
		return diags
	} else {
		_ = d.Set("gateway_id", resp.Settings.GatewayID)
	}

	if labels, ok := d.GetOk("labels"); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}
	return diags
}

func resourceHostedIbmConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("cloud_circuit_id", resp.CloudCircuitID)
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("vlan", resp.Settings.VlanIDCust)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("pop", resp.CloudProvider.Pop)
		_ = d.Set("ibm_account_id", resp.Settings.AccountID)
		_ = d.Set("ibm_bgp_asn", resp.Settings.BgpAsn)
		if _, ok := d.GetOk("ibm_bgp_cer_cidr"); ok {
			_ = d.Set("ibm_bgp_cer_cidr", resp.Settings.BgpCerCidr)
		}
		if _, ok := d.GetOk("ibm_bgp_ibm_cidr"); ok {
			_ = d.Set("ibm_bgp_ibm_cidr", resp.Settings.BgpIbmCidr)
		}
		if _, ok := d.GetOk("po_number"); ok {
			_ = d.Set("po_number", resp.PONumber)
		}
		if resp.Settings.GatewayID == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Incomplete Cloud Information",
				Detail:   "The gateway_id is currently unavailable.",
			})
			return diags
		} else {
			_ = d.Set("gateway_id", resp.Settings.GatewayID)
		}
	}
	resp2, err2 := c.GetBackboneByVcCID(d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	if resp2 != nil {
		_ = d.Set("port", resp2.Interfaces[0].PortCircuitID) // Port A
		if resp2.Interfaces[0].Svlan != 0 {
			_ = d.Set("src_svlan", resp2.Interfaces[0].Svlan) // Port A if ENNI
		}
		_ = d.Set("zone", resp2.Interfaces[1].Zone) // Port Z
	}
	// unsetFields: published_quote_line_uuid

	labels, err3 := getLabels(c, d.Id())
	if err3 != nil {
		return diag.FromErr(err3)
	}
	_ = d.Set("labels", labels)
	return diags
}

func resourceHostedIbmConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesHostedUpdate(ctx, d, m)
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
	if poNumber, ok := d.GetOk("po_number"); ok {
		hostedConn.PONumber = poNumber.(string)
	}
	return hostedConn
}
