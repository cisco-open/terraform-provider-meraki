
data "meraki_networks_switch_storm_control" "example" {

  network_id = "string"
}

output "meraki_networks_switch_storm_control_example" {
  value = data.meraki_networks_switch_storm_control.example.item
}
