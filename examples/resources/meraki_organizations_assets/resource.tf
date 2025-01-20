
resource "meraki_organizations_assets" "example" {

  organization_id  = "string"
  theme_identifier = "string"
  parameters = {

    content = "string"
    name    = "string"
  }
}

output "meraki_organizations_assets_example" {
  value = meraki_organizations_assets.example
}