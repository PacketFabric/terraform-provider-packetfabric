data "packetfabric_documents" "current" {
  provider = packetfabric
}
output "my-documents" {
  value = data.packetfabric_documents.current
}