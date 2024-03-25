
data "meraki_devices_switch_routing_static_routes" "example" {

  serial = "string"
}

output "meraki_devices_switch_routing_static_routes_example" {
  value = data.meraki_devices_switch_routing_static_routes.example.items
}

data "meraki_devices_switch_routing_static_routes" "example" {

  serial = "string"
}

output "meraki_devices_switch_routing_static_routes_example" {
  value = data.meraki_devices_switch_routing_static_routes.example.item
}
