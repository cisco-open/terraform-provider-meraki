terraform {
  required_providers {
    meraki = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_appliance_firewall_l7_firewall_rules" "this" {
  network_id = "L_828099381482771185"
  rules      = []
}

