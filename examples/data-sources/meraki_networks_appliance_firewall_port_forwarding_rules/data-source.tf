
data "meraki_networks_appliance_firewall_port_forwarding_rules" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_firewall_port_forwarding_rules_example" {
  value = data.meraki_networks_appliance_firewall_port_forwarding_rules.example.item
}
