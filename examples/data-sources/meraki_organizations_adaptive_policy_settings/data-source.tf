
data "meraki_organizations_adaptive_policy_settings" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_settings_example" {
  value = data.meraki_organizations_adaptive_policy_settings.example.item
}
