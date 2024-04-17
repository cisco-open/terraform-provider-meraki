
data "meraki_organizations_appliance_uplinks_usage_by_network" "example" {

  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_appliance_uplinks_usage_by_network_example" {
  value = data.meraki_organizations_appliance_uplinks_usage_by_network.example.items
}
