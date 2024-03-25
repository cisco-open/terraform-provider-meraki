
data "meraki_devices_wireless_status" "example" {

  serial = "string"
}

output "meraki_devices_wireless_status_example" {
  value = data.meraki_devices_wireless_status.example.item
}
