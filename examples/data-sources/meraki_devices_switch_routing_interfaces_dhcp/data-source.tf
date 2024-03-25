
data "meraki_devices_switch_routing_interfaces_dhcp" "example" {

  interface_id = "string"
  serial       = "string"
}

output "meraki_devices_switch_routing_interfaces_dhcp_example" {
  value = data.meraki_devices_switch_routing_interfaces_dhcp.example.item
}
