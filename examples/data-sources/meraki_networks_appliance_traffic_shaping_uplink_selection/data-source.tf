
data "meraki_networks_appliance_traffic_shaping_uplink_selection" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_traffic_shaping_uplink_selection_example" {
  value = data.meraki_networks_appliance_traffic_shaping_uplink_selection.example.item
}
