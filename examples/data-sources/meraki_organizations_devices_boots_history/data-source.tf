
data "meraki_organizations_devices_boots_history" "example" {

  ending_before          = "string"
  most_recent_per_device = false
  organization_id        = "string"
  per_page               = 1
  serials                = ["string"]
  sort_order             = "string"
  starting_after         = "string"
  t0                     = "string"
  t1                     = "string"
  timespan               = 1.0
}

output "meraki_organizations_devices_boots_history_example" {
  value = data.meraki_organizations_devices_boots_history.example.items
}
