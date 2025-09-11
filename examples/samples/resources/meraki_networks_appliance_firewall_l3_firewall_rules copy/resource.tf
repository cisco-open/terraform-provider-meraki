terraform {
  required_providers {
    meraki = {
      version = "1.2.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_appliance_firewall_l3_firewall_rules" "example" {

  network_id = "L_828099381482777130"
  rules = [{

    comment        = "Allow TCP traffic to subnet with HTTP servers."
    dest_cidr      = "200.168.1.0/24"
    dest_port      = "443"
    policy         = "allow"
    protocol       = "tcp"
    src_cidr       = "icmp"
    src_port       = "8080"
    syslog_enabled = false
  }]
}

output "meraki_networks_appliance_firewall_l3_firewall_rules_example" {
  value = meraki_networks_appliance_firewall_l3_firewall_rules.example
}