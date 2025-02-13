
terraform {
  required_providers {
    meraki = {
      version = "1.0.2-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_wireless_rf_profiles" "example" {

  ap_band_settings = {

    band_operation_mode   = "dual"
    band_steering_enabled = true
    bands = {

      enabled = ["2.4", "5"]
    }
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
  # flex_radios = {

  #   by_model = [{

  #     bands = ["5"]
  #     model = "MR34"
  #   }]
  # }
  min_bitrate_type = "band"
  name             = "Main Office"
  network_id       = "L_828099381482771185"
  per_ssid_settings = {

    status_0 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_1 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_10 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_11 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_12 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_13 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_14 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_2 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_3 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_4 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_5 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_6 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_7 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_8 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
    status_9 = {

      band_operation_mode   = "dual"
      band_steering_enabled = true
      bands = {

        enabled = ["2.4", "5"]
      }
      min_bitrate = 11.0
    }
  }
  six_ghz_settings = {

    channel_width       = "auto"
    max_power           = 30
    min_bitrate         = 12
    min_power           = 8
    rxsop               = -95
    valid_auto_channels = [1, 5, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233]
  }
  transmission = {

    enabled = true
  }
  two_four_ghz_settings = {

    ax_enabled          = true
    max_power           = 30
    min_bitrate         = 11.0
    min_power           = 5
    rxsop               = -95
    valid_auto_channels = [1, 6]
  }
}

output "meraki_networks_wireless_rf_profiles_example" {
  value = meraki_networks_wireless_rf_profiles.example
}