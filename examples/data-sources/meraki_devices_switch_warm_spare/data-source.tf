
data "meraki_devices_switch_warm_spare" "example" {

  serial = "string"
}

output "meraki_devices_switch_warm_spare_example" {
  value = data.meraki_devices_switch_warm_spare.example.item
}
