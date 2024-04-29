
data "meraki_organizations_summary_top_networks_by_status" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_summary_top_networks_by_status_example" {
  value = data.meraki_organizations_summary_top_networks_by_status.example.items
}
