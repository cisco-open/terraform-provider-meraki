
resource "meraki_networks_webhooks_payload_templates" "example" {

  body      = "{'event_type':'{{alertTypeId}}','client_payload':{'text':'{{alertData}}'}}"
  body_file = "Qm9keSBGaWxl"
  headers = [{

    name     = "Authorization"
    template = "Bearer {{sharedSecret}}"
  }]
  headers_file = "SGVhZGVycyBGaWxl"
  name         = "Custom Template"
  network_id   = "string"
}

output "meraki_networks_webhooks_payload_templates_example" {
  value = meraki_networks_webhooks_payload_templates.example
}