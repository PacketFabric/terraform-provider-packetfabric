resource "packetfabric_streaming_events" "example" {
  provider = packetfabric
  streams {
    type   = "customer"
    events = ["auth", "physical_interface"]
  }
  streams {
    type   = "port"
    ifds   = ["PF-AP-LAB1-2999387", "PF-AP-LAB6-3001683"]
    events = ["errors", "etherstats", "metrics", "optical"]
  }
  # streams {
  #   type   = "vc"
  #   vcs    = ["PF-BC-NYC-LAB-3011206-PF&PF-AP-LAB6-3001683", "PF-BC-NYC-LAB-3011206-PF&PF-AP-LAB1-2999387"]
  #   events = ["metrics"]
  # }
}

resource "null_resource" "stream_events" {
  provisioner "local-exec" {
    command = "python3 packetfabric_streaming_events.py --subscription_uuid ${packetfabric_streaming_events.example.id} --duration_seconds 60 --output_file pf_events.json"
  }
}