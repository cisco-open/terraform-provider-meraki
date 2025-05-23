
data "meraki_organizations_summary_top_clients_manufacturers_by_usage" "example" {

  device_tag      = "string"
  network_tag     = "string"
  organization_id = "string"
  quantity        = 1
  ssid_name       = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  usage_uplink    = "string"
}

output "meraki_organizations_summary_top_clients_manufacturers_by_usage_example" {
  value = data.meraki_organizations_summary_top_clients_manufacturers_by_usage.example.items
}
