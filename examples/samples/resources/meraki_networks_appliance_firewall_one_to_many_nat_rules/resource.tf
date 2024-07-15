terraform {
  required_providers {
    meraki = {
      version = "0.2.6-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_appliance_firewall_one_to_many_nat_rules" "this" {
  network_id = "L_828099381482771185"
  rules = [
    
  ]
}