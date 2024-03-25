
data "meraki_networks_wireless_ssids_traffic_shaping_rules" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_traffic_shaping_rules_example" {
  value = data.meraki_networks_wireless_ssids_traffic_shaping_rules.example.item
}
