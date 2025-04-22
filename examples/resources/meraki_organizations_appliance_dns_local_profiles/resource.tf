
resource "meraki_organizations_appliance_dns_local_profiles" "example" {

  name            = "Default profile"
  organization_id = "string"
}

output "meraki_organizations_appliance_dns_local_profiles_example" {
  value = meraki_organizations_appliance_dns_local_profiles.example
}