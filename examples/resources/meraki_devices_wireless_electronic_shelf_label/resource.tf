
resource "meraki_devices_wireless_electronic_shelf_label" "example" {

  channel = "1"
  enabled = true
  serial  = "string"
}

output "meraki_devices_wireless_electronic_shelf_label_example" {
  value = meraki_devices_wireless_electronic_shelf_label.example
}