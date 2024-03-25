
data "meraki_organizations_devices_statuses" "example" {

  ending_before    = "string"
  models           = ["string"]
  network_ids      = ["string"]
  organization_id  = "string"
  per_page         = 1
  product_types    = ["string"]
  serials          = ["string"]
  starting_after   = "string"
  statuses         = ["string"]
  tags             = ["string"]
  tags_filter_type = "string"
}

output "meraki_organizations_devices_statuses_example" {
  value = data.meraki_organizations_devices_statuses.example.item
}
