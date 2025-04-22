
data "meraki_organizations_devices_system_memory_usage_history_by_interval" "example" {

  ending_before   = "string"
  interval        = 1
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  product_types   = ["string"]
  serials         = ["string"]
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_devices_system_memory_usage_history_by_interval_example" {
  value = data.meraki_organizations_devices_system_memory_usage_history_by_interval.example.item
}
