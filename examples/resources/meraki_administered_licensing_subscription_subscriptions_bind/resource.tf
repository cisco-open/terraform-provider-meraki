
resource "meraki_administered_licensing_subscription_subscriptions_bind" "example" {

  subscription_id = "string"
  validate        = false
  parameters = {

    network_ids = ["L_1234", "N_5678"]
  }
}

output "meraki_administered_licensing_subscription_subscriptions_bind_example" {
  value = meraki_administered_licensing_subscription_subscriptions_bind.example
}