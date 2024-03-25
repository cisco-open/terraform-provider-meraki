
resource "meraki_organizations_saml" "example" {

  enabled         = true
  organization_id = "string"
}

output "meraki_organizations_saml_example" {
  value = meraki_organizations_saml.example
}