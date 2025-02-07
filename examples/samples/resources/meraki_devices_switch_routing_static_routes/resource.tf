terraform {
  required_providers {
    meraki = {
      version = "1.0.1-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_devices_switch_routing_static_routes" "example" {

  # advertise_via_ospf_enabled      = false
  name        = "My route"
  next_hop_ip = "192.168.1.4"
  serial      = "QBSB-VLHZ-JQCN"
  subnet      = "192.168.2.0/24"
}

output "meraki_devices_switch_routing_static_routes_example" {
  value = meraki_devices_switch_routing_static_routes.example
}