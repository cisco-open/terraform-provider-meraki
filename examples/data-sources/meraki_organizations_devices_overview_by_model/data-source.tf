
data "meraki_organizations_devices_overview_by_model" "example" {

  models          = ["string"]
  network_ids     = ["string"]
  organization_id = "string"
  product_types   = ["string"]
}

output "meraki_organizations_devices_overview_by_model_example" {
  value = data.meraki_organizations_devices_overview_by_model.example.item
}
