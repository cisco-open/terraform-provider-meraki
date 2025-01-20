
data "meraki_networks_vlan_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_vlan_profiles_example" {
  value = data.meraki_networks_vlan_profiles.example.items
}

data "meraki_networks_vlan_profiles" "example" {

  network_id = "string"
}

output "meraki_networks_vlan_profiles_example" {
  value = data.meraki_networks_vlan_profiles.example.item
}
