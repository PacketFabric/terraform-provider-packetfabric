# Use Case: PacketFabric Cloud Router Terraform Module with AWS and Google

This use case builds a PacketFabric Cloud Router between AWS and Google Cloud Platform using a custom Terraform module.
Terraform providers used: PacketFabric.


## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [Creating Modules in Terraform](https://developer.hashicorp.com/terraform/language/modules/develop)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

- "packetfabric_cloud_router"
- "packetfabric_cloud_router_connection_aws"
- "packetfabric_cloud_router_connection_google"
- "packetfabric_cloud_router_bgp_session"
- "random_pet"

## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)
- Don't have an Google Account? [Get Started](https://cloud.google.com/free)

## Prerequisites

Make sure you have installed all of the following prerequisites on your machine:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Make sure you have the following items available:

- [AWS Account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)

## Quick Start

1. Set PacketFabric API key, Account ID and AWS Account ID in environment variables and update each variables as needed (edit ``tfvars.json``).

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
export PF_AWS_ACCOUNT_ID="123456789"
```

2. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan --var-file=tfvars.json
```

Apply the plan:

```sh
terraform apply --var-file=tfvars.json
```

3. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy --var-file=tfvars.json
```
