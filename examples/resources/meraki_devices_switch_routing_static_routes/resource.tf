
resource "meraki_devices_switch_routing_static_routes" "example" {

  advertise_via_ospf_enabled      = false
  name                            = "My route"
  next_hop_ip                     = "1.2.3.4"
  prefer_over_ospf_routes_enabled = false
  serial                          = "string"
  subnet                          = "192.168.1.0/24"
}

output "meraki_devices_switch_routing_static_routes_example" {
  value = meraki_devices_switch_routing_static_routes.example
}