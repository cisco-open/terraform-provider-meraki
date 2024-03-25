
data "meraki_organizations_adaptive_policy_groups" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_groups_example" {
  value = data.meraki_organizations_adaptive_policy_groups.example.items
}

data "meraki_organizations_adaptive_policy_groups" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_groups_example" {
  value = data.meraki_organizations_adaptive_policy_groups.example.item
}
