
data "meraki_networks_vlan_profiles_assignments_by_device" "example" {

  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  product_types  = ["string"]
  serials        = ["string"]
  stack_ids      = ["string"]
  starting_after = "string"
}

output "meraki_networks_vlan_profiles_assignments_by_device_example" {
  value = data.meraki_networks_vlan_profiles_assignments_by_device.example.items
}
