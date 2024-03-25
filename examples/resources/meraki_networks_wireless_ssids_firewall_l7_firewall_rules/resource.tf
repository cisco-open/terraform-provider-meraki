
resource "meraki_networks_wireless_ssids_firewall_l7_firewall_rules" "example" {

  network_id = "string"
  number     = "string"
  rules = [{

    policy = "deny"
    type   = "host"
    value  = "google.com"
  }]
}

output "meraki_networks_wireless_ssids_firewall_l7_firewall_rules_example" {
  value = meraki_networks_wireless_ssids_firewall_l7_firewall_rules.example
}