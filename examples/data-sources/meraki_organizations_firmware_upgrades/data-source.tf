
data "meraki_organizations_firmware_upgrades" "example" {

  organization_id = "string"
  product_type    = ["string"]
  status          = ["string"]
}

output "meraki_organizations_firmware_upgrades_example" {
  value = data.meraki_organizations_firmware_upgrades.example.items
}
