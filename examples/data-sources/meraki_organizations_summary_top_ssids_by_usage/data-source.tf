
data "meraki_organizations_summary_top_ssids_by_usage" "example" {

  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_summary_top_ssids_by_usage_example" {
  value = data.meraki_organizations_summary_top_ssids_by_usage.example.items
}
