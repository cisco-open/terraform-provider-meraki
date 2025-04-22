
data "meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  ssids           = ["string"]
  starting_after  = "string"
}

output "meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_example" {
  value = data.meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries.example.item
}
