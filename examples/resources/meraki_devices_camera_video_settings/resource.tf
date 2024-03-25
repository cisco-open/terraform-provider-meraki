
resource "meraki_devices_camera_video_settings" "example" {

  external_rtsp_enabled = true
  serial                = "string"
}

output "meraki_devices_camera_video_settings_example" {
  value = meraki_devices_camera_video_settings.example
}