
data "meraki_organizations_wireless_controller_devices_interfaces_l2_by_device" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_wireless_controller_devices_interfaces_l2_by_device_example" {
  value = data.meraki_organizations_wireless_controller_devices_interfaces_l2_by_device.example.item
}
