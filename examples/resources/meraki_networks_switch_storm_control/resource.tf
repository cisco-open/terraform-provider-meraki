
resource "meraki_networks_switch_storm_control" "example" {

  broadcast_threshold       = 30
  multicast_threshold       = 30
  network_id                = "string"
  unknown_unicast_threshold = 30
}

output "meraki_networks_switch_storm_control_example" {
  value = meraki_networks_switch_storm_control.example
}