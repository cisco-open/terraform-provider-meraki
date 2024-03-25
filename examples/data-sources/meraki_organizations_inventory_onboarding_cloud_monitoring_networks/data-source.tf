
data "meraki_organizations_inventory_onboarding_cloud_monitoring_networks" "example" {

  device_type     = "string"
  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_inventory_onboarding_cloud_monitoring_networks_example" {
  value = data.meraki_organizations_inventory_onboarding_cloud_monitoring_networks.example.items
}
