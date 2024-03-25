
data "meraki_devices_switch_ports" "example" {

  serial = "string"
}

output "meraki_devices_switch_ports_example" {
  value = data.meraki_devices_switch_ports.example.items
}

data "meraki_devices_switch_ports" "example" {

  serial = "string"
}

output "meraki_devices_switch_ports_example" {
  value = data.meraki_devices_switch_ports.example.item
}
