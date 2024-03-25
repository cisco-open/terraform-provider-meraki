
data "meraki_networks_wireless_ssids_bonjour_forwarding" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_bonjour_forwarding_example" {
  value = data.meraki_networks_wireless_ssids_bonjour_forwarding.example.item
}
