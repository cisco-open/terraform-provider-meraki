
data "meraki_networks_sm_devices_certs" "example" {

  device_id  = "string"
  network_id = "string"
}

output "meraki_networks_sm_devices_certs_example" {
  value = data.meraki_networks_sm_devices_certs.example.items
}
