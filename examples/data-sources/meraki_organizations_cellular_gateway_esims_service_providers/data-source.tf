
data "meraki_organizations_cellular_gateway_esims_service_providers" "example" {

  organization_id = "string"
}

output "meraki_organizations_cellular_gateway_esims_service_providers_example" {
  value = data.meraki_organizations_cellular_gateway_esims_service_providers.example.item
}
