[![Release](https://img.shields.io/github/v/release/PacketFabric/terraform-provider-packetfabric?display_name=tag)](https://github.com/PacketFabric/terraform-provider-packetfabric/releases)
![release-date](https://img.shields.io/github/release-date/PacketFabric/terraform-provider-packetfabric)
![contributors](https://img.shields.io/github/contributors/PacketFabric/terraform-provider-packetfabric)
![commit-activity](https://img.shields.io/github/commit-activity/m/PacketFabric/terraform-provider-packetfabric)
[![License](https://img.shields.io/github/license/PacketFabric/terraform-provider-packetfabric)](https://github.com/PacketFabric/terraform-provider-packetfabric)

## How to contribute to PacketFabric Terraform Provider

#### **Did you find a bug?**

* **Ensure the bug was not already reported** by searching on GitHub under [Issues](https://github.com/PacketFabric/terraform-provider-packetfabric/issues).

* If you're unable to find an open issue addressing the problem, [open a new one](https://github.com/PacketFabric/terraform-provider-packetfabric/issues/new?assignees=&labels=bug&template=bug-report.md). Be sure to include a **title and clear description**, as much relevant information as possible, and a **Terraform code sample** or an **executable test case** demonstrating the expected behavior that is not occurring.

#### **Did you write a patch that fixes a bug?**

* Open a new GitHub pull request with the patch.

* Ensure the PR description clearly describes the problem and solution. Include the relevant issue number if applicable.

#### **Do you intend to add a new feature or change an existing one?**

* Open a new [feature request](https://github.com/PacketFabric/terraform-provider-packetfabric/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=), start writing code and submit it for review.

* Adding or Updating existing Terraform resources/data-sources consists of:

    * Add/Update Terraform Go Code located under the internal folder 
    * Add/Update Tests 
        * using mock data `<file>_test.go` under `internal/packetfabric`
        * using real data `<resource_name>_test.go` under `internal/provider` (see [ACC](https://github.com/PacketFabric/terraform-provider-packetfabric#acceptance-tests))
    * Add/Update examples under `examples/resources` and/or `examples/data-sources` (used for the documentation)
    * Add/Update the templates used to generate the docs  under `templates`
    * Generate the docs using [tfplugindocs](https://github.com/hashicorp/terraform-plugin-docs) and verify `*.md` under `docs/`
    * Find more details on the [Readme](https://github.com/PacketFabric/terraform-provider-packetfabric)

* Create your own branch with your updates including code changes, test, examples and documentation. 

* Make sure you test it locally, then create a PR so it can be reviewed by our team.

#### **Do you have questions about the source code or the provider in general?**

* Ask any question [here](https://github.com/PacketFabric/terraform-provider-packetfabric/issues/new?assignees=&labels=help+wanted&template=terraform-packetfabric-provider-questions.md&title=).

Thanks! :rocket: :smile:

The PacketFabric Team
