
data "meraki_organizations_inventory_onboarding_cloud_monitoring_imports" "example" {

  import_ids      = ["string"]
  organization_id = "string"
}

output "meraki_organizations_inventory_onboarding_cloud_monitoring_imports_example" {
  value = data.meraki_organizations_inventory_onboarding_cloud_monitoring_imports.example.items
}
