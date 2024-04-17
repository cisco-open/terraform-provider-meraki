
data "meraki_organizations_openapi_spec" "example" {

  organization_id = "string"
  version         = 1
}

output "meraki_organizations_openapi_spec_example" {
  value = data.meraki_organizations_openapi_spec.example.item
}
