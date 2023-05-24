# Use Case: PacketFabric Hosted cloud connection to IBM

This use case shows an example on how to use the PacketFabric & IBM Terraform providers 
to automate the creation of a Hosted Cloud Connection between PacketFabric and IBM in a Cloud On-Ramps facility.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Hosted IBM Connection](https://docs.packetfabric.com/cloud/ibm/hosted/create/)
- [IBM Direct Link Connect providers and locations](https://cloud.ibm.com/docs/dl?topic=dl-locations)
- [PacketFabric Cloud On-Ramps Locations](https://packetfabric.com/locations/cloud-on-ramps)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [IBM Cloud Terraform Provider](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources & data-sources used

- "random_pet"
- "ibm_resource_group"
- "ibm_is_vpc"
- "ibm_is_vpc_address_prefix"
- "ibm_is_subnet"
- "ibm_dl_virtual_connection"
- "packetfabric_cs_ibm_hosted_connection"
- "time_sleep"
- "ibm_dl_gateway"
- "ibm_dl_gateway_action"

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an IBM Account? [Get Started](https://www.ibm.com/cloud/free)

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Ensure you have the following items available:

- [IBM Credentials](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#environment-variables)
- [PacketFabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)

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


Set additional environment variables for IBM:

```sh
# https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#environment-variables
export PF_IBM_ACCOUNT_ID="123456789"
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
export TF_VAR_public_key="ssh-rsa AAAA..." # see link Create an SSH key pair in the pre-req
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
