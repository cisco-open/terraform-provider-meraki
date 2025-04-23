
resource "meraki_devices_switch_routing_interfaces" "example" {

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
  ospf_settings = {

    area               = "0"
    cost               = 1
    is_passive_enabled = true
  }
  serial  = "string"
  subnet  = "192.168.1.0/24"
  vlan_id = 100
}

output "meraki_devices_switch_routing_interfaces_example" {
  value = meraki_devices_switch_routing_interfaces.example
}