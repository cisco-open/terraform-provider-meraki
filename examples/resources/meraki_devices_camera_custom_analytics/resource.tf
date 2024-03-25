
resource "meraki_devices_camera_custom_analytics" "example" {

  artifact_id = "1"
  enabled     = true
  parameters = [{

    name  = "detection_threshold"
    value = "0.5"
  }]
  serial = "string"
}

output "meraki_devices_camera_custom_analytics_example" {
  value = meraki_devices_camera_custom_analytics.example
}