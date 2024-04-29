
resource "meraki_networks_sm_devices_shutdown" "example" {

  network_id = "string"
  parameters = {

    ids       = ["1284392014819", "2983092129865"]
    scope     = ["withAny", "tag1", "tag2"]
    serials   = ["XY0XX0Y0X0", "A01B01CD00E", "X02YZ1ZYZX"]
    wifi_macs = ["00:11:22:33:44:55"]
  }
}

output "meraki_networks_sm_devices_shutdown_example" {
  value = meraki_networks_sm_devices_shutdown.example
}