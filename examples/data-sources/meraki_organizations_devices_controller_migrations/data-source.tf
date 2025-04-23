
data "meraki_organizations_devices_controller_migrations" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
  target          = "string"
}

output "meraki_organizations_devices_controller_migrations_example" {
  value = data.meraki_organizations_devices_controller_migrations.example.item
}
