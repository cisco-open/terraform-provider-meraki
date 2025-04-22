
data "meraki_networks_switch_stacks_routing_interfaces" "example" {

  network_id      = "string"
  protocol        = "string"
  switch_stack_id = "string"
}

output "meraki_networks_switch_stacks_routing_interfaces_example" {
  value = data.meraki_networks_switch_stacks_routing_interfaces.example.items
}

data "meraki_networks_switch_stacks_routing_interfaces" "example" {

  network_id      = "string"
  protocol        = "string"
  switch_stack_id = "string"
}

output "meraki_networks_switch_stacks_routing_interfaces_example" {
  value = data.meraki_networks_switch_stacks_routing_interfaces.example.item
}
