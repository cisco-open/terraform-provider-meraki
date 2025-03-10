terraform {
  required_providers {
    meraki = {
      version = "1.0.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_appliance_firewall_one_to_many_nat_rules" "this" {
  network_id = "L_828099381482771185"
  rules = [{
    port_rules = [{

      allowed_ips = ["any"]
      local_ip    = "192.168.128.1"
      local_port  = "443"
      name        = "Rule 2"
      protocol    = "tcp"
      public_port = "9443"
    }]
    public_ip = "146.11.11.13"
    uplink    = "internet2"
    },

    { port_rules = [{

      allowed_ips = ["any"]
      local_ip    = "202.168.128.1"
      local_port  = "443"
      name        = "Rule 1"
      protocol    = "tcp"
      public_port = "9443"
      }]
      public_ip = "10.11.11.13"
      uplink    = "internet1"
    },
  ]
}