
data "meraki_devices_live_tools_throughput_test" "example" {

  serial             = "string"
  throughput_test_id = "string"
}

output "meraki_devices_live_tools_throughput_test_example" {
  value = data.meraki_devices_live_tools_throughput_test.example.item
}
