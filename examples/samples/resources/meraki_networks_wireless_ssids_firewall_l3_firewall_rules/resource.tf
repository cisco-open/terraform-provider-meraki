terraform {
  required_providers {
    meraki = {
      version = "1.2.2-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_wireless_ssids_firewall_l3_firewall_rules" "test" {
  network_id       = "L_828099381482771185"
  number           = 0
  allow_lan_access = false
  rules = [
    {
      dest_cidr = "any"
      comment   = "Allow the Accessable VLAN"
      policy    = "allow"
      protocol  = "any"
      dest_port = "any"
    },
    {
      dest_cidr = "any"
      comment   = "Allow the Accessable VLAN 2"
      policy    = "allow"
      protocol  = "any"
      dest_port = "any"
      ip_ver    = "ipv6"
    },
  ]
}

