
data "meraki_networks_sm_target_groups" "example" {

  network_id   = "string"
  with_details = false
}

output "meraki_networks_sm_target_groups_example" {
  value = data.meraki_networks_sm_target_groups.example.items
}

data "meraki_networks_sm_target_groups" "example" {

  network_id   = "string"
  with_details = false
}

output "meraki_networks_sm_target_groups_example" {
  value = data.meraki_networks_sm_target_groups.example.item
}
