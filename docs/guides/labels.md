---
page_title: "Service Labels"
subcategory: "Guides"
---

You can use labels to organize services. A service can have multiple labels. 

Note the following when writing label names:

* The name cannot include the following characters: `" ' ,` (spaces are allowed).
* The name cannot exceed 40 characters.
* Once created, you cannot edit the label name.

Once a label is added to a service, the label is listed next to the service in the PacketFabric portal. You can also see all members within a service using the portal. 

Labels are available in the PacketFabric portal from the admin menu in the upper right of the portal. For more information, see [Service Labels in the PacketFabric documentation](https://docs.packetfabric.com/admin/labels/).

For example:

```terraform
resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US"]
  labels   = ["terraform", "dev"]
}
```

Please note, when importing existing infrastructure into Terraform using `terraform import`, the labels attribute will not be imported. However, you can safely reapply labels using Terraform without causing service disruptions. Terraform will detect and resolve any discrepancies between your configuration and the actual state of your resources during the next terraform apply operation.