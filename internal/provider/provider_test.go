package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccExternalProviders = map[string]resource.ExternalProvider{
	"time": {
		VersionConstraint: "0.9.1",
		Source:            "hashicorp/time",
	},
	"google": {
		VersionConstraint: "4.61.0",
		Source:            "hashicorp/google",
	},
<<<<<<< HEAD
	"azurerm": {
		VersionConstraint: "3.56.0",
		Source:            "hashicorp/azurerm",
		Components: []resource.ExternalComponent{
			{
				ProviderName: "azurerm",
				Config: map[string]interface{}{
					"features": []map[string]interface{}{
						{
							"resource_group": []map[string]interface{}{
								{
									"prevent_deletion_if_contains_resources": false,
								},
							},
						},
					},
				},
			},
	},
	"ibm": {
		VersionConstraint: "1.53.0",
		Source:            "IBM-Cloud/ibm",
	},
	"oci": {
		VersionConstraint: "4.111.0",
		Source:            "oracle/oci",
	},
=======
>>>>>>> main
}

var testAccProviders = map[string]*schema.Provider{
	"packetfabric": Provider(),
}

var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"packetfabric": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}
