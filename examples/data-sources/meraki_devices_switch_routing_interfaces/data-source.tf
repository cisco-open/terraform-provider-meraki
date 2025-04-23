
data "meraki_devices_switch_routing_interfaces" "example" {

  protocol = "string"
  serial   = "string"
}

output "meraki_devices_switch_routing_interfaces_example" {
  value = data.meraki_devices_switch_routing_interfaces.example.items
}

data "meraki_devices_switch_routing_interfaces" "example" {

  protocol = "string"
  serial   = "string"
}

output "meraki_devices_switch_routing_interfaces_example" {
  value = data.meraki_devices_switch_routing_interfaces.example.item
}
