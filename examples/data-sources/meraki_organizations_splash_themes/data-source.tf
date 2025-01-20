
data "meraki_organizations_splash_themes" "example" {

  organization_id = "string"
}

output "meraki_organizations_splash_themes_example" {
  value = data.meraki_organizations_splash_themes.example.items
}
