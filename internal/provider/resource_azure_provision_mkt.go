package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAzureProvision() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAzureProvisionCreate,
		ReadContext:   resourceAzureProvisionRead,
		UpdateContext: resourceAzureProvisionUpdate,
		DeleteContext: resourceAzureProvisionDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: resourceProvisionAzure(),
	}
}

func resourceAzureProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceProvisionCreate(ctx, d, m, c.CreateMktProvisionReq, azureProvider)
}

func resourceAzureProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesHostedRead(ctx, d, m)
}

func resourceAzureProvisionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesHostedUpdate(ctx, d, m, c.UpdateServiceHostedConn)
}

func resourceAzureProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudSourceDelete(ctx, d, m, "Azure Service Delete")
}
