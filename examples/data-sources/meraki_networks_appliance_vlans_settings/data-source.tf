
data "meraki_networks_appliance_vlans_settings" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_vlans_settings_example" {
  value = data.meraki_networks_appliance_vlans_settings.example.item
}
