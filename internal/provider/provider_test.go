package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

const testPrefix = "tf-acc"

var testAccProviders = map[string]*schema.Provider{
	"packetfabric": Provider(),
}
