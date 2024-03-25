
resource "meraki_networks_appliance_vlans_settings" "example" {

  network_id    = "string"
  vlans_enabled = true
}

output "meraki_networks_appliance_vlans_settings_example" {
  value = meraki_networks_appliance_vlans_settings.example
}