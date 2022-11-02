# Use Case: 

This use case shows an example on how to use the PacketFabric Terraform provider 
to automate the creation of a Virtual Circuit between 2 ports owned by 2 different PacketFabric accounts (A and Z).

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [PacketFabric Ports Overview](https://docs.packetfabric.com/ports/)
- [PacketFabric Virtual Circuits Overview](https://docs.packetfabric.com/vc/)
- [PacketFabric Marketplace Overview](https://docs.packetfabric.com/eco/overview/)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

**A side**

- resource **"random_pet"**: Get a random pet name (use to name objects created)
- resource **"packetfabric_backbone_virtual_circuit_marketplace"**: Create the Backbone Virtual Circuit Request to the Z side
- resource **"packetfabric_backbone_virtual_circuit"**: Once the Request has been accepted, import and manage the new resource in Terraform

**Z side**

- resource **"random_pet"**: Get a random pet name (use to name objects created)
- resource **"packetfabric_marketplace_service_accept_request"**: Accept the Backbone Virtual Circuit Request from the A side
- resource **"packetfabric_marketplace_service_reject_request"**: Reject the Backbone Virtual Circuit Request from the A side

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

1. Create the file ``secret.tfvars`` and update each variables as needed for each A and Z sides.

**A side**

```sh
cp secret.tfvars.sample secret.tfvars
```

**B side**

```sh
cp secret.tfvars.sample secret.tfvars
```

2. Initialize Terraform, create an execution plan and execute the plan.

**A side** in `a_side` folder:

```sh
terraform init
terraform plan -var-file="secret.tfvars"
```

Apply the plan

```sh
terraform apply -var-file="secret.tfvars"
```

**B side** in `a_side` folder:

```sh
terraform init
terraform plan -var-file="secret.tfvars"
```

Update the `pf_a_side_vc_request_uuid` with the **A Side** Virtual Circuit Request UUID.
You can either Accept or Reject the request (comment/comment out as desire).

Apply the plan

```sh
terraform apply -var-file="secret.tfvars"
```

3. Destroy all remote objects managed by the Terraform configuration on both sides.

```sh
terraform destroy -var-file="secret.tfvars"
```
