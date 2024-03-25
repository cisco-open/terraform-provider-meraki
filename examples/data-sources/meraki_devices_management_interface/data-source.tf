
data "meraki_devices_management_interface" "example" {

  serial = "string"
}

output "meraki_devices_management_interface_example" {
  value = data.meraki_devices_management_interface.example.item
}
