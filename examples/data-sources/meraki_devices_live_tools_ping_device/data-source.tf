
data "meraki_devices_live_tools_ping_device" "example" {

  id     = "string"
  serial = "string"
}

output "meraki_devices_live_tools_ping_device_example" {
  value = data.meraki_devices_live_tools_ping_device.example.item
}
