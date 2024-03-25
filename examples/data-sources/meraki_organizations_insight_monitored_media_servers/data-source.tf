
data "meraki_organizations_insight_monitored_media_servers" "example" {

  organization_id = "string"
}

output "meraki_organizations_insight_monitored_media_servers_example" {
  value = data.meraki_organizations_insight_monitored_media_servers.example.items
}

data "meraki_organizations_insight_monitored_media_servers" "example" {

  organization_id = "string"
}

output "meraki_organizations_insight_monitored_media_servers_example" {
  value = data.meraki_organizations_insight_monitored_media_servers.example.item
}
