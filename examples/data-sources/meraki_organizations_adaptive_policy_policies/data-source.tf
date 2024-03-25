
data "meraki_organizations_adaptive_policy_policies" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_policies_example" {
  value = data.meraki_organizations_adaptive_policy_policies.example.items
}

data "meraki_organizations_adaptive_policy_policies" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_policies_example" {
  value = data.meraki_organizations_adaptive_policy_policies.example.item
}
