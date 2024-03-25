
data "meraki_networks_sm_devices_connectivity" "example" {

  device_id      = "string"
  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
}

output "meraki_networks_sm_devices_connectivity_example" {
  value = data.meraki_networks_sm_devices_connectivity.example.items
}
