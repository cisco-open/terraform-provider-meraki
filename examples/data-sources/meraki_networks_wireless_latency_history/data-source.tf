
data "meraki_networks_wireless_latency_history" "example" {

  access_category = "string"
  ap_tag          = "string"
  auto_resolution = false
  band            = "string"
  client_id       = "string"
  device_serial   = "string"
  network_id      = "string"
  resolution      = 1
  ssid            = 1
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_networks_wireless_latency_history_example" {
  value = data.meraki_networks_wireless_latency_history.example.items
}
