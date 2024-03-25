
resource "meraki_networks_switch_dscp_to_cos_mappings" "example" {

  mappings = [{

    cos   = 1
    dscp  = 1
    title = "Video"
  }]
  network_id = "string"
}

output "meraki_networks_switch_dscp_to_cos_mappings_example" {
  value = meraki_networks_switch_dscp_to_cos_mappings.example
}