
data "meraki_networks_wireless_ssids_eap_override" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_eap_override_example" {
  value = data.meraki_networks_wireless_ssids_eap_override.example.item
}
