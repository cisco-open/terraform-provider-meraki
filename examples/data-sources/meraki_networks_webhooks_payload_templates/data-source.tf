
data "meraki_networks_webhooks_payload_templates" "example" {

  network_id = "string"
}

output "meraki_networks_webhooks_payload_templates_example" {
  value = data.meraki_networks_webhooks_payload_templates.example.items
}

data "meraki_networks_webhooks_payload_templates" "example" {

  network_id = "string"
}

output "meraki_networks_webhooks_payload_templates_example" {
  value = data.meraki_networks_webhooks_payload_templates.example.item
}
