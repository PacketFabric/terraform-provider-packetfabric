# Use Case: 

This use case shows an example on how to use the PacketFabric Terraform provider 
to automate the creation of a mesh of Backbone Virtual Circuits between 6 ports.

![Deployment Diagram](./images/diagram_mesh.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

- "random_pet"
- "packetfabric_backbone_virtual_circuit"

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

1. Set PacketFabric API key and Account ID in environment variables and update each variables as needed (edit ``variables.tf``).

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

2. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan
```

Apply the plan:

```sh
terraform apply
```

3. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy
```
