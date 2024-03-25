
data "meraki_networks_firmware_upgrades_staged_events" "example" {

  network_id = "string"
}

output "meraki_networks_firmware_upgrades_staged_events_example" {
  value = data.meraki_networks_firmware_upgrades_staged_events.example.item
}
