
data "meraki_networks_appliance_settings" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_settings_example" {
  value = data.meraki_networks_appliance_settings.example.item
}
