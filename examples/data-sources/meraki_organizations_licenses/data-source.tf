
data "meraki_organizations_licenses" "example" {

  license_id      = "string"
  organization_id = "string"
}

output "meraki_organizations_licenses_example" {
  value = data.meraki_organizations_licenses.example.item
}
