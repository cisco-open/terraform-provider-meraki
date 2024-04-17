
resource "meraki_networks_wireless_billing" "example" {

  currency   = "USD"
  network_id = "string"
  plans = [{

    bandwidth_limits = {

      limit_down = 1000000
      limit_up   = 1000000
    }
    id         = "1"
    price      = 5.0
    time_limit = "1 hour"
  }]
}

output "meraki_networks_wireless_billing_example" {
  value = meraki_networks_wireless_billing.example
}