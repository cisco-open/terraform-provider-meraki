
resource "meraki_devices_wireless_radio_settings" "example" {

  five_ghz_settings = {

    channel       = 149
    channel_width = 20
    target_power  = 15
  }
  rf_profile_id = "1234"
  serial        = "string"
  two_four_ghz_settings = {

    channel      = 11
    target_power = 21
  }
}

output "meraki_devices_wireless_radio_settings_example" {
  value = meraki_devices_wireless_radio_settings.example
}