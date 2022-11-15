# Use Case: 

This use case shows an example on how to use the PacketFabric Terraform provider 
to automate the creation of 2 ports in PacketFabric, a Backbone Virtual Circuit between the 2 ports and 
the outbound Cross connect for those 2 ports.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [PacketFabric Ports Overview](https://docs.packetfabric.com/ports/)
- [PacketFabric Virtual Circuits Overview](https://docs.packetfabric.com/vc/)
- [PacketFabric Cross Connects Overview](https://docs.packetfabric.com/xconnect/)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

- "random_pet"
- "packetfabric_port"
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

1. Create the file ``secret.tfvars`` and update each variables as needed (edit ``variables.tf``).

```sh
cp secret.tfvars.sample secret.tfvars
```

2. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan -var-file="secret.tfvars"
```

Apply the plan:

```sh
terraform apply -var-file="secret.tfvars"
```

3. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy -var-file="secret.tfvars"
```
