
data "meraki_networks_sm_devices_network_adapters" "example" {

  device_id  = "string"
  network_id = "string"
}

output "meraki_networks_sm_devices_network_adapters_example" {
  value = data.meraki_networks_sm_devices_network_adapters.example.items
}
