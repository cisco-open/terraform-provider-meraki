
data "meraki_networks_wireless_ssids_firewall_l7_firewall_rules" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_firewall_l7_firewall_rules_example" {
  value = data.meraki_networks_wireless_ssids_firewall_l7_firewall_rules.example.item
}
