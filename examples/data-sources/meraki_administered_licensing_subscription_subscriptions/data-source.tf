
data "meraki_administered_licensing_subscription_subscriptions" "example" {

  end_date         = "string"
  ending_before    = "string"
  name             = "string"
  organization_ids = ["string"]
  per_page         = 1
  product_types    = ["string"]
  start_date       = "string"
  starting_after   = "string"
  statuses         = ["string"]
  subscription_ids = ["string"]
}

output "meraki_administered_licensing_subscription_subscriptions_example" {
  value = data.meraki_administered_licensing_subscription_subscriptions.example.items
}
