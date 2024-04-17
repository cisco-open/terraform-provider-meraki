
data "meraki_administered_licensing_subscription_subscriptions_compliance_statuses" "example" {

  organization_ids = ["string"]
  subscription_ids = ["string"]
}

output "meraki_administered_licensing_subscription_subscriptions_compliance_statuses_example" {
  value = data.meraki_administered_licensing_subscription_subscriptions_compliance_statuses.example.items
}
