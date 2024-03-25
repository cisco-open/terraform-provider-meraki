
data "meraki_networks_wireless_ssids_identity_psks" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_identity_psks_example" {
  value = data.meraki_networks_wireless_ssids_identity_psks.example.items
}

data "meraki_networks_wireless_ssids_identity_psks" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_identity_psks_example" {
  value = data.meraki_networks_wireless_ssids_identity_psks.example.item
}
