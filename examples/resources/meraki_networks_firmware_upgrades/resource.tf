
resource "meraki_networks_firmware_upgrades" "example" {

  network_id = "string"
  products = {

    appliance = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        to_version = {

          id = "1001"
        }
      }
      participate_in_next_beta_release = false
    }
    camera = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        to_version = {

          id = "1003"
        }
      }
      participate_in_next_beta_release = false
    }
    cellular_gateway = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        to_version = {

          id = "1004"
        }
      }
      participate_in_next_beta_release = false
    }
    sensor = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        to_version = {

          id = "1005"
        }
      }
      participate_in_next_beta_release = false
    }
    switch = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        to_version = {

          id = "1002"
        }
      }
      participate_in_next_beta_release = false
    }
    switch_catalyst = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        to_version = {

          id = "1234"
        }
      }
      participate_in_next_beta_release = false
    }
    wireless = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        to_version = {

          id = "1000"
        }
      }
      participate_in_next_beta_release = false
    }
  }
  timezone = "America/Los_Angeles"
  upgrade_window = {

    day_of_week = "sun"
    hour_of_day = "4:00"
  }
}

output "meraki_networks_firmware_upgrades_example" {
  value = meraki_networks_firmware_upgrades.example
}