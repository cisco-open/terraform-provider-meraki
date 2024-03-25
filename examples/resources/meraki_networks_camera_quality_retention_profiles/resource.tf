
resource "meraki_networks_camera_quality_retention_profiles" "example" {

  name       = "Sample quality retention profile"
  network_id = "string"
}

output "meraki_networks_camera_quality_retention_profiles_example" {
  value = meraki_networks_camera_quality_retention_profiles.example
}