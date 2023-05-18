# Use Case: PacketFabric Cloud Router with Google and a VPN Connection

This use case builds a PacketFabric Cloud Router between Google Cloud Platform and a branch location, on-premises users, or a remote data center.
Terraform providers used: PacketFabric IPsec connection and Google.

:rocket: You can simplify the configuration and management of PacketFabric Cloud Routers by utilizing the [PacketFabric Terraform Cloud Router Module](https://registry.terraform.io/modules/PacketFabric/cloud-router-module/connectivity/latest). This module provides pre-defined configurations and workflows for provisioning cloud routers on the PacketFabric platform.

![Deployment Diagram](./images/diagram_cloud_router_google_vpn.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Google Terraform Provider](https://registry.terraform.io/providers/hashicorp/google)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources & data-sources used

- "packetfabric_cloud_router"
- "google_compute_router"
- "packetfabric_cloud_router_connection_google"
- "packetfabric_cloud_router_connection_ipsec"
- "packetfabric_cloud_router_bgp_session"
- "google_compute_firewall"
- "google_compute_instance"
- "google_compute_network"
- "google_compute_subnetwork"
- "random_pet"

**Estimated time:** ~5 min for Google & PacketFabric resources

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
- [IPsec information for the Site-to-Site VPN](https://docs.packetfabric.com/cr/vpn/)
- [PacketFabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)
- [Create an SSH key pair](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-key-pairs.html)

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

export TF_VAR_public_key="ssh-rsa AAAA...= user@mac.lan" # see link Create an SSH key pair in the pre-req
```

**Note**: To convert a pretty-printed JSON into a single line JSON string: `jq -c '.' google_credentials.json`.

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

In case you get the following error:

```
╷
│ Error: Error when reading or editing InterconnectAttachment: googleapi: Error 400: The resource 'projects/prefab-setting-123456/regions/us-west1/interconnectAttachments/demo-pf-gcp-vpn-master-cricket' is not ready, resourceNotReady
│ 
│ 
```

This seems to be a problem with Google Terraform Provider, run again the terraform destroy command and the destroy will complete correctly the 2nd try.
Please [vote](https://github.com/hashicorp/terraform-provider-google/issues/12631) for this issue on GitHub.

## Screenshots

Example Google Interconnect (VLAN attachment) in Google Cloud Console:

![VLAN attachment in Google Cloud Console](./images/google_interconnect.png)

Example Google Cloud Router in Google Cloud Console:

![VLAN attachment in Google Cloud Console](./images/google_cloud_router.png)
