
data "meraki_organizations_firmware_upgrades_by_device" "example" {

  ending_before              = "string"
  firmware_upgrade_batch_ids = ["string"]
  firmware_upgrade_ids       = ["string"]
  macs                       = ["string"]
  network_ids                = ["string"]
  organization_id            = "string"
  per_page                   = 1
  serials                    = ["string"]
  starting_after             = "string"
}

output "meraki_organizations_firmware_upgrades_by_device_example" {
  value = data.meraki_organizations_firmware_upgrades_by_device.example.items
}
