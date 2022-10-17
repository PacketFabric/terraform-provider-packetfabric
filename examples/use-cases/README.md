# List of PacketFabric Use Cases

- [backbone_virtual_circuit](./backbone_virtual_circuit): create ports and virtual circuit
- [backbone_virtual_circuit_mesh](./backbone_virtual_circuit_mesh): create a mesh between virtual circuits
- [cloud_router_aws](./cloud_router_aws): connect 2 AWS region using a Cloud Router
- [cloud_router_google_ipsec](./cloud_router_google_ipsec): connect GCP and a branch location, on-premises users, or a remote data center using a Cloud Router
- [cloud_router_google_azure](./cloud_router_google_azure): connect GCP and Azure Clouds using a Cloud Router
- [cloud_router_ibm_oracle](./cloud_router_ibm_oracle): connect Oracle and IBM Clouds using a Cloud Router
- [hosted_cloud_aws](./hosted_cloud_aws): create an AWS Hosted Cloud Connection
- [hosted_cloud_google](./hosted_cloud_google): create a Google Hosted Cloud Connection
- [hosted_cloud_azure](./hosted_cloud_azure): create an Azure Hosted Cloud Connection
- [dedicated_cloud_aws](./dedicated_cloud_aws): create Dedicated Cloud Connection on AWS

# List of known issues on other terraform providers, please vote!

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/25989)](https://github.com/hashicorp/terraform-provider-aws/issues/25989) aws_dx_public_virtual_interface does do dependency checks for amazon_address and customer_address

**Impact**: AWS Hosted Cloud and Cloud Router using Public VIF

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26335)](https://github.com/hashicorp/terraform-provider-aws/issues/26335) aws_dx_connection_confirmation add timeout and do not fail when state is available 

**Impact**: AWS Hosted Cloud and Cloud Router

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26432)](https://github.com/hashicorp/terraform-provider-aws/issues/26432) New aws_dx_virtual_interface_router_configuration data source

**Impact**: AWS Hosted Cloud

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26919)](https://github.com/hashicorp/terraform-provider-aws/issues/26919) aws_dx_connection: Error: 2 Direct Connect Connections matched (add filter)

**Impact**: AWS Hosted Cloud and Cloud Router

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26461)](https://github.com/hashicorp/terraform-provider-aws/issues/26461) aws_dx_connection data source: add VLAN ID

**Impact**: AWS Hosted Cloud

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26436)](https://github.com/hashicorp/terraform-provider-aws/issues/26436) aws_dx_connection data source: add PDF LOA in base64 encoded

**Impact**: AWS Dedicated Cloud

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-aws/26438)](https://github.com/hashicorp/terraform-provider-aws/issues/26438) aws_dx_locations: add Direct Connect Locations & Speed + filter capability

**Impact**: AWS Dedicated Cloud

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/11458)](https://github.com/hashicorp/terraform-provider-google/issues/11458) Expose bgpPeers from google_compute_router

**Impact**: Google Hosted Cloud and Cloud Router

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/12624)](https://github.com/hashicorp/terraform-provider-google/issues/12624) New data source for google_compute_interconnect_attachment 

**Impact**: Google Hosted Cloud and Cloud Router

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/12630)](https://github.com/hashicorp/terraform-provider-google/issues/12630) New google_compute_router_peer_asn_update resource for Partner Interconnect

**Impact**: Google Hosted Cloud and Cloud Router

- [![Issues](https://img.shields.io/github/issues/detail/state/hashicorp/terraform-provider-google/12631)](https://github.com/hashicorp/terraform-provider-google/issues/12631) google_compute_interconnect_attachment  Error 400: The resource is not ready

**Impact**: Google Hosted Cloud and Cloud Router

- [![Issues](https://img.shields.io/github/issues/detail/state/IBM-Cloud/terraform-provider-ibm/3978)](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3978) New resource to accept a direct link creation request: ibm_dl_gateway_accept

**Impact**: IBM Hosted Cloud and Cloud Router
