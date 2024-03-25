
data "meraki_organizations_appliance_vpn_vpn_firewall_rules" "example" {

  organization_id = "string"
}

output "meraki_organizations_appliance_vpn_vpn_firewall_rules_example" {
  value = data.meraki_organizations_appliance_vpn_vpn_firewall_rules.example.item
}
