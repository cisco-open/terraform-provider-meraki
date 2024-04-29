
data "meraki_organizations_devices_availabilities_change_history" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  product_types   = ["string"]
  serials         = ["string"]
  starting_after  = "string"
  statuses        = ["string"]
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_devices_availabilities_change_history_example" {
  value = data.meraki_organizations_devices_availabilities_change_history.example.items
}
