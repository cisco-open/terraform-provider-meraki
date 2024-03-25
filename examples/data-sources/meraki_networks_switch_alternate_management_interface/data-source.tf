
data "meraki_networks_switch_alternate_management_interface" "example" {

  network_id = "string"
}

output "meraki_networks_switch_alternate_management_interface_example" {
  value = data.meraki_networks_switch_alternate_management_interface.example.item
}
