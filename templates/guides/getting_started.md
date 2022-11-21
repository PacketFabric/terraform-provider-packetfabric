---
page_title: "Getting Started"
subcategory: "Guides"
---



## Setting up API keys

You must authenticate your Terraform actions with an API key. You can use the PacketFabric portal to create and manage keys. 

1. Click the account icon in the upper right (your username initials) and select **API keys**. 

2. Click **Add API key**. 

3. Provide a meaningful name for the key and, optionally, an expiration term for it. If you leave the expiration field empty, the key does not expire.

Once you create the key, you will only have one chance to copy it. The key is enabled until you delete it or it expires.

For more information, see [API Keys in the PacketFabric documentation](https://docs.packetfabric.com/admin/my_account/keys/).

->**Note:** API keys will still work for users with MFA enabled.

## Account ID

When you provision a new service through Terraform, you will need to provide an account ID. The account ID is mapped to a billing account. 


You can find this ID via the portal by navigating to **Billing > Accounts**. The ID is listed in the table. 

To find this information via the API, use the following call: `https://api.packetfabric.com/v2/contacts?billing=true`. This API call must be authorized.

For more information, see [Get the Account UUID in the PacketFabric documentation](https://docs.packetfabric.com/api/examples/account_uuid/).

Account ID can be provided by using the `PF_ACCOUNT_ID` environment variable.

For example:

$ export PF_ACCOUNT_ID="123456789"

## Getting location information

When provisioning new services, you may be required to provide a point of presence (POP), market, or region. For this, you can use the `packetfabric_locations` data source. See [packetfabric_locations (Data Source)](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/data-sources/packetfabric_locations).


However, the simplest way to find location information is from the [PacketFabric website](https://packetfabric.com/locations). 


If you prefer to use the API, the following GET call will return all PacketFabric locations: `https://api.packetfabric.com/v2.1/locations` 

This list is extensive, so you should add one or more location filters: `pop`, `city`, `state`, `market`, `region`. For example, to see all POPs in the the New York City metro area, you could use `https://api.packetfabric.com/v2.1/locations?market=nyc`.

->**Tip:** You can use `/v2/locations/markets` and `/v2/locations/regions` to get lists of our markets and regions.




## Getting port information

When provisioning new port-based services, you may need to specify information such as media and availability zone. 

Once you have decided on a POP, check port availability using the following GET call:

`https://api.packetfabric.com/v2/locations/{pop}/port-availability`

This returns information that includes the available media types and zones.

->**Note:** The `{pop}` parameter is case sensitive. For example, `https://api.packetfabric.com/v2/locations/DAL1/port-availability` returns information about the DAL1 pop, but `https://api.packetfabric.com/v2/locations/dal1/port-availability` returns an error stating the location can’t be found.


## Getting support

If you are experiencing issues with the PacketFabric Terraform provider, open an [issue](https://github.com/PacketFabric/terraform-provider-packetfabric/issues) via Github. 

If you are experiencing issues with your PacketFabric services, open a support ticket by emailing [support AT packetfabric.com](mailto:support AT packetfabric.com).