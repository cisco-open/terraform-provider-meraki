
resource "meraki_organizations_insight_monitored_media_servers" "example" {

  address                        = "123.123.123.1"
  best_effort_monitoring_enabled = true
  name                           = "Sample VoIP Provider"
  organization_id                = "string"
}

output "meraki_organizations_insight_monitored_media_servers_example" {
  value = meraki_organizations_insight_monitored_media_servers.example
}