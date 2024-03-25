
resource "meraki_devices_live_tools_ping" "example" {

  serial = "string"
  parameters = {

    count  = 2
    target = "75.75.75.75"
  }
}

output "meraki_devices_live_tools_ping_example" {
  value = meraki_devices_live_tools_ping.example
}