
data "meraki_devices_wireless_connection_stats" "example" {

  ap_tag   = "string"
  band     = "string"
  serial   = "string"
  ssid     = 1
  t0       = "string"
  t1       = "string"
  timespan = 1.0
  vlan     = 1
}

output "meraki_devices_wireless_connection_stats_example" {
  value = data.meraki_devices_wireless_connection_stats.example.item
}
