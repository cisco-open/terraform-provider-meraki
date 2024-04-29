
data "meraki_networks_appliance_rf_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_rf_profiles_example" {
  value = data.meraki_networks_appliance_rf_profiles.example.item
}
