
resource "meraki_devices_cellular_gateway_port_forwarding_rules" "example" {

  rules = [{

    access      = "any"
    lan_ip      = "172.31.128.5"
    local_port  = "4"
    name        = "test"
    protocol    = "tcp"
    public_port = "11-12"
  }]
  serial = "string"
}

output "meraki_devices_cellular_gateway_port_forwarding_rules_example" {
  value = meraki_devices_cellular_gateway_port_forwarding_rules.example
}