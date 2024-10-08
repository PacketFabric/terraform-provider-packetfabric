---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_port Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_port (Resource)

A port on the PacketFabric network. For more information, see [Ports in the PacketFabric documentation](https://docs.packetfabric.com/ports/).

## Example Usage

```terraform
resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  enabled           = true
  autoneg           = true
  description       = "hello world"
  media             = "LX"
  nni               = false
  pop               = "SEA2"
  speed             = "1Gbps"
  subscription_term = 1
  zone              = "A"
  labels            = ["terraform", "dev"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_uuid` (String) The UUID for the billing account that should be billed. Can also be set with the PF_ACCOUNT_ID environment variable.
- `description` (String) A brief description of the port.
- `media` (String) Optic media type.

	Enum: ["LX" "EX" "ZX" "LR" "ER" "ER DWDM" "ZR" "ZR DWDM" "LR4" "ER4" "CWDM4" "LR4" "ER4 Lite"]
- `pop` (String) Point of presence in which the port should be located.
- `speed` (String) Speed of the port.

	Enum: ["1Gbps" "10Gbps" "40Gbps" "100Gbps"]
- `subscription_term` (Number) Duration of the subscription in months

	Enum ["1" "12" "24" "36"]
- `zone` (String) The desired availability zone of the port.

	Example: "A"

### Optional

- `autoneg` (Boolean) Only applicable to 1Gbps ports. Controls whether auto negotiation is on (true) or off (false). Defaults: true
- `enabled` (Boolean) Change Port Admin Status. Set it to true when port is enabled, false when port is disabled. Defaults: true
- `labels` (Set of String) Label value linked to an object.
- `nni` (Boolean) Set this to true to provision an ENNI port. ENNI ports will use a nni_svlan_tpid value of 0x8100.

	By default, ENNI ports are not available to all users. If you are provisioning your first ENNI port and are unsure if you have permission, contact support@packetfabric.com. Defaults: false
- `po_number` (String) Purchase order number or identifier of a service.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `etl` (Number) Early Termination Liability (ETL) fees apply when terminating a service before its term ends. ETL is prorated to the remaining contract days.
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)




## Import

Import a port using its circuit ID. 

```bash
terraform import packetfabric_port.port_1 PF-AP-WDC1-1726464
```