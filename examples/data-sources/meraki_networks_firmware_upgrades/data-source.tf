
data "meraki_networks_firmware_upgrades" "example" {

  network_id = "string"
}

output "meraki_networks_firmware_upgrades_example" {
  value = data.meraki_networks_firmware_upgrades.example.item
}
