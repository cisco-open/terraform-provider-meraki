
data "meraki_organizations_wireless_ssids_statuses_by_device" "example" {

  bssids          = ["string"]
  ending_before   = "string"
  hide_disabled   = false
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
}

output "meraki_organizations_wireless_ssids_statuses_by_device_example" {
  value = data.meraki_organizations_wireless_ssids_statuses_by_device.example.item
}
