terraform {
  required_providers {
    meraki = {
      version = "1.2.4-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_devices_wireless_radio_settings" "example" {

  five_ghz_settings = {

    channel       = 149
    channel_width = 20
    target_power  = 15
  }
  # rf_profile_id = "1234"
  serial = "Q2FV-DJ6J-4QHD"
  two_four_ghz_settings = {

    channel      = 11
    target_power = 21
  }
}

output "meraki_devices_wireless_radio_settings_example" {
  value = meraki_devices_wireless_radio_settings.example
  sensitive = true
}