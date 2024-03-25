
data "meraki_organizations_appliance_vpn_third_party_vpnpeers" "example" {

  organization_id = "string"
}

output "meraki_organizations_appliance_vpn_third_party_vpnpeers_example" {
  value = data.meraki_organizations_appliance_vpn_third_party_vpnpeers.example.item
}
