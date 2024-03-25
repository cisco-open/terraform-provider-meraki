
resource "meraki_networks_sm_devices_unenroll" "example" {

  device_id  = "string"
  network_id = "string"

}

output "meraki_networks_sm_devices_unenroll_example" {
  value = meraki_networks_sm_devices_unenroll.example
}