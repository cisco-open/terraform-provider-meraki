
resource "meraki_networks_wireless_ssids_device_type_group_policies" "example" {

  device_type_policies = [{

    device_policy = "Allowed"
    device_type   = "Android"
  }]
  enabled    = true
  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_device_type_group_policies_example" {
  value = meraki_networks_wireless_ssids_device_type_group_policies.example
}