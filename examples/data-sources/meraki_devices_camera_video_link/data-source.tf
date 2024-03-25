
data "meraki_devices_camera_video_link" "example" {

  serial    = "string"
  timestamp = "string"
}

output "meraki_devices_camera_video_link_example" {
  value = data.meraki_devices_camera_video_link.example.item
}
