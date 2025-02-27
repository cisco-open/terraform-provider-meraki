
data "meraki_organizations_wireless_controller_clients_overview_history_by_device_by_interval" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  resolution      = 1
  serials         = ["string"]
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_wireless_controller_clients_overview_history_by_device_by_interval_example" {
  value = data.meraki_organizations_wireless_controller_clients_overview_history_by_device_by_interval.example.item
}
