
resource "meraki_organizations_appliance_vpn_vpn_firewall_rules" "example" {

  organization_id = "string"
  rules = [{

    comment        = "Allow TCP traffic to subnet with HTTP servers."
    dest_cidr      = "192.168.1.0/24"
    dest_port      = "443"
    policy         = "allow"
    protocol       = "tcp"
    src_cidr       = "Any"
    src_port       = "Any"
    syslog_enabled = false
  }]
  syslog_default_rule = false
}

output "meraki_organizations_appliance_vpn_vpn_firewall_rules_example" {
  value = meraki_organizations_appliance_vpn_vpn_firewall_rules.example
}