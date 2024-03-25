
data "meraki_networks_wireless_ssids_device_type_group_policies" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_device_type_group_policies_example" {
  value = data.meraki_networks_wireless_ssids_device_type_group_policies.example.item
}
