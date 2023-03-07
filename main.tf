terraform {
 required_providers {
   packetfabric = {
     source  = "PacketFabric/packetfabric"
     version = "0.6.0"
   }
 }
}

data "packetfabric_locations_cloud" "cloud_location_1" {
  provider              = packetfabric
  cloud_provider        = "aws"
  cloud_connection_type = "hosted"
}

output "packetfabric_locations_cloud" {
  value = data.packetfabric_locations_cloud.cloud_location_1.cloud_locations[0].pop
}

# data "packetfabric_locations_cloud" "pf_testacc_cr_locations_a110fb7f_36ea_4f84_95fc_1bde1af8bf3c" {
#   provider              = packetfabric
#   cloud_provider        = "aws"
#   cloud_connection_type = "hosted"
# }

# data "packetfabric_locations_port_availability" "p" {
#     provider = packetfabric
#     pop = data.packetfabric_locations_cloud.pf_testacc_cr_locations_a110fb7f_36ea_4f84_95fc_1bde1af8bf3c.cloud_locations[0].pop
# }

# output "packetfabric_locations_port_availability" {
#   value = data.packetfabric_locations_port_availability.p
# }

# resource "packetfabric_port" "pf_62a145f3_4da4_41d1_96a6_86fad72313af" {
#   provider          = packetfabric
#   description       = "pf_testacc_packetfabric_port_9c4982ac_5b12_43a4_b90c_7c138f90f5f4"
#   media             = "ZX"
#   pop               = "LAB1"
#   speed             = "1Gbps"
#   subscription_term = 1
#   enabled          = false
# }
# resource "packetfabric_cs_aws_hosted_connection" "pf_70ed7341_bd0e_407a_81a1_8efc1c607be9" {
#   provider       = packetfabric
#   description    = "pf_testacc_packetfabric_cs_aws_hosted_connection_7b606c53_c750_4ff7_8b81_9d9de51dc098"
#   aws_account_id = "777804547227"
#   port           = packetfabric_port.pf_62a145f3_4da4_41d1_96a6_86fad72313af.id
#   speed          = "50Mbps"
#   pop            = "LAB1"
#   vlan           = 100
# }

# resource "packetfabric_cs_azure_dedicated_connection" "pf_9f6d1e2b_d47f_4ca7_8817_618b9a14a6ff" {
#   provider          = packetfabric
#   description       = "pf_testacc_packetfabric_cs_azure_dedicated_connection_2f6b57ea_6719_4ab3_8ce4_4d05f3139e97"
#   pop               = "LAB4"
#   subscription_term = 1
#   service_class     = "metro"
#   encapsulation     = "qinq"
#   port_category     = "primary"
#   speed             = "10Gbps"
# }

resource "time_sleep" "wait_30_seconds" {
  depends_on = [data.packetfabric_locations_cloud.cloud_location_1]

  destroy_duration = "30s"
}