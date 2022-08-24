# Use Case: PacketFabric’s Hosted cloud connection to AWS

This use case shows an example on how to use the PacketFabric & AWS Terraform providers 
to automate the creation of a Hosted Cloud Connection between PacketFabric and AWS in a Cloud On-Ramps facility.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Hosted AWS Connection](https://docs.packetfabric.com/cloud/aws/hosted/create/)
- [AWS Direct Connect Locations](https://aws.amazon.com/directconnect/locations/)
- [PacketFabric Cloud On-Ramps Locations](https://packetfabric.com/locations/cloud-on-ramps)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp AWS Terraform Provider](https://registry.terraform.io/providers/hashicorp/aws)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)
- [Best practices for using Terraform](https://cloud.google.com/docs/terraform/best-practices-for-terraform)
- [Automate Terraform with GitHub Actions](https://learn.hashicorp.com/tutorials/terraform/github-actions?in=terraform/automation)

## Terraform resources deployed

- resource **"random_pet"**: Get a random pet name (use to name objects created)
- resource **"aws_vpc"**: Create a VPC
- resource **"aws_subnet"**: Create subnet in the VPC
- resource **"aws_internet_gateway"**: Create internet gateway (used to access future EC2 instances)
- resource **"aws_vpn_gateway"**: Create Virtual Private Gateway (or Private VIF - Virtual Interface)
- resource **"aws_route_table"**: Create route table for the VPCs
- resource **"aws_route_table_association"**: Associate Route Table to the VPCs subnets
- resource & data source **"packetfabric_cs_aws_hosted_connection"**: Create a AWS Hosted Cloud Connection 
- resource **"time_sleep" "wait_60_seconds"**: Wait few seconds for the Connections to appear on AWS side
- data source **"aws_dx_connection"**: Retreive Direct Connect Connection details
- resource **"aws_dx_connection_confirmation"**: Accept the connections coming from PacketFabric
<!--  - resource **"aws_dx_gateway"**: Create Direct Connect Gateways -->
<!--  - resource **"aws_dx_private_virtual_interface"**: Create Direct Connect Private Virtual interfaces -->
<!--  - resource **"aws_dx_gateway_association"**: Associates a Direct Connect Gateway with a Virtual Private Gateways (VPG)  -->

**Estimated time:** ~15 min for AWS & PacketFabric resources + ~10-15 min for AWS Direct Connect Gateway association with AWS Virtual Private Gateways

## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)
    - Permissions required: VPC, EC2, Direct Connect

## Prerequisites

Make sure you have installed all of the following prerequisites on your machine:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Make sure you have the following items available:

- [AWS Account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
- [AWS Access and Secret Keys](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)
- 1 [PacketFabric Port](https://docs.packetfabric.com/ports/) [cross connected](https://docs.packetfabric.com/xconnect/) to your network infrastructure (update the ``pf_port_circuit_id`` in ``variables.tf``)

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

3. Cleanup/Remove all in both PacketFabric and AWS.

```sh
terraform destroy -var-file="secret.tfvars"
```
