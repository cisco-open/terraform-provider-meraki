
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

resource "meraki_networks_appliance_firewall_l7_firewall_rules" "example2" {

  network_id = "string"
  rules = [{
    policy = "deny"
    type = "blockedCountries"
    value_list = ["IT", "IL", "US"]
  }]
}

output "meraki_networks_appliance_firewall_l7_firewall_rules_example2" {
  value = meraki_networks_appliance_firewall_l7_firewall_rules.example2
}

resource "meraki_networks_appliance_firewall_l7_firewall_rules" "example3" {

  network_id = "string"
  rules = [{

    policy = "deny"
    type   = "applicationCategory"
    # value                   = "10.11.12.00/24"
    value_obj = {
      name = "Sports"
      id   = "meraki:layer7/category/5"
    }
  }]
}

output "meraki_networks_appliance_firewall_l7_firewall_rules_example3" {
  value = meraki_networks_appliance_firewall_l7_firewall_rules.example3
}