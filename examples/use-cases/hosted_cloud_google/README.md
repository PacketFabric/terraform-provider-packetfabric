# Use Case: PacketFabric Hosted cloud connection to Google Cloud

This use case shows an example on how to use the PacketFabric & Google Terraform providers 
to automate the creation of a Hosted Cloud Connection between PacketFabric and Google in a Cloud On-Ramps facility.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Google Partner Interconnect Process Overview](https://docs.packetfabric.com/cloud/google/hosted/process_overview/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Google Terraform Provider](https://registry.terraform.io/providers/hashicorp/google)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

- "random_pet"
- "google_compute_network"
- "google_compute_subnetwork"
- "google_compute_router"
- "google_compute_interconnect_attachment"
- "packetfabric_cs_google_hosted_connection"

**Estimated time:** ~5 min

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have a Google Account? [Get Started](https://cloud.google.com/free)

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Ensure you have the following items available:

- [Google Service Account](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)
- 1 [PacketFabric Port](https://docs.packetfabric.com/ports/) [cross connected](https://docs.packetfabric.com/xconnect/) to your network infrastructure  (update the ``pf_port_circuit_id`` in ``variables.tf``)

## Quick start

1. Set the PacketFabric API key and Account ID in the environment variables and update each variables as needed (edit ``variables.tf``).

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

Set additional environment variables for Google:

```sh
export TF_VAR_gcp_project_id="my-project-id"
export TF_VAR_gcp_credentials_path="/Users/myuser/secret.json"
```

2. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan
```

3. Apply the plan:

```sh
terraform apply
```

4. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy
```