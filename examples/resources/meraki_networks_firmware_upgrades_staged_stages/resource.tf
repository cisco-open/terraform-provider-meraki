
resource "meraki_networks_firmware_upgrades_staged_stages" "example" {

  json = [{

    group = {

      id = "1234"
    }
  }]
  network_id = "string"
}

output "meraki_networks_firmware_upgrades_staged_stages_example" {
  value = meraki_networks_firmware_upgrades_staged_stages.example
}