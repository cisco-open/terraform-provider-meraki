
data "meraki_networks_wireless_bluetooth_settings" "example" {

  network_id = "string"
}

output "meraki_networks_wireless_bluetooth_settings_example" {
  value = data.meraki_networks_wireless_bluetooth_settings.example.item
}
