terraform {
  required_providers {
    meraki = {
      version = "1.1.3-beta"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}


provider "meraki" {
  meraki_debug    = "true"
  meraki_base_url = "http://localhost:3001"
}

data "meraki_organizations_appliance_vpn_third_party_vpnpeers" "example" {

  organization_id = "828099381482762270"
}

# output "meraki_organizations_appliance_vpn_third_party_vpnpeers_example" {
#   value = data.meraki_organizations_appliance_vpn_third_party_vpnpeers.example.item
# }
