
data "meraki_networks_wireless_settings" "example" {

  network_id = "string"
}

output "meraki_networks_wireless_settings_example" {
  value = data.meraki_networks_wireless_settings.example.item
}
