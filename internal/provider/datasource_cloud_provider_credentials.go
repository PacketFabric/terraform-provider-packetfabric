package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceCloudProviderCredentials() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceCloudProviderCredentialsRead,
		Schema: map[string]*schema.Schema{
			"cloud_credentials": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud_provider_credential_uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the Cloud Provider Credentials.",
						},
						"cloud_provider": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "For cloud connections, this is the cloud provider: \"aws\", \"google\", \"oracle\", \"azure\"",
						},
						"is_unused": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time of connection creation",
						},
						"time_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time connection was last updated",
						},
					},
				},
			},
		},
	}
}

func datasourceCloudProviderCredentialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudProviderCredentials, err := c.ListCloudProviderCredentials()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("cloud_credentials", flattenCloudProviderCredentials(cloudProviderCredentials))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenCloudProviderCredentials(credentials []packetfabric.CloudProviderCredentialResponse) []interface{} {
	if credentials != nil {
		flattens := make([]interface{}, len(credentials))
		for i, credential := range credentials {
			flatten := make(map[string]interface{})
			flatten["cloud_provider_credential_uuid"] = credential.CloudProviderCredentialUUID
			flatten["description"] = credential.Description
			flatten["cloud_provider"] = credential.CloudProvider
			flatten["is_unused"] = credential.IsUnused
			flatten["time_created"] = credential.TimeCreated
			flatten["time_updated"] = credential.TimeUpdated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
