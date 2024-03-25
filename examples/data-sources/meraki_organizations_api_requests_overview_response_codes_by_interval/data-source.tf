
data "meraki_organizations_api_requests_overview_response_codes_by_interval" "example" {

  admin_ids       = ["string"]
  interval        = 1
  operation_ids   = ["string"]
  organization_id = "string"
  source_ips      = ["string"]
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  user_agent      = "string"
  version         = 1
}

output "meraki_organizations_api_requests_overview_response_codes_by_interval_example" {
  value = data.meraki_organizations_api_requests_overview_response_codes_by_interval.example.items
}
