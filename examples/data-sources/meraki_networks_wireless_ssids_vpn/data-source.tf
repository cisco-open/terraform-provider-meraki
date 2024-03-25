
data "meraki_networks_wireless_ssids_vpn" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_vpn_example" {
  value = data.meraki_networks_wireless_ssids_vpn.example.item
}
