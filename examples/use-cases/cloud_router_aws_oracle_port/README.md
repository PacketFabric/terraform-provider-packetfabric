# Use Case: PacketFabric Cloud Router with AWS, Oracle and a PacketFabric Port

This use case builds a PacketFabric Cloud Router between AWS, Oracle and a PacketFabric Port.
Terraform providers used: PacketFabric, AWS and Oracle. This example uses AWS Transit VIF & Gateway.

:rocket: You can simplify the configuration and management of PacketFabric Cloud Routers by utilizing the [PacketFabric Terraform Cloud Router Module](https://registry.terraform.io/modules/PacketFabric/cloud-router-module/connectivity/latest). This module provides pre-defined configurations and workflows for provisioning cloud routers on the PacketFabric platform.

![Deployment Diagram](./images/diagram_cloud_router_aws_oracle_port.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp AWS Terraform Provider](https://registry.terraform.io/providers/hashicorp/aws)
- [Oracle Cloud Terraform Provider](https://registry.terraform.io/providers/oracle/oci/latest)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources & data-sources used

- "aws_dx_gateway"
- "aws_dx_gateway_association"
- "aws_ec2_transit_gateway"
- "aws_ec2_transit_gateway_vpc_attachment"
- "aws_vpc"
- "aws_subnet"
- "aws_internet_gateway"
- "aws_route_table_association"
- "oci_identity_compartment"
- "oci_core_vcn"
- "oci_core_drg"
- "oci_core_fast_connect_provider_services"
- "oci_core_virtual_circuit"
- "packetfabric_cloud_router"
- "packetfabric_cloud_router_connection_aws"
- "time_sleep"
- "aws_dx_connection_confirmation"
- "packetfabric_port"
- "packetfabric_cloud_router_connection_oracle"
- "packetfabric_cloud_router_connection_port"
- "packetfabric_cloud_router_bgp_session"
- "random_pet"

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)
- Don't have an Oracle Account? [Get Started](https://www.oracle.com/cloud/free/)

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Ensure you have the following items available:

- [AWS Account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
- [AWS Access and Secret Keys](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html)
- [Oracle Credentials](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformproviderconfiguration.htm)
- [PacketFabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)
- Create an SSH key pair (Download [PuttyGen](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html) for Windows or use [ssh-keygen](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent?platform=mac) on Mac)

## Quick start

1. Set the PacketFabric API key and Account ID in your terminal as environment variables.

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

Windows PowerShell:
```powershell
PS C:\> $Env:PF_TOKEN="secret"
PS C:\> $Env:PF_ACCOUNT_ID="123456789"
```

Set additional environment variables for AWS and Oracle:

```sh
### AWS
export PF_AWS_ACCOUNT_ID="98765432"
export AWS_ACCESS_KEY_ID="ABCDEFGH"
export AWS_SECRET_ACCESS_KEY="secret"

### Oracle
# https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformproviderconfiguration.htm
export TF_VAR_tenancy_ocid="ocid1.tenancy.oc1..1234"
export TF_VAR_user_ocid="ocid1.user.oc1.1234"
export TF_VAR_fingerprint="AA:aa:a1:12:34:56..."
export TF_VAR_private_key="-----BEGIN PRIVATE KEY-----\nsecret\n-----END PRIVATE KEY-----"
export TF_VAR_parent_compartment_id="ocid1.tenancy.oc1.1234" # Parent comportment
```

2. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan
```

**Note:** you can update terraform variables in the ``variables.tf``.

3. Apply the plan:

```sh
terraform apply
```

4. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy
```