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
}

var testAccProviders = map[string]*schema.Provider{
	"packetfabric": Provider(),
}

var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"packetfabric": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}
