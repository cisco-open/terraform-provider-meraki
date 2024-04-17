
resource "meraki_networks_sm_devices_wipe" "example" {

  network_id = "string"
  parameters = {

    id       = "1284392014819"
    pin      = 123456
    serial   = "XY0XX0Y0X0"
    wifi_mac = "00:11:22:33:44:55"
  }
}

output "meraki_networks_sm_devices_wipe_example" {
  value = meraki_networks_sm_devices_wipe.example
}