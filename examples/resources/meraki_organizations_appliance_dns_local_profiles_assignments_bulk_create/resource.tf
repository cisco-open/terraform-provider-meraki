
resource "meraki_organizations_appliance_dns_local_profiles_assignments_bulk_create" "example" {

  organization_id = "string"
  parameters = {

    items = [{

      network = {

        id = "N_123456"
      }
      profile = {

        id = "1234"
      }
    }]
  }
}

output "meraki_organizations_appliance_dns_local_profiles_assignments_bulk_create_example" {
  value = meraki_organizations_appliance_dns_local_profiles_assignments_bulk_create.example
}