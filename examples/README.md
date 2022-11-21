# Examples

- [Get Started with Terraform with PacketFabric](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric documentation in the Terraform registry](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs)
- [Use case examples of using the PacketFabric provider](./use-cases)

## Quick Start

1. Create the file ``secret.tfvars`` and update each variables as needed (edit ``variables.tf``).

```sh
cp secret.tfvars.sample secret.tfvars
```

**Note:** As an alternative, you can also use ``env_vars_source.sh.sample`` (rename to ``env_vars_source.sh`` and update each variables as needed). Then run ``./env_vars_source.sh`` or run:

```
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
```

2. Edit the ``main.tf`` and ``variables.tf`` files and uncomment/comment out sections as needed. It is highly recommended to use the [PacketFabric documentation in the Terraform registry](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs).

3. Initialize Terraform, create an execution plan and execute the plan (``-var-file="secret.tfvars"`` can be removed if you used ``env_vars_source.sh``)

```sh
terraform init
terraform plan -var-file="secret.tfvars"
```

Apply the plan:

```sh
terraform apply -var-file="secret.tfvars"
```

4. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy -var-file="secret.tfvars"
```