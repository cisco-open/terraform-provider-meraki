
data "meraki_networks_appliance_firewall_one_to_many_nat_rules" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_firewall_one_to_many_nat_rules_example" {
  value = data.meraki_networks_appliance_firewall_one_to_many_nat_rules.example.item
}
