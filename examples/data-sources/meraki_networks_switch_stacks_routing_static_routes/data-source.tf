
data "meraki_networks_switch_stacks_routing_static_routes" "example" {

  network_id      = "string"
  switch_stack_id = "string"
}

output "meraki_networks_switch_stacks_routing_static_routes_example" {
  value = data.meraki_networks_switch_stacks_routing_static_routes.example.items
}

data "meraki_networks_switch_stacks_routing_static_routes" "example" {

  network_id      = "string"
  switch_stack_id = "string"
}

output "meraki_networks_switch_stacks_routing_static_routes_example" {
  value = data.meraki_networks_switch_stacks_routing_static_routes.example.item
}
