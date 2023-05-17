package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudProviderCredentialGoogle() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudProviderCredentialGoogleCreate,
		ReadContext:   resourceCloudProviderCredentialRead,
		UpdateContext: resourceCloudProviderCredentialGoogleUpdate,
		DeleteContext: resourceCloudProviderCredentialDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Description of the Cloud Provider Credentials.",
			},
			"google_service_account": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GOOGLE_CREDENTIALS", nil),
				Description: "The Google service account JSON you want to save. " +
					"Can also be set with the GOOGLE_CREDENTIALS environment variable.",
			},
		},
	}
}

func resourceCloudProviderCredentialGoogleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	cpc := extractCloudProviderCredentialsGoogle(d)

	resp, err := c.CreateCloudProviderCredential(cpc)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.CloudProviderCredentialUUID)
	return diags
}

func resourceCloudProviderCredentialGoogleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	cpcID := d.Id()

	cpc := packetfabric.CloudProviderCredentialUpdate{}

	if d.HasChange("description") {
		cpc.Description = d.Get("description").(string)
	}
	credentials := packetfabric.CloudCredentials{}
	if googleServiceAccount, ok := d.GetOk("google_service_account"); ok {
		credentials.GoogleServiceAccount = googleServiceAccount.(string)
	}
	cpc.CloudCredentials = credentials

	_, err := c.UpdateCloudProviderCredential(cpc, cpcID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(cpcID)
	return diags
}

func extractCloudProviderCredentialsGoogle(d *schema.ResourceData) packetfabric.CloudProviderCredentialCreate {
	cpc := packetfabric.CloudProviderCredentialCreate{}
	cpc.CloudProvider = "google"
	if description, ok := d.GetOk("description"); ok {
		cpc.Description = description.(string)
	}
	cloudCredentials := packetfabric.CloudCredentials{}
	if googleServiceAccount, ok := d.GetOk("google_service_account"); ok {
		cloudCredentials.GoogleServiceAccount = googleServiceAccount.(string)
	}
	cpc.CloudCredentials = cloudCredentials
	return cpc
}
