
resource "meraki_devices_switch_warm_spare" "example" {

  enabled      = true
  serial       = "string"
  spare_serial = "Q234-ABCD-0002"
}

output "meraki_devices_switch_warm_spare_example" {
  value = meraki_devices_switch_warm_spare.example
}