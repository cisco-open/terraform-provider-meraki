
data "meraki_devices_switch_ports_statuses" "example" {

  serial   = "string"
  t0       = "string"
  timespan = 1.0
}

output "meraki_devices_switch_ports_statuses_example" {
  value = data.meraki_devices_switch_ports_statuses.example.items
}
