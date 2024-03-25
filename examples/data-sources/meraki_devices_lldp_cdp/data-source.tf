
data "meraki_devices_lldp_cdp" "example" {

  serial = "string"
}

output "meraki_devices_lldp_cdp_example" {
  value = data.meraki_devices_lldp_cdp.example.item
}
