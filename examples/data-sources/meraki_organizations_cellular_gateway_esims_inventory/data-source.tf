
data "meraki_organizations_cellular_gateway_esims_inventory" "example" {

  eids            = ["string"]
  organization_id = "string"
}

output "meraki_organizations_cellular_gateway_esims_inventory_example" {
  value = data.meraki_organizations_cellular_gateway_esims_inventory.example.item
}
