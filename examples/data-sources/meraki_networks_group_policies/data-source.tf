
data "meraki_networks_group_policies" "example" {

  network_id = "string"
}

output "meraki_networks_group_policies_example" {
  value = data.meraki_networks_group_policies.example.items
}

data "meraki_networks_group_policies" "example" {

  network_id = "string"
}

output "meraki_networks_group_policies_example" {
  value = data.meraki_networks_group_policies.example.item
}
