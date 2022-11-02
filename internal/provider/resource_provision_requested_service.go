package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceProvisionRequestedService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProvisionRequestedServiceCreate,
		ReadContext:   resourceRequestedServiceRead,
		UpdateContext: resourceRequestedServiceUpdate,
		DeleteContext: resourceRequestedServiceDelete,
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
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"backbone", "ix", "cloud"}, true),
				Description:  "The service type.",
			},
			"cloud_provider": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"aws", "google", "oracle", "azure"}, true),
				Description:  "The cloud provider.",
			},
			"vc_request_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The VC Request UUID to be provisioned.",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Provision request description.",
			},
			"interface": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_circuit_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The circuit ID for the port. This starts with \"PF-AP-\"",
						},
						"vlan": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntBetween(4, 4094),
							Description:  "Valid VLAN range is from 4-4094, inclusive.",
						},
						"svlan": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntBetween(4, 4094),
							Description:  "Valid S-VLAN range is from 4-4094, inclusive.",
						},
						"vlan_private": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IntBetween(4, 4094),
							Description:  "Valid VLAN Private range is from 4-4094, inclusive.",
						},
						"vlan_microsoft": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IntBetween(4, 4094),
							Description:  "Valid VLAN Microsoft range is from 4-4094, inclusive.",
						},
						"untagged": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether the interface should be untagged.",
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

func resourceProvisionRequestedServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	provisionReq := extractProvisionRequest(d)
	vcReqUUID := d.Get("vc_request_uuid")
	reqType := d.Get("type")
	provDebug := make(map[string]interface{})
	provDebug["provision_req"] = provisionReq
	tflog.Debug(ctx, "\n### REQUEST PROVISION ###", provDebug)
	_, err := c.RequestServiceProvision(vcReqUUID.(string), reqType.(string), provisionReq)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())

	return diags
}

func resourceRequestedServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	vcCID := d.Id()
	if _, err := c.GetBackboneByVcCID(vcCID); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceRequestedServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceUpdateMarketplace(ctx, d, m)
}

func resourceRequestedServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Marketplace Request delete.",
		Detail: `Will not delete the Marketplace Request. 
		Terraform will remove this resource from the state file, 
		however Marketplace Request may remain.`,
	}}
}

func extractProvisionRequest(d *schema.ResourceData) packetfabric.ServiceProvision {
	provisionReq := packetfabric.ServiceProvision{}
	if desc, ok := d.GetOk("description"); ok {
		provisionReq.Description = desc.(string)
	}
	cloudProvider := d.Get("cloud_provider").(string)
	provisionReq.Provider = cloudProvider
	for _, interfA := range d.Get("interface").(*schema.Set).List() {
		provisionReq.Interface = extractProvisionInterf(cloudProvider, interfA.(map[string]interface{}))
	}
	return provisionReq
}

func extractProvisionInterf(cloudProvider string, interf map[string]interface{}) packetfabric.Interface {
	provisionInterf := packetfabric.Interface{}
	provisionInterf.PortCircuitID = interf["port_circuit_id"].(string)
	switch cloudProvider {
	case awsProvider, googleProvider, oracleProvider:
		provisionInterf.Vlan = interf["vlan"].(int)
	case azureProvider:
		provisionInterf.VlanMicrosoft = interf["vlan_microsoft"].(int)
		provisionInterf.VlanPrivate = interf["vlan_private"].(int)
	default:
		provisionInterf.Vlan = interf["vlan"].(int)
		provisionInterf.Svlan = interf["svlan"].(int)
		provisionInterf.Untagged = interf["untagged"].(bool)
	}
	return provisionInterf
}