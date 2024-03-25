
data "meraki_networks_camera_quality_retention_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_camera_quality_retention_profiles_example" {
  value = data.meraki_networks_camera_quality_retention_profiles.example.items
}

data "meraki_networks_camera_quality_retention_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_camera_quality_retention_profiles_example" {
  value = data.meraki_networks_camera_quality_retention_profiles.example.item
}
