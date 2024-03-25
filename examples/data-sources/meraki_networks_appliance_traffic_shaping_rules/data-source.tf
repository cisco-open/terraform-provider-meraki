
data "meraki_networks_appliance_traffic_shaping_rules" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_traffic_shaping_rules_example" {
  value = data.meraki_networks_appliance_traffic_shaping_rules.example.item
}
