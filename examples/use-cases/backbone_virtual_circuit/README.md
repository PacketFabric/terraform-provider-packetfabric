# Use Case: 

This use case shows an example on how to use the PacketFabric Terraform provider 
to automate the creation of 2 ports in PacketFabric, a Backbone Virtual Circuit between the 2 ports and 
the outbound Cross connect for those 2 ports.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [PacketFabric Ports/Interfaces Overview](https://docs.packetfabric.com/ports/)
- [PacketFabric Cross Connects Overview](https://docs.packetfabric.com/xconnect/)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)
- [Best practices for using Terraform](https://cloud.google.com/docs/terraform/best-practices-for-terraform)
- [Automate Terraform with GitHub Actions](https://learn.hashicorp.com/tutorials/terraform/github-actions?in=terraform/automation)

## Terraform resources deployed

- resource **"random_pet"**: Get a random pet name (use to name objects created)
- resource **"packetfabric_interface"**: Create 2 ports in 2 PacketFabric Point of Presence (PoP)
- data source **"packetfabric_billing"**: Get the billing details for those 2 ports
- data source **"packetfabric_locations"**: Get PacketFabric available locations
- resource **"packetfabric_outbound_cross_connect"**: Customer Inbound/PacketFabric Outbound Cross Connect (You provide PacketFabric with an LOA/CFA authorizing us to make the connection)
- resource **"packetfabric_create_backbone_virtual_circuit"**: Create a Backbone Virtual Circuit between the 2 ports


## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)

## Prerequisites

Make sure you have installed all of the following prerequisites on your machine:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Make sure you have the following items available:

- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)

## Quick Start

1. Create the file ``secret.tfvars`` and update each variables.

```sh
cp secret.tfvars.sample secret.tfvars
```

2. Create resources 
```sh
terraform init
terraform plan -var-file="secret.tfvars"
```

Apply the plan:

```sh
terraform apply -var-file="secret.tfvars"
```

3. Cleanup/Remove all in both PacketFabric.

```sh
terraform destroy -var-file="secret.tfvars"
```
