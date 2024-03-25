
data "meraki_networks_appliance_firewall_cellular_firewall_rules" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_firewall_cellular_firewall_rules_example" {
  value = data.meraki_networks_appliance_firewall_cellular_firewall_rules.example.item
}
