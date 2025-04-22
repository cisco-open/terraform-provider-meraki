
resource "meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_update" "example" {

  entry_id        = "string"
  organization_id = "string"
  parameters = {

    client = {

      mac = "A1:B2:C3:D4:E5:F6"
    }
    description = "Example mac address"
  }
}

output "meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_update_example" {
  value = meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_update.example
}