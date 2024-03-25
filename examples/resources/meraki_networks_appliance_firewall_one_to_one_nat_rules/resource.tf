
resource "meraki_networks_appliance_firewall_one_to_one_nat_rules" "example" {

  network_id = "string"
  rules = [{

    allowed_inbound = [{

      allowed_ips       = ["10.82.112.0/24", "10.82.0.0/16"]
      destination_ports = ["80"]
      protocol          = "tcp"
    }]
    lan_ip    = "192.168.128.22"
    name      = "Service behind NAT"
    public_ip = "146.12.3.33"
    uplink    = "internet1"
  }]
}

output "meraki_networks_appliance_firewall_one_to_one_nat_rules_example" {
  value = meraki_networks_appliance_firewall_one_to_one_nat_rules.example
}