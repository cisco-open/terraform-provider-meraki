
data "meraki_administered_licensing_subscription_entitlements" "example" {

  skus = ["string"]
}

output "meraki_administered_licensing_subscription_entitlements_example" {
  value = data.meraki_administered_licensing_subscription_entitlements.example.item
}
