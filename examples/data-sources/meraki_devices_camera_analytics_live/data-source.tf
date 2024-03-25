
data "meraki_devices_camera_analytics_live" "example" {

  serial = "string"
}

output "meraki_devices_camera_analytics_live_example" {
  value = data.meraki_devices_camera_analytics_live.example.item
}
