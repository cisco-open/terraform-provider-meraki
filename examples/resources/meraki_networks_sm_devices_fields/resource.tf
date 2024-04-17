
resource "meraki_networks_sm_devices_fields" "example" {

  network_id = "string"
  parameters = {

    device_fields = {

      name  = "Miles's phone"
      notes = "Here's some info about my device"
    }
    id       = "1284392014819"
    serial   = "XY0XX0Y0X0"
    wifi_mac = "00:11:22:33:44:55"
  }
}

output "meraki_networks_sm_devices_fields_example" {
  value = meraki_networks_sm_devices_fields.example
}