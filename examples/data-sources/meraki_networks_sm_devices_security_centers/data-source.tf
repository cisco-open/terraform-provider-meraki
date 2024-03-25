
data "meraki_networks_sm_devices_security_centers" "example" {

  device_id  = "string"
  network_id = "string"
}

output "meraki_networks_sm_devices_security_centers_example" {
  value = data.meraki_networks_sm_devices_security_centers.example.items
}
