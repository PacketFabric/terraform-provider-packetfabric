---
page_title: "Importing Resources"
subcategory: "Guides"
---


You can import certain resources that were initially created outside the Terraform provider (for example, services provisioned through the portal).


Once imported, you can manage these resources through Terraform. 


The following example walks you through importing a Cloud Router. 

->**Note:** At this time, only certain resources can be imported. Check the documentation for that resource; importable resources have an "Import" section with an applicable example.


### 1: Add the resource you want to import to main.tf 

```terraform
resource "packetfabric_cloud_router" "awesome_cloud_routers" {
  provider     = packetfabric
  asn          = 4556
  name         = "Awesome Cloud Routers"
  capacity     = "5Gbps"
  regions      = ["US"]
}
```

### 2: Run the 'import' command

In this step, you specify the resource type, its name, and circuit ID. 

```bash
$ terraform import packetfabric_cloud_router.awesome_cloud_routers PF-L3-CUST-1700239 
```


### 3: Confirm the resource is now managed by Terraform

```bash
$ terraform state list 
packetfabric_cloud_router.awesome_cloud_routers

$ terraform state show packetfabric_cloud_router.awesome_cloud_routers
# packetfabric_cloud_router.awesome_cloud_routers:
resource "packetfabric_cloud_router" "awesome_cloud_routers" {
    account_uuid = "a2115890-ed02-4795-a6dd-c485bec3529c"
    asn          = 4556
    capacity     = "5Gbps"
    id           = "PF-L3-CUST-1700239"
    name         = "Awesome Cloud Routers"
    regions      = [
        "US",
    ]

    timeouts {}
}
```

### 4: Run Terraform Plan

```bash
$ terraform plan

packetfabric_cloud_router.awesome_cloud_routers: Refreshing state... [id=PF-L3-CUST-1700239]

No changes. Your infrastructure matches the configuration.

Terraform has compared your real infrastructure against your configuration and found no differences, so no changes are needed.
```

You may need to adjust your resource defintion in your Terraform HCL code or Terraform state file based on the output.