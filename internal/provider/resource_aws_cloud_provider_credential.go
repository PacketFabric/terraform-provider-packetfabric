package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudProviderCredentialAws() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudProviderCredentialAwsCreate,
		ReadContext:   resourceCloudProviderCredentialRead,
		UpdateContext: resourceCloudProviderCredentialAwsUpdate,
		DeleteContext: resourceCloudProviderCredentialDelete,
		Schema: map[string]*schema.Schema{
			PfId:           schemaStringComputedPlain(),
			PfDescription:  schemaStringRequired(PfCloudProviderCredentials),
			PfAwsAccessKey: schemaStringEnvSensitive(PfeAwsAccessKeyId, PfAwsAccessKeyDescription),
			PfAwsSecretKey: schemaStringEnvSensitive(PfeAwsSecretAccessKey, PfAwsSecretKeyDescription),
		},
	}
}

func resourceCloudProviderCredentialAwsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	cpc := extractCloudProviderCredentialsAws(d)

	resp, err := c.CreateCloudProviderCredential(cpc)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.CloudProviderCredentialUUID)
	return diags
}

func resourceCloudProviderCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	cpcID := d.Id()

	_, err := c.ReadCloudProviderCredential(cpcID)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceCloudProviderCredentialAwsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	cpcID := d.Id()

	cpc := packetfabric.CloudProviderCredentialUpdate{}

	if d.HasChange(PfDescription) {
		cpc.Description = d.Get(PfDescription).(string)
	}
	credentials := packetfabric.CloudCredentials{}
	if accessKey, ok := d.GetOk(PfAwsAccessKey); ok {
		credentials.AWSAccessKey = accessKey.(string)
	}
	if secretKey, ok := d.GetOk(PfAwsSecretKey); ok {
		credentials.AWSSecretKey = secretKey.(string)
	}
	cpc.CloudCredentials = credentials

	_, err := c.UpdateCloudProviderCredential(cpc, cpcID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(cpcID)
	return diags
}

func resourceCloudProviderCredentialDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	cpcID := d.Id()

	resp, err := c.ReadCloudProviderCredential(cpcID)
	if err != nil {
		return diag.FromErr(err)
	}

	if !resp.IsUnused {
		return diag.Errorf(MessageCredentialsInUse)
	}

	_, err2 := c.DeleteCloudProviderCredential(cpcID)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(PfEmptyString)
	return diags
}

func extractCloudProviderCredentialsAws(d *schema.ResourceData) packetfabric.CloudProviderCredentialCreate {
	cpc := packetfabric.CloudProviderCredentialCreate{}
	cpc.CloudProvider = PfAws
	if description, ok := d.GetOk(PfDescription); ok {
		cpc.Description = description.(string)
	}
	cloudCredentials := packetfabric.CloudCredentials{}
	if awsAccessKey, ok := d.GetOk(PfAwsAccessKey); ok {
		cloudCredentials.AWSAccessKey = awsAccessKey.(string)
	}
	if awsSecretKey, ok := d.GetOk(PfAwsSecretKey); ok {
		cloudCredentials.AWSSecretKey = awsSecretKey.(string)
	}
	cpc.CloudCredentials = cloudCredentials
	return cpc
}
