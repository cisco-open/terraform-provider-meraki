
data "meraki_networks_settings" "example" {

  network_id = "string"
}

output "meraki_networks_settings_example" {
  value = data.meraki_networks_settings.example.item
}
