---
page_title: "Getting Started"
subcategory: "Guides"
---

## Video demo

The following video provides a walk-through on how to get started using the PacketFabric Terraform provider: [YouTube - Get Started with Terraform | PacketFabric](https://www.youtube.com/watch?v=UnK9jslY3jg)


## Setting up API keys

You must authenticate your Terraform actions with an API key. You can use the PacketFabric portal to create and manage keys. 

1. Click the account icon in the upper right (your username initials) and select **API keys**. 

2. Click **Add API key**. 

3. Provide a meaningful name for the key and, optionally, an expiration term for it. If you leave the expiration field empty, the key does not expire.

Once you create the key, you will only have one chance to copy it. The key is enabled until you delete it or it expires.

For more information, see [API Keys in the PacketFabric documentation](https://docs.packetfabric.com/admin/my_account/keys/).

->**Note:** API keys will still work for users with MFA enabled.

API Key can be provided by using the `PF_TOKEN` environment variable.

For example:

```sh
$ export PF_TOKEN="secret"
```

## Account ID

When you provision a new service through Terraform, you will need to provide an account ID. The account ID is mapped to a billing account. 


You can find this ID via the portal by navigating to **Billing > Accounts**. The ID is listed in the table. 

To find this information via the API, use the following call: `https://api.packetfabric.com/v2/contacts?billing=true`. This API call must be authorized.

For more information, see [Get the Account UUID in the PacketFabric documentation](https://docs.packetfabric.com/api/examples/account_uuid/).

Account ID can be provided by using the `PF_ACCOUNT_ID` environment variable.

For example:

```sh
$ export PF_ACCOUNT_ID="123456789"
```

## Getting location information

When provisioning new services, you may be required to provide a point of presence (POP), market, or region. The simplest way to find location information is from the [PacketFabric website](https://packetfabric.com/locations). 

#### API requests

If you prefer to use the API, the following GET call will return all PacketFabric locations: `https://api.packetfabric.com/v2.1/locations` 

This list is extensive, so you should add one or more location filters: `pop`, `city`, `state`, `market`, `region`. For example, to see all POPs in the the New York City metro area, you could use `https://api.packetfabric.com/v2.1/locations?market=nyc`.

->**Tip:** You can use `/v2/locations/markets` and `/v2/locations/regions` to get lists of our markets and regions.

#### Terraform data sources

If you want to stay within Terraform, there are several data sources that you can use to gather location information:

*`[packetfabric_locations (Data Source)](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/data-sources/packetfabric_locations) - This returns the full PacketFabric location list with all the attributes for each location. 
*`[packetfabric_locations_markets (Data Source)](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/data-sources/packetfabric_locations_markets) - This returns a list of PacketFabric markets, including the market code to use for filtering. 
*`[packetfabric_locations_regions (Data Source)](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/data-sources/packetfabric_locations_regions) - This returns a list of PacketFabric regions, including the region code to use for filtering. 
*`[packetfabric_locations_regions (Data Source)](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/data-sources/packetfabric_locations_cloud) - This returns a list of PacketFabric cloud on-ramps.



## Getting port information

When provisioning new port-based services, you may need to specify information such as media and availability zone. 

#### API requests

Once you have decided on a POP, you can check port availability using the following GET call:

`https://api.packetfabric.com/v2/locations/{pop}/port-availability`

 This returns information about includes the available media types, speeds, and zones.

->**Note:** The `{pop}` parameter is case sensitive. For example, `https://api.packetfabric.com/v2/locations/DAL1/port-availability` returns information about the DAL1 pop, but `https://api.packetfabric.com/v2/locations/dal1/port-availability` returns an error stating the location can’t be found.


#### Terraform data sources

If you want to stay within Terraform, you can use the following data sources:

*`[packetfabric_locations_port_availability (Data Source)](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/data-sources/packetfabric_locations_port_availability) - This returns information about includes the available media types, speeds, and zones within a given POP.
*`[packetfabric_locations_pop_zones (Data Source)](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/data-sources/packetfabric_locations_pop_zones) - This returns a list of availability zones within the given POP.


## Getting support

If you are experiencing issues with the PacketFabric Terraform provider, open an [issue](https://github.com/PacketFabric/terraform-provider-packetfabric/issues) via Github. 

If you are experiencing issues with your PacketFabric services, open a support ticket by emailing [support AT packetfabric.com](mailto:support AT packetfabric.com).