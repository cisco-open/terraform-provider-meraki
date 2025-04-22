
resource "meraki_networks_switch_stacks_routing_interfaces" "example" {

  default_gateway = "192.168.1.1"
  interface_ip    = "192.168.1.2"
  ipv6 = {

    address         = "2001:db8::1"
    assignment_mode = "static"
    gateway         = "2001:db8::2"
    prefix          = "2001:db8::/32"
  }
  multicast_routing = "disabled"
  name              = "L3 interface"
  network_id        = "string"
  ospf_settings = {

    area               = "0"
    cost               = 1
    is_passive_enabled = true
  }
  subnet          = "192.168.1.0/24"
  switch_stack_id = "string"
  vlan_id         = 100
}

output "meraki_networks_switch_stacks_routing_interfaces_example" {
  value = meraki_networks_switch_stacks_routing_interfaces.example
}