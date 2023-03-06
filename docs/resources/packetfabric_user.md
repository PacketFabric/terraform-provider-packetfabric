---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_user Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_user (Resource)

For more information, see the [PacketFabric user documentation](https://docs.packetfabric.com/admin/user/).

## Example Usage

```terraform
resource "packetfabric_user" "user1" {
  provider   = packetfabric
  first_name = "Alice"
  last_name  = "Thomas"
  email      = "alice@mycompany.com"
  phone      = "2065434573"
  login      = "alice@mycompany.com"
  password   = "secret"
  timezone   = "America/Vancouver"
  group      = "read-only"
}

output "packetfabric_user" {
  value = packetfabric_user.user1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) User e-mail.
- `first_name` (String) User first name.
- `group` (String) User group. Available options are admin, regular, read-only, support, and sales.
- `last_name` (String) User last name.
- `login` (String) User login.
- `password` (String) User password. Keep it in secret.
- `phone` (String) User phone number.
- `timezone` (String) User time-zone. You can find the list of available timezones [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)




## Import

Import a User using the User login.

```bash
terraform import packetfabric_user.user1 mylogin1
```