
data "meraki_organizations_switch_ports_usage_history_by_device_by_interval" "example" {

  configuration_updated_after = "string"
  ending_before               = "string"
  interval                    = 1
  mac                         = "string"
  macs                        = ["string"]
  name                        = "string"
  network_ids                 = ["string"]
  organization_id             = "string"
  per_page                    = 1
  port_profile_ids            = ["string"]
  serial                      = "string"
  serials                     = ["string"]
  starting_after              = "string"
  t0                          = "string"
  t1                          = "string"
  timespan                    = 1.0
}

output "meraki_organizations_switch_ports_usage_history_by_device_by_interval_example" {
  value = data.meraki_organizations_switch_ports_usage_history_by_device_by_interval.example.item
}
