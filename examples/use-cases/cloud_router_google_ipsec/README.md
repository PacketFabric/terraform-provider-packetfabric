# Use Case: PacketFabric Cloud Router with Google and a VPN Connection

This use case builds a PacketFabric Cloud Router VPN connection between Google Cloud Platform and a branch location, on-premises users, or a remote data center.
Terraform providers used: PacketFabric and Google.

![Deployment Diagram](./images/diagram_cloud_router_google_vpn.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Google Terraform Provider](https://registry.terraform.io/providers/hashicorp/google)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

- resource **"random_pet"**: Get a random pet name (use to name objects created)
- resource **"google_compute_network"**: Create a VPC
- resource **"google_compute_subnetwork"**: Create a subnet in the VPC
- resource & data source **"google_compute_router"**: Create a Google Cloud Router used for the Interconnect
- resource **"google_compute_interconnect_attachment"**: Create a Google Interconnect
- resource **"packetfabric_cloud_router"**: Create the Cloud Router in PacketFabric NaaS
- resource & data source **"packetfabric_google_cloud_router_connection"**: Add a Google Partner Interconnect to the Cloud Router
- resource & data source **"packetfabric_ipsec_cloud_router_connection"**: Add a VPN Connection to the Cloud Router
- module **"terraform-google-gcloud"**: Get the BGP Peer Addresses and set the PacketFabric Cloud Router ASN to the BGP settings in the Google Cloud Router
- resource **"packetfabric_cloud_router_bgp_session"**: Create BGP sessions in PacketFabric
- resource **"packetfabric_cloud_router_bgp_prefixes"**: Add BGP Prefixes to the BGP sessions in PacketFabric

**Estimated time:** ~5 min for Google & PacketFabric resources

**Note**: Because the BGP session is created automatically, we use gcloud terraform module to retreive the BGP addresses and set the PacketFabric Cloud Router ASN in the BGP settings in the Google Cloud Router. Please [vote](https://github.com/hashicorp/terraform-provider-google/issues/11458), [vote](https://github.com/hashicorp/terraform-provider-google/issues/12624) and [vote](https://github.com/hashicorp/terraform-provider-google/issues/12630) for these issues on GitHub.

## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an Google Account? [Get Started](https://cloud.google.com/free)

## Prerequisites

Make sure you have installed all of the following prerequisites on your machine:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)
- [gcloud](https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest)
- [jq](https://stedolan.github.io/jq/download/)

Make sure you have the following items available:

- [Google Service Account](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances)
- [IPsec information for the Site-to-Site VPN](https://docs.packetfabric.com/cr/vpn/)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)

## Quick Start

1. Create the file ``secret.tfvars`` and update each variables.

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

## Troubleshooting

1. In case you get the following error:

```
╷
│ Error: Error when reading or editing InterconnectAttachment: googleapi: Error 400: The resource 'projects/prefab-setting-123456/regions/us-west1/interconnectAttachments/demo-pf-gcp-vpn-master-cricket' is not ready, resourceNotReady
│ 
│ 
```

This seems to be a problem with Google Terraform Provider, run again the terraform destroy command and the destroy will complete correctly the 2nd try.
Please [vote](https://github.com/hashicorp/terraform-provider-google/issues/12631) for this issue on GitHub.

2. In case the gcloud_bgp_address fails, check the error, fix it and manually remove the state before re-running the terraform config.

```sh
terraform state rm module.gcloud_bgp_addresses
terraform state rm module.gcloud_bgp_peer_update
```

## Screenshots

Example Google Interconnect (VLAN attachment) in Google Cloud Console:

![VLAN attachment in Google Cloud Console](./images/google_interconnect.png)

Example Google Cloud Router in Google Cloud Console:

![VLAN attachment in Google Cloud Console](./images/google_cloud_router.png)
