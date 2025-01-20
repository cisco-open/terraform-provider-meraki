
data "meraki_organizations_wireless_controller_devices_redundancy_statuses" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
}

output "meraki_organizations_wireless_controller_devices_redundancy_statuses_example" {
  value = data.meraki_organizations_wireless_controller_devices_redundancy_statuses.example.item
}
