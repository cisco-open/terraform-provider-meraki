
data "meraki_networks_switch_settings" "example" {

  network_id = "string"
}

output "meraki_networks_switch_settings_example" {
  value = data.meraki_networks_switch_settings.example.item
}
