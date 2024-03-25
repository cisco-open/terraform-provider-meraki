
data "meraki_networks_wireless_failed_connections" "example" {

  ap_tag     = "string"
  band       = "string"
  client_id  = "string"
  network_id = "string"
  serial     = "string"
  ssid       = 1
  t0         = "string"
  t1         = "string"
  timespan   = 1.0
  vlan       = 1
}

output "meraki_networks_wireless_failed_connections_example" {
  value = data.meraki_networks_wireless_failed_connections.example.items
}
