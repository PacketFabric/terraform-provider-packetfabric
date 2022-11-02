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

- resource **"random_pet"**: Get a random pet name (use to name objects created)
- resource **"ibm_resource_group"**: Create an IBM resource group
- resource **"ibm_is_vpc"**: Create an IBM VPC
- resource **"ibm_is_vpc_address_prefix"**: Create an IP address prefix.
- resource **"ibm_is_subnet"**: Create an IBM Subnet in the VPC
- resource **"oci_identity_compartment"**: Create an Oracle Compartment resource
- resource **"oci_core_vcn"**: Create a new Virtual Cloud Network (VCN) in Oracle
- resource **"oci_core_drg"**: Create an Oracle Dynamic Routing Gateway
- resource **"oci_core_virtual_circuit""**: Create an Oracle FastConnect Connection
- resource **"packetfabric_cloud_router"**: Create the Cloud Router in PacketFabric NaaS
- resource & data source **"packetfabric_cloud_router_connection_ibm"**: Add an IBM Partner Connection to the Cloud Router
- resource & data source **"packetfabric_cloud_router_connection_oracle"**: Add an Oracle FastConnect to the Cloud Router
- resource **"packetfabric_cloud_router_bgp_session"**: Create BGP sessions in PacketFabric

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

1. Create the file ``secret.tfvars`` and update each variables as needed.

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
