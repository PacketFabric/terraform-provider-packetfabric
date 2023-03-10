package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudProviderCredentialAws() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudProviderCredentialAwsCreate,
		ReadContext:   resourceCloudProviderCredentialRead,
		UpdateContext: resourceCloudProviderCredentialAwsUpdate,
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
			"aws_access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_ACCESS_KEY_ID", nil),
				Description: "The AWS access key you want to save. " +
					"Can also be set with the AWS_ACCESS_KEY_ID environment variable.",
			},
			"aws_secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_SECRET_ACCESS_KEY", nil),
				Description: "The AWS secret key you want to save. " +
					"Can also be set with the AWS_SECRET_ACCESS_KEY environment variable.",
			},
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

	if d.HasChange("description") {
		cpc.Description = d.Get("description").(string)
	}
	credentials := packetfabric.CloudCredentials{}
	if accessKey, ok := d.GetOk("aws_access_key"); ok {
		credentials.AWSAccessKey = accessKey.(string)
	}
	if secretKey, ok := d.GetOk("aws_secret_key"); ok {
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
		return diag.Errorf("cannot delete cloud provider credential as it is currently in use")
	}

	_, err2 := c.DeleteCloudProviderCredential(cpcID)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId("")
	return diags
}

func extractCloudProviderCredentialsAws(d *schema.ResourceData) packetfabric.CloudProviderCredentialCreate {
	cpc := packetfabric.CloudProviderCredentialCreate{}
	cpc.CloudProvider = "aws"
	if description, ok := d.GetOk("description"); ok {
		cpc.Description = description.(string)
	}
	cloudCredentials := packetfabric.CloudCredentials{}
	if awsAccessKey, ok := d.GetOk("aws_access_key"); ok {
		cloudCredentials.AWSAccessKey = awsAccessKey.(string)
	}
	if awsSecretKey, ok := d.GetOk("aws_secret_key"); ok {
		cloudCredentials.AWSSecretKey = awsSecretKey.(string)
	}
	cpc.CloudCredentials = cloudCredentials
	return cpc
}
