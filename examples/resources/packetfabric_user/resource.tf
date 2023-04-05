resource "packetfabric_user" "user1" {
  provider   = packetfabric
  first_name = "Alice"
  last_name  = "Thomas"
  email      = "alice@mycompany.com"
  phone      = "2065434573"
  login      = "alice@mycompany.com"
  password   = "secret"
  timezone   = "America/Vancouver"
  group      = "read-only"
}