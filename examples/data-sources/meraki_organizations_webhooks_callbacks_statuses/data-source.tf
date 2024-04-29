
data "meraki_organizations_webhooks_callbacks_statuses" "example" {

  callback_id     = "string"
  organization_id = "string"
}

output "meraki_organizations_webhooks_callbacks_statuses_example" {
  value = data.meraki_organizations_webhooks_callbacks_statuses.example.item
}
