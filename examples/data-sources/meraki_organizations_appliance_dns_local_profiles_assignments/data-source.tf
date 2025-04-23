
data "meraki_organizations_appliance_dns_local_profiles_assignments" "example" {

  network_ids     = ["string"]
  organization_id = "string"
  profile_ids     = ["string"]
}

output "meraki_organizations_appliance_dns_local_profiles_assignments_example" {
  value = data.meraki_organizations_appliance_dns_local_profiles_assignments.example.item
}
