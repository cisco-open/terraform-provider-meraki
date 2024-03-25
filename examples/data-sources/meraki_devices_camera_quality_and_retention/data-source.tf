
data "meraki_devices_camera_quality_and_retention" "example" {

  serial = "string"
}

output "meraki_devices_camera_quality_and_retention_example" {
  value = data.meraki_devices_camera_quality_and_retention.example.item
}
