
data "meraki_networks_wireless_ssids_splash_settings" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_splash_settings_example" {
  value = data.meraki_networks_wireless_ssids_splash_settings.example.item
}
