
data "meraki_networks_switch_access_control_lists" "example" {

  network_id = "string"
}

output "meraki_networks_switch_access_control_lists_example" {
  value = data.meraki_networks_switch_access_control_lists.example.item
}
