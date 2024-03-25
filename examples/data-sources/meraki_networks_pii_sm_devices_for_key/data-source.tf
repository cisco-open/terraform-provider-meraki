
data "meraki_networks_pii_sm_devices_for_key" "example" {

  bluetooth_mac = "string"
  email         = "string"
  imei          = "string"
  mac           = "string"
  network_id    = "string"
  serial        = "string"
  username      = "string"
}

output "meraki_networks_pii_sm_devices_for_key_example" {
  value = data.meraki_networks_pii_sm_devices_for_key.example.item
}
