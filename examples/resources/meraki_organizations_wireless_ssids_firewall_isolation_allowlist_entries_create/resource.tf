
resource "meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_create" "example" {

  organization_id = "string"
  parameters = {

    client = {

      mac = "A1:B2:C3:D4:E5:F6"
    }
    description = "Example mac address"
    network = {

      id = "N_123"
    }
    ssid = {

      number = 2
    }
  }
}

output "meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_create_example" {
  value = meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_create.example
}