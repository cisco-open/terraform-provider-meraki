
data "meraki_organizations_saml_roles" "example" {

  organization_id = "string"
}

output "meraki_organizations_saml_roles_example" {
  value = data.meraki_organizations_saml_roles.example.items
}

data "meraki_organizations_saml_roles" "example" {

  organization_id = "string"
}

output "meraki_organizations_saml_roles_example" {
  value = data.meraki_organizations_saml_roles.example.item
}
