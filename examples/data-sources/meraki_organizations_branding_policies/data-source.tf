
data "meraki_organizations_branding_policies" "example" {

  organization_id = "string"
}

output "meraki_organizations_branding_policies_example" {
  value = data.meraki_organizations_branding_policies.example.items
}

data "meraki_organizations_branding_policies" "example" {

  organization_id = "string"
}

output "meraki_organizations_branding_policies_example" {
  value = data.meraki_organizations_branding_policies.example.item
}
