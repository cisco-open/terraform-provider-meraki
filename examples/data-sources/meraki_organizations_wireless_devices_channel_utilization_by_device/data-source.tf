
data "meraki_organizations_wireless_devices_channel_utilization_by_device" "example" {

  ending_before   = "string"
  interval        = 1
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_wireless_devices_channel_utilization_by_device_example" {
  value = data.meraki_organizations_wireless_devices_channel_utilization_by_device.example.items
}
