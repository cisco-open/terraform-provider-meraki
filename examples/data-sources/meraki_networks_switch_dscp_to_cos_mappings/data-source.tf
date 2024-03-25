
data "meraki_networks_switch_dscp_to_cos_mappings" "example" {

  network_id = "string"
}

output "meraki_networks_switch_dscp_to_cos_mappings_example" {
  value = data.meraki_networks_switch_dscp_to_cos_mappings.example.item
}
