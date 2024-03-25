
data "meraki_networks_wireless_ssids_hotspot20" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_hotspot20_example" {
  value = data.meraki_networks_wireless_ssids_hotspot20.example.item
}
