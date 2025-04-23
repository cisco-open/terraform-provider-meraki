
resource "meraki_organizations_config_templates" "example" {

  copy_from_network_id = "N_24329156"
  name                 = "My config template"
  organization_id      = "string"
  time_zone            = "America/Los_Angeles"
}

output "meraki_organizations_config_templates_example" {
  value = meraki_organizations_config_templates.example
}