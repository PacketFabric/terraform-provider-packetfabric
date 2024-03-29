resource "packetfabric_cloud_provider_credential_aws" "aws_creds_staged" {
  provider       = packetfabric
  description    = "AWS Staging Environement"
  aws_access_key = var.pf_aws_key    # or use env var AWS_ACCESS_KEY_ID
  aws_secret_key = var.pf_aws_secret # or use env var AWS_SECRET_ACCESS_KEY
}