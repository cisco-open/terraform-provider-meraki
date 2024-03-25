
data "meraki_networks_sm_devices_device_command_logs" "example" {

  device_id      = "string"
  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
}

output "meraki_networks_sm_devices_device_command_logs_example" {
  value = data.meraki_networks_sm_devices_device_command_logs.example.items
}
