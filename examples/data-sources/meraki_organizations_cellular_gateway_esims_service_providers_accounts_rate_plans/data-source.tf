
data "meraki_organizations_cellular_gateway_esims_service_providers_accounts_rate_plans" "example" {

  account_ids     = ["string"]
  organization_id = "string"
}

output "meraki_organizations_cellular_gateway_esims_service_providers_accounts_rate_plans_example" {
  value = data.meraki_organizations_cellular_gateway_esims_service_providers_accounts_rate_plans.example.item
}
