
resource "meraki_networks_appliance_firewall_l7_firewall_rules" "example" {

  network_id = "string"
  rules = [{

    policy = "deny"
    type   = "host"
    value  = "google.com"
  }]
}

output "meraki_networks_appliance_firewall_l7_firewall_rules_example" {
  value = meraki_networks_appliance_firewall_l7_firewall_rules.example
}