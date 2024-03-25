
data "meraki_organizations_branding_policies_priorities" "example" {

  organization_id = "string"
}

output "meraki_organizations_branding_policies_priorities_example" {
  value = data.meraki_organizations_branding_policies_priorities.example.item
}
