package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccExternalProviders = map[string]resource.ExternalProvider{
	"time": {
		Source:            "hashicorp/time",
		VersionConstraint: "0.9.1",
	},
	"google": {
		Source:            "hashicorp/google",
		VersionConstraint: "4.61.0",
	},
	"azurerm": {
		Source:            "hashicorp/azurerm",
		VersionConstraint: "3.56.0",
	},
	"ibm": {
		Source:            "IBM-Cloud/ibm",
		VersionConstraint: "1.53.0",
	},
	"oci": {
		Source:            "oracle/oci",
		VersionConstraint: "4.111.0",
	},
}

var testAccProviders = map[string]*schema.Provider{
	"packetfabric": Provider(),
}

var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"packetfabric": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}
