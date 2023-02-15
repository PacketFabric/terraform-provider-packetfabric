# Use Case: 

This use case shows an example on how to use the PacketFabric Terraform provider 
to automate the creation of a Virtual Circuit between 2 ports owned by 2 different PacketFabric accounts (A and Z).

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [PacketFabric Ports Overview](https://docs.packetfabric.com/ports/)
- [PacketFabric Virtual Circuits Overview](https://docs.packetfabric.com/vc/)
- [PacketFabric Marketplace Overview](https://docs.packetfabric.com/eco/overview/)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

**A side**

- "random_pet"
- "packetfabric_backbone_virtual_circuit_marketplace"
- "packetfabric_backbone_virtual_circuit"

**Z side**

- "random_pet"
- "packetfabric_marketplace_service_port_accept_request"
- "packetfabric_marketplace_service_port_reject_request"

## Before you begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)

## Prerequisites

Ensure you have installed the following prerequisites:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Ensure you have the following items available:

- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)

## Quick start

1. Set the PacketFabric API key and Account ID for **A side** in your terminal as environment variables.

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

Windows PowerShell:
```powershell
PS C:\> $Env:PF_TOKEN="secret"
PS C:\> $Env:PF_ACCOUNT_ID="123456789"
```

2. Initialize Terraform, create an execution plan and execute the plan.

**A side** in `a_side` folder:

```sh
terraform init
terraform plan
```

**Note:** you can update terraform variables in the ``variables.tf``.

3. Apply the plan

```sh
terraform apply
```

4. **B side** in `b_side` folder:

Update API key and Account ID for **B side**:

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

Windows PowerShell:
```powershell
PS C:\> $Env:PF_TOKEN="secret"
PS C:\> $Env:PF_ACCOUNT_ID="123456789"
```

5. Then run:

```sh
terraform init
terraform plan
```

6. Update the `pf_a_side_vc_request_uuid` with the **A Side** Virtual Circuit Request UUID in the `variables.tf`.
You can either Accept or Reject the request (comment/comment out as desire).

7. Update API key and Account ID for **A side**:

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

Windows PowerShell:
```powershell
PS C:\> $Env:PF_TOKEN="secret"
PS C:\> $Env:PF_ACCOUNT_ID="123456789"
```

8. Apply the plan

```sh
terraform apply
```

9. **A side**, comment `packetfabric_backbone_virtual_circuit_marketplace` and comment out `packetfabric_backbone_virtual_circuit` resources.

10. **A side**, import the new Marketplace backbone Virtual Circuit (replace with correct VC ID).

```sh
terraform import packetfabric_backbone_virtual_circuit.vc_marketplace PF-DC-PHX-NYC-1751589-PF 
```

11. **A side**, apply the plan to confirm the resource is correctly imported and managed by Terraform.

```sh
terraform plan
terraform apply
```

12. Destroy all remote objects managed by the Terraform configuration on both sides (in `a_side` and `b_side` folders).

```sh
terraform destroy
```
