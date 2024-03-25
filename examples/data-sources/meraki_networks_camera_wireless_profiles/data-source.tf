
data "meraki_networks_camera_wireless_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_camera_wireless_profiles_example" {
  value = data.meraki_networks_camera_wireless_profiles.example.items
}

data "meraki_networks_camera_wireless_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_camera_wireless_profiles_example" {
  value = data.meraki_networks_camera_wireless_profiles.example.item
}
