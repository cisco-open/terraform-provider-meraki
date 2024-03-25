
resource "meraki_networks_firmware_upgrades_staged_events_rollbacks" "example" {

  network_id = "string"
  parameters = {

    reasons = [{

      category = "performance"
      comment  = "Network was slower with the upgrade"
    }]
    stages = [{

      group = {

        id = "1234"
      }
      milestones = {

        scheduled_for = "2018-02-11T00:00:00Z"
      }
    }]
  }
}

output "meraki_networks_firmware_upgrades_staged_events_rollbacks_example" {
  value = meraki_networks_firmware_upgrades_staged_events_rollbacks.example
}