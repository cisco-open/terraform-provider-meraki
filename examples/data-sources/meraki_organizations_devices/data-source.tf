
data "meraki_organizations_devices" "example" {

  configuration_updated_after = "string"
  ending_before               = "string"
  mac                         = "string"
  macs                        = ["string"]
  model                       = "string"
  models                      = ["string"]
  name                        = "string"
  network_ids                 = ["string"]
  organization_id             = "string"
  per_page                    = 1
  product_types               = ["string"]
  sensor_alert_profile_ids    = ["string"]
  sensor_metrics              = ["string"]
  serial                      = "string"
  serials                     = ["string"]
  starting_after              = "string"
  tags                        = ["string"]
  tags_filter_type            = "string"
}

output "meraki_organizations_devices_example" {
  value = data.meraki_organizations_devices.example.items
}
