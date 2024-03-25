
data "meraki_networks_switch_access_policies" "example" {

  network_id = "string"
}

output "meraki_networks_switch_access_policies_example" {
  value = data.meraki_networks_switch_access_policies.example.items
}

data "meraki_networks_switch_access_policies" "example" {

  network_id = "string"
}

output "meraki_networks_switch_access_policies_example" {
  value = data.meraki_networks_switch_access_policies.example.item
}
