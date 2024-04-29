
resource "meraki_administered_licensing_subscription_subscriptions_claim_key_validate" "example" {

  parameters = {

    claim_key = "S2345-6789A-BCDEF-GHJKM"
  }
}

output "meraki_administered_licensing_subscription_subscriptions_claim_key_validate_example" {
  value = meraki_administered_licensing_subscription_subscriptions_claim_key_validate.example
}