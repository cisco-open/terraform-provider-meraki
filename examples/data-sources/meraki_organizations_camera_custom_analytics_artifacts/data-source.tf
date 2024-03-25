
data "meraki_organizations_camera_custom_analytics_artifacts" "example" {

  organization_id = "string"
}

output "meraki_organizations_camera_custom_analytics_artifacts_example" {
  value = data.meraki_organizations_camera_custom_analytics_artifacts.example.items
}

data "meraki_organizations_camera_custom_analytics_artifacts" "example" {

  organization_id = "string"
}

output "meraki_organizations_camera_custom_analytics_artifacts_example" {
  value = data.meraki_organizations_camera_custom_analytics_artifacts.example.item
}
