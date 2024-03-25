
resource "meraki_organizations_alerts_profiles" "example" {

  alert_condition = {

    bit_rate_bps = 10000
    duration     = 60
    interface    = "wan1"
    window       = 600
  }
  description     = "WAN 1 high utilization"
  network_tags    = ["tag1", "tag2"]
  organization_id = "string"
  recipients = {

    emails          = ["admin@example.org"]
    http_server_ids = ["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vcGF0aA=="]
  }
  type = "wanUtilization"
}

output "meraki_organizations_alerts_profiles_example" {
  value = meraki_organizations_alerts_profiles.example
}