
resource "meraki_organizations_licenses_assign_seats" "example" {

  organization_id = "string"
  parameters = {

    license_id = "1234"
    network_id = "N_24329156"
    seat_count = 20
  }
}

output "meraki_organizations_licenses_assign_seats_example" {
  value = meraki_organizations_licenses_assign_seats.example
}