
resource "meraki_networks_wireless_settings" "example" {

  ipv6_bridge_enabled        = false
  led_lights_on              = false
  location_analytics_enabled = false
  meshing_enabled            = true
  named_vlans = {

    pool_dhcp_monitoring = {

      duration = 5
      enabled  = true
    }
  }
  network_id      = "string"
  upgradestrategy = "minimizeUpgradeTime"
}

output "meraki_networks_wireless_settings_example" {
  value = meraki_networks_wireless_settings.example
}