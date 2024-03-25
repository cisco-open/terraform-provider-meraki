
data "meraki_organizations_insight_applications" "example" {

  organization_id = "string"
}

output "meraki_organizations_insight_applications_example" {
  value = data.meraki_organizations_insight_applications.example.items
}
