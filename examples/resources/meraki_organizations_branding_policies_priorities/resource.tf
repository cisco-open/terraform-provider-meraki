
resource "meraki_organizations_branding_policies_priorities" "example" {

  branding_policy_ids = ["123", "456", "789"]
  organization_id     = "string"
}

output "meraki_organizations_branding_policies_priorities_example" {
  value = meraki_organizations_branding_policies_priorities.example
}