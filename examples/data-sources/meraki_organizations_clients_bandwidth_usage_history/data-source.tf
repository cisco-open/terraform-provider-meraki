
data "meraki_organizations_clients_bandwidth_usage_history" "example" {

  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_clients_bandwidth_usage_history_example" {
  value = data.meraki_organizations_clients_bandwidth_usage_history.example.items
}
