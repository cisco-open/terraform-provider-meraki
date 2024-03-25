
resource "meraki_devices_live_tools_ping_device" "example" {

  serial = "string"
  parameters = {

    count = 3
  }
}

output "meraki_devices_live_tools_ping_device_example" {
  value = meraki_devices_live_tools_ping_device.example
}