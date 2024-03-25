
data "meraki_networks_switch_stacks_routing_interfaces_dhcp" "example" {

  interface_id    = "string"
  network_id      = "string"
  switch_stack_id = "string"
}

output "meraki_networks_switch_stacks_routing_interfaces_dhcp_example" {
  value = data.meraki_networks_switch_stacks_routing_interfaces_dhcp.example.item
}
