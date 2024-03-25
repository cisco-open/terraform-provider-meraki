
data "meraki_networks_appliance_single_lan" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_single_lan_example" {
  value = data.meraki_networks_appliance_single_lan.example.item
}
