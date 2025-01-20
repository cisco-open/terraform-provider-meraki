
data "meraki_organizations_wireless_rf_profiles_assignments_by_device" "example" {

  ending_before   = "string"
  mac             = "string"
  macs            = ["string"]
  model           = "string"
  models          = ["string"]
  name            = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  product_types   = ["string"]
  serial          = "string"
  serials         = ["string"]
  starting_after  = "string"
}

output "meraki_organizations_wireless_rf_profiles_assignments_by_device_example" {
  value = data.meraki_organizations_wireless_rf_profiles_assignments_by_device.example.items
}
