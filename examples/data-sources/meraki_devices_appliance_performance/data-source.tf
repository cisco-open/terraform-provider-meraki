
data "meraki_devices_appliance_performance" "example" {

  serial = "string"
}

output "meraki_devices_appliance_performance_example" {
  value = data.meraki_devices_appliance_performance.example.item
}
