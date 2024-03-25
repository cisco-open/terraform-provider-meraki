
resource "meraki_networks_appliance_settings" "example" {

  client_tracking_method = "MAC address"
  deployment_mode        = "routed"
  dynamic_dns = {

    enabled = true
    prefix  = "test"
  }
  network_id = "string"
}

output "meraki_networks_appliance_settings_example" {
  value = meraki_networks_appliance_settings.example
}