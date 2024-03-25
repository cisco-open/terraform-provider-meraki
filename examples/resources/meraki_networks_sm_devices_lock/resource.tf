
resource "meraki_networks_sm_devices_lock" "example" {

  network_id = "string"
  parameters = {

    ids       = ["1284392014819", "2983092129865"]
    pin       = 123456
    scope     = ["withAny", "tag1", "tag2"]
    serials   = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
    wifi_macs = ["00:11:22:33:44:55"]
  }
}

output "meraki_networks_sm_devices_lock_example" {
  value = meraki_networks_sm_devices_lock.example
}