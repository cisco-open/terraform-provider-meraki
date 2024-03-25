
data "meraki_organizations_inventory_devices" "example" {

  ending_before    = "string"
  macs             = ["string"]
  models           = ["string"]
  network_ids      = ["string"]
  order_numbers    = ["string"]
  organization_id  = "string"
  per_page         = 1
  product_types    = ["string"]
  search           = "string"
  serials          = ["string"]
  starting_after   = "string"
  tags             = ["string"]
  tags_filter_type = "string"
  used_state       = "string"
}

output "meraki_organizations_inventory_devices_example" {
  value = data.meraki_organizations_inventory_devices.example.items
}

data "meraki_organizations_inventory_devices" "example" {

  ending_before    = "string"
  macs             = ["string"]
  models           = ["string"]
  network_ids      = ["string"]
  order_numbers    = ["string"]
  organization_id  = "string"
  per_page         = 1
  product_types    = ["string"]
  search           = "string"
  serials          = ["string"]
  starting_after   = "string"
  tags             = ["string"]
  tags_filter_type = "string"
  used_state       = "string"
}

output "meraki_organizations_inventory_devices_example" {
  value = data.meraki_organizations_inventory_devices.example.item
}
