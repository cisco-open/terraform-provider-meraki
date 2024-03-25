
data "meraki_networks_webhooks_webhook_tests" "example" {

  network_id      = "string"
  webhook_test_id = "string"
}

output "meraki_networks_webhooks_webhook_tests_example" {
  value = data.meraki_networks_webhooks_webhook_tests.example.item
}
