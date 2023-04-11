resource "packetfabric_cloud_provider_credential_google" "google_creds_staged" {
  provider               = packetfabric
  description            = "Google Staging Environement"
  google_service_account = var.google_service_account # or use env var GOOGLE_CREDENTIALS
}