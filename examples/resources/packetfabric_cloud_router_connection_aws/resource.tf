resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = var.pf_cr_asn
  name     = var.pf_cr_name
  capacity = var.pf_cr_capacity
  regions  = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_aws" "crc1" {
  provider    = packetfabric
  circuit_id  = packetfabric_cloud_router.cr1.id
  maybe_nat   = var.pf_crc_maybe_nat
  description = var.pf_crc_description
  pop         = var.pf_crc_pop
  zone        = var.pf_crc_zone
  is_public   = var.pf_crc_is_public
  speed       = var.pf_crc_speed
}

output "packetfabric_cloud_router_connection_aws" {
  value = packetfabric_cloud_router_connection_aws.crc1
}

# Wait for the connection to show up in AWS
resource "time_sleep" "wait_aws_connection" {
  create_duration = "2m"
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc1
  ]
}

# Retrieve the Direct Connect connections in AWS
data "aws_dx_connection" "current" {
  provider = aws
  name     = "${var.pf_crc_description}"
  depends_on = [time_sleep.wait_aws_connection]
}

# Accept the connection
resource "aws_dx_connection_confirmation" "confirmation" {
  provider      = aws
  connection_id = data.aws_dx_connection.current.id
}