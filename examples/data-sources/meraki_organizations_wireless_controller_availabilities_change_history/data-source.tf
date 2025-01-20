
data "meraki_organizations_wireless_controller_availabilities_change_history" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_wireless_controller_availabilities_change_history_example" {
  value = data.meraki_organizations_wireless_controller_availabilities_change_history.example.item
}
