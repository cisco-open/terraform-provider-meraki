
data "meraki_organizations_wireless_controller_overview_by_device" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
}

output "meraki_organizations_wireless_controller_overview_by_device_example" {
  value = data.meraki_organizations_wireless_controller_overview_by_device.example.item
}
