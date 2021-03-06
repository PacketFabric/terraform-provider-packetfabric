---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_aws_cloud_router_connection Data Source - terraform-provider-packetfabric"
subcategory: ""
description: |-

---

# packetfabric_aws_cloud_router_connection (Data Source)



## Data Example

```terraform
[
  {
    "port_type": "hosted",
    "connection_type": "cloud_hosted",
    "port_circuit_id": "PF-AE-1234",
    "pending_delete": true,
    "deleted": true,
    "speed": "1Gbps",
    "state": "Requested",
    "cloud_circuit_id": "PF-AP-LAX1-1002",
    "account_uuid": "a2115890-ed02-4795-a6dd-c485bec3529c",
    "service_class": "metro",
    "service_provider": "aws",
    "service_type": "cr_connection",
    "description": "AWS connection for Foo Corp.",
    "uuid": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "cloud_provider_connection_id": "dxcon-fgadaaa1",
    "cloud_settings": {},
    "user_uuid": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "customer_uuid": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "time_created": "2022-05-25T06:04:49.342Z",
    "time_updated": "2022-05-25T06:04:49.342Z",
    "cloud_provider": {
      "pop": "LAX1",
      "site": "us-west-1"
    },
    "pop": "LAX1",
    "site": "us-west-1",
    "bgp_state": "string",
    "cloud_router_circuit_id": "PF-L3-CUST-2001",
    "nat_capable": true
  }
]
```

## Schema

### Required

- `circuit_id` (String)

### Read-Only

- `packetfabric_aws_cloud_connections` (List of Object) (see [below for nested schema](#nestedatt--packetfabric_aws_cloud_connections))
- `id` (String) The ID of this resource.

<a id="nestedatt--packetfabric_aws_cloud_connections"></a>
### Nested Schema for `packetfabric_aws_cloud_connections`

Read-Only:

- `account_uuid` (String) The UUID of the PacketFabric contact that will be billed.
      Example: a2115890-ed02-4795-a6dd-c485bec12345
- `bgp_state` (String) The status of the BGP session
      Enum: established, configuring, fetching, etc.
- `cloud_circuit_id` (String) The unique PF circuit ID for this connection.
      Example: \"PF-AP-LAX1-1002\"
- `cloud_provider` (Set of Object) (see [below for nested schema](#nestedobjatt--packetfabric_aws_cloud_connections--cloud_provider))
- `cloud_provider_connection_id` (String) The cloud provider specific connection ID, eg. the Amazon connection ID of the cloud router connection.
      Example: dxcon-fgadaaa1
- `cloud_router_circuit_id` (String) The circuit ID of the cloud router this connection is associated with.
      Example: PF-L3-CUST-2001
- `cloud_settings` (Set of Object) (see [below for nested schema](#nestedobjatt--packetfabric_aws_cloud_connections--cloud_settings))
- `connection_type` (String) The type of the connection.
      Enum: cloud_hosted, cloud_dedicated, ipsec, packetfabric
- `customer_uuid` (String) The UUID for the customer this connection belongs to
- `deleted` (Boolean) Whether or not the connection has been fully deleted.
- `description` (String) The description of this connection.
- `nat_capable` (Boolean) Indicates whether this connection supports NAT
- `pending_delete` (Boolean) Whether or not the connection is currently deleting.
- `pop` (String) Point of Presence for the cloud provider location
      Example: LAX1
- `port_circuit_id` (String) The circuit ID of the port to connect to the cloud router.
      Example "PF-AE-1234"
- `port_type` (String)
- `service_class` (String) The service class of the connection.
      Enum: metro, longhaul
- `service_provider` (String) The service provider of the connection.
      Enum: aws, azure, packet, google, ibm, salesforce, webex
- `service_type` (String) The type of connection, this will currently always be cr_connection.
      Enum: cr_connection
- `site` (String) Region short name
      Example: us-west-1
- `speed` (String) The speed of the connection.
      Enum: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps
- `state` (String) The state of the connection
      Enum: Requested, Active, Inactive, PendingDelete
- `time_created` (String) Date and time of connection creation
- `time_updated` (String) Date and time connection was last updated
- `user_uuid` (String) The UUID for the user this connection belongs to
- `uuid` (String) The UUID of the connection.

<a id="nestedobjatt--packetfabric_aws_cloud_connections--cloud_provider"></a>
### Nested Schema for `packetfabric_aws_cloud_connections.cloud_provider`

Read-Only:

- `pop` (String) Point of Presence for the cloud provider location
      Example: LAX1
- `site` (String) Region short name
      Example: us-west-1


<a id="nestedobjatt--packetfabric_aws_cloud_connections--cloud_settings"></a>
### Nested Schema for `packetfabric_aws_cloud_connections.cloud_settings`

Read-Only:

- `aws_account_id` (String)
- `aws_connection_id` (String)
- `aws_hosted_type` (String)
- `aws_region` (String)
- `nat_public_ip` (String)
- `public_ip` (String)
- `vlan_id_cust` (Number)
- `vlan_id_pf` (Number)
