
data "meraki_networks_pii_pii_keys" "example" {

  bluetooth_mac = "string"
  email         = "string"
  imei          = "string"
  mac           = "string"
  network_id    = "string"
  serial        = "string"
  username      = "string"
}

output "meraki_networks_pii_pii_keys_example" {
  value = data.meraki_networks_pii_pii_keys.example.item
}
