
data "meraki_networks_webhooks_http_servers" "example" {

  network_id = "string"
}

output "meraki_networks_webhooks_http_servers_example" {
  value = data.meraki_networks_webhooks_http_servers.example.items
}

data "meraki_networks_webhooks_http_servers" "example" {

  network_id = "string"
}

output "meraki_networks_webhooks_http_servers_example" {
  value = data.meraki_networks_webhooks_http_servers.example.item
}
