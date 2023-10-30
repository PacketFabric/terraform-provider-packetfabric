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
			PfCloudCredentials: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfCloudProviderCredentialUuid: schemaStringComputedPlain(),
						PfDescription:                 schemaStringComputed(PfCloudProviderCredentials),
						PfCloudProvider:               schemaStringComputed(PfCloudProviderDescription3),
						PfIsUnused:                    schemaBoolComputedPlain(),
						PfTimeCreated:                 schemaStringComputed(PfTimeCreatedDescription),
						PfTimeUpdated:                 schemaStringComputed(PfTimeUpdatedDescription),
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
	cloudProviderCredentials, err := c.ListCloudProviderCredential()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfCloudCredentials, flattenCloudProviderCredentials(cloudProviderCredentials))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenCloudProviderCredentials(credentials *[]packetfabric.CloudProviderCredentialResponse) []interface{} {
	if credentials != nil {
		flattens := make([]interface{}, len(*credentials))
		for i, credential := range *credentials {
			flattens[i] = structToMapAll(&credential)
		}
		return flattens
	}
	return make([]interface{}, 0)
}
