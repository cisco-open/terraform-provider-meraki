terraform {
  required_providers {
    meraki = {
      version = "1.0.5-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}



resource "meraki_networks_firmware_upgrades" "example" {

  network_id = "L_828099381482775374"
  products = {

    appliance = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        # to_version = {

        #   id = "1004"
        # }
      }
      participate_in_next_beta_release = false
    }

    # sensor = {

    #   next_upgrade = {

    #     time = "2019-03-17T17:22:52Z"
    #     # to_version = {

    #     #   id = "1005"
    #     # }
    #   }
    #   participate_in_next_beta_release = false
    # }

    switch_catalyst = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        # to_version = {

        #   id = "1004"
        # }
      }
      participate_in_next_beta_release = false
    }
    wireless = {

      next_upgrade = {

        time = "2019-03-17T17:22:52Z"
        # to_version = {

        #   id = "1004"
        # }
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