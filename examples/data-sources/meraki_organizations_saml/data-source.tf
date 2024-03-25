
data "meraki_organizations_saml" "example" {

  organization_id = "string"
}

output "meraki_organizations_saml_example" {
  value = data.meraki_organizations_saml.example.item
}
