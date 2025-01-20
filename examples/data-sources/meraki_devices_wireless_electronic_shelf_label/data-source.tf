
data "meraki_devices_wireless_electronic_shelf_label" "example" {

  serial = "string"
}

output "meraki_devices_wireless_electronic_shelf_label_example" {
  value = data.meraki_devices_wireless_electronic_shelf_label.example.item
}
