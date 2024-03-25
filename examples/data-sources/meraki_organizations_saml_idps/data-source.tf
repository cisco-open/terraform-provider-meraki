
data "meraki_organizations_saml_idps" "example" {

  organization_id = "string"
}

output "meraki_organizations_saml_idps_example" {
  value = data.meraki_organizations_saml_idps.example.items
}

data "meraki_organizations_saml_idps" "example" {

  organization_id = "string"
}

output "meraki_organizations_saml_idps_example" {
  value = data.meraki_organizations_saml_idps.example.item
}
