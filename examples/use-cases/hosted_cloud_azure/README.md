# Use Case: PacketFabric Hosted cloud connection to Microsoft Azure ExpressRoute

This use case shows an example on how to use the PacketFabric & Azure Terraform providers 
to automate the creation of a Hosted Cloud Connection between PacketFabric and Azure in a Cloud On-Ramps facility.

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Hosted Microsoft ExpressRoute Process Overview](https://docs.packetfabric.com/cloud/microsoft/hosted/process/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp Microsoft Azure Terraform Provider](https://registry.terraform.io/providers/hashicorp/azurerm)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

> **Warning**: Microsoft begins billing as soon as the service key is created, which is why we advise that you wait until your cross connect is established first.

- "random_pet"
- "azurerm_resource_group"
- "azurerm_virtual_network"
- "azurerm_subnet"
- "azurerm_express_route_circuit"
- "packetfabric_cs_azure_hosted_connection"
- "azurerm_express_route_circuit_peering"

**Estimated time:** ~5 min for Azure & PacketFabric resources + up to 50 min for Azure Virtual Network Gateway (deletion up to 12min)

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an Azure Account? [Get Started](https://azure.microsoft.com/en-us/free/)

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Ensure you have the following items available:

- [Microsoft Azure Credentials](https://docs.microsoft.com/en-us/azure/developer/terraform/authenticate-to-azure?tabs=bash)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)
- 1 [PacketFabric Port](https://docs.packetfabric.com/ports/) [cross connected](https://docs.packetfabric.com/xconnect/) to your network infrastructure  (update the ``pf_port_circuit_id`` in ``variables.tf``)

Also:

- [Review the ExpressRoute network requirements](https://docs.microsoft.com/en-us/azure/expressroute/expressroute-prerequisites#network-requirements)
- Enable AzureExpressRoute in the Azure Subscription:

```sh
az feature register --namespace Microsoft.Network --name AllowExpressRoutePorts
az provider register -n Microsoft.Network
```

## Quick start

1. Set the PacketFabric API key and Account ID in your terminal as environment variables.

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

Set additional environment variables for Azure:

```sh
export ARM_CLIENT_ID="00000000-0000-0000-0000-000000000000"
export ARM_CLIENT_SECRET="00000000-0000-0000-0000-000000000000"
export ARM_SUBSCRIPTION_ID="00000000-0000-0000-0000-000000000000"
export ARM_TENANT_ID="00000000-0000-0000-0000-000000000000"
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
terraform state rm azurerm_express_route_circuit.azure_express_route_1
terraform destroy
```

The ExpressRoute Circuit needs to be in a deprovisioned state before being deleted.

Because **pf_cs_conn1** depends on **azure_express_route_1**, ``terraform destroy`` will try to delete **azure_express_route_1** first. By removing the state of **azure_express_route_1**, the **pf_cs_conn1** object is deleted, then the deletion of the Azure ExpressRoute circuit **azure_express_route_1** will happen part of the resource group **resource_group_1** deletion.
