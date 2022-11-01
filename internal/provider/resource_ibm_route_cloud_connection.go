package provider

import (
        "context"
        "time"

        "github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
        "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIBMCloudRouteConn() *schema.Resource {
        return &schema.Resource{
                CreateContext: resourceIBMCloudRouteConnCreate,
                ReadContext:   resourceIBMCloudRouteConnRead,
                UpdateContext: resourceIBMCloudRouteConnUpdate,
                DeleteContext: resourceIBMCloudRouteConnDelete,
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
                        "maybe_nat": {
                                Type:        schema.TypeBool,
                                Optional:    true,
                                Default:     false,
                                Description: "Circuit ID of the target cloud router.",
                        },
                        "circuit_id": {
                                Type:         schema.TypeString,
                                Required:     true,
                                ValidateFunc: validation.StringIsNotEmpty,
                                Description:  "Circuit ID of the target cloud router.",
                        },
                        "account_uuid": {
                                Type:         schema.TypeString,
                                Required:     true,
                                ValidateFunc: validation.IsUUID,
                                Description:  "The UUID for the billing account that should be billed.",
                        },
                        "ibm_account_id": {
                                Type:         schema.TypeString,
                                Required:     true,
                                ValidateFunc: validation.StringLenBetween(1, 32),
                                Description:  "The customer IBM account ID to connect with.",
                        },
                        "ibm_bgp_asn": {
                                Type:        schema.TypeInt,
                                Required:    true,
                                Description: "The customers IP address for the BGP session with IBM..",
                        },
                        "ibm_bgp_cer_cidr": {
                                Type:         schema.TypeString,
                                Optional:     true,
                                ValidateFunc: validation.StringIsNotEmpty,
                                Description:  "The customer IP address for the BGP session with IBM.",
                        },
                        "ibm_bgp_ibm_cidr": {
                                Type:         schema.TypeString,
                                Optional:     true,
                                ValidateFunc: validation.StringIsNotEmpty,
                                Description:  "IBM's IP address for the BGP session.",
                        },
                        "description": {
                                Type:         schema.TypeString,
                                Required:     true,
                                ValidateFunc: validation.StringIsNotEmpty,
                                Description:  "The description of this connection. Will be shown as name field in IBM portal.",
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
                                ValidateFunc: validation.StringInSlice(speedOptions(), true),
                                Description:  "The peed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
                        },
                        "zone": {
                                Type:         schema.TypeString,
                                Optional:     true,
                                ValidateFunc: validation.StringIsNotEmpty,
                                Description:  "The desired zone of the new connection.",
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

func resourceIBMCloudRouteConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        c := m.(*packetfabric.PFClient)
        c.Ctx = ctx

        var diags diag.Diagnostics
        if cid, ok := d.GetOk("circuit_id"); ok {
                ibmRouter := extractIBMCloudRouterConn(d)
                resp, err := c.CreateIBMCloudRouteConn(ibmRouter, cid.(string))
                if err != nil {
                        return diag.FromErr(err)
                }
                createOkCh := make(chan bool)
                defer close(createOkCh)
                fn := func() (*packetfabric.ServiceState, error) {
                        return c.GetCloudConnectionStatus(cid.(string), resp.CloudCircuitID)
                }
                go c.CheckServiceStatus(createOkCh, fn)
                if !<-createOkCh {
                        return diag.FromErr(err)
                }
                if resp != nil {
                        d.SetId(resp.CloudCircuitID)
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

func resourceIBMCloudRouteConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceIBMCloudRouteConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        c := m.(*packetfabric.PFClient)
        c.Ctx = ctx
        var diags diag.Diagnostics
        if cid, ok := d.GetOk("circuit_id"); ok {
                cloudConnCID := d.Get("id")
                desc := d.Get("description")
                descUpdate := packetfabric.DescriptionUpdate{
                        Description: desc.(string),
                }
                if _, err := c.UpdateCloudRouterConnection(cid.(string), cloudConnCID.(string), descUpdate); err != nil {
                        diags = diag.FromErr(err)
                }
        }
        return diags
}

func resourceIBMCloudRouteConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractIBMCloudRouterConn(d *schema.ResourceData) packetfabric.IBMCloudRouterConn {
        ibmRouter := packetfabric.IBMCloudRouterConn{}
        if maybeNat, ok := d.GetOk("maybe_nat"); ok {
                ibmRouter.MaybeNat = maybeNat.(bool)
        }
        if ibmAccountID, ok := d.GetOk("ibm_account_id"); ok {
                ibmRouter.IbmAccountID = ibmAccountID.(string)
        }
        if ibmBgpAsn, ok := d.GetOk("ibm_bgp_asn"); ok {
                ibmRouter.IbmBgpAsn = ibmBgpAsn.(int)
        }
        if ibmBgpCerCidr, ok := d.GetOk("ibm_bgp_cer_cidr"); ok {
                ibmRouter.IbmBgpCerCidr = ibmBgpCerCidr.(string)
        }
        if ibmBgpIbmCidr, ok := d.GetOk("ibm_bgp_ibm_cidr"); ok {
                ibmRouter.IbmBgpIbmCidr = ibmBgpIbmCidr.(string)
        }
        if desc, ok := d.GetOk("description"); ok {
                ibmRouter.Description = desc.(string)
        }
        if accountUUID, ok := d.GetOk("account_uuid"); ok {
                ibmRouter.AccountUUID = accountUUID.(string)
        }
        if pop, ok := d.GetOk("pop"); ok {
                ibmRouter.Pop = pop.(string)
        }
        if zone, ok := d.GetOk("zone"); ok {
                ibmRouter.Zone = zone.(string)
        }
        if speed, ok := d.GetOk("speed"); ok {
                ibmRouter.Speed = speed.(string)
        }
        if publishedQuote, ok := d.GetOk("published_quote_line_uuid"); ok {
                ibmRouter.PublishedQuoteLineUUID = publishedQuote.(string)
        }
        return ibmRouter
}
