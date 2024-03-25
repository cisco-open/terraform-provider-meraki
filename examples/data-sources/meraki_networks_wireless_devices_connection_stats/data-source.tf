
data "meraki_networks_wireless_devices_connection_stats" "example" {

  ap_tag     = "string"
  band       = "string"
  network_id = "string"
  ssid       = 1
  t0         = "string"
  t1         = "string"
  timespan   = 1.0
  vlan       = 1
}

output "meraki_networks_wireless_devices_connection_stats_example" {
  value = data.meraki_networks_wireless_devices_connection_stats.example.items
}
