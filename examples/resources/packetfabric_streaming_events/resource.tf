resource "packetfabric_streaming_events" "example" {
  provider = packetfabric
  streams {
    type   = "customer"
    events = ["auth", "physical_interface"]
  }

  streams {
    type   = "port"
    ifds   = ["PF-AP-ABC-1234", "PF-AE-DEF-5678"]
    events = ["errors", "etherstats", "metrics", "optical"]
  }

  streams {
    type   = "vc"
    vcs    = ["PF-BC-ABC-QRS-84274-PF&PF-AP-ABC-1234", "PF-CC-JKL-DEF-13758-PF&PF-AE-DEF-5678"]
    events = ["metrics"]
  }
}

data "packetfabric_streaming_events" "example" {
  provider        = packetfabric
  subscription_id = packetfabric_streaming_events.example.id
  stream_time     = 1 # min
}
output "packetfabric_streaming_events_result" {
  value = data.packetfabric_streaming_events.example.events
}
resource "local_file" "packetfabric_streaming_events_output" {
  content  = jsonencode(data.packetfabric_streaming_events.example.events)
  filename = "packetfabric_streaming_events_output.json"
}