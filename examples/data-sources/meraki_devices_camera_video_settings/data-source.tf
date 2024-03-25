
data "meraki_devices_camera_video_settings" "example" {

  serial = "string"
}

output "meraki_devices_camera_video_settings_example" {
  value = data.meraki_devices_camera_video_settings.example.item
}
