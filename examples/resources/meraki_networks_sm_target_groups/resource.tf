
resource "meraki_networks_sm_target_groups" "example" {

  name       = "Target group name"
  network_id = "string"
  scope      = "withAny, tag1, tag2"
}

output "meraki_networks_sm_target_groups_example" {
  value = meraki_networks_sm_target_groups.example
}