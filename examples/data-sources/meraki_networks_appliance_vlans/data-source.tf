
data "meraki_networks_appliance_vlans" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_vlans_example" {
  value = data.meraki_networks_appliance_vlans.example.items
}

data "meraki_networks_appliance_vlans" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_vlans_example" {
  value = data.meraki_networks_appliance_vlans.example.item
}
