
data "meraki_devices_wireless_bluetooth_settings" "example" {

  serial = "string"
}

output "meraki_devices_wireless_bluetooth_settings_example" {
  value = data.meraki_devices_wireless_bluetooth_settings.example.item
}
