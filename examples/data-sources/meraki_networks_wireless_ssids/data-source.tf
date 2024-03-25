
data "meraki_networks_wireless_ssids" "example" {

  network_id = "string"
}

output "meraki_networks_wireless_ssids_example" {
  value = data.meraki_networks_wireless_ssids.example.items
}

data "meraki_networks_wireless_ssids" "example" {

  network_id = "string"
}

output "meraki_networks_wireless_ssids_example" {
  value = data.meraki_networks_wireless_ssids.example.item
}
