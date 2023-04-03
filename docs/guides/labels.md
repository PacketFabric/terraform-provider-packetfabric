---
page_title: "Labels"
subcategory: "Guides"
---

You can use labels to organize services.

Labels are available in the PacketFabric Portal from the admin menu in the upper right of the portal. For more information, see [Service Labels in the PacketFabric documentation](https://docs.packetfabric.com/admin/labels/).

For example:

```terraform
resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = ["terraform", "dev"]
}
```