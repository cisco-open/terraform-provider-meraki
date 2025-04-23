
resource "meraki_organizations_integrations_xdr_networks_enable" "example" {

  organization_id = "string"
  parameters = {

    networks = [{

      network_id    = "N_1234567"
      product_types = ["appliance"]
    }]
  }
}

output "meraki_organizations_integrations_xdr_networks_enable_example" {
  value = meraki_organizations_integrations_xdr_networks_enable.example
}