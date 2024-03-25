
resource "meraki_organizations_camera_custom_analytics_artifacts" "example" {

  name            = "example"
  organization_id = "string"
}

output "meraki_organizations_camera_custom_analytics_artifacts_example" {
  value = meraki_organizations_camera_custom_analytics_artifacts.example
}