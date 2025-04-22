
resource "meraki_organizations_cellular_gateway_esims_service_providers_accounts" "example" {

  account_id      = 987654321
  api_key         = "foobarfoobarfoobarfoobarfoobarfoobar"
  organization_id = "string"
  service_provider = {

    name = "ATT"
  }
  title    = "My AT&T account"
  username = "MerakiUser"
}

output "meraki_organizations_cellular_gateway_esims_service_providers_accounts_example" {
  value = meraki_organizations_cellular_gateway_esims_service_providers_accounts.example
}