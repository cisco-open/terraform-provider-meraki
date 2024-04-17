
data "meraki_organizations_appliance_uplinks_statuses_overview" "example" {

  organization_id = "string"
}

output "meraki_organizations_appliance_uplinks_statuses_overview_example" {
  value = data.meraki_organizations_appliance_uplinks_statuses_overview.example.item
}
