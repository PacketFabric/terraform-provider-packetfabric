data "packetfabric_document" "current" {
  provider = packetfabric
}
output "my-documents" {
  value = data.packetfabric_document.current
}