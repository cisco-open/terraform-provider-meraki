terraform {
  required_providers {
    meraki = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_administered_licensing_subscription_subscriptions_claim" "example" {

  validate = false
  parameters = {

    claim_key       = "S2345-6789A-BCDEF-GHJKM"
    description     = "Subscription for all main offices"
    name            = "Corporate subscription"
    organization_id = "828099381482762270"
  }
}

output "meraki_administered_licensing_subscription_subscriptions_claim_example" {
  value = meraki_administered_licensing_subscription_subscriptions_claim.example
}