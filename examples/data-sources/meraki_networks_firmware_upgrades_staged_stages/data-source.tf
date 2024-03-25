
data "meraki_networks_firmware_upgrades_staged_stages" "example" {

  network_id = "string"
}

output "meraki_networks_firmware_upgrades_staged_stages_example" {
  value = data.meraki_networks_firmware_upgrades_staged_stages.example.items
}
