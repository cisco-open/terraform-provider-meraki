
resource "meraki_networks_appliance_firewall_one_to_many_nat_rules" "example" {

  network_id = "string"
  rules = [{

    port_rules = [{

      allowed_ips = ["any"]
      local_ip    = "192.168.128.1"
      local_port  = "443"
      name        = "Rule 1"
      protocol    = "tcp"
      public_port = "9443"
    }]
    public_ip = "146.11.11.13"
    uplink    = "internet1"
  }]
}

output "meraki_networks_appliance_firewall_one_to_many_nat_rules_example" {
  value = meraki_networks_appliance_firewall_one_to_many_nat_rules.example
}