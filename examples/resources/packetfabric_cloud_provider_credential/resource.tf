resource "packetfabric_cloud_provider_credential" "aws_creds_staged" {
  provider       = packetfabric
  cloud_provider = "aws"
  description    = "AWS Staging Environement"
  aws_access_key = var.pf_aws_key
  aws_secret_key = var.pf_aws_secret
}

output "packetfabric_cloud_provider_credential" {
  value     = packetfabric_cloud_provider_credential.aws_creds_staged
  sensitive = true
}
