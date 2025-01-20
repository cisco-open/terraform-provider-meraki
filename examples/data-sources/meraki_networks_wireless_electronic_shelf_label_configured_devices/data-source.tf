
data "meraki_networks_wireless_electronic_shelf_label_configured_devices" "example" {

  network_id = "string"
}

output "meraki_networks_wireless_electronic_shelf_label_configured_devices_example" {
  value = data.meraki_networks_wireless_electronic_shelf_label_configured_devices.example.items
}
