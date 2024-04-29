
resource "meraki_administered_licensing_subscription_subscriptions_claim" "example" {

  validate = false
  parameters = {

    claim_key       = "S2345-6789A-BCDEF-GHJKM"
    description     = "Subscription for all main offices"
    name            = "Corporate subscription"
    organization_id = "12345678910"
  }
}

output "meraki_administered_licensing_subscription_subscriptions_claim_example" {
  value = meraki_administered_licensing_subscription_subscriptions_claim.example
}