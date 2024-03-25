
resource "meraki_organizations_licenses" "example" {

  device_serial   = "Q234-ABCD-5678"
  license_id      = "string"
  organization_id = "string"
}

output "meraki_organizations_licenses_example" {
  value = meraki_organizations_licenses.example
}