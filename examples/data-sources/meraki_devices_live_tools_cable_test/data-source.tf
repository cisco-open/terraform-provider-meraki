
data "meraki_devices_live_tools_cable_test" "example" {

  id     = "string"
  serial = "string"
}

output "meraki_devices_live_tools_cable_test_example" {
  value = data.meraki_devices_live_tools_cable_test.example.item
}
