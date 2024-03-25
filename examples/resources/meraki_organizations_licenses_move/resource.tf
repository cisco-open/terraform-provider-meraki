
resource "meraki_organizations_licenses_move" "example" {

  organization_id = "string"
  parameters = {

    dest_organization_id = "2930418"
    license_ids          = ["123", "456"]
  }
}

output "meraki_organizations_licenses_move_example" {
  value = meraki_organizations_licenses_move.example
}