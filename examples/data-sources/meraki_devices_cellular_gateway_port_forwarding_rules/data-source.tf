
data "meraki_devices_cellular_gateway_port_forwarding_rules" "example" {

  serial = "string"
}

output "meraki_devices_cellular_gateway_port_forwarding_rules_example" {
  value = data.meraki_devices_cellular_gateway_port_forwarding_rules.example.item
}
