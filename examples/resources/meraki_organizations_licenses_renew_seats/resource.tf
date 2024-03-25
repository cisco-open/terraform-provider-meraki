
resource "meraki_organizations_licenses_renew_seats" "example" {

  organization_id = "string"
  parameters = {

    license_id_to_renew = "123"
    unused_license_id   = "1234"
  }
}

output "meraki_organizations_licenses_renew_seats_example" {
  value = meraki_organizations_licenses_renew_seats.example
}