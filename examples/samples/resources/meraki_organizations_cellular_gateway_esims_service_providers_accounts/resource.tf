terraform {
  required_providers {
    meraki = {
      version = "1.2.4-beta"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}


resource "meraki_organizations_cellular_gateway_esims_service_providers_accounts" "example" {

  account_id      = "0987654321"
  api_key         = "foobarfoobarfoobarfoobarfoobarfoobar"
  organization_id = "828099381482762270"
  service_provider = {

    name = "ATT"
  }
  title    = "My AT&T account"
  username = "MerakiUser"
}

output "meraki_organizations_cellular_gateway_esims_service_providers_accounts_example" {
  value = meraki_organizations_cellular_gateway_esims_service_providers_accounts.example
  sensitive = true
}