
resource "meraki_networks_wireless_electronic_shelf_label" "example" {

  enabled    = true
  hostname   = "example.com"
  mode       = "high frequency"
  network_id = "string"
}

output "meraki_networks_wireless_electronic_shelf_label_example" {
  value = meraki_networks_wireless_electronic_shelf_label.example
}