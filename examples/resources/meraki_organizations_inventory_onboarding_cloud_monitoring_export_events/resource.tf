
resource "meraki_organizations_inventory_onboarding_cloud_monitoring_export_events" "example" {

  organization_id = "string"
  parameters = {

    log_event = "download"
    request   = "r=cb"
    target_os = "mac"
    timestamp = 1526087474
  }
}

output "meraki_organizations_inventory_onboarding_cloud_monitoring_export_events_example" {
  value = meraki_organizations_inventory_onboarding_cloud_monitoring_export_events.example
}