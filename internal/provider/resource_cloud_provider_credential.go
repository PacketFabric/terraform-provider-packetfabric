package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudProviderCredential() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudProviderCredentialCreate,
		ReadContext:   resourceCloudProviderCredentialRead,
		UpdateContext: resourceCloudProviderCredentialUpdate,
		DeleteContext: resourceCloudProviderCredentialDelete,
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
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Description of the Cloud Provider Credentials.",
			},
			"cloud_provider": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"aws", "google"}, true),
				Description:  "The cloud provider of this cloud provider credential.\n\n\tEnum: [\"aws\" or \"google\"]",
			},
			"aws_access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PF_AWS_ACCESS_KEY_ID", nil),
				Description: "The AWS access key you want to save. " +
					"Can also be set with the PF_AWS_ACCESS_KEY_ID environment variable.",
			},
			"aws_secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PF_AWS_SECRET_ACCESS_KEY", nil),
				Description: "The AWS secret key you want to save. " +
					"Can also be set with the PF_AWS_SECRET_ACCESS_KEY environment variable.",
			},
			"service_account": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PF_GOOGLE_CREDENTIALS", nil),
				Description: "The Google service account JSON you want to save. " +
					"Can also be set with the PF_GOOGLE_CREDENTIALS environment variable.",
			},
		},
	}
}

func resourceCloudProviderCredentialCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	cpc := extractCloudProviderCredentials(d)

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

func resourceCloudProviderCredentialUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	cpcID := d.Id()

	cpc := packetfabric.CloudProviderCredentialUpdate{}

	if d.HasChange("description") {
		cpc.Description = d.Get("description").(string)
	}

	credentials := packetfabric.CloudCredentials{}
	if cloudProvider, ok := d.GetOk("cloud_provider"); ok {
		if cloudProvider == "aws" {
			if accessKey, ok := d.GetOk("aws_access_key"); ok {
				credentials.AWSAccessKey = accessKey.(string)
			}
			if secretKey, ok := d.GetOk("aws_secret_key"); ok {
				credentials.AWSSecretKey = secretKey.(string)
			}
		} else if cloudProvider == "google" {
			if serviceAccount, ok := d.GetOk("service_account"); ok {
				credentials.ServiceAccount = serviceAccount.(string)
			}
		}
	}
	cpc.CloudCredentials = credentials

	_, err := c.UpdateCloudProviderCredential(cpc, cpcID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(cpcID)
	return diags
}

func extractCloudProviderCredentials(d *schema.ResourceData) packetfabric.CloudProviderCredentialCreate {
	cpc := packetfabric.CloudProviderCredentialCreate{}
	if description, ok := d.GetOk("description"); ok {
		cpc.Description = description.(string)
	}
	if cloudProvider, ok := d.GetOk("cloud_provider"); ok {
		cpc.CloudProvider = cloudProvider.(string)
	}
	cloudCredentials := packetfabric.CloudCredentials{}
	if cpc.CloudProvider == "aws" {
		if awsAccessKey, ok := d.GetOk("aws_access_key"); ok {
			cloudCredentials.AWSAccessKey = awsAccessKey.(string)
		}
		if awsSecretKey, ok := d.GetOk("aws_secret_key"); ok {
			cloudCredentials.AWSSecretKey = awsSecretKey.(string)
		}
	}
	if cpc.CloudProvider == "google" {
		if serviceAccount, ok := d.GetOk("service_account"); ok {
			cloudCredentials.ServiceAccount = serviceAccount.(string)
		}
	}
	cpc.CloudCredentials = cloudCredentials
	return cpc
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
