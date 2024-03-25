
data "meraki_networks_appliance_traffic_shaping" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_traffic_shaping_example" {
  value = data.meraki_networks_appliance_traffic_shaping.example.item
}
