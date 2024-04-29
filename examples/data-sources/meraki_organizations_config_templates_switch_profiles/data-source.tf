
data "meraki_organizations_config_templates_switch_profiles" "example" {

  config_template_id = "string"
  organization_id    = "string"
}

output "meraki_organizations_config_templates_switch_profiles_example" {
  value = data.meraki_organizations_config_templates_switch_profiles.example.items
}
