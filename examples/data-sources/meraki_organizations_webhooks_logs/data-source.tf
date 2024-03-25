
data "meraki_organizations_webhooks_logs" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  url             = "string"
}

output "meraki_organizations_webhooks_logs_example" {
  value = data.meraki_organizations_webhooks_logs.example.items
}
