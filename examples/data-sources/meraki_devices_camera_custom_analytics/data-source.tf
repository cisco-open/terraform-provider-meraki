
data "meraki_devices_camera_custom_analytics" "example" {

  serial = "string"
}

output "meraki_devices_camera_custom_analytics_example" {
  value = data.meraki_devices_camera_custom_analytics.example.item
}
