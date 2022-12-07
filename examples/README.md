# Examples

- [Get Started with Terraform with PacketFabric](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric documentation in the Terraform registry](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs)
- [Use case examples of using the PacketFabric provider](./use-cases)

## Quick start

1. Set the PacketFabric API key and Account ID in the environment variables and update each variables as needed (edit ``variables.tf``).

```sh
export PF_TOKEN="secret"
export PF_ACCOUNT_ID="123456789"
export PF_AWS_ACCOUNT_ID="98765432"
```

**Note:** you can also use ``source_env_var.sh.sample`` (rename to ``source_env_var.sh`` and update each variables as needed). Then run ``source source_env_var.sh``.

2. Edit the ``main.tf`` and ``variables.tf`` files and uncomment/comment out sections as needed. It is highly recommended to use the [PacketFabric documentation in the Terraform registry](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs).

3. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan
```

4. Apply the plan:

```sh
terraform apply
```

5. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy
```