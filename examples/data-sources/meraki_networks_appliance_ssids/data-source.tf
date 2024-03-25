
data "meraki_networks_appliance_ssids" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_ssids_example" {
  value = data.meraki_networks_appliance_ssids.example.items
}

data "meraki_networks_appliance_ssids" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_ssids_example" {
  value = data.meraki_networks_appliance_ssids.example.item
}
