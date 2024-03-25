
resource "meraki_networks_wireless_bluetooth_settings" "example" {

  advertising_enabled         = true
  major                       = 1
  major_minor_assignment_mode = "Non-unique"
  minor                       = 1
  network_id                  = "string"
  scanning_enabled            = true
  uuid                        = "00000000-0000-0000-000-000000000000"
}

output "meraki_networks_wireless_bluetooth_settings_example" {
  value = meraki_networks_wireless_bluetooth_settings.example
}