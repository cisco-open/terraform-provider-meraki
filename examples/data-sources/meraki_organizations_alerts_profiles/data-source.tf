
data "meraki_organizations_alerts_profiles" "example" {

  organization_id = "string"
}

output "meraki_organizations_alerts_profiles_example" {
  value = data.meraki_organizations_alerts_profiles.example.items
}
