
resource "meraki_organizations_splash_themes" "example" {

  base_theme      = "string"
  id              = "string"
  name            = "string"
  organization_id = "string"
}

output "meraki_organizations_splash_themes_example" {
  value = meraki_organizations_splash_themes.example
}