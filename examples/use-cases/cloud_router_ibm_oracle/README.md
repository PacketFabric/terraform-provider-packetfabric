# Use Case: PacketFabric Cloud Router with IBM and Oracle Cloud

This use case builds a PacketFabric Cloud Router between IBM Cloud Platform and Oracle CLoud.
Terraform providers used: PacketFabric, IBM and Oracle.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [IBM Cloud Terraform Provider](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest)
- [Oracle Cloud Terraform Provider](https://registry.terraform.io/providers/oracle/oci/latest)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

- "packetfabric_cloud_router"
- "packetfabric_cloud_router_connection_ibm"
- "ibm_dl_virtual_connection"
- oci_core_drg"
- "oci_core_virtual_circuit"
- "packetfabric_cloud_router_connection_oracle"
- "packetfabric_cloud_router_bgp_session"
- "ibm_resource_group"
- "ibm_is_vpc"
- "ibm_is_vpc_address_prefix"
- "ibm_is_subnet"
- "oci_identity_compartment"
- "oci_core_vcn"
- "random_pet"

**Estimated time:** ~15 min

## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an IBM Account? [Get Started](https://www.ibm.com/cloud/free)
- Don't have an Oracle Account? [Get Started](https://www.oracle.com/cloud/free/)

## Prerequisites

Make sure you have installed all of the following prerequisites on your machine:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Make sure you have the following items available:

- [IBM Credentials](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#environment-variables)
- [Oracle Credentials](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformproviderconfiguration.htm)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)

## Quick Start

1. Set PacketFabric API key and Account ID in environment variables and update each variables as needed (edit ``variables.tf``).

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

Set additional environment variables for Oracle and IBM:

```sh
### Oracle
# https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformproviderconfiguration.htm
export TF_VAR_tenancy_ocid="ocid1.tenancy.oc1..1234"
export TF_VAR_user_ocid="ocid1.user.oc1.1234"
export TF_VAR_fingerprint="AA:aa:a1:12:34:56..."
export TF_VAR_private_key="-----BEGIN PRIVATE KEY-----\nsecret\n-----END PRIVATE KEY-----"
export TF_VAR_parent_compartment_id="ocid1.tenancy.oc1.1234" # Parent comportment

### IBM
# https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#environment-variables
export PF_IBM_ACCOUNT_ID="123456789"
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
export TF_VAR_public_key="ssh-rsa AAAA...= user@mac.lan"
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
