
data "meraki_networks_sm_user_access_devices" "example" {

  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
}

output "meraki_networks_sm_user_access_devices_example" {
  value = data.meraki_networks_sm_user_access_devices.example.items
}
