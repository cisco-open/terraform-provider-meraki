
data "meraki_organizations_summary_top_appliances_by_utilization" "example" {

  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_summary_top_appliances_by_utilization_example" {
  value = data.meraki_organizations_summary_top_appliances_by_utilization.example.items
}
