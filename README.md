# PacketFabric Terraform Provider

- Documentation: https://registry.terraform.io/providers/packetfabric/packetfabric/latest/docs

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) v1.2.2
- [Go](https://golang.org/doc/install) 1.18.2 (to build the provider plugin)

## Standard Provider usage

See the [PacketFabric Provider documentation](https://registry.terraform.io/providers/packetfabric/packetfabric/latest/docs) for resource definition and data-source structure and examples.

## Building and installing the Provider locally

```sh
$ git clone git@github.com:packetfabric/terraform-provider-packetfabric
$ make install
go build -o terraform-provider-packetfabric
mkdir -p ~/.terraform.d/plugins/[YOURHOSTNAME]/packetfabric/packetfabric/0.0.1/linux_amd64
mv terraform-provider-packetfabric ~/.terraform.d/plugins/[YOURHOSTNAME]/packetfabric/packetfabric/0.0.1/linux_amd64

```

## Using the local build/installed provider

```terraform
terraform {
  required_providers {
    packetfabric = {
      source  = "[YOURHOSTNAME]/packetfabric/packetfabric"
      version = "~> 0.0.1"
    }
  }
}

```

## Contributing Documentation

Markdown documents found in this repository are the source of the [PacketFabric Provider documentation](https://registry.terraform.io/providers/packetfabric/packetfabric/latest/docs). These source documents are generated using the [Terraform-Plugin-Docs Tools](https://github.com/hashicorp/terraform-plugin-docs).

Data-Source and Resource field descriptions are pulled from the [internal/provider](https://github.com/packetfabric/terraform-provider-packetfabric/tree/main/internal/provider) source-code. It's organized based on the resource and data-source matching markdown template found in [templates](https://github.com/packetfabric/terraform-provider-packetfabric/tree/main/templates) and merged with examples found in the [examples](https://github.com/packetfabric/terraform-provider-packetfabric/tree/main/examples)

Updating the provider function field descriptions should be done in [internal/provider](https://github.com/packetfabric/terraform-provider-packetfabric/tree/main/internal/provider) code schema descriptions. Structural changes to the MD formating should be done in the [templates](https://github.com/packetfabric/terraform-provider-packetfabric/tree/main/templates) tmpls files.

Caveat: As of tfplugindocs 0.10.1, Nested schema elements are not properly discovered and inserted at generation time. This means data-source function field descriptions must be manually managed in the [templates](https://github.com/packetfabric/terraform-provider-packetfabric/tree/main/templates) tmpls files for data-sources.

## Developing The Provider

To work on the provider, you'll need [Go](http://www.golang.org) installed on your machine (version 1.11+ is _required_). You'll need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile, run `make build`. To compile and install, run `make install` . This will build (and install) the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-packetfabric
...
```

To test the provider, run `make test`.

```sh
$ make test
```

To check changes you made locally to the provider, you can use the binary you compiled by adding the following
to your `~/.terraformrc` file. This is valid for Terraform 0.14+. See
[Terraform's documentation](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers) for more details.

```
provider_installation {

  # Use /home/developer/go/bin as an overridden package directory
  # for the PacketFabric/packetfabric provider. This disables the version and checksum
  # verifications for this provider and forces Terraform to look for the
  # packetfabric provider plugin in the given directory.
  dev_overrides {
    "PacketFabric/packetfabric" = "/home/developer/go/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

For information about writing acceptance tests, see the main Terraform [contributing guide](https://github.com/hashicorp/terraform/blob/master/.github/CONTRIBUTING.md#writing-acceptance-tests).

## Acceptance Tests

To run acceptance tests on your local machine, you have to set the following
environmental variables:

```shell
export PF_HOST="https://api.packetfabric.com"
export PF_TOKEN="api-secret"
export PF_ACCOUNT_UUID="1234"
export PF_ACC_TEST_ROUTING_ID="PD-WUY-9VB0"
export PF_ACC_TEST_MARKET="HOU"
```

> **Warning**: Running below command will order various PacketFabric products, then delete them.

Then you can run the following command:

```shell
make testacc
```

If you want to know the current list of acceptance tests available without executing them, run the following command:

```
cd ./internal/provider
go test -cover -v | grep -v testutil.go | grep -v github.com
```

## Releasing the Provider

This provider is published using GitHub Actions triggered by tagging a branch using semantic versioning with the pattern `v*`(Example: `v0.1.3`)

Once the branch is tagged the release is built and publish via the Terraform Registry.

Provider release candidates will be based on main-branch and be committed on their own, dedicated, dev branch. The release branch will be qualified and UAT then merged with main. A new Release branch will be created from main at the merge point. The Release branch will be tagged for publishing. This process allows us to support multiple versions of the provider simultaneously if desired.
