
data "meraki_organizations_config_templates_switch_profiles_ports" "example" {

  config_template_id = "string"
  organization_id    = "string"
  profile_id         = "string"
}

output "meraki_organizations_config_templates_switch_profiles_ports_example" {
  value = data.meraki_organizations_config_templates_switch_profiles_ports.example.items
}

data "meraki_organizations_config_templates_switch_profiles_ports" "example" {

  config_template_id = "string"
  organization_id    = "string"
  profile_id         = "string"
}

output "meraki_organizations_config_templates_switch_profiles_ports_example" {
  value = data.meraki_organizations_config_templates_switch_profiles_ports.example.item
}
