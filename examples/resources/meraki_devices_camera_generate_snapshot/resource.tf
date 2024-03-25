
resource "meraki_devices_camera_generate_snapshot" "example" {

  serial = "string"
  parameters = {

    fullframe = false
    timestamp = "2021-04-30T15:18:08Z"
  }
}

output "meraki_devices_camera_generate_snapshot_example" {
  value = meraki_devices_camera_generate_snapshot.example
}