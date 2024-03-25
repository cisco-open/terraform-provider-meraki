
resource "meraki_networks_wireless_ssids" "example" {

  enabled    = true
  name       = "My SSID"
  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_example" {
  value = meraki_networks_wireless_ssids.example
}