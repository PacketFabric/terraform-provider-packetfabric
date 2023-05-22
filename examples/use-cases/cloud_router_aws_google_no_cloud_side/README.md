# Use Case: PacketFabric Cloud Router with AWS and Google - No Cloud Side Provisioning

This use case builds a PacketFabric Cloud Router between AWS and Google Cloud Platform.
Terraform providers used: PacketFabric, AWS and Google. This example uses AWS Transit VIF & Gateway.

**Note:** This example does not utilize PacketFabric's Cloud Side provisioning feature. For an example that demonstrates the use of the Cloud Side provisioning feature, please refer to this [alternate example](../cloud_router_aws_google).

:rocket: You can simplify the configuration and management of PacketFabric Cloud Routers by utilizing the [PacketFabric Terraform Cloud Router Module](https://registry.terraform.io/modules/PacketFabric/cloud-router-module/connectivity/latest). This module provides pre-defined configurations and workflows for provisioning cloud routers on the PacketFabric platform.

![Deployment Diagram](./images/diagram_cloud_router_aws_google.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp AWS Terraform Provider](https://registry.terraform.io/providers/hashicorp/aws)
- [HashiCorp Google Terraform Provider](https://registry.terraform.io/providers/hashicorp/google)
- [Google Cloud CLI Terraform Module](https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Demo video

Check out our demo video to see how the PacketFabric Terraform Provider can be used to automate the provisioning and management of PacketFabric resources.

<p align="center"><a href="https://www.youtube.com/watch?v=EblOg0Uaf8Q" target=”_blank”><img width="60%" height="60%" src="https://img.youtube.com/vi/EblOg0Uaf8Q/1.jpg"></a></p>

## Terraform resources & data-sources used

- "aws_dx_gateway"
- "aws_dx_transit_virtual_interface"
- "aws_dx_gateway_association"
- "aws_security_group"
- "aws_network_interface"
- "aws_key_pair"
- "aws_instance"
- "aws_eip"
- "aws_ec2_transit_gateway"
- "aws_ec2_transit_gateway_vpc_attachment"
- "aws_vpc"
- "aws_subnet"
- "aws_internet_gateway"
- "aws_route_table_association"
- "packetfabric_cloud_router"
- "packetfabric_cloud_router_connection_aws"
- "aws_dx_connection_confirmation"
- "google_compute_router"
- "google_compute_interconnect_attachment"
- "packetfabric_cloud_router_connection_google"
- "packetfabric_cloud_router_bgp_session"
- "google_compute_firewall"
- "google_compute_instance"
- "google_compute_network"
- "google_compute_subnetwork"
- "random_pet"

**Estimated time:** ~10 min for Google, AWS & PacketFabric resources + ~10-15 min for AWS Direct Connect Gateway association with AWS Transit Gateway

**Note**: Because the BGP session is created automatically, we use gcloud terraform module to retreive the BGP addresses and set the PacketFabric Cloud Router ASN in the BGP settings in the Google Cloud Router. Please [vote](https://github.com/hashicorp/terraform-provider-google/issues/11458), [vote](https://github.com/hashicorp/terraform-provider-google/issues/12624) and [vote](https://github.com/hashicorp/terraform-provider-google/issues/12630) for these issues on GitHub.

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)
- Don't have a Google Account? [Get Started](https://cloud.google.com/free)

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)
- [gcloud](https://cloud.google.com/sdk/docs/install)
- [jq](https://stedolan.github.io/jq/download/)

Ensure you have the following items available:

- [AWS Account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
- [AWS Access and Secret Keys](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html)
- [Google Service Account](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances)
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

Set additional environment variables for AWS and Google:

```sh
### AWS
export PF_AWS_ACCOUNT_ID="98765432"
export AWS_ACCESS_KEY_ID="ABCDEFGH"
export AWS_SECRET_ACCESS_KEY="secret"

### Google
export TF_VAR_gcp_project_id="my-project-id" # used for bash script used with gcloud module
export GOOGLE_CREDENTIALS='{ "type": "service_account", "project_id": "demo-setting-1234", "private_key_id": "1234", "private_key": "-----BEGIN PRIVATE KEY-----\nsecret\n-----END PRIVATE KEY-----\n", "client_email": "demoapi@demo-setting-1234.iam.gserviceaccount.com", "client_id": "102640829015169383380", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "token_uri": "https://oauth2.googleapis.com/token", "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/demoapi%40demo-setting-1234.iam.gserviceaccount.com" }'

export TF_VAR_public_key="ssh-rsa AAAA..." # see link Create an SSH key pair in the pre-req
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

1. In case you get the following error:

```
╷
│ Error: Error when reading or editing InterconnectAttachment: googleapi: Error 400: The resource 'projects/prefab-setting-123456/regions/us-west1/interconnectAttachments/demo-pf-gcp-vpn-master-cricket' is not ready, resourceNotReady
│ 
│ 
```

This seems to be a problem with Google Terraform Provider, run again the terraform destroy command and the destroy will complete correctly the 2nd try.
Please [vote](https://github.com/hashicorp/terraform-provider-google/issues/12631) for this issue on GitHub.

2. In case the ``gcloud_bgp_address`` module fails, check the error, fix it and manually remove the state before re-running the terraform config.

```sh
terraform state rm module.gcloud_bgp_addresses
terraform state rm module.gcloud_bgp_peer_update
```

3. In case you get the following error:

```
╷
│ Error: error waiting for Direct Connection Connection (dxcon-fgohxwui) confirm: timeout while waiting for state to become 'available' (last state: 'pending', timeout: 10m0s)
│ 
│   with aws_dx_connection_confirmation.confirmation,
│   on cloud_router_connection_aws.tf line 46, in resource "aws_dx_connection_confirmation" "confirmation":
│   46: resource "aws_dx_connection_confirmation" "confirmation" {
│ 
```

You are hitting a timeout issue in AWS [aws_dx_connection_confirmation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/dx_connection_confirmation) resource. Please [vote](https://github.com/hashicorp/terraform-provider-aws/issues/26335) for this issue on GitHub.

As a workaround, edit the `cloud_router_connection_aws.tf` and comment out the following resource:

```
# resource "aws_dx_connection_confirmation" "confirmation" {
#   provider      = aws
#   connection_id = data.aws_dx_connection.current.id

#   lifecycle {
#     ignore_changes = [
#       connection_id
#     ]
#   }
# }
```

Edit the `aws_dx_transit_vif.tf` and comment out the dependency with `confirmation` in `packetfabric_cloud_router_connections` data source: 

```
data "packetfabric_cloud_router_connections" "current" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  # depends_on = [
  #   aws_dx_connection_confirmation.confirmation
  # ]
}
```

Then remove the `confirmation` state, check the Direct Connect connection is **available** and re-apply the terraform plan:
```
terraform apply
```