
data "meraki_devices_camera_wireless_profiles" "example" {

  serial = "string"
}

output "meraki_devices_camera_wireless_profiles_example" {
  value = data.meraki_devices_camera_wireless_profiles.example.item
}
