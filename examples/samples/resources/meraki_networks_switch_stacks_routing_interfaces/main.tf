terraform {
  required_providers {
    meraki = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_switch_stacks_routing_interfaces" "example" {

  default_gateway = "192.168.1.1"
  interface_ip    = "192.168.1.2"
  ipv6 = {

    address         = "1:2:3:4::1"
    assignment_mode = "static"
    gateway         = "1:2:3:4::2"
    prefix          = "1:2:3:4::/48"
  }
  multicast_routing = "disabled"
  name              = "L3 interface"
  network_id        = "L_828099381482775374"
  ospf_settings = {

    area               = "0"
    cost               = 1
    is_passive_enabled = true
  }
  subnet          = "192.168.1.0/24"
  switch_stack_id = "string"
  vlan_id         = 100
}