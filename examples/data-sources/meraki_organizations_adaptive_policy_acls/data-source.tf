
data "meraki_organizations_adaptive_policy_acls" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_acls_example" {
  value = data.meraki_organizations_adaptive_policy_acls.example.items
}

data "meraki_organizations_adaptive_policy_acls" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_acls_example" {
  value = data.meraki_organizations_adaptive_policy_acls.example.item
}
