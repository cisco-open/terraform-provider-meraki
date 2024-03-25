
resource "meraki_networks_sm_user_access_devices_delete" "example" {

  network_id            = "string"
  user_access_device_id = "string"
}

output "meraki_networks_sm_user_access_devices_delete_example" {
  value = meraki_networks_sm_user_access_devices_delete.example
}