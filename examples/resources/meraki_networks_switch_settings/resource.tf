
resource "meraki_networks_switch_settings" "example" {

  mac_blocklist = {

    enabled = false
  }
  network_id = "string"
  power_exceptions = [{

    power_type = "string"
    serial     = "string"
  }]
  uplink_client_sampling = {

    enabled = false
  }
  use_combined_power = false
  vlan               = 1
}

output "meraki_networks_switch_settings_example" {
  value = meraki_networks_switch_settings.example
}