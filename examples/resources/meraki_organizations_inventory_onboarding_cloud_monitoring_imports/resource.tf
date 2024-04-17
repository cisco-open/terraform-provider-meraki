
resource "meraki_organizations_inventory_onboarding_cloud_monitoring_imports" "example" {

  organization_id = "string"
  parameters = {

    devices = [{

      device_id  = "161b2602-a713-4aac-b1eb-d9b55205353d"
      network_id = "1338481"
      udi        = "PID:C9200L-24P-4G SN:JAE25220R2K"
    }]
  }
}

output "meraki_organizations_inventory_onboarding_cloud_monitoring_imports_example" {
  value = meraki_organizations_inventory_onboarding_cloud_monitoring_imports.example
}