
data "meraki_devices_appliance_performance" "example" {

  serial   = "string"
  t0       = "string"
  t1       = "string"
  timespan = 1.0
}

output "meraki_devices_appliance_performance_example" {
  value = data.meraki_devices_appliance_performance.example.item
}
