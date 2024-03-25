
data "meraki_organizations_adaptive_policy_overview" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_overview_example" {
  value = data.meraki_organizations_adaptive_policy_overview.example.item
}
