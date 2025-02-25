
data "meraki_organizations_firmware_upgrades_by_device" "example" {

  current_upgrades_only      = false
  ending_before              = "string"
  firmware_upgrade_batch_ids = ["string"]
  macs                       = ["string"]
  network_ids                = ["string"]
  organization_id            = "string"
  per_page                   = 1
  serials                    = ["string"]
  starting_after             = "string"
  upgradestatuses            = ["string"]
}

output "meraki_organizations_firmware_upgrades_by_device_example" {
  value = data.meraki_organizations_firmware_upgrades_by_device.example.items
}
