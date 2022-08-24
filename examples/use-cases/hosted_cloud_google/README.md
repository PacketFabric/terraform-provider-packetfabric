# Use Case: PacketFabric’s Hosted cloud connection to Google Cloud

This use case shows an example on how to use the PacketFabric & Google Terraform providers 
to automate the creation of a Hosted Cloud Connection between PacketFabric and Google in a Cloud On-Ramps facility.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Google Partner Interconnect Process Overview](https://docs.packetfabric.com/cloud/google/hosted/process_overview/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Microsoft Google Terraform Provider](https://registry.terraform.io/providers/hashicorp/google)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)
- [Best practices for using Terraform](https://cloud.google.com/docs/terraform/best-practices-for-terraform)
- [Automate Terraform with GitHub Actions](https://learn.hashicorp.com/tutorials/terraform/github-actions?in=terraform/automation)

## Terraform resources deployed

- resource **"random_pet"**: Get a random pet name (use to name objects created)
- resource **"google_compute_network"**: Create a VPC
- resource **"google_compute_subnetwork"**: Create a subnet in the VPC
- resource & data source **"google_compute_router"**: Create a Google Cloud Router used for the Interconnect
- resource **"google_compute_interconnect_attachment"**: Create a Google Interconnect
- resource **"packetfabric_cs_google_hosted_connection"**: Create a Google Hosted Cloud Connection 

**Estimated time:** ~5 min

## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an Google Account? [Get Started](https://cloud.google.com/free)

## Prerequisites

Make sure you have installed all of the following prerequisites on your machine:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Make sure you have the following items available:

- [Google Service Account](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)
- 1 [PacketFabric Port](https://docs.packetfabric.com/ports/) [cross connected](https://docs.packetfabric.com/xconnect/) to your network infrastructure  (update the ``pf_port_circuit_id`` in ``variables.tf``)

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

3. Cleanup/Remove all in both PacketFabric and Google.

```sh
terraform destroy -var-file="secret.tfvars"
```
