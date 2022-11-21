# Use Case: PacketFabric Hosted cloud connection to AWS

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

## Terraform resources deployed

- "random_pet"
- "aws_vpc"
- "aws_subnet"
- "aws_internet_gateway"
- "aws_vpn_gateway"
- "aws_vpn_gateway_attachment"
- "aws_route_table"
- "aws_route_table_association"
- "packetfabric_cs_aws_hosted_connection"
- "time_sleep"
- "aws_dx_connection_confirmation"
- "aws_dx_gateway"
- "aws_dx_private_virtual_interface"

**Estimated time:** ~15 min for AWS & PacketFabric resources + ~10-15 min for AWS Direct Connect Gateway association with AWS Virtual Private Gateways

## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)

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

1. Set PacketFabric API key and Account ID in environment variables and update each variables as needed (edit ``variables.tf``).

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

Set additional environment variables for AWS:

```sh
export TF_VAR_pf_aws_account_id="123456789"
export AWS_ACCESS_KEY_ID = "ABCDEFGH"
export AWS_SECRET_ACCESS_KEY = "secret"
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
