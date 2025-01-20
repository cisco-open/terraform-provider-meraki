
data "meraki_organizations_clients_bandwidth_usage_history" "example" {

  device_tag      = "string"
  network_tag     = "string"
  organization_id = "string"
  ssid_name       = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  usage_uplink    = "string"
}

output "meraki_organizations_clients_bandwidth_usage_history_example" {
  value = data.meraki_organizations_clients_bandwidth_usage_history.example.items
}
