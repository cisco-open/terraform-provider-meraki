
data "meraki_devices_cellular_gateway_lan" "example" {

  serial = "string"
}

output "meraki_devices_cellular_gateway_lan_example" {
  value = data.meraki_devices_cellular_gateway_lan.example.item
}
