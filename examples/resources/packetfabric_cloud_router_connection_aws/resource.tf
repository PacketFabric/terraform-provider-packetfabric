resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
}

resource "packetfabric_cloud_router_connection_aws" "crc1" {
  provider    = packetfabric
  circuit_id  = packetfabric_cloud_router.cr1.id
  maybe_nat   = false
  description = "hello world"
  pop         = "PDX2"
  zone        = "A"
  is_public   = false
  speed       = "1Gbps"
  labels      = sort(["terraform", "dev"])
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
  provider   = aws
  name       = "hello world"
  depends_on = [time_sleep.wait_aws_connection]
}

# Accept the connection
resource "aws_dx_connection_confirmation" "confirmation" {
  provider      = aws
  connection_id = data.aws_dx_connection.current.id
}