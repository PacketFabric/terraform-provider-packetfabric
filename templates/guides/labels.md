---
page_title: "Label Resources"
subcategory: "Guides"
---

TBD

```terraform
resource "packetfabric_cloud_router" "awesome_cloud_routers" {
  provider     = packetfabric
  asn          = 4556
  name         = "Awesome Cloud Routers"
  capacity     = "5Gbps"
  regions      = ["US"]
  labels       = ["prod", "awesome"]
}
```
