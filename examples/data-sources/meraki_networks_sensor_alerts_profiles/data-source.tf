
data "meraki_networks_sensor_alerts_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_sensor_alerts_profiles_example" {
  value = data.meraki_networks_sensor_alerts_profiles.example.items
}

data "meraki_networks_sensor_alerts_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_sensor_alerts_profiles_example" {
  value = data.meraki_networks_sensor_alerts_profiles.example.item
}
