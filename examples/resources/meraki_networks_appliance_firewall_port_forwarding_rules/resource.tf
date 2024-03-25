
resource "meraki_networks_appliance_firewall_port_forwarding_rules" "example" {

  network_id = "string"
  rules = [{

    allowed_ips = ["any"]
    lan_ip      = "192.168.128.1"
    local_port  = "442-443"
    name        = "Description of Port Forwarding Rule"
    protocol    = "tcp"
    public_port = "8100-8101"
    uplink      = "both"
  }]
}

output "meraki_networks_appliance_firewall_port_forwarding_rules_example" {
  value = meraki_networks_appliance_firewall_port_forwarding_rules.example
}