
resource "meraki_devices_camera_quality_and_retention" "example" {

  audio_recording_enabled           = false
  motion_based_retention_enabled    = false
  motion_detector_version           = 2
  profile_id                        = "1234"
  quality                           = "Standard"
  resolution                        = "1280x720"
  restricted_bandwidth_mode_enabled = false
  serial                            = "string"
}

output "meraki_devices_camera_quality_and_retention_example" {
  value = meraki_devices_camera_quality_and_retention.example
}