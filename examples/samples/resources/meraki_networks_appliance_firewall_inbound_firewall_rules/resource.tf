terraform {
  required_providers {
    meraki = {
      version = "0.2.12-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_appliance_firewall_inbound_firewall_rules" "this" {
  network_id = "L_828099381482775375"
  rules = [
    {
      comment        = "Allow TCP traffic to subnet with HTTP servers."
      dest_cidr      = "2001:db8::/64"
      dest_port      = "443"
      policy         = "allow"
      protocol       = "tcp"
      src_cidr       = "Any"
      src_port       = "Any"
      syslog_enabled = false
    }
  ]
  syslog_default_rule = true
}