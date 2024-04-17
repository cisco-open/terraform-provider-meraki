
data "meraki_devices_appliance_radio_settings" "example" {

  serial = "string"
}

output "meraki_devices_appliance_radio_settings_example" {
  value = data.meraki_devices_appliance_radio_settings.example.item
}
