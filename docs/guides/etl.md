---
page_title: "Early Termination Liability (ETL) Fees"
subcategory: "Guides"
---

Early Termination Liability (ETL) fees are applied when a service is terminated before the end of its contract term. ETL is calculated based on the remaining days of the contract and is prorated accordingly. This guide will help you understand how ETL fees are managed in Terraform when working with PacketFabric resources.

For more information, see the [Canceling Services](https://docs.packetfabric.com/billing/services/cancel/).

## Understanding ETL in Terraform

In Terraform, you can create, modify, and delete resources using various providers. When working with PacketFabric resources, you may encounter situations where you need to delete a service that is still under contract. In such cases, ETL fees apply.

To help manage ETL fees in Terraform, the PacketFabric provider includes a computed field named `etl` for relevant resources. This field represents the current ETL fee for a resource and can be accessed in your Terraform configuration.

For example:

```terraform
resource "packetfabric_cloud_router" "cr1" {
  provider     = packetfabric
  asn          = 4556
  name         = "hello world"
  capacity     = "10Gbps"
  regions      = ["US", "UK"]
  etl          = 0  # This is the computed ETL field
}
```

The `etl` field will be updated during the terraform apply process, reflecting the current ETL fee for the resource.

## Monitoring ETL fees

It's important to monitor ETL fees and factor them into your infrastructure planning. You can use the PacketFabric portal or API to check the current ETL fees for your services. Additionally, you can access the etl field of resources in your Terraform configuration to keep track of ETL fees programmatically.

## Handling ETL fees when deleting resources

When deleting a PacketFabric resource in Terraform, it's crucial to consider the ETL fees associated with the resource. Be sure to review your contract terms and ETL fees before proceeding with the deletion.

!> **Warning:** You will have an hourly trial period for the first 24 hours after provisioning a service. This means that any order can be canceled in that 24-hour window, and you will only be charged for the hours in which the service was active. Any NRC (None-Recurrent Cost) charges are also dropped. The exception to this rule is cross connects, Colt ports, Colt VCs, VC marketplace and IX which will always trigger a full ETL fee. 

You can check the ETL fee for a resource by looking at its `etl` field. If the ETL fee is greater than 0, it indicates that there are remaining contract days, and terminating the service will incur the ETL fee. Ensure that you are aware of these fees and account for them in your budget before proceeding with the deletion.
