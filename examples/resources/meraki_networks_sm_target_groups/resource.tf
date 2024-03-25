
resource "meraki_networks_sm_target_groups" "example" {

  name       = "My target group"
  network_id = "string"
  scope      = "none"
}

output "meraki_networks_sm_target_groups_example" {
  value = meraki_networks_sm_target_groups.example
}