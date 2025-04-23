
data "meraki_organizations_integrations_xdr_networks" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_integrations_xdr_networks_example" {
  value = data.meraki_organizations_integrations_xdr_networks.example.item
}
