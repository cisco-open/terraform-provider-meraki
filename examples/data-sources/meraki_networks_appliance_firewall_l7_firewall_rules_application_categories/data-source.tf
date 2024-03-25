
data "meraki_networks_appliance_firewall_l7_firewall_rules_application_categories" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_firewall_l7_firewall_rules_application_categories_example" {
  value = data.meraki_networks_appliance_firewall_l7_firewall_rules_application_categories.example.item
}
