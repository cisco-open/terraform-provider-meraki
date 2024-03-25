
data "meraki_organizations_api_requests" "example" {

  admin_id        = "string"
  ending_before   = "string"
  method          = "string"
  operation_ids   = ["string"]
  organization_id = "string"
  path            = "string"
  per_page        = 1
  response_code   = 1
  source_ip       = "string"
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  user_agent      = "string"
  version         = 1
}

output "meraki_organizations_api_requests_example" {
  value = data.meraki_organizations_api_requests.example.items
}
