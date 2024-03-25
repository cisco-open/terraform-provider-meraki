
data "meraki_networks_sm_devices_cellular_usage_history" "example" {

  device_id  = "string"
  network_id = "string"
}

output "meraki_networks_sm_devices_cellular_usage_history_example" {
  value = data.meraki_networks_sm_devices_cellular_usage_history.example.items
}
