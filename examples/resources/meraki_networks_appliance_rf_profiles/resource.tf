
resource "meraki_networks_appliance_rf_profiles" "example" {

  five_ghz_settings = {

    ax_enabled  = true
    min_bitrate = 48
  }
  name       = "MX RF Profile"
  network_id = "string"
  per_ssid_settings = {

    status_1 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
    }
    status_2 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
    }
    status_3 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
    }
    status_4 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
    }
  }
  two_four_ghz_settings = {

    ax_enabled  = true
    min_bitrate = 12.0
  }
}

output "meraki_networks_appliance_rf_profiles_example" {
  value = meraki_networks_appliance_rf_profiles.example
}