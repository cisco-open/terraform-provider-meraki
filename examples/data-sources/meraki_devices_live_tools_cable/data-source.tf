
data "meraki_devices_live_tools_cable" "example" {

  id     = "string"
  serial = "string"
}

output "meraki_devices_live_tools_cable_example" {
  value = data.meraki_devices_live_tools_cable.example.item
}
