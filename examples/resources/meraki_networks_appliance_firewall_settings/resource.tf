
resource "meraki_networks_appliance_firewall_settings" "example" {

  network_id = "string"
  spoofing_protection = {

    ip_source_guard = {

      mode = "block"
    }
  }
}

output "meraki_networks_appliance_firewall_settings_example" {
  value = meraki_networks_appliance_firewall_settings.example
}