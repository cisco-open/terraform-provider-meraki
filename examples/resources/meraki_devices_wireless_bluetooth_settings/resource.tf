
resource "meraki_devices_wireless_bluetooth_settings" "example" {

  major  = 13
  minor  = 125
  serial = "string"
  uuid   = "00000000-0000-0000-000-000000000000"
}

output "meraki_devices_wireless_bluetooth_settings_example" {
  value = meraki_devices_wireless_bluetooth_settings.example
}