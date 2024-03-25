
data "meraki_devices_appliance_uplinks_settings" "example" {

  serial = "string"
}

output "meraki_devices_appliance_uplinks_settings_example" {
  value = data.meraki_devices_appliance_uplinks_settings.example.item
}
