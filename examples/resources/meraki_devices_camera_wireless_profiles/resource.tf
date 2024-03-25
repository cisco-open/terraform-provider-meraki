
resource "meraki_devices_camera_wireless_profiles" "example" {

  ids = {

    backup    = "1"
    primary   = "3"
    secondary = "2"
  }
  serial = "string"
}

output "meraki_devices_camera_wireless_profiles_example" {
  value = meraki_devices_camera_wireless_profiles.example
}