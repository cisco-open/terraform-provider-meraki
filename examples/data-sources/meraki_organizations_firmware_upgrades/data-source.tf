
data "meraki_organizations_firmware_upgrades" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  product_types   = ["string"]
  starting_after  = "string"
  status          = ["string"]
}

output "meraki_organizations_firmware_upgrades_example" {
  value = data.meraki_organizations_firmware_upgrades.example.items
}
