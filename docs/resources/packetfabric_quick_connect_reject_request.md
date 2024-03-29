---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_quick_connect_reject_request Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_quick_connect_reject_request (Resource)



## Example Usage

```terraform
resource "packetfabric_quick_connect_reject_request" "reject_request_quick_connect" {
  provider          = packetfabric
  import_circuit_id = "PF-L3-IMP-2896010"
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `import_circuit_id` (String) Circuit ID of the Quick Connect import.

### Optional

- `rejection_reason` (String) The reason that you are rejecting the request.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)




