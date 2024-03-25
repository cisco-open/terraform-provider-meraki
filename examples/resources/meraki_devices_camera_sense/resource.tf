
resource "meraki_devices_camera_sense" "example" {

  audio_detection = {

    enabled = false
  }
  mqtt_broker_id = "1234"
  sense_enabled  = true
  serial         = "string"
}

output "meraki_devices_camera_sense_example" {
  value = meraki_devices_camera_sense.example
}