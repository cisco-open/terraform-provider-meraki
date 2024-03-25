
resource "meraki_networks_wireless_rf_profiles" "example" {

  ap_band_settings = {

    band_operation_mode   = "dual"
    band_steering_enabled = true
  }
  band_selection_type      = "ap"
  client_balancing_enabled = true
  five_ghz_settings = {

    channel_width       = "auto"
    max_power           = 30
    min_bitrate         = 12
    min_power           = 8
    rxsop               = -95
    valid_auto_channels = [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165]
  }
  min_bitrate_type = "band"
  name             = "Main Office"
  network_id       = "string"
  per_ssid_settings = {

    status_0 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_1 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_10 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_11 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_12 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_13 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_14 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_2 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_3 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_4 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_5 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_6 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_7 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_8 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
    status_9 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      min_bitrate           = 11
    }
  }
  six_ghz_settings = {

    channel_width       = "auto"
    max_power           = 30
    min_bitrate         = 12
    min_power           = 8
    rxsop               = -95
    valid_auto_channels = [1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233]
  }
  transmission = {

    enabled = true
  }
  two_four_ghz_settings = {

    ax_enabled          = true
    max_power           = 30
    min_bitrate         = 11
    min_power           = 5
    rxsop               = -95
    valid_auto_channels = [1, 6, 11]
  }
}

output "meraki_networks_wireless_rf_profiles_example" {
  value = meraki_networks_wireless_rf_profiles.example
}