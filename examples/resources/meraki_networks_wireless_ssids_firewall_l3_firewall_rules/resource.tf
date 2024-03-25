
resource "meraki_networks_wireless_ssids_firewall_l3_firewall_rules" "example" {

  network_id = "string"
  number     = "string"
  rules = [{

    comment   = "Allow TCP traffic to subnet with HTTP servers."
    dest_cidr = "192.168.1.0/24"
    dest_port = "443"
    policy    = "allow"
    protocol  = "tcp"
  }]
}

output "meraki_networks_wireless_ssids_firewall_l3_firewall_rules_example" {
  value = meraki_networks_wireless_ssids_firewall_l3_firewall_rules.example
}