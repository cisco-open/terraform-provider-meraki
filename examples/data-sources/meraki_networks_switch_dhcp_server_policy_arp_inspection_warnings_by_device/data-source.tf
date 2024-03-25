
data "meraki_networks_switch_dhcp_server_policy_arp_inspection_warnings_by_device" "example" {

  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
}

output "meraki_networks_switch_dhcp_server_policy_arp_inspection_warnings_by_device_example" {
  value = data.meraki_networks_switch_dhcp_server_policy_arp_inspection_warnings_by_device.example.items
}
