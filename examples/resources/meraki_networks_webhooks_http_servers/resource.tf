
resource "meraki_networks_webhooks_http_servers" "example" {

  name       = "Example Webhook Server"
  network_id = "string"
  payload_template = {

    name                = "Meraki (included)"
    payload_template_id = "wpt_00001"
  }
  shared_secret = "shhh"
  url           = "https://example.com"
}

output "meraki_networks_webhooks_http_servers_example" {
  value = meraki_networks_webhooks_http_servers.example
}