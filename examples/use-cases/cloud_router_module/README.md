# Use Case: Custom PacketFabric Cloud Router Terraform Module with AWS and Google

This use case builds a PacketFabric Cloud Router between AWS and Google Cloud Platform using a custom Terraform module.

:rocket: You can simplify the configuration and management of PacketFabric Cloud Routers by utilizing the [PacketFabric Terraform Cloud Router Module](https://registry.terraform.io/modules/PacketFabric/cloud-router-module/connectivity/latest). This module provides pre-defined configurations and workflows for provisioning cloud routers on the PacketFabric platform.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [Creating Modules in Terraform](https://developer.hashicorp.com/terraform/language/modules/develop)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources & data-sources used

- "packetfabric_cloud_router"
- "packetfabric_locations_pop_zones"
- "packetfabric_cloud_router_connection_aws"
- "packetfabric_cloud_router_connection_google"
- "packetfabric_cloud_router_bgp_session"
- "random_pet"

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)
- Don't have a Google Account? [Get Started](https://cloud.google.com/free)

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Ensure you have the following items available:

- [AWS Account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
- [PacketFabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)

## Quick start

1. Set the PacketFabric API key, Account ID, and AWS Account ID in your terminal as environment variables.

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
export PF_AWS_ACCOUNT_ID="123456789"
```

Windows PowerShell:
```powershell
PS C:\> $Env:PF_TOKEN="secret"
PS C:\> $Env:PF_ACCOUNT_ID="123456789"
PS C:\> $Env:PF_AWS_ACCOUNT_ID="123456789"
```

2. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan --var-file=tfvars_dev.json
```

3. Apply the plan:

```sh
terraform apply --var-file=tfvars_dev.json
```

**Note:** you can update terraform variables in the ``tfvars_dev.json``.

4. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy --var-file=tfvars_dev.json
```
