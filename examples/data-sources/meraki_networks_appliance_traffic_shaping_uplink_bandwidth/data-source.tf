
data "meraki_networks_appliance_traffic_shaping_uplink_bandwidth" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_traffic_shaping_uplink_bandwidth_example" {
  value = data.meraki_networks_appliance_traffic_shaping_uplink_bandwidth.example.item
}
