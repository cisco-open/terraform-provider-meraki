
data "meraki_devices_camera_sense" "example" {

  serial = "string"
}

output "meraki_devices_camera_sense_example" {
  value = data.meraki_devices_camera_sense.example.item
}
