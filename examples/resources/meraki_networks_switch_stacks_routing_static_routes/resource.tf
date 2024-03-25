
resource "meraki_networks_switch_stacks_routing_static_routes" "example" {

  advertise_via_ospf_enabled      = false
  name                            = "My route"
  network_id                      = "string"
  next_hop_ip                     = "1.2.3.4"
  prefer_over_ospf_routes_enabled = false
  subnet                          = "192.168.1.0/24"
  switch_stack_id                 = "string"
}

output "meraki_networks_switch_stacks_routing_static_routes_example" {
  value = meraki_networks_switch_stacks_routing_static_routes.example
}