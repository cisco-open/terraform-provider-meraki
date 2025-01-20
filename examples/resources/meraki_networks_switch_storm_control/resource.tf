
resource "meraki_networks_switch_storm_control" "example" {

  broadcast_threshold                        = 30
  multicast_threshold                        = 30
  network_id                                 = "string"
  treat_these_traffic_types_as_one_threshold = ["broadcast", "multicast"]
  unknown_unicast_threshold                  = 30
}

output "meraki_networks_switch_storm_control_example" {
  value = meraki_networks_switch_storm_control.example
}