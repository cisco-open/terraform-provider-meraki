
resource "meraki_organizations_devices_controller_migrations_create" "example" {

  organization_id = "string"
  parameters = {

    serials = ["QABC-DEFG-HIJK"]
    target  = "wirelessController"
  }
}

output "meraki_organizations_devices_controller_migrations_create_example" {
  value = meraki_organizations_devices_controller_migrations_create.example
}