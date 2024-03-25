
data "meraki_networks_wireless_billing" "example" {

  network_id = "string"
}

output "meraki_networks_wireless_billing_example" {
  value = data.meraki_networks_wireless_billing.example.item
}
