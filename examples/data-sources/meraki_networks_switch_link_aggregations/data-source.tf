
data "meraki_networks_switch_link_aggregations" "example" {

  network_id = "string"
}

output "meraki_networks_switch_link_aggregations_example" {
  value = data.meraki_networks_switch_link_aggregations.example.items
}
