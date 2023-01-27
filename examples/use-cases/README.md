# List of PacketFabric Use Cases

Use Cases | Description
--- | --- 
[backbone_virtual_circuit](./backbone_virtual_circuit) | Create Ports and Virtual Circuit
[backbone_virtual_circuit_mesh](./backbone_virtual_circuit_mesh) | Create a mesh between Virtual Circuits
[cloud_router_module](./cloud_router_module) | Create a Cloud Router Terraform Module with AWS and Google
[cloud_router_aws](./cloud_router_aws) | Connect 2 AWS regions using a Cloud Router (private VIF)
[cloud_router_aws_google](./cloud_router_aws_google) | Connect AWS (transit VIF) and Google Clouds using a Cloud Router
[cloud_router_google_ipsec](./cloud_router_google_ipsec) | Connect Google Cloud and a branch location, on-premises users, or a remote data center using a Cloud Router
[cloud_router_google_azure](./cloud_router_google_azure) | Connect Google and Azure Clouds using a Cloud Router
[cloud_router_ibm_oracle](./cloud_router_ibm_oracle) | Connect Oracle and IBM Clouds using a Cloud Router
[dedicated_cloud_aws](./dedicated_cloud_aws) | Create Dedicated Cloud Connection on AWS
[hosted_cloud_aws](./hosted_cloud_aws) | Create a Port and an AWS Hosted Cloud Connection
[hosted_cloud_google](./hosted_cloud_google) | Create a Port and a Google Hosted Cloud Connection
[hosted_cloud_azure](./hosted_cloud_azure) | Create a Port and an Azure Hosted Cloud Connection
[marketplace](./marketplace) | Create a Marketplace Request and Accept/Reject it

# List of known issues on other terraform providers, please up vote!

**[AWS](https://registry.terraform.io/providers/hashicorp/aws/latest)**

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/25989)](https://github.com/hashicorp/terraform-provider-aws/issues/25989) aws_dx_public_virtual_interface does do dependency checks for amazon_address and customer_address (**impact**: PacketFabric AWS Hosted Cloud and Cloud Router using Public VIF)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26335)](https://github.com/hashicorp/terraform-provider-aws/issues/26335) aws_dx_connection_confirmation add timeout and do not fail when state is available (**impact**: PacketFabric AWS Hosted Cloud and Cloud Router)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26919)](https://github.com/hashicorp/terraform-provider-aws/issues/26919) aws_dx_connection: Error: 2 Direct Connect Connections matched (add filter) (**impact**: PacketFabric AWS Hosted Cloud and Cloud Router)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26436)](https://github.com/hashicorp/terraform-provider-aws/issues/26436) aws_dx_connection data source: add PDF LOA in base64 encoded (**impact**: PacketFabric AWS Dedicated Cloud

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26438)](https://github.com/hashicorp/terraform-provider-aws/issues/26438) aws_dx_locations: add Direct Connect Locations & Speed + filter capability (**impact**: PacketFabric AWS Dedicated Cloud)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26461)](https://github.com/hashicorp/terraform-provider-aws/issues/26461) aws_dx_connection data source: add VLAN ID (**impact**: PacketFabric AWS Hosted Cloud)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26432)](https://github.com/hashicorp/terraform-provider-aws/issues/26432) New aws_dx_virtual_interface_router_configuration data source (**impact**: PacketFabric AWS Hosted Cloud)

**[Google Cloud](https://registry.terraform.io/providers/hashicorp/google/latest)**

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/11458)](https://github.com/hashicorp/terraform-provider-google/issues/11458) Expose bgpPeers from google_compute_router (**impact**: PacketFabric Google Hosted Cloud and Cloud Router)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/12624)](https://github.com/hashicorp/terraform-provider-google/issues/12624) New data source for google_compute_interconnect_attachment (**impact**: PacketFabric Google Hosted Cloud and Cloud Router)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/12630)](https://github.com/hashicorp/terraform-provider-google/issues/12630) New google_compute_router_peer_asn_update resource for Partner Interconnect (**impact**: PacketFabric Google Hosted Cloud and Cloud Router)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/12631)](https://github.com/hashicorp/terraform-provider-google/issues/12631) google_compute_interconnect_attachment  Error 400: The resource is not ready (**impact**: PacketFabric Google Hosted Cloud and Cloud Router)

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/13103)](https://github.com/hashicorp/terraform-provider-google/issues/13103) Add md5_authentication_key in google_compute_router (**impact**: PacketFabric Google Hosted Cloud and Cloud Router)

**[IBM](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest)**

- [![Issues](https://img.shields.io/github/issues/detail/state/IBM-Cloud/terraform-provider-ibm/3978)](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3978) New resource to accept a direct link creation request: ibm_dl_gateway_accept (**impact**: PacketFabric IBM Hosted Cloud and Cloud Router)


**[Terraform](https://www.terraform.io/)**

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform/27360)](https://github.com/hashicorp/terraform/issues/27360) Add support for lifecycle meta-argument in modules (**impact**: PacketFabric Google Cloud Router)
