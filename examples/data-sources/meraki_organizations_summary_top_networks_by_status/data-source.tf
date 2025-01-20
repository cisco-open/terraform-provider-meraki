
data "meraki_organizations_summary_top_networks_by_status" "example" {

  device_tag      = "string"
  ending_before   = "string"
  network_tag     = "string"
  organization_id = "string"
  per_page        = 1
  quantity        = 1
  ssid_name       = "string"
  starting_after  = "string"
  usage_uplink    = "string"
}

output "meraki_organizations_summary_top_networks_by_status_example" {
  value = data.meraki_organizations_summary_top_networks_by_status.example.items
}
