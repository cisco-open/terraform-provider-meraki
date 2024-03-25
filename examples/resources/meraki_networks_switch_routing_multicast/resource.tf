
resource "meraki_networks_switch_routing_multicast" "example" {

  default_settings = {

    flood_unknown_multicast_traffic_enabled = true
    igmp_snooping_enabled                   = true
  }
  network_id = "string"
  overrides = [{

    flood_unknown_multicast_traffic_enabled = true
    igmp_snooping_enabled                   = true
    switches                                = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }]
}

output "meraki_networks_switch_routing_multicast_example" {
  value = meraki_networks_switch_routing_multicast.example
}