
data "meraki_organizations_devices_statuses_overview" "example" {

  network_ids     = ["string"]
  organization_id = "string"
  product_types   = ["string"]
}

output "meraki_organizations_devices_statuses_overview_example" {
  value = data.meraki_organizations_devices_statuses_overview.example.item
}
