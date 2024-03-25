
data "meraki_organizations_licenses_overview" "example" {

  organization_id = "string"
}

output "meraki_organizations_licenses_overview_example" {
  value = data.meraki_organizations_licenses_overview.example.item
}
