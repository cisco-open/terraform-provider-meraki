terraform {
  required_providers {
    meraki = {
      version = "1.1.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_devices_switch_routing_interfaces_dhcp" "example" {
  dhcp_lease_time = "1 day"
  dhcp_mode       = "dhcpServer"
  interface_id    = "828099381482756008"
  serial          = "QBSB-VLHZ-JQCN"
}

output "meraki_devices_switch_routing_interfaces_dhcp_example" {
  value = meraki_devices_switch_routing_interfaces_dhcp.example
}