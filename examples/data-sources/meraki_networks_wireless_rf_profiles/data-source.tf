
data "meraki_networks_wireless_rf_profiles" "example" {

  include_template_profiles = false
  network_id                = "string"
}

output "meraki_networks_wireless_rf_profiles_example" {
  value = data.meraki_networks_wireless_rf_profiles.example.items
}

data "meraki_networks_wireless_rf_profiles" "example" {

  include_template_profiles = false
  network_id                = "string"
}

output "meraki_networks_wireless_rf_profiles_example" {
  value = data.meraki_networks_wireless_rf_profiles.example.item
}
