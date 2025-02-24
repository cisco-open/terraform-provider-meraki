terraform {
  required_providers {
    meraki = {
      version = "1.0.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_networks_wireless_ssids_firewall_l3_firewall_rules" "test-Terraform" {
  network_id       = "L_828099381482776088"
  number           = "1"
  allow_lan_access = true
  rules = [
    {
      # comment = ""
      dest_cidr = "0.0.0.0/22"
      dest_port = "any"
      policy    = "allow"
      protocol  = "any"
      ip_ver    = "ipv4"
    },
    {
      # comment = ""
      dest_cidr = "0.0.0.0/12"
      dest_port = "any"
      policy    = "deny"
      protocol  = "any"
      ip_ver    = "ipv4"
    },
    {
      # comment = ""
      dest_cidr = "0.0.0.0/16"
      dest_port = "any"
      policy    = "allow"
      protocol  = "any"
      ip_ver    = "ipv4"
    },
    {
      # comment = ""
      dest_cidr = "10.0.0.0/8"
      dest_port = "any"
      policy    = "deny"
      protocol  = "any"
      ip_ver    = "ipv4"
    }
  ]
}
