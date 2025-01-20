
resource "meraki_networks_wireless_air_marshal_settings" "example" {

  network_id = "string"
  parameters = {

    default_policy = "allow"
  }
}

output "meraki_networks_wireless_air_marshal_settings_example" {
  value = meraki_networks_wireless_air_marshal_settings.example
}