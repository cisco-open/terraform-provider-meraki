terraform {
  required_providers {
    meraki = {
      version = "1.0.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_appliance_ports" "my_mx" {

  allowed_vlans = "all"
  enabled       = true
  network_id    = "L_828099381482771185"
  port_id       = "4"
  type          = "trunk"
  vlan          = 1234
}

