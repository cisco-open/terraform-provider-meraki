
resource "meraki_organizations_licenses_move_seats" "example" {

  organization_id = "string"
  parameters = {

    dest_organization_id = "2930418"
    license_id           = "1234"
    seat_count           = 20
  }
}

output "meraki_organizations_licenses_move_seats_example" {
  value = meraki_organizations_licenses_move_seats.example
}