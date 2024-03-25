
data "meraki_organizations_sm_vpp_accounts" "example" {

  organization_id = "string"
}

output "meraki_organizations_sm_vpp_accounts_example" {
  value = data.meraki_organizations_sm_vpp_accounts.example.items
}

data "meraki_organizations_sm_vpp_accounts" "example" {

  organization_id = "string"
}

output "meraki_organizations_sm_vpp_accounts_example" {
  value = data.meraki_organizations_sm_vpp_accounts.example.item
}
