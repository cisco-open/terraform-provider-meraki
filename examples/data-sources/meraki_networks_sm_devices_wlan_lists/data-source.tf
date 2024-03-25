
data "meraki_networks_sm_devices_wlan_lists" "example" {

  device_id  = "string"
  network_id = "string"
}

output "meraki_networks_sm_devices_wlan_lists_example" {
  value = data.meraki_networks_sm_devices_wlan_lists.example.items
}
