# Use Case: PacketFabric Cloud Router with Google and a VPN Connection

This use case builds a PacketFabric Cloud Router between Google Cloud Platform and a branch location, on-premises users, or a remote data center.
Terraform providers used: PacketFabric VPN connection and Google.

![Deployment Diagram](./images/diagram_cloud_router_google_vpn.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Google Terraform Provider](https://registry.terraform.io/providers/hashicorp/google)
<!-- - [Google Cloud CLI Terraform Module](https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest) -->
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources & data-sources used

- "packetfabric_cloud_router"
- "google_compute_router"
<!-- - "google_compute_interconnect_attachment" -->
- "packetfabric_cloud_router_connection_google"
- "packetfabric_cloud_router_connection_ipsec"
- "packetfabric_cloud_router_bgp_session"
- "google_compute_firewall"
- "google_compute_instance"
- "google_compute_network"
- "google_compute_subnetwork"
- "random_pet"

**Estimated time:** ~5 min for Google & PacketFabric resources
<!-- 
**Note**: Because the BGP session is created automatically, we use gcloud terraform module to retreive the BGP addresses and set the PacketFabric Cloud Router ASN in the BGP settings in the Google Cloud Router. Please [vote](https://github.com/hashicorp/terraform-provider-google/issues/11458), [vote](https://github.com/hashicorp/terraform-provider-google/issues/12624) and [vote](https://github.com/hashicorp/terraform-provider-google/issues/12630) for these issues on GitHub. -->

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have a Google Account? [Get Started](https://cloud.google.com/free)

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)
<!-- - [gcloud](https://cloud.google.com/sdk/docs/install) -->
<!-- - [jq](https://stedolan.github.io/jq/download/) -->

Ensure you have the following items available:

- [Google Service Account](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances)
- [IPsec information for the Site-to-Site VPN](https://docs.packetfabric.com/cr/vpn/)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
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

Set additional environment variables for Google:

```sh
### Google
export TF_VAR_gcp_project_id="my-project-id" # used for bash script used with gcloud module
export GOOGLE_CREDENTIALS='{ "type": "service_account", "project_id": "demo-setting-1234", "private_key_id": "1234", "private_key": "-----BEGIN PRIVATE KEY-----\nsecret\n-----END PRIVATE KEY-----\n", "client_email": "demoapi@demo-setting-1234.iam.gserviceaccount.com", "client_id": "102640829015169383380", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "token_uri": "https://oauth2.googleapis.com/token", "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/demoapi%40demo-setting-1234.iam.gserviceaccount.com" }'

export TF_VAR_public_key="ssh-rsa AAAA...= user@mac.lan"
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

**Note:** Default login/password for Locust is ``demo:packetfabric`` edit ``user-data-ubuntu.sh`` script to change it.

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
<!-- 
2. In case the ``gcloud_bgp_address`` module fails, check the error, fix it and manually remove the state before re-running the terraform config.

```sh
terraform state rm module.gcloud_bgp_addresses
terraform state rm module.gcloud_bgp_peer_update
``` -->

## Screenshots

Example Google Interconnect (VLAN attachment) in Google Cloud Console:

![VLAN attachment in Google Cloud Console](./images/google_interconnect.png)

Example Google Cloud Router in Google Cloud Console:

![VLAN attachment in Google Cloud Console](./images/google_cloud_router.png)
