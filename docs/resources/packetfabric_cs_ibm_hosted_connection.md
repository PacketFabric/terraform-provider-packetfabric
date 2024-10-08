---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_cs_ibm_hosted_connection Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_cs_ibm_hosted_connection (Resource)

 A hosted cloud connection to your IBM cloud environment. For more information, see [Cloud Connections in the PacketFabric documentation](https://docs.packetfabric.com/cloud/).



## Example Usage

```terraform
resource "packetfabric_cs_ibm_hosted_connection" "cs_conn1_hosted_ibm" {
  provider    = packetfabric
  ibm_bgp_asn = 64536
  description = "hello world"
  pop         = "WDC1"
  port        = packetfabric_port.port_1.id
  vlan        = 102
  speed       = "10Gbps"
  labels      = ["terraform", "dev"]
}

resource "time_sleep" "wait_ibm_connection" {
  create_duration = "5m"
}
data "ibm_dl_gateway" "current" {
  provider   = ibm
  name       = "hello world" # same as the PacketFabric IBM Hosted Cloud description
  depends_on = [time_sleep.wait_ibm_connection]
}
data "ibm_resource_group" "existing_rg" {
  provider = ibm
  name     = "My Resource Group"
}

resource "ibm_dl_gateway_action" "confirmation" {
  provider                    = ibm
  gateway                     = data.ibm_dl_gateway.current.id
  resource_group              = data.ibm_resource_group.existing_rg.id
  action                      = "create_gateway_approve"
  global                      = true
  metered                     = true # If set true gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway
  bgp_asn                     = 64536
  default_export_route_filter = "permit"
  default_import_route_filter = "permit"
  speed_mbps                  = 10000 # must match PacketFabric speed

  provisioner "local-exec" {
    when    = destroy
    command = "sleep 30"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_uuid` (String) The UUID for the billing account that should be billed. Can also be set with the PF_ACCOUNT_ID environment variable.
- `description` (String) The description of this connection. This will appear as the connection name from the IBM side. Allows only numbers, letters, underscores and dashes.
- `ibm_account_id` (String) Your IBM account ID. Can also be set with the PF_IBM_ACCOUNT_ID environment variable.
- `ibm_bgp_asn` (Number) Enter an ASN to use with your BGP session. This should be the same ASN you used for your Cloud Router.
- `pop` (String) The POP in which you want to provision the connection (the on-ramp).
- `port` (String) The port to connect to IBM.
- `speed` (String) The speed of the new connection.

	Enum: ["50Mbps" "100Mbps" "200Mbps" "300Mbps" "400Mbps" "500Mbps" "1Gbps" "2Gbps" "5Gbps" "10Gbps"]
- `vlan` (Number) Valid VLAN range is from 4-4094, inclusive.
- `zone` (String) The desired availability zone of the connection.

	Example: "A"

### Optional

- `ibm_bgp_cer_cidr` (String) The IP address in CIDR format for the PacketFabric-side router in the BGP session. If you do not specify an address, IBM will assign one on your behalf.
- `ibm_bgp_ibm_cidr` (String) The IP address in CIDR format for the IBM-side router in the BGP session. If you do not specify an address, IBM will assign one on your behalf. See the documentation for information on which IP ranges are allowed.
- `labels` (Set of String) Label value linked to an object.
- `po_number` (String) Purchase order number or identifier of a service.
- `published_quote_line_uuid` (String) UUID of the published quote line with which this connection should be associated.
- `src_svlan` (Number) Valid S-VLAN range is from 4-4094, inclusive.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `etl` (Number) Early Termination Liability (ETL) fees apply when terminating a service before its term ends. ETL is prorated to the remaining contract days.
- `gateway_id` (String) The IBM Gateway ID.
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)




## Import

Import an IBM hosted connection using its circuit ID.

```bash
terraform import packetfabric_cs_ibm_hosted_connection.cs_conn1_hosted_ibm PF-CC-WDC-NYC-1726496-PF
```