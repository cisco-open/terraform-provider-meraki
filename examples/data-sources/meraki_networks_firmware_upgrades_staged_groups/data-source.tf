
data "meraki_networks_firmware_upgrades_staged_groups" "example" {

  network_id = "string"
}

output "meraki_networks_firmware_upgrades_staged_groups_example" {
  value = data.meraki_networks_firmware_upgrades_staged_groups.example.items
}

data "meraki_networks_firmware_upgrades_staged_groups" "example" {

  network_id = "string"
}

output "meraki_networks_firmware_upgrades_staged_groups_example" {
  value = data.meraki_networks_firmware_upgrades_staged_groups.example.item
}
