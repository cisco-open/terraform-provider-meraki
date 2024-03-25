
data "meraki_networks_traffic_shaping_dscp_tagging_options" "example" {

  network_id = "string"
}

output "meraki_networks_traffic_shaping_dscp_tagging_options_example" {
  value = data.meraki_networks_traffic_shaping_dscp_tagging_options.example.items
}
