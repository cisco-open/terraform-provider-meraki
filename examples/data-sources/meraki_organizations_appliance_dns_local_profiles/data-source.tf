
data "meraki_organizations_appliance_dns_local_profiles" "example" {

  organization_id = "string"
  profile_ids     = ["string"]
}

output "meraki_organizations_appliance_dns_local_profiles_example" {
  value = data.meraki_organizations_appliance_dns_local_profiles.example.items
}
