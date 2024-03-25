
data "meraki_organizations_config_templates" "example" {

  organization_id = "string"
}

output "meraki_organizations_config_templates_example" {
  value = data.meraki_organizations_config_templates.example.items
}

data "meraki_organizations_config_templates" "example" {

  organization_id = "string"
}

output "meraki_organizations_config_templates_example" {
  value = data.meraki_organizations_config_templates.example.item
}
