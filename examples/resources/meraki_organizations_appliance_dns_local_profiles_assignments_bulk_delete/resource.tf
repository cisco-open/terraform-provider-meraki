
resource "meraki_organizations_appliance_dns_local_profiles_assignments_bulk_delete" "example" {

  organization_id = "string"
  parameters = {

    items = [{

      assignment_id = "123456"
    }]
  }
}

output "meraki_organizations_appliance_dns_local_profiles_assignments_bulk_delete_example" {
  value = meraki_organizations_appliance_dns_local_profiles_assignments_bulk_delete.example
}