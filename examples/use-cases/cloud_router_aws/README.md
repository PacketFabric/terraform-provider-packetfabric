# Use Case: PacketFabric Cloud Router with AWS

This use case builds a connection between two AWS regions using the PacketFabric Cloud Router.
Terraform providers used: PacketFabric and AWS. This example uses AWS Private VIF & Gateway.

:rocket: You can simplify the configuration and management of PacketFabric Cloud Routers by utilizing the [PacketFabric Terraform Cloud Router Module](https://registry.terraform.io/modules/PacketFabric/cloud-router-module/connectivity/latest). This module provides pre-defined configurations and workflows for provisioning cloud routers on the PacketFabric platform.

![Deployment Diagram](./images/diagram_cloud_router_aws.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp AWS Terraform Provider](https://registry.terraform.io/providers/hashicorp/aws)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources & data-sources used

- "aws_dx_gateway"
- "aws_dx_gateway_association"
- "aws_security_group"
- "aws_network_interface"
- "aws_key_pair"
- "aws_instance"
- "aws_eip"
- "aws_vpn_gateway"
- "aws_route_table"
- "aws_vpc"
- "aws_subnet"
- "aws_internet_gateway"
- "aws_route_table_association"
- "packetfabric_cloud_router"
- "packetfabric_cloud_router_connection_aws"
- "random_pet"

**Estimated time:** ~15 min for AWS & PacketFabric resources + ~10-15 min for AWS Direct Connect Gateway association with AWS Virtual Private Gateways

**Note**: Make sure you set the correct AWS region based on the PacketFabric pop selected (find details on location [here](https://packetfabric.com/locations/cloud-on-ramps) and [here](https://aws.amazon.com/directconnect/locations/). Essentially, select the PacketFabric pop the closest to the AWS region you want to connect to. Example: AWS region ``us-west-1`` is the closest to PacketFabric pop ``LAX1``.

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)
    - Permissions required: VPC, EC2, Direct Connect

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Ensure you have the following items available:

- [AWS Account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
- [AWS Access and Secret Keys](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html)
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

Set additional environment variables for AWS:

```sh
export PF_AWS_ACCOUNT_ID="98765432"
export AWS_ACCESS_KEY_ID="ABCDEFGH"
export AWS_SECRET_ACCESS_KEY="secret"

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

4. Either use and [locust](https://locust.io/) or [iperf3](https://github.com/esnet/iperf) to simulate traffic between the 2 EC2 instances in the 2 AWS regions.

- iperf3: (replace ``<ec2_private_ip_2>`` with the EC2 private IP from instance 2)

    - Server side (on instance 1): ``iperf3 -s -p 5001``
    - Client side (on instance 2): ``iperf3 -c <ec2_private_ip_1> -p 5001``

- locust: (replace ``<ec2_public_ip_1>`` with the EC2 public IP)

In a browser, on instance 1), open ``http://<ec2_public_ip_1>:8089/``, then update the host with the correct IP (using `` <ec2_private_ip_2>`` from instance 2), set the number of users to ``500`` and spawn rate to ``50``.

If you want to use iperf3, open a ssh session using the user ``ubuntu`` and the ssh private key linked to the public key you specified in the ``secret.tfvars`` file.

5. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy
```

**Note:** Default login/password for Locust is ``demo:packetfabric`` edit ``user-data-ubuntu.sh`` script to change it.

## Screenshots

Example AWS Direct Connect Connections:

![AWS Direct Connect Connections](./images/aws_direct_connect_connections.png)

Example AWS Direct Connect gateway:

![AWS Direct Connect gateway](./images/aws_direct_connect_gateway.png)

Example Direct Connect Private Virtual interfaces:

![AWS Direct Connect Private Virtual interfaces](./images/aws_direct_connect_private_virtual_interfaces.png)

Example AWS VPC Routing Table:

![AWS VPC Routing Table](./images/aws_vpc_routing_table.png)

Traffic Generator using Locust: *Response Time ~77ms*

![Demo Locust AWS PacketFabric](./images/demo_aws_traffic_direct_connect_through_PacketFabric_500_users_locust.png)
