
resource "meraki_networks_firmware_upgrades_rollbacks" "example" {

  network_id = "string"
  parameters = {

    product = "switch"
    reasons = [{

      category = "performance"
      comment  = "Network was slower with the upgrade"
    }]
    time = "2020-10-21T02:00:00Z"
    to_version = {

      id = "7857"
    }
  }
}

output "meraki_networks_firmware_upgrades_rollbacks_example" {
  value = meraki_networks_firmware_upgrades_rollbacks.example
}