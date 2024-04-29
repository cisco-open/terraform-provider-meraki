
data "meraki_devices_live_tools_ping" "example" {

  id     = "string"
  serial = "string"
}

output "meraki_devices_live_tools_ping_example" {
  value = data.meraki_devices_live_tools_ping.example.item
}
