
resource "meraki_networks_firmware_upgrades_staged_events" "example" {

  network_id = "string"
  products = {

    switch = {

      next_upgrade = {

        to_version = {

          id = "1234"
        }
      }
    }
  }
  stages = [{

    group = {

      id = "1234"
    }
    milestones = {

      scheduled_for = "2018-02-11T00:00:00Z"
    }
  }]
}

output "meraki_networks_firmware_upgrades_staged_events_example" {
  value = meraki_networks_firmware_upgrades_staged_events.example
}