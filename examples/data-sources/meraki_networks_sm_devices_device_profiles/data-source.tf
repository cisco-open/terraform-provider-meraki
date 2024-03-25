
data "meraki_networks_sm_devices_device_profiles" "example" {

  device_id  = "string"
  network_id = "string"
}

output "meraki_networks_sm_devices_device_profiles_example" {
  value = data.meraki_networks_sm_devices_device_profiles.example.items
}
